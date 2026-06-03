package netbox

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestDataSourceNetboxDevicesReadNilDeviceType verifies that reading a device
// whose device_type is null does not panic. NetBox can return devices without a
// device type, and the data source must skip the device_type-derived attributes
// instead of dereferencing a nil pointer.
func TestDataSourceNetboxDevicesReadNilDeviceType(t *testing.T) {
	device := minimalDeviceJSON(1, "device-without-type")
	device["device_type"] = nil

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/dcim/devices/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count":    1,
			"next":     nil,
			"previous": nil,
			"results":  []map[string]interface{}{device},
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
	if len(devices) != 1 {
		t.Fatalf("got %d devices, want 1", len(devices))
	}
	if mapping := devices[0].(map[string]interface{}); mapping["device_type_id"] != 0 {
		t.Errorf("expected no device_type_id for device without device type, got %v", mapping["device_type_id"])
	}
}
