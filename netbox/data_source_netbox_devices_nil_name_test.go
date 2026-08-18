package netbox

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestDataSourceNetboxDevicesReadNameRegexNilName verifies that the name_regex
// filter does not panic when the API returns a device with a null name. NetBox
// can return unnamed devices (identified only by an asset tag), and such devices
// can never match a regex, so they must be skipped instead of dereferencing a
// nil pointer.
func TestDataSourceNetboxDevicesReadNameRegexNilName(t *testing.T) {
	named := minimalDeviceJSON(1, "device-with-name")
	unnamed := minimalDeviceJSON(2, "")
	unnamed["name"] = nil

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/dcim/devices/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count":    2,
			"next":     nil,
			"previous": nil,
			"results":  []map[string]interface{}{named, unnamed},
		})
	}))
	defer ts.Close()

	cfg := Config{APIToken: "test-token", ServerURL: ts.URL}
	client, err := cfg.Client()
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	state := &providerState{NetBoxAPI: client}
	d := schema.TestResourceDataRaw(t, dataSourceNetboxDevices().Schema, map[string]interface{}{
		"name_regex": "device-with-name",
	})

	if err := dataSourceNetboxDevicesRead(d, state); err != nil {
		t.Fatalf("dataSourceNetboxDevicesRead returned error: %v", err)
	}

	devices := d.Get("devices").([]interface{})
	if len(devices) != 1 {
		t.Fatalf("got %d devices, want 1 (the named device matching the regex)", len(devices))
	}
	if mapping := devices[0].(map[string]interface{}); mapping["name"] != "device-with-name" {
		t.Errorf("expected matched device name %q, got %v", "device-with-name", mapping["name"])
	}
}
