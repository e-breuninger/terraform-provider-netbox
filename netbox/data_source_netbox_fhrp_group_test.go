package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxFhrpGroupSetUp(testName string) string {
	return fmt.Sprintf(`

resource "netbox_fhrp_group" "test" {
  protocol    = "other"
  group_id    = 1234
  auth_type   = "md5"
  auth_key    = "%[1]s"
  name        = "test"
  description = "test"
  comments    = "test"
}
resource "netbox_fhrp_group" "test2" {
  protocol    = "other"
  group_id    = 1235
  auth_type   = "md5"
  auth_key    = "%[1]s"
  name        = "test2"
  description = "test"
  comments    = "test"
}
resource "netbox_fhrp_group" "test3" {
  protocol    = "vrrp3"
  group_id    = 1234
  auth_type   = "md5"
  auth_key    = "%[1]s"
  name        = "test3"
  description = "test"
  comments    = "test"
}`, testName)
}

const testAccNetboxFhrpGroupNoResult = `
data "netbox_fhrp_group" "test" {
  group_id = 27
  protocol = "other"
}`

const testAccNetboxFhrpGroupTooManyResult = `
data "netbox_fhrp_group" "test" {
  group_id = 1234
}`

func testAccNetboxFhrpGroupByGroupIdProtocol() string {
	return `
data "netbox_fhrp_group" "test" {
  group_id = 1234
  protocol = "other"
}`
}

func TestAccNetboxFhrpGroupDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("asn_ds_basic")
	setUp := testAccNetboxFhrpGroupSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
			},
			{
				Config: setUp + testAccNetboxFhrpGroupByGroupIdProtocol(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_fhrp_group.test", "id", "netbox_fhrp_group.test", "id"),
				),
			},
			{
				Config:      setUp + testAccNetboxFhrpGroupNoResult,
				ExpectError: regexp.MustCompile("no group found matching filter"),
			},
			{
				Config:      setUp + testAccNetboxFhrpGroupTooManyResult,
				ExpectError: regexp.MustCompile("more than one group returned, specify a more narrow filter"),
			},
		},
	})
}
