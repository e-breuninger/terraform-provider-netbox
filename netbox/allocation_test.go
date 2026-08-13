package netbox

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

// codeError stands in for the go-netbox "Default" error types, which all
// expose the HTTP status through a Code method.
type codeError struct {
	code int
}

func (e codeError) Code() int     { return e.code }
func (e codeError) Error() string { return fmt.Sprintf("netbox returned %d", e.code) }

// timeoutError stands in for a client side timeout.
type timeoutError struct{}

func (timeoutError) Error() string   { return "request timed out" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }

func TestLockAllocationSerializesSameKey(t *testing.T) {
	state := &providerState{}
	ctx := context.Background()

	const goroutines = 20
	// Deliberately not atomic: the race detector and the check below only
	// catch overlap if the increment is unguarded.
	counter := 0
	inside := false

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			unlock, err := state.lockAllocation(ctx, allocationLockKeyPrefix(1))
			if err != nil {
				t.Errorf("unexpected error taking lock: %s", err)
				return
			}
			defer unlock()

			if inside {
				t.Error("two goroutines held the same allocation lock at once")
			}
			inside = true
			counter++
			inside = false
		}()
	}
	wg.Wait()

	if counter != goroutines {
		t.Errorf("expected counter to be %d, got %d", goroutines, counter)
	}
}

// TestLockAllocationDifferentKeysAreIndependent deadlocks if the lock is
// global rather than per parent, so it is the test that actually proves
// allocations from different pools still run in parallel.
func TestLockAllocationDifferentKeysAreIndependent(t *testing.T) {
	state := &providerState{}
	ctx := context.Background()

	first := make(chan struct{})
	second := make(chan struct{})
	done := make(chan struct{})

	go func() {
		unlock, err := state.lockAllocation(ctx, allocationLockKeyPrefix(1))
		if err != nil {
			t.Errorf("unexpected error taking lock: %s", err)
			return
		}
		defer unlock()

		close(first)
		<-second
		done <- struct{}{}
	}()

	go func() {
		unlock, err := state.lockAllocation(ctx, allocationLockKeyPrefix(2))
		if err != nil {
			t.Errorf("unexpected error taking lock: %s", err)
			return
		}
		defer unlock()

		<-first
		close(second)
		done <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("timed out: locks for different parents blocked each other")
		}
	}
}

func TestLockAllocationKeysAreNamespaced(t *testing.T) {
	// A prefix and an IP range with the same id must not share a lock.
	if allocationLockKeyPrefix(7) == allocationLockKeyIPRange(7) {
		t.Error("prefix and ip range keys collide")
	}
	if allocationLockKeyPrefix(7) == allocationLockKeyVLANGroup(7) {
		t.Error("prefix and vlan group keys collide")
	}
}

func TestLockAllocationRespectsCancelledContext(t *testing.T) {
	state := &providerState{}

	// Hold the lock so the second caller has to wait for it.
	unlock, err := state.lockAllocation(context.Background(), allocationLockKeyPrefix(1))
	if err != nil {
		t.Fatalf("unexpected error taking lock: %s", err)
	}
	defer unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	if _, err := state.lockAllocation(ctx, allocationLockKeyPrefix(1)); err == nil {
		t.Error("expected an error when the context is cancelled while waiting")
	}
}

func TestIsRetryableAllocationError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"conflict", codeError{409}, true},
		{"too many requests", codeError{429}, true},
		{"bad request", codeError{400}, false},
		{"not found", codeError{404}, false},
		{"forbidden", codeError{403}, false},
		{"server error", codeError{500}, false},
		{"wrapped conflict", fmt.Errorf("creating prefix: %w", codeError{409}), true},
		{"plain error", errors.New("boom"), false},
		{"deadline exceeded", context.DeadlineExceeded, false},
		{"client timeout", timeoutError{}, false},
		{"wrapped timeout", fmt.Errorf("posting: %w", timeoutError{}), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isRetryableAllocationError(test.err); got != test.want {
				t.Errorf("isRetryableAllocationError(%v) = %t, want %t", test.err, got, test.want)
			}
		})
	}
}

// A timeout must stay terminal even though net.Error is also how some
// transport errors surface: Netbox may have committed the allocation already.
func TestRetryAllocationDoesNotRetryTimeouts(t *testing.T) {
	attempts := 0
	err := retryAllocation(context.Background(), func() error {
		attempts++
		return timeoutError{}
	})

	if err == nil {
		t.Fatal("expected an error")
	}
	if attempts != 1 {
		t.Errorf("expected a timeout not to be retried, got %d attempts", attempts)
	}
}

func TestRetryAllocationSucceedsImmediately(t *testing.T) {
	attempts := 0
	err := retryAllocation(context.Background(), func() error {
		attempts++
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if attempts != 1 {
		t.Errorf("expected 1 attempt, got %d", attempts)
	}
}

func TestRetryAllocationRetriesConflictThenSucceeds(t *testing.T) {
	attempts := 0
	err := retryAllocation(context.Background(), func() error {
		attempts++
		if attempts < 3 {
			return codeError{409}
		}
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}

func TestRetryAllocationDoesNotRetryTerminalError(t *testing.T) {
	attempts := 0
	want := codeError{400}
	err := retryAllocation(context.Background(), func() error {
		attempts++
		return want
	})

	if !errors.Is(err, want) {
		t.Errorf("expected the original error back, got %v", err)
	}
	if attempts != 1 {
		t.Errorf("expected 1 attempt, got %d", attempts)
	}
}

func TestRetryAllocationGivesUp(t *testing.T) {
	// Cancel well before allocationRetryTimeout so the test stays quick.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	attempts := 0
	err := retryAllocation(ctx, func() error {
		attempts++
		return codeError{409}
	})

	if err == nil {
		t.Fatal("expected an error once the budget ran out")
	}
	if attempts < 2 {
		t.Errorf("expected several attempts before giving up, got %d", attempts)
	}
}

// Guard against the interface assertion in isRetryableAllocationError silently
// failing to match net.Error.
var _ net.Error = timeoutError{}
