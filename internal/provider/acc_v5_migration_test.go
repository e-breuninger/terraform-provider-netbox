//go:build acctest

// Upgrading a workspace from provider v5: state written by v5.8.0 must be read
// by v6 without manual state surgery. Covers the resources whose v5 state v6
// could not read (netbox_device and netbox_custom_field with an empty JSON
// attribute, netbox_primary_ip with a bare VM id) and the netbox_interface ->
// netbox_virtual_machine_interface rename through a moved block.
package provider_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// v5 only sends v1 tokens ("Token <key>"); create one for the v5 step with the
// v2 token the suite runs with, and delete it afterwards.
func testAccV1Token(t *testing.T) string {
	t.Helper()
	server := strings.TrimRight(os.Getenv("NETBOX_SERVER_URL"), "/")
	auth := "Bearer " + os.Getenv("NETBOX_API_TOKEN")
	do := func(method, url string, body any) map[string]any {
		var reader *bytes.Reader
		if body != nil {
			raw, _ := json.Marshal(body)
			reader = bytes.NewReader(raw)
		} else {
			reader = bytes.NewReader(nil)
		}
		req, _ := http.NewRequest(method, server+url, reader)
		req.Header.Set("Authorization", auth)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("%s %s: %v", method, url, err)
		}
		defer res.Body.Close()
		out := map[string]any{}
		_ = json.NewDecoder(res.Body).Decode(&out)
		if res.StatusCode >= 300 {
			t.Fatalf("%s %s: HTTP %d %v", method, url, res.StatusCode, out)
		}
		return out
	}
	users := do(http.MethodGet, "/api/users/users/?username=admin", nil)
	results, _ := users["results"].([]any)
	if len(results) == 0 {
		t.Skip("no admin user to own a v1 token")
	}
	userID := results[0].(map[string]any)["id"]
	token := do(http.MethodPost, "/api/users/tokens/", map[string]any{"user": userID, "version": 1, "description": "v5 migration acceptance test"})
	t.Cleanup(func() { do(http.MethodDelete, fmt.Sprintf("/api/users/tokens/%v/", token["id"]), nil) })
	return token["token"].(string)
}

func TestAccNetboxMigrationFromV5(t *testing.T) {
	// The provider under test must answer for the address the v5 state is
	// recorded under, registry.terraform.io/e-breuninger/netbox.
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "e-breuninger")
	testAccPreCheck(t)
	v1Token := testAccV1Token(t)
	testName := testAccGetTestName("v5migr")
	address := fmt.Sprintf("198.18.%d.%d/24", acctest.RandIntRange(0, 250), acctest.RandIntRange(1, 250))
	common := `
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.test.id
}
resource "netbox_primary_ip" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  ip_address_id      = netbox_ip_address.test.id
}
`
	var deviceID, vmIfaceID, fieldID string
	remember := func(target *string, address string) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			rs, ok := s.RootModule().Resources[address]
			if !ok {
				return fmt.Errorf("%s not in state", address)
			}
			*target = rs.Primary.ID
			return nil
		}
	}
	same := func(want *string, address string) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			rs, ok := s.RootModule().Resources[address]
			if !ok {
				return fmt.Errorf("%s not in state", address)
			}
			if rs.Primary.ID != *want {
				return fmt.Errorf("%s: id %q, want %q (object replaced instead of migrated)", address, rs.Primary.ID, *want)
			}
			return nil
		}
	}
	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"netbox": {Source: "registry.terraform.io/e-breuninger/netbox", VersionConstraint: "5.8.0"},
				},
				Config: fmt.Sprintf(`
provider "netbox" {
  server_url = "%[2]s"
  api_token  = "%[3]s"
}
resource "netbox_custom_field" "test" {
  name          = "%[4]s"
  type          = "text"
  content_types = ["dcim.device"]
}
resource "netbox_interface" "test" {
  name               = "eth0"
  virtual_machine_id = netbox_virtual_machine.test.id
}
resource "netbox_ip_address" "test" {
  ip_address   = "%[5]s"
  status       = "active"
  object_type  = "virtualization.vminterface"
  interface_id = netbox_interface.test.id
}
`+common, testName, os.Getenv("NETBOX_SERVER_URL"), v1Token, strings.ReplaceAll(testName, "-", "_"), address),
				Check: resource.ComposeTestCheckFunc(
					remember(&deviceID, "netbox_device.test"),
					remember(&vmIfaceID, "netbox_interface.test"),
					remember(&fieldID, "netbox_custom_field.test"),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name         = "%[2]s"
  type         = "text"
  object_types = ["dcim.device"]
}
moved {
  from = netbox_interface.test
  to   = netbox_virtual_machine_interface.test
}
resource "netbox_virtual_machine_interface" "test" {
  name               = "eth0"
  virtual_machine_id = netbox_virtual_machine.test.id
}
resource "netbox_ip_address" "test" {
  ip_address                   = "%[3]s"
  status                       = "active"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}
`+common, testName, strings.ReplaceAll(testName, "-", "_"), address),
				Check: resource.ComposeTestCheckFunc(
					same(&deviceID, "netbox_device.test"),
					same(&vmIfaceID, "netbox_virtual_machine_interface.test"),
					same(&fieldID, "netbox_custom_field.test"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "local_context_data"),
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttrWith("netbox_primary_ip.test", "id", func(v string) error {
						if !strings.HasPrefix(v, "vm/") {
							return fmt.Errorf("id %q, want vm/<id>", v)
						}
						return nil
					}),
				),
			},
		},
	})
}
