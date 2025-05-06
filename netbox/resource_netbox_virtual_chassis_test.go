package netbox

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNetboxVirtualChassis_basic(t *testing.T) {
	testSlug := "virtual_chassis"
	testName := testAccGetTestName(testSlug)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualChassisDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "tag_a" {
	name = "[%[1]s_a]"
	color_hex = "123456"
}
resource "netbox_virtual_chassis" "test" {
	name = "%[1]s"
	domain = "domain"
	description = "description"
	comments = "comment"
	tags = [netbox_tag.tag_a.name]
}
				`, testName),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "tag_a" {
	name = "[%[1]s_a]"
	color_hex = "123456"
}
resource "netbox_virtual_chassis" "test" {
	name = "%[1]s_updated"
	domain = "domain_updated"
	description = "description updated"
	comments = "comment updated"
	tags = [netbox_tag.tag_a.name]
}
				`, testName),
			},
			{
				ResourceName:      "netbox_virtual_chassis.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckVirtualChassisDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each virtual machine
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_virtual_chassis" {
			continue
		}

		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimVirtualChassisReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimVirtualChassisRead(params, nil)

		if err == nil {
			return fmt.Errorf("virtual chassis (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimVirtualChassisReadDefault); ok {
				errorcode := errresp.Code()
				if errorcode == 404 {
					return nil
				}
			}
			return err
		}
	}
	return nil
}
