package netboxapi

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// TestRetryTransportRepeatsTransientFailures: a 500 (NetBox's answer to a Postgres deadlock
// between a device PATCH and a parallel cable delete) is retried with the body replayed; a POST is
// not, and a 4xx is not.
func TestRetryTransportRepeatsTransientFailures(t *testing.T) {
	var calls atomic.Int32
	var lastBody atomic.Value
	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		n := calls.Add(1)
		buf := new(bytes.Buffer)
		buf.ReadFrom(request.Body)
		lastBody.Store(buf.String())
		switch {
		case request.URL.Path == "/flaky" && n < 3:
			writer.WriteHeader(http.StatusInternalServerError)
		case request.URL.Path == "/bad":
			writer.WriteHeader(http.StatusBadRequest)
		default:
			writer.WriteHeader(http.StatusOK)
		}
	}))
	defer srv.Close()
	client := &http.Client{Transport: &retryTransport{next: http.DefaultTransport, attempts: 3, backoff: time.Millisecond}}
	call := func(method, path string) int {
		t.Helper()
		req, _ := http.NewRequestWithContext(context.Background(), method, srv.URL+path, bytes.NewBufferString(`{"primary_ip4":null}`))
		res, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		return res.StatusCode
	}

	calls.Store(0)
	if code := call(http.MethodPatch, "/flaky"); code != http.StatusOK {
		t.Errorf("PATCH after two 500s: got %d, want 200", code)
	}
	if calls.Load() != 3 {
		t.Errorf("PATCH attempts: got %d, want 3", calls.Load())
	}
	if lastBody.Load() != `{"primary_ip4":null}` {
		t.Errorf("retried request lost its body: %q", lastBody.Load())
	}

	calls.Store(0)
	if code := call(http.MethodPost, "/flaky"); code != http.StatusInternalServerError {
		t.Errorf("POST: got %d, want the 500 passed through", code)
	}
	if calls.Load() != 1 {
		t.Errorf("POST attempts: got %d, want 1", calls.Load())
	}

	calls.Store(0)
	if code := call(http.MethodPatch, "/bad"); code != http.StatusBadRequest {
		t.Errorf("400: got %d", code)
	}
	if calls.Load() != 1 {
		t.Errorf("400 attempts: got %d, want 1", calls.Load())
	}
}
