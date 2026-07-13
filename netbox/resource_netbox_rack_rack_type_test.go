package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccNetboxRack_rackType verifies that a rack can reference a rack type and
// that the physical dimensions are inherited from the type (NetBox owns them
// and ignores/overrides any explicitly supplied values), read back without
// producing a perpetual diff.
func TestAccNetboxRack_rackType(t *testing.T) {
	testSlug := "rack_rack_type"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRackDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_rack_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
  form_factor     = "4-post-frame"
  width           = 19
  u_height        = 47
  starting_unit   = 1
}

resource "netbox_rack" "test" {
  name         = "%[1]s"
  site_id      = netbox_site.test.id
  status       = "active"
  rack_type_id = netbox_rack_type.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_rack.test", "rack_type_id", "netbox_rack_type.test", "id"),
					// dimensions are inherited from the rack type
					resource.TestCheckResourceAttr("netbox_rack.test", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack.test", "u_height", "47"),
					resource.TestCheckResourceAttr("netbox_rack.test", "form_factor", "4-post-frame"),
				),
			},
			{
				ResourceName:      "netbox_rack.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
