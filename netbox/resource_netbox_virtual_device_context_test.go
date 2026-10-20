package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVirtualDeviceContextDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name            = "%[1]s"
  device_type_id  = netbox_device_type.test.id
  role_id         = netbox_device_role.test.id
  site_id         = netbox_site.test.id
  status          = "active"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxVirtualDeviceContext_basic(t *testing.T) {
	testSlug := "vdc_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualDeviceContextDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_device_context" "test" {
  name        = "%[1]s"
  device_id   = netbox_device.test.id
  identifier  = 1
  tenant_id   = netbox_tenant.test.id
  status      = "active"
  description = "This is my virtual device context"
  comments    = "Virtual device context comments"
  tags        = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_device_context.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "identifier", "1"),
					resource.TestCheckResourceAttrPair("netbox_virtual_device_context.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "description", "This is my virtual device context"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "comments", "Virtual device context comments"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_virtual_device_context.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_virtual_device_context", &resource.Sweeper{
		Name:         "netbox_virtual_device_context",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimVirtualDeviceContextsListParams()
			res, err := api.Dcim.DcimVirtualDeviceContextsList(params, nil)
			if err != nil {
				return err
			}
			for _, vdc := range res.GetPayload().Results {
				if strings.HasPrefix(*vdc.Name, testPrefix) {
					deleteParams := dcim.NewDcimVirtualDeviceContextsDeleteParams().WithID(vdc.ID)
					_, err := api.Dcim.DcimVirtualDeviceContextsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a virtual_device_context")
				}
			}
			return nil
		},
	})
}
