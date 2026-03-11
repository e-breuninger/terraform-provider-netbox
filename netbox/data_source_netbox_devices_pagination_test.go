package netbox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestDataSourceNetboxDevicesPagination verifies that the devices data source fetches
// all pages when total results exceed the page size, using a mock HTTP server.
//
// Without pagination the data source would return at most DefaultPageSize items;
// this test uses 130 items (100 + 30, not divisible by page size) so the last page
// is always partial, exercising the boundary condition.
func TestDataSourceNetboxDevicesPagination(t *testing.T) {
	const totalDevices = 130

	requestCount := 0

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/dcim/devices/" {
			http.NotFound(w, r)
			return
		}

		requestCount++

		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit := int(DefaultPageSize)
		if s := r.URL.Query().Get("limit"); s != "" {
			limit, _ = strconv.Atoi(s)
		}

		end := offset + limit
		if end > totalDevices {
			end = totalDevices
		}

		var nextURL interface{}
		if end < totalDevices {
			nextURL = fmt.Sprintf("http://%s/api/dcim/devices/?limit=%d&offset=%d", r.Host, limit, end)
		}

		results := make([]map[string]interface{}, 0, end-offset)
		for i := offset; i < end; i++ {
			results = append(results, minimalDeviceJSON(i+1, fmt.Sprintf("device-%d", i+1)))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count":    totalDevices,
			"next":     nextURL,
			"previous": nil,
			"results":  results,
		})
	}))
	defer ts.Close()

	cfg := Config{APIToken: "test-token", ServerURL: ts.URL}
	client, err := cfg.Client()
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	state := &providerState{NetBoxAPI: client}
	d := schema.TestResourceDataRaw(t, dataSourceNetboxDevices().Schema, map[string]interface{}{})

	if err := dataSourceNetboxDevicesRead(d, state); err != nil {
		t.Fatalf("dataSourceNetboxDevicesRead returned error: %v", err)
	}

	devices := d.Get("devices").([]interface{})
	if len(devices) != totalDevices {
		t.Errorf("got %d devices, want %d (pagination likely broken)", len(devices), totalDevices)
	}
	if requestCount < 2 {
		t.Errorf("expected multiple API requests (pagination), got %d", requestCount)
	}
}

// minimalDeviceJSON returns the minimum fields required by the go-netbox swagger
// validator for a DeviceWithConfigContext response item.
func minimalDeviceJSON(id int, name string) map[string]interface{} {
	return map[string]interface{}{
		"id":   id,
		"name": name,
		"device_type": map[string]interface{}{
			"id":    1,
			"model": "test-model",
			"slug":  "test-model",
			"manufacturer": map[string]interface{}{
				"id":   1,
				"name": "test-manufacturer",
				"slug": "test-manufacturer",
			},
		},
		"role": map[string]interface{}{
			"id":   1,
			"name": "test-role",
			"slug": "test-role",
		},
		"site": map[string]interface{}{
			"id":   1,
			"name": "test-site",
			"slug": "test-site",
		},
		"tags": []interface{}{},
	}
}
