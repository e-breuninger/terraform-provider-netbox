package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVlanDataSource_basic(t *testing.T) {
	testVid := 4092
	testName := testAccGetTestName("vlan")
	setUp := testAccNetboxVlanSetUp(testVid, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
			},
			{
				Config:      setUp + testAccNetboxVlanDataNoResult,
				ExpectError: regexp.MustCompile("expected one device type, but got 0"),
			},
			{
				Config: setUp + testAccNetboxVlanDataByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "id", "netbox_vlan.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanDataByVid(testVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_vlan.test", "status", "active"),
					resource.TestCheckResourceAttr("data.netbox_vlan.test", "description", "Test"),
				),
			},
		},
	})
}

func testAccNetboxVlanSetUp(testVid int, testName string) string {
	return fmt.Sprintf(`
resource "netbox_vlan" "test" {
	vid = %[1]d
	name = "%[2]s"
	status = "active"
	description = "Test"
	tags = []
}
`, testVid, testName)
}

const testAccNetboxVlanDataNoResult = `
data "netbox_vlan" "no_result" {
	name = "_no_result_"
}`

func testAccNetboxVlanDataByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_vlan" "test" {
	name = "%[1]s"
}`, testName)
}

func testAccNetboxVlanDataByVid(testVid int) string {
	return fmt.Sprintf(`
data "netbox_vlan" "test" {
	vid = "%[1]d"
}`, testVid)
}
