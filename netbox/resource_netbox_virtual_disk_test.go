package netbox

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetboxVirtualDisk_basic(t *testing.T) {
	testSlug := "virtual_disk"
	testName := testAccGetTestName(testSlug)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualDiskDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "tag_a" {
	name = "[%[1]s_a]"
	color_hex = "123456"
}
resource "netbox_site" "test" {
	name = "%[1]s"
	status = "active"
}
resource "netbox_virtual_machine" "test" {
	name = "%[1]s"
	site_id = netbox_site.test.id
}
resource "netbox_virtual_disk" "test" {
	name = "%[1]s"
	description = "description"
	size_gb = 30
	virtual_machine_id = netbox_virtual_machine.test.id
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
resource "netbox_site" "test" {
	name = "%[1]s"
	status = "active"
}
resource "netbox_virtual_machine" "test" {
	name = "%[1]s"
	site_id = netbox_site.test.id
}
resource "netbox_virtual_disk" "test" {
	name = "%[1]s_updated"
	description = "description updated"
	size_gb = 60
	virtual_machine_id = netbox_virtual_machine.test.id
	tags = [netbox_tag.tag_a.name]
}
				`, testName),
			},
			{
				ResourceName:      "netbox_virtual_disk.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckVirtualDiskDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each virtual machine
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_virtual_disk" {
			continue
		}

		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := virtualization.NewVirtualizationVirtualDisksReadParams().WithID(stateID)
		_, err := conn.Virtualization.VirtualizationVirtualDisksRead(params, nil)

		if err == nil {
			return fmt.Errorf("virtual disk (%s) still exists", rs.Primary.ID)
		}

		if errresp, ok := err.(*virtualization.VirtualizationVirtualDisksReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				return nil
			}
		}
		return err
	}
	return nil
}
