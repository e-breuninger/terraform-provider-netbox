package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxAsnsSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_asn" "test_1" {
  asn    = "123"
  rir_id = netbox_rir.test.id
  tags   = [netbox_tag.test.slug]
}

resource "netbox_asn" "test_2" {
	asn    = "1234"
	rir_id = netbox_rir.test.id
	tags   = [netbox_tag.test.slug]
  }`, testName)
}

func testAccNetboxAsnsByAsn() string {
	return fmt.Sprintf(`
data "netbox_asns" "test" {
  filter {
	name = "asn"
	value = "123"
  }
}`)
}

func testAccNetboxAsnsByRange(testName string) string {
	return fmt.Sprintf(`
data "netbox_asns" "test" {
  filter {
	name = "asn__gte"
	value = "100"
  }

  filter {
	name = "asn__lte"
	value = "2000"
  }
}`)
}

func TestAccNetboxAsnsDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("asns_ds_basic")
	setUp := testAccNetboxAsnsSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_asn.test_1", "asn", "123"),
				),
			},
			{
				Config: setUp + testAccNetboxAsnsByAsn(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_asns.test", "asns.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_asns.test", "asns.0.id", "netbox_asn.test_1", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxAsnsByRange(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_asns.test", "asns.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_asns.test", "asns.0.id", "netbox_asn.test_1", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_asns.test", "asns.1.id", "netbox_asn.test_2", "id"),
				),
			},
		},
	})
}
