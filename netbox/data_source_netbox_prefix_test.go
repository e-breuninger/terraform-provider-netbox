package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func resources() string {
	return `
`
}

func TestAccNetboxPrefixDataSource_basic(t *testing.T) {

	testPrefix := "10.0.0.0/24"
	testSlug := "prefix_ds_basic"
	testVlanVid := 4090
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name = "%[1]s_vrf"
}

resource "netbox_vlan" "test" {
  name = "%[1]s_vlan_test_id"
  vid  = %[3]d
}

resource "netbox_site" "test" {
  name = "%[1]s_site"
}

resource "netbox_prefix" "test" {
  prefix = "%[2]s"
  status = "active"
  vrf_id = netbox_vrf.test.id
  vlan_id = netbox_vlan.test.id
  site_id = netbox_site.test.id
  description = "%[1]s_description_test_id"
}

data "netbox_prefix" "by_description" {
  description = netbox_prefix.test.description
}

data "netbox_prefix" "by_cidr" {
  depends_on = [netbox_prefix.test]
  cidr = "%[2]s"
}

data "netbox_prefix" "by_vrf_id" {
  depends_on = [netbox_prefix.test]
  vrf_id = netbox_vrf.test.id
}

data "netbox_prefix" "by_vlan_id" {
  depends_on = [netbox_prefix.test]
  vlan_id = netbox_vlan.test.id
}

data "netbox_prefix" "by_vlan_vid" {
  depends_on = [netbox_prefix.test]
  vlan_vid = %[3]d
}

data "netbox_prefix" "by_prefix" {
  depends_on = [netbox_prefix.test]
  prefix = "%[2]s"
}

data "netbox_prefix" "by_site_id" {
  depends_on = [netbox_prefix.test]
  site_id = netbox_site.test.id
}
`, testName, testPrefix, testVlanVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_prefix", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_description", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_cidr", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vrf_id", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_id", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_vid", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_site_id", "id", "netbox_prefix.test", "id"),
				),
			},
		},
	})
}
