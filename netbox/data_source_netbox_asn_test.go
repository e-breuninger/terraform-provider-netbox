package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxAsnSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_asn" "test" {
  asn    = "456"
  rir_id = netbox_rir.test.id
  tags   = [netbox_tag.test.slug]
}`, testName)
}

const testAccNetboxAsnNoResult = `
data "netbox_asn" "test" {
  asn = "1337"
}`

func testAccNetboxAsnByAsn() string {
	return `
data "netbox_asn" "test" {
  asn = "456"
}`
}

func testAccNetboxAsnByTag(testName string) string {
	return fmt.Sprintf(`
data "netbox_asn" "test" {
  tag = "%[1]s"
}`, testName)
}

func TestAccNetboxAsnDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("asn_ds_basic")
	setUp := testAccNetboxAsnSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_asn.test", "asn", "456"),
				),
			},
			{
				Config:      setUp + testAccNetboxAsnNoResult,
				ExpectError: regexp.MustCompile("expected one ASN, but got 0"),
			},
			{
				Config: setUp + testAccNetboxAsnByAsn(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_asn.test", "id", "netbox_asn.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxAsnByTag(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_asn.test", "id", "netbox_asn.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_asn.test", "asn", "456"),
				),
			},
		},
	})
}
