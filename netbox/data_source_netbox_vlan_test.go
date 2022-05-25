package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVlanDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("ds_dv_tp")
	setUp := testAccNetboxVlanSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
			},
			{
				Config:      testAccNetboxVlanDataNoResult,
				ExpectError: regexp.MustCompile("no result"),
			},
			{
				Config: setUp + testAccNetboxVlanDataByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.netbox_vlan.by_name", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan.by_name", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_vlan.by_name", "vid", "1234"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanDataByVid,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.netbox_vlan.by_vid", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan.by_vid", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_vlan.by_vid", "vid", "1234"),
				),
			},
		},
	})
}

func testAccNetboxVlanSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_vlan" "test" {
	vid = 1234
	name = "%[1]s"
	status = "active"
	tags = []
}
`, testName)
}

const testAccNetboxVlanDataNoResult = `
data "netbox_vlan" "no_result" {
	name = "_no_result_"
}`

func testAccNetboxVlanDataByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_vlan" "by_name" {
	name = "%[1]s"
}`, testName)
}

const testAccNetboxVlanDataByVid = `
data "netbox_vlan" "by_vid" {
	vid = 1234
}`
