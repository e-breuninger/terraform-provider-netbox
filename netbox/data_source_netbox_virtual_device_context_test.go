package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualDeviceContextDataSource_basic(t *testing.T) {
	testSlug := "vdc_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxVirtualDeviceContextDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_device_context" "test" {
  name        = "%[1]s"
  device_id   = netbox_device.test.id
  identifier  = 1
  tenant_id   = netbox_tenant.test.id
  status      = "active"
  description = "This is my virtual device context"
  comments    = "Virtual device context comments"
  tags        = [netbox_tag.test.name]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_virtual_device_context" "test" {
  depends_on = []

  name = netbox_virtual_device_context.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_virtual_device_context.test", "id", "netbox_virtual_device_context.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_context.test", "name", testName),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_device_context.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_context.test", "identifier", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_device_context.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_context.test", "status", "active"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_context.test", "description", "This is my virtual device context"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_context.test", "comments", "Virtual device context comments"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_context.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_context.test", "tags.0", testName),
				),
			},
		},
	})
}
