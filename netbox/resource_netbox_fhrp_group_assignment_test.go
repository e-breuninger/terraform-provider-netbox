package netbox

import (
	"fmt"
	"log"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxFhrpGroupAssignment_basic(t *testing.T) {
	testSlug := "fhrp_group_assignment_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_fhrp_group" "test" {
  protocol    = "other"
  group_id    = 1234
  auth_type   = "md5"
  auth_key    = "test"
  name        = "test"
  description = "test"
  comments    = "test"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "test"
}

resource "netbox_device_type" "test" {
  model           = "test"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_interface" "test" {
  name      = "testinterface"
  device_id = netbox_device.test.id
  type      = "1000base-t"
}

resource "netbox_fhrp_group_assignment" "test" {
  group_id = netbox_fhrp_group.test.id
  interface_id = netbox_device_interface.test.id
  interface_type = "dcim.interface"
  priority = 150
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_fhrp_group_assignment.test", "group_id", "netbox_fhrp_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_fhrp_group_assignment.test", "interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttr("netbox_fhrp_group_assignment.test", "interface_type", "dcim.interface"),
					resource.TestCheckResourceAttr("netbox_fhrp_group_assignment.test", "priority", "150"),
				),
			},
			{
				ResourceName:      "netbox_fhrp_group_assignment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_fhrp_group_assignment", &resource.Sweeper{
		Name:         "netbox_fhrp_group_assigment",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamFhrpGroupAssignmentsListParams()
			res, err := api.Ipam.IpamFhrpGroupAssignmentsList(params, nil)
			if err != nil {
				return err
			}
			for _, asn := range res.GetPayload().Results {
				deleteParams := ipam.NewIpamFhrpGroupAssignmentsDeleteParams().WithID(asn.ID)
				_, err := api.Ipam.IpamFhrpGroupAssignmentsDelete(deleteParams, nil)
				if err != nil {
					return err
				}
				log.Print("[DEBUG] Deleted an asn")
			}
			return nil
		},
	})
}
