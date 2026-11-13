package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// Exercise the real resource methods and wire requests without a live NetBox server.
// The fixture keeps both primary assignments so deletion of the wrong family is visible.
func TestPrimaryIPImportLifecycle(t *testing.T) {
	for _, kind := range []string{"device", "vm"} {
		for _, family := range []int64{4, 6} {
			t.Run(fmt.Sprintf("%s/IPv%d", kind, family), func(t *testing.T) {
				ctx := context.Background()
				var mu sync.Mutex
				primaries := map[string]any{"primary_ip4": float64(104), "primary_ip6": float64(106)}
				var patches []map[string]any
				endpoint := "/api/dcim/devices/12/"
				if kind == "vm" {
					endpoint = "/api/virtualization/virtual-machines/12/"
				}
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					mu.Lock()
					defer mu.Unlock()
					w.Header().Set("Content-Type", "application/json")
					switch {
					case req.Method == http.MethodGet && req.URL.Path == endpoint:
						body := map[string]any{"id": 12}
						for field, id := range primaries {
							if id == nil {
								body[field] = nil
							} else {
								body[field] = map[string]any{"id": id}
							}
						}
						_ = json.NewEncoder(w).Encode(body)
					case req.Method == http.MethodGet && req.URL.Path == fmt.Sprintf("/api/ipam/ip-addresses/%d/", 100+family):
						_ = json.NewEncoder(w).Encode(map[string]any{"id": 100 + family, "family": map[string]any{"value": family}})
					case req.Method == http.MethodPatch && req.URL.Path == endpoint:
						var body map[string]any
						if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
							t.Error(err)
						}
						patches = append(patches, body)
						for field, id := range body {
							primaries[field] = id
						}
						_, _ = w.Write([]byte(`{"id":12}`))
					default:
						t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
						w.WriteHeader(http.StatusBadRequest)
					}
				}))
				defer server.Close()
				r, state := primaryIPTestResource(t, server.URL)
				imported := resource.ImportStateResponse{State: state}
				r.ImportState(ctx, resource.ImportStateRequest{ID: fmt.Sprintf("%s/12/%d", kind, family)}, &imported)
				if imported.Diagnostics.HasError() {
					t.Fatal(imported.Diagnostics)
				}
				read := resource.ReadResponse{State: imported.State}
				r.Read(ctx, resource.ReadRequest{State: imported.State}, &read)
				if read.Diagnostics.HasError() {
					t.Fatal(read.Diagnostics)
				}
				var model primaryIPResourceModel
				if diags := read.State.Get(ctx, &model); diags.HasError() {
					t.Fatal(diags)
				}
				if model.ID.ValueString() != kind+"/12" || model.Family.ValueInt64() != family || model.IPAddressID.ValueInt64() != 100+family {
					t.Fatalf("wrong imported binding: %+v", model)
				}
				mu.Lock()
				if len(patches) != 0 {
					t.Error("import/refresh must not write to NetBox")
				}
				mu.Unlock()

				// An unchanged configuration matches the imported state. If Update is
				// called, it must only reassert this family, never clear its sibling.
				plan := tfsdk.Plan{Schema: state.Schema}
				if diags := plan.Set(ctx, &model); diags.HasError() {
					t.Fatal(diags)
				}
				updated := resource.UpdateResponse{State: state}
				r.Update(ctx, resource.UpdateRequest{Plan: plan, State: read.State}, &updated)
				if updated.Diagnostics.HasError() {
					t.Fatal(updated.Diagnostics)
				}
				if !updated.State.Raw.Equal(read.State.Raw) {
					t.Error("unchanged update altered imported state")
				}
				deleted := resource.DeleteResponse{}
				r.Delete(ctx, resource.DeleteRequest{State: updated.State}, &deleted)
				if deleted.Diagnostics.HasError() {
					t.Fatal(deleted.Diagnostics)
				}
				mu.Lock()
				want := []map[string]any{{familyField(family): float64(100 + family)}, {familyField(family): nil}}
				if !reflect.DeepEqual(patches, want) {
					t.Errorf("PATCH bodies = %v, want %v", patches, want)
				}
				sibling := int64(10) - family
				if primaries[familyField(sibling)] != float64(100+sibling) {
					t.Error("sibling primary was changed")
				}
				mu.Unlock()
				// Refresh must not fall back to the sibling after this family disappears.
				refreshed := resource.ReadResponse{State: updated.State}
				r.Read(ctx, resource.ReadRequest{State: updated.State}, &refreshed)
				if refreshed.Diagnostics.HasError() || !refreshed.State.Raw.IsNull() {
					t.Fatalf("deleted family was not removed from state: %v", refreshed.Diagnostics)
				}
			})
		}
	}
}

func TestPrimaryIPImportRead(t *testing.T) {
	for _, kind := range []string{"device", "vm"} {
		for _, tc := range []struct {
			name, suffix, body string
			family, ipID       int64
			ambiguous, absent  bool
		}{
			{name: "legacy IPv4", body: `{"primary_ip4":{"id":104}}`, family: 4, ipID: 104},
			{name: "legacy IPv6", body: `{"primary_ip6":{"id":106}}`, family: 6, ipID: 106},
			{name: "legacy dual stack", body: `{"primary_ip4":{"id":104},"primary_ip6":{"id":106}}`, ambiguous: true},
			{name: "legacy empty", body: `{}`, absent: true},
			{name: "explicit IPv4 absent", suffix: "/4", body: `{"primary_ip6":{"id":106}}`, absent: true},
			{name: "explicit IPv6 absent", suffix: "/6", body: `{"primary_ip4":{"id":104}}`, absent: true},
		} {
			t.Run(kind+"/"+tc.name, func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					if req.Method != http.MethodGet {
						t.Errorf("import must be read-only, got %s", req.Method)
					}
					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(tc.body))
				}))
				defer server.Close()
				r, state := primaryIPTestResource(t, server.URL)
				ctx := context.Background()
				imported := resource.ImportStateResponse{State: state}
				r.ImportState(ctx, resource.ImportStateRequest{ID: kind + "/12" + tc.suffix}, &imported)
				if imported.Diagnostics.HasError() {
					t.Fatal(imported.Diagnostics)
				}
				read := resource.ReadResponse{State: imported.State}
				r.Read(ctx, resource.ReadRequest{State: imported.State}, &read)
				if tc.ambiguous {
					if !read.Diagnostics.HasError() || !strings.Contains(read.Diagnostics[0].Detail(), kind+"/12/6") {
						t.Fatalf("expected actionable ambiguity error, got %v", read.Diagnostics)
					}
					if !read.State.Raw.Equal(imported.State.Raw) {
						t.Error("ambiguous import adopted a primary address")
					}
					return
				}
				if read.Diagnostics.HasError() {
					t.Fatal(read.Diagnostics)
				}
				if tc.absent {
					if !read.State.Raw.IsNull() {
						t.Error("absent primary must not be adopted")
					}
					return
				}
				var model primaryIPResourceModel
				if diags := read.State.Get(ctx, &model); diags.HasError() {
					t.Fatal(diags)
				}
				if model.Family.ValueInt64() != tc.family || model.IPAddressID.ValueInt64() != tc.ipID {
					t.Fatalf("wrong legacy binding: %+v", model)
				}
			})
		}
	}
}

func TestPrimaryIPImportInvalidID(t *testing.T) {
	for _, id := range []string{"", "12", "device", "host/12/4", "device/no/6", "device/0/4", "vm/-1/6", "vm/9223372036854775808/4", "device/12/", "device/12/5", "vm/12/04", "vm/12/ipv6", "vm/12/6/extra"} {
		t.Run(id, func(t *testing.T) {
			r, state := primaryIPTestResource(t, "http://127.0.0.1")
			resp := resource.ImportStateResponse{State: state}
			r.ImportState(context.Background(), resource.ImportStateRequest{ID: id}, &resp)
			if !resp.Diagnostics.HasError() {
				t.Fatal("expected invalid import ID diagnostic")
			}
		})
	}
}

func primaryIPTestResource(t *testing.T, serverURL string) (*primaryIPResource, tfsdk.State) {
	t.Helper()
	api, err := netboxapi.New(serverURL, "dummy", netboxapi.Options{})
	if err != nil {
		t.Fatal(err)
	}
	r := &primaryIPResource{client: &netboxapi.Client{NetBoxAPI: api}}
	var schema resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &schema)
	return r, tfsdk.State{Schema: schema.Schema}
}
