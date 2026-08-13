package netbox

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

// Netbox hands out "next available" objects (prefixes, IPs, VLANs) from a
// shared pool. Terraform creates resources in parallel, so several resources
// carving out of the same parent would otherwise race each other on that pool
// and get conflicts back from Netbox.
//
// Two mechanisms guard against this:
//
//   - lockAllocation serializes allocations that target the same parent, so
//     this provider never competes with itself.
//   - retryAllocation retries conflicts that come from outside this process,
//     such as a second terraform run or somebody using the Netbox UI.

const (
	// allocationRetryTimeout bounds the total time spent retrying a single
	// allocation. Requests are cut off by the provider's request_timeout
	// (10 seconds by default), so this is a budget for several attempts
	// rather than one long wait.
	allocationRetryTimeout = 2 * time.Minute

	// allocationRetryMinDelay and allocationRetryMaxDelay bound the backoff
	// between attempts.
	allocationRetryMinDelay = 500 * time.Millisecond
	allocationRetryMaxDelay = 10 * time.Second
)

// allocationLockKeyPrefix builds the lock key for allocations out of a parent
// prefix, either for a child prefix or for an IP address.
func allocationLockKeyPrefix(prefixID int64) string {
	return fmt.Sprintf("prefix:%d", prefixID)
}

// allocationLockKeyIPRange builds the lock key for allocations out of an IP range.
func allocationLockKeyIPRange(rangeID int64) string {
	return fmt.Sprintf("iprange:%d", rangeID)
}

// allocationLockKeyVLANGroup builds the lock key for allocations out of a VLAN group.
func allocationLockKeyVLANGroup(groupID int64) string {
	return fmt.Sprintf("vlangroup:%d", groupID)
}

// lockAllocation serializes allocations against a single parent pool. Keys are
// namespaced by object type, so a prefix and an IP range that happen to share
// an ID do not block each other. Allocations from different parents are
// unaffected and still run in parallel.
//
// It returns the unlock function, so callers can write:
//
//	unlock, err := api.lockAllocation(ctx, key)
//	if err != nil {
//		return diag.FromErr(err)
//	}
//	defer unlock()
//
// Waiting for the lock respects ctx, so a cancelled apply does not sit in the
// queue. The returned error is non-nil only when ctx is done before the lock
// was taken, in which case unlock is a no-op and must still not be called.
func (s *providerState) lockAllocation(ctx context.Context, key string) (func(), error) {
	value, _ := s.allocationLocks.LoadOrStore(key, make(chan struct{}, 1))
	lock := value.(chan struct{})

	select {
	case lock <- struct{}{}:
		var once sync.Once
		return func() { once.Do(func() { <-lock }) }, nil
	case <-ctx.Done():
		return func() {}, ctx.Err()
	}
}

// retryAllocation runs fn, retrying while Netbox reports that the pool is
// contended. It only makes sense around the call that actually allocates the
// object: anything after that point has already consumed a Netbox object, and
// running it twice would leak one.
func retryAllocation(ctx context.Context, fn func() error) error {
	ctx, cancel := context.WithTimeout(ctx, allocationRetryTimeout)
	defer cancel()

	started := time.Now()
	delay := allocationRetryMinDelay
	for {
		err := fn()
		if err == nil || !isRetryableAllocationError(err) {
			return err
		}

		// Jitter keeps callers released by the same conflict from retrying in
		// lockstep and colliding all over again.
		wait := time.Duration(rand.Int63n(int64(delay)) + int64(delay)/2)

		timer := time.NewTimer(wait)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()
			// Report the last error from Netbox rather than the bare deadline,
			// since that is what the user needs to act on.
			return fmt.Errorf("gave up allocating after %s: %w", time.Since(started).Round(time.Second), err)
		}

		if delay < allocationRetryMaxDelay {
			delay *= 2
			if delay > allocationRetryMaxDelay {
				delay = allocationRetryMaxDelay
			}
		}
	}
}

// isRetryableAllocationError reports whether err looks like contention on the
// pool rather than a permanent failure.
//
// Only the HTTP status is used. Netbox runs behind different application
// servers depending on version and deployment, so no assumption is made about
// server side timeouts or error bodies.
func isRetryableAllocationError(err error) bool {
	if err == nil {
		return false
	}

	// A client side timeout is deliberately not retried. Netbox may have
	// committed the allocation and only the response went missing, so retrying
	// would quietly leak an object on every attempt.
	if errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return false
	}

	// Every go-netbox endpoint has its own error type, but they all expose the
	// status code the same way.
	var coder interface{ Code() int }
	if !errors.As(err, &coder) {
		return false
	}

	switch coder.Code() {
	case 409, 429:
		return true
	default:
		return false
	}
}
