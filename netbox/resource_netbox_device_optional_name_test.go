package netbox

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceNetboxDeviceOptionalNameValidation(t *testing.T) {
	for _, tc := range []struct {
		name      string
		config    map[string]interface{}
		wantError string
	}{
		{
			name:   "omitted name",
			config: map[string]interface{}{"device_type_id": 1, "role_id": 1, "site_id": 1},
		},
		{
			name:   "named device",
			config: map[string]interface{}{"name": "device-name", "device_type_id": 1, "role_id": 1, "site_id": 1},
		},
		{
			name:   "explicit empty name",
			config: map[string]interface{}{"name": "", "device_type_id": 1, "role_id": 1, "site_id": 1},
		},
		{
			name:      "site remains required",
			config:    map[string]interface{}{"name": "device-name", "device_type_id": 1, "role_id": 1},
			wantError: `"site_id"`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			diags := schema.InternalMap(resourceNetboxDevice().Schema).Validate(terraform.NewResourceConfigRaw(tc.config))
			if tc.wantError == "" {
				if diags.HasError() {
					t.Fatalf("configuration should be valid: %v", diags)
				}
			} else if !diags.HasError() || len(diags) != 1 || !strings.Contains(diags[0].Detail, tc.wantError) {
				t.Fatalf("expected a validation error for %s, got %v", tc.wantError, diags)
			}
		})
	}
}

func TestResourceNetboxDeviceOptionalNameLifecycle(t *testing.T) {
	for _, tc := range []struct {
		name         string
		initialName  string
		updatedName  string
		initialReply interface{}
		updatedReply interface{}
	}{
		{"unnamed to named", "", "device-name", nil, "device-name"},
		{"named to omitted", "device-name", "", "device-name", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			api := deviceNameTestAPI(t, []deviceNameRequest{
				{http.MethodPost, tc.initialName, tc.initialReply},
				{http.MethodGet, "", tc.initialReply},
				{http.MethodPut, tc.updatedName, tc.updatedReply},
				{http.MethodGet, "", tc.updatedReply},
			})
			resource := Provider().ResourcesMap["netbox_device"]
			ctx := context.Background()
			var state *terraform.InstanceState

			for _, name := range []string{tc.initialName, tc.updatedName} {
				config := terraform.NewResourceConfigRaw(deviceNameTestConfig(name))
				if diags := resource.Validate(config); diags.HasError() {
					t.Fatalf("configuration should be valid: %v", diags)
				}
				diff, err := resource.Diff(ctx, state, config, api)
				if err != nil {
					t.Fatalf("diff failed: %v", err)
				}
				if diff == nil {
					t.Fatal("expected a device create or name change")
				}
				if state != nil {
					change := diff.Attributes["name"]
					if change == nil || change.Old != tc.initialName || change.New != name {
						t.Fatalf("expected name change %q -> %q, got %#v", tc.initialName, name, change)
					}
					if diff.RequiresNew() {
						t.Fatal("changing the name must not replace the device")
					}
				}
				next, diags := resource.Apply(ctx, state, diff, api)
				if diags.HasError() {
					t.Fatalf("apply failed: %v", diags)
				}
				state = next
				checkDeviceNameState(t, state, name)
				checkDeviceNameNoDiff(t, resource, state, config, api)
			}
		})
	}
}

func TestResourceNetboxDeviceOptionalNameRefreshAndImport(t *testing.T) {
	for _, tc := range []struct {
		name  string
		reply interface{}
	}{
		{"null", nil},
		{"empty", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			api := deviceNameTestAPI(t, []deviceNameRequest{
				{http.MethodGet, "", tc.reply},
				{http.MethodGet, "", tc.reply},
			})
			resource := Provider().ResourcesMap["netbox_device"]
			ctx := context.Background()

			state, diags := resource.RefreshWithoutUpgrade(ctx, &terraform.InstanceState{
				ID: "42", Attributes: map[string]string{"name": "old-name"},
			}, api)
			if diags.HasError() {
				t.Fatalf("refresh failed: %v", diags)
			}
			checkDeviceNameState(t, state, "")
			config := terraform.NewResourceConfigRaw(deviceNameTestConfig(""))
			checkDeviceNameNoDiff(t, resource, state, config, api)

			data := resource.Data(nil)
			data.SetId("42")
			imported, err := resource.Importer.StateContext(ctx, data, api)
			if err != nil {
				t.Fatalf("import failed: %v", err)
			}
			if len(imported) != 1 || imported[0].Id() != "42" {
				t.Fatalf("expected imported device 42, got %#v", imported)
			}
			state, diags = resource.RefreshWithoutUpgrade(ctx, imported[0].State(), api)
			if diags.HasError() {
				t.Fatalf("import refresh failed: %v", diags)
			}
			checkDeviceNameState(t, state, "")
			checkDeviceNameNoDiff(t, resource, state, config, api)
		})
	}
}

func deviceNameTestConfig(name string) map[string]interface{} {
	config := map[string]interface{}{"device_type_id": 1, "role_id": 1, "site_id": 1}
	if name != "" {
		config["name"] = name
	}
	return config
}

func checkDeviceNameState(t *testing.T, state *terraform.InstanceState, name string) {
	t.Helper()
	if state == nil || state.ID != "42" {
		t.Fatalf("expected device 42 in state, got %#v", state)
	}
	if got := state.Attributes["name"]; got != name {
		t.Errorf("expected name %q in state, got %q", name, got)
	}
	for _, key := range []string{"device_type_id", "role_id", "site_id"} {
		if got := state.Attributes[key]; got != "1" {
			t.Errorf("expected %s = 1 in state, got %q", key, got)
		}
	}
}

func checkDeviceNameNoDiff(t *testing.T, resource *schema.Resource, state *terraform.InstanceState, config *terraform.ResourceConfig, api *providerState) {
	t.Helper()
	diff, err := resource.Diff(context.Background(), state, config, api)
	if err != nil {
		t.Fatalf("diff after read failed: %v", err)
	}
	if diff != nil && diff.Attributes["name"] != nil {
		t.Errorf("unexpected name diff after read: %#v", diff.Attributes["name"])
	}
}

type deviceNameRequest struct {
	method    string
	wantName  string
	replyName interface{}
}

func deviceNameTestAPI(t *testing.T, requests []deviceNameRequest) *providerState {
	t.Helper()
	var count atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		i := int(count.Add(1)) - 1
		if i >= len(requests) {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		want := requests[i]
		path := "/api/dcim/devices/42/"
		if want.method == http.MethodPost {
			path = "/api/dcim/devices/"
		}
		if r.Method != want.method || r.URL.Path != path {
			t.Errorf("request %d: got %s %s, want %s %s", i+1, r.Method, r.URL.Path, want.method, path)
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		if r.Method != http.MethodGet {
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode request: %v", err)
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			if name, ok := body["name"]; !ok || name != want.wantName {
				t.Errorf("request %d: expected explicit name %q, got %#v (present: %t)", i+1, want.wantName, name, ok)
			}
		}
		device := minimalDeviceJSON(42, "")
		device["name"] = want.replyName
		device["status"] = map[string]interface{}{"value": "active", "label": "Active"}
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			w.WriteHeader(http.StatusCreated)
		}
		if err := json.NewEncoder(w).Encode(device); err != nil {
			t.Errorf("encode response: %v", err)
		}
	}))
	t.Cleanup(func() {
		server.Close()
		if got := int(count.Load()); got != len(requests) {
			t.Errorf("expected %d API requests, got %d", len(requests), got)
		}
	})
	cfg := Config{APIToken: "test-token", ServerURL: server.URL}
	client, err := cfg.Client()
	if err != nil {
		t.Fatalf("create client: %v", err)
	}
	return &providerState{
		NetBoxAPI:   client,
		defaultTags: schema.NewSet(schema.HashString, nil),
		tagCache:    map[string]*models.NestedTag{},
	}
}
