//go:build acctest

// Object and multiobject custom fields: configured as related object ids, and
// carried through updates when NetBox (not Terraform) set them.
package provider_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNetboxObjectCustomFields(t *testing.T) {
	testName := testAccGetTestName("objcf")
	field := strings.ReplaceAll(testName, "-", "_")
	deps := fmt.Sprintf(`
resource "netbox_custom_field" "object" {
  name                = "%[2]s_o"
  type                = "object"
  related_object_type = "dcim.site"
  object_types        = ["dcim.device"]
}
resource "netbox_custom_field" "multiobject" {
  name                = "%[2]s_m"
  type                = "multiobject"
  related_object_type = "dcim.site"
  object_types        = ["dcim.device"]
}
resource "netbox_site" "a" {
  name = "%[1]s-a"
}
resource "netbox_site" "b" {
  name = "%[1]s-b"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
`, testName, field)
	device := func(description, customFields string) string {
		return fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.a.id
  description    = "%[2]s"
  %[3]s
  depends_on     = [netbox_custom_field.object, netbox_custom_field.multiobject]
}`, testName, description, customFields)
	}
	configured := fmt.Sprintf(`custom_fields = {
    %[1]s_o = tostring(netbox_site.a.id)
    %[1]s_m = jsonencode([netbox_site.b.id, netbox_site.a.id])
  }`, field)

	var deviceID, siteID string
	remember := func(s *terraform.State) error {
		deviceID = s.RootModule().Resources["netbox_device.test"].Primary.ID
		siteID = s.RootModule().Resources["netbox_site.b"].Primary.ID
		return nil
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{ // configured as ids, NetBox answers nested objects: no diff
				Config: deps + device("one", configured),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device.test", "custom_fields."+field+"_o", "netbox_site.a", "id"),
					remember,
				),
			},
			{ // an update re-sends the object fields
				Config: deps + device("two", configured),
				Check:  resource.TestCheckResourceAttr("netbox_device.test", "description", "two"),
			},
			{ // custom_fields not configured, an object field set in NetBox: updates still work
				PreConfig: func() {
					testAccPatchJSON(t, "/api/dcim/devices/"+deviceID+"/", map[string]any{
						"custom_fields": map[string]any{field + "_o": siteID},
					})
				},
				Config: deps + device("three", ""),
				Check:  resource.TestCheckResourceAttr("netbox_device.test", "description", "three"),
			},
		},
	})
}

// testAccPatchJSON changes an object behind Terraform's back.
func testAccPatchJSON(t *testing.T, path string, body any) {
	t.Helper()
	raw, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPatch, strings.TrimRight(os.Getenv("NETBOX_SERVER_URL"), "/")+path, bytes.NewReader(raw))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("NETBOX_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("PATCH %s: %v", path, err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		t.Fatalf("PATCH %s: HTTP %d", path, res.StatusCode)
	}
}
