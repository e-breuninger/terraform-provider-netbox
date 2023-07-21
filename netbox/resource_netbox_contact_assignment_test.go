package netbox

import (
	"fmt"
	"log"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxContactAssignment_basic(t *testing.T) {
	testSlug := "contactassign"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}
resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}
resource "netbox_device" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
}
resource "netbox_contact" "test" {
  name = "%[1]s"
}
resource "netbox_contact_role" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_contact_assignment" "test" {
  content_type = "dcim.device"
  object_id = netbox_device.test.id
  contact_id = netbox_contact.test.id
  role_id = netbox_contact_role.test.id
  priority = "primary"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact_assignment.test", "content_type", "dcim.device"),
					resource.TestCheckResourceAttrPair("netbox_contact_assignment.test", "object_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_contact_assignment.test", "contact_id", "netbox_contact.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_contact_assignment.test", "role_id", "netbox_contact_role.test", "id"),
					resource.TestCheckResourceAttr("netbox_contact_assignment.test", "priority", "primary"),
				),
			},
			{
				ResourceName:      "netbox_contact_assignment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_contact_assignment", &resource.Sweeper{
		Name:         "netbox_contact_assignment",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := tenancy.NewTenancyContactAssignmentsListParams()
			res, err := api.Tenancy.TenancyContactAssignmentsList(params, nil)
			if err != nil {
				return err
			}
			for _, contactassignment := range res.GetPayload().Results {
				deleteParams := tenancy.NewTenancyContactAssignmentsDeleteParams().WithID(contactassignment.ID)
				_, err := api.Tenancy.TenancyContactAssignmentsDelete(deleteParams, nil)
				if err != nil {
					return err
				}
				log.Print("[DEBUG] Deleted a contact assignment")
			}
			return nil
		},
	})
}
