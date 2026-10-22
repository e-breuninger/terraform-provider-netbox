//go:build acctest

// Contact assignments: a contact attached to an object in a role. No sweeper: an assignment goes
// with the site or contact it hangs off, which the catalog sweepers already remove.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxContactAssignment_basic(t *testing.T) {
	testName := testAccGetTestName("contactassign")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_contact" "test" {
  name = "%[1]s"
}
resource "netbox_contact_role" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_contact_assignment" "test" {
  object_type = "dcim.site"
  object_id   = netbox_site.test.id
  contact_id  = netbox_contact.test.id
  role_id     = netbox_contact_role.test.id
  priority    = "primary"
}
data "netbox_contact_assignment" "test" {
  id = netbox_contact_assignment.test.id
}
data "netbox_contact_assignments" "test" {
  filters = [
    { name = "contact_id", value = netbox_contact.test.id },
  ]
  depends_on = [netbox_contact_assignment.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact_assignment.test", "object_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_contact_assignment.test", "object_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_contact_assignment.test", "contact_id", "netbox_contact.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_contact_assignment.test", "role_id", "netbox_contact_role.test", "id"),
					resource.TestCheckResourceAttr("netbox_contact_assignment.test", "priority", "primary"),
					resource.TestCheckResourceAttrPair("data.netbox_contact_assignment.test", "id", "netbox_contact_assignment.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_contact_assignments.test", "contact_assignments.#", "1"),
				),
			},
			{
				// Shrink: the priority clears. The role does not — NetBox refuses a null there — so
				// role_id is optional+computed and an unset one keeps NetBox's value. That also means
				// the role is still referenced with no config edge left to it, so destroy needs the
				// depends_on to take the assignment down before the role.
				Config: deps + `
resource "netbox_contact_assignment" "test" {
  object_type = "dcim.site"
  object_id   = netbox_site.test.id
  contact_id  = netbox_contact.test.id
  depends_on  = [netbox_contact_role.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_contact_assignment.test", "role_id", "netbox_contact_role.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_contact_assignment.test", "priority"),
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
