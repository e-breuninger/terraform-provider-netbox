package netboxapi

import (
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
)

// TestNewBasePath: the API base path is the server URL's path prefix plus /api, a trailing slash
// on the prefix (strip_trailing_slashes_from_url off) notwithstanding.
func TestNewBasePath(t *testing.T) {
	for serverURL, want := range map[string]string{
		"https://netbox.example.com":          "/api",
		"https://netbox.example.com/":         "/api",
		"https://intranet.example.com/netbox": "/netbox/api",
		"http://localhost:8001/netbox/":       "/netbox/api",
	} {
		api, err := New(serverURL, "token", Options{})
		if err != nil {
			t.Fatalf("New(%s): %v", serverURL, err)
		}
		runtime, isRuntime := api.Transport.(*httptransport.Runtime)
		if !isRuntime {
			t.Fatalf("New(%s): transport is %T, not *httptransport.Runtime", serverURL, api.Transport)
		}
		if runtime.BasePath != want {
			t.Errorf("New(%s): base path %q, want %q", serverURL, runtime.BasePath, want)
		}
	}
}
