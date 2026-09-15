package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxL2VPNTerminationDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
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
}

resource "netbox_device_interface" "test" {
  name      = "%[1]s"
  device_id = netbox_device.test.id
  type      = "1000base-t"
}

resource "netbox_l2vpn" "test" {
  name = "%[1]s"
  slug = "%[1]s"
  type = "vxlan"
}`, testName)
}

func TestAccNetboxL2VPNTermination_basic(t *testing.T) {
	testSlug := "l2vpn_term_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxL2VPNTerminationDependencies(testName) + `
resource "netbox_l2vpn_termination" "test" {
  l2vpn_id              = netbox_l2vpn.test.id
  assigned_object_type  = "dcim.interface"
  assigned_object_id    = netbox_device_interface.test.id
  tags                  = [netbox_tag.test.name]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "l2vpn_id", "netbox_l2vpn.test", "id"),
					resource.TestCheckResourceAttr("netbox_l2vpn_termination.test", "assigned_object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "assigned_object_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttr("netbox_l2vpn_termination.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_l2vpn_termination.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_l2vpn_termination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_l2vpn_termination", &resource.Sweeper{
		Name: "netbox_l2vpn_termination",
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamL2vpnTerminationsListParams()
			res, err := api.Ipam.IpamL2vpnTerminationsList(params, nil)
			if err != nil {
				return err
			}
			for _, termination := range res.GetPayload().Results {
				if termination.L2vpn != nil && termination.L2vpn.Name != nil && strings.HasPrefix(*termination.L2vpn.Name, testPrefix) {
					deleteParams := ipam.NewIpamL2vpnTerminationsDeleteParams().WithID(termination.ID)
					_, err := api.Ipam.IpamL2vpnTerminationsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an l2vpn termination")
				}
			}
			return nil
		},
	})
}
