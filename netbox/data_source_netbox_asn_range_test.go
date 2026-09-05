package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxAsnRangeDataSource_basic(t *testing.T) {
	testSlug := "asn_range_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxAsnRangeDependencies(testName) + fmt.Sprintf(`
resource "netbox_asn_range" "test" {
  name        = "%[1]s"
  slug        = "%[1]s"
  rir_id      = netbox_rir.test.id
  start       = 65000
  end         = 65100
  tenant_id   = netbox_tenant.test.id
  description = "asn range for acctest"
  comments    = "some comments"
  tags        = [netbox_tag.test.name]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_asn_range" "test" {
  depends_on = []

  name = netbox_asn_range.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_asn_range.test", "id", "netbox_asn_range.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "slug", testName),
					resource.TestCheckResourceAttrPair("data.netbox_asn_range.test", "rir_id", "netbox_rir.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "start", "65000"),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "end", "65100"),
					resource.TestCheckResourceAttrPair("data.netbox_asn_range.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "description", "asn range for acctest"),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "comments", "some comments"),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_asn_range.test", "tags.0", testName),
				),
			},
		},
	})
}
