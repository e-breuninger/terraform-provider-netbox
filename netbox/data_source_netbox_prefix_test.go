package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixDataSource_basic(t *testing.T) {
	testv4Prefix := "10.0.0.0/24"
	testv6Prefix := "2000::/64"
	testSlug := "prefix_ds_basic"
	testVlanVid := 4090
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name = "%[1]s_vrf"
}

resource "netbox_vlan" "test" {
  name = "%[1]s_vlan_test_id"
  vid  = %[4]d
}

resource "netbox_tenant" "test" {
  name = "%[1]s_tenant"
}

resource "netbox_site" "test" {
  name = "%[1]s_site"
}

resource "netbox_ipam_role" "test" {
  name = "%[1]s_role"
}

resource "netbox_prefix" "testv4" {
  prefix      = "%[2]s"
  status      = "active"
  vrf_id      = netbox_vrf.test.id
  vlan_id     = netbox_vlan.test.id
  tenant_id   = netbox_tenant.test.id
  site_id     = netbox_site.test.id
  role_id     = netbox_ipam_role.test.id
  description = "%[1]s_description_test_idv4"
}

resource "netbox_prefix" "testv6" {
  prefix      = "%[3]s"
  status      = "container"
  vrf_id      = netbox_vrf.test.id
  vlan_id     = netbox_vlan.test.id
  tenant_id   = netbox_tenant.test.id
  site_id     = netbox_site.test.id
  description = "%[1]s_description_test_idv6"
}

data "netbox_prefix" "by_description" {
  description = netbox_prefix.testv4.description
}

data "netbox_prefix" "by_cidr" {
  depends_on = [netbox_prefix.testv4]
  cidr       = "%[2]s"
}

data "netbox_prefix" "by_vrf_id" {
  depends_on = [netbox_prefix.testv4]
  vrf_id     = netbox_vrf.test.id
  family     = 4
}

data "netbox_prefix" "by_vlan_id" {
  depends_on = [netbox_prefix.testv4]
  vlan_id    = netbox_vlan.test.id
  family     = 4
}

data "netbox_prefix" "by_vlan_vid" {
  depends_on = [netbox_prefix.testv4]
  vlan_vid   = %[4]d
  family     = 4
}

data "netbox_prefix" "by_prefix" {
  depends_on = [netbox_prefix.testv4]
  prefix     = "%[2]s"
}

data "netbox_prefix" "by_tenant_id" {
  depends_on = [netbox_prefix.testv4]
  tenant_id  = netbox_tenant.test.id
  family     = 4
}

data "netbox_prefix" "by_site_id" {
  depends_on = [netbox_prefix.testv4]
  site_id    = netbox_site.test.id
  family     = 4
}

data "netbox_prefix" "by_role_id" {
  depends_on = [netbox_prefix.testv4]
  role_id    = netbox_ipam_role.test.id
}

data "netbox_prefix" "by_status" {
  depends_on = [netbox_prefix.testv4]
  status     = "active"
}

data "netbox_prefix" "by_family" {
  depends_on = [netbox_prefix.testv6]
	family   = 6
}`, testName, testv4Prefix, testv6Prefix, testVlanVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_prefix", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_description", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_cidr", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vrf_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_vid", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_tenant_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_site_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_role_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_status", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_family", "id", "netbox_prefix.testv6", "id"),
				),
			},
		},
	})
}

func TestAccNetboxPrefixDataSource_customFields(t *testing.T) {
	testSlug := "prefix_customfields"
	testPrefix := "10.0.0.0/24"
	testField := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name = "%[1]s"
  type = "text"
  content_types = ["ipam.prefix"]
  weight        = 100
}

resource "netbox_prefix" "test" {
  prefix = "%[2]s"
  status = "active"
  custom_fields = {
    "${netbox_custom_field.test.name}" = "test value"
  }
}

data "netbox_prefix" "test_output" {
  depends_on = [netbox_prefix.test]
  prefix = "%[2]s"
}`, testField, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_prefix.test_output", "status", "active"),
					resource.TestCheckResourceAttr("data.netbox_prefix.test_output", "prefix", testPrefix),
					resource.TestCheckResourceAttr("data.netbox_prefix.test_output", "custom_fields."+testField, "test value"),
				),
			},
		},
	})
}
