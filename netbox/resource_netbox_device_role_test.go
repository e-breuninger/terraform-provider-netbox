package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceRole_basic(t *testing.T) {

	testSlug := "dvcrl_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_device_role" "test" {
  name = "%s"
  slug = "%s"
  color_hex = "111111"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_role.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_role.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_role.test", "color_hex", "111111"),
				),
			},
			{
				ResourceName:      "netbox_device_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxDeviceRole_defaultSlug(t *testing.T) {

	testSlug := "device_role_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_device_role" "test" {
  name = "%s"
  color_hex = "111111"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_role.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_role.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_device_role", &resource.Sweeper{
		Name:         "netbox_device_role",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimDeviceRolesListParams()
			res, err := api.Dcim.DcimDeviceRolesList(params, nil)
			if err != nil {
				return err
			}
			for _, device_role := range res.GetPayload().Results {
				if strings.HasPrefix(*device_role.Name, testPrefix) {
					deleteParams := dcim.NewDcimDeviceRolesDeleteParams().WithID(device_role.ID)
					_, err := api.Dcim.DcimDeviceRolesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device_role")
				}
			}
			return nil
		},
	})
}
