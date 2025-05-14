package netbox

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxPrefixFullDependencies(testName string, testSlug string, testVid string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_vrf" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_ipam_role" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}

resource "netbox_vlan" "test" {
  name = "%[1]s"
  vid = "%[3]s"
  status = "active"
  description = "Test"
  tags = []
}
`, testName, testSlug, testVid)
}

func TestAccNetboxPrefix_basic(t *testing.T) {
	testPrefix := "1.1.1.128/25"
	testSlug := "prefix"
	testVid := "123"
	randomSlug := testAccGetTestName(testSlug)
	testDesc := "test prefix"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s"
  status = "active"
  tags = [netbox_tag.test.name]
  mark_utilized = true
  is_pool = true
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "description", testDesc),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_prefix.test", "mark_utilized", "true"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "is_pool", "true"),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s"
  status = "provoke_error"
  tags = [netbox_tag.test.name]
  mark_utilized = true
}`, testPrefix, testDesc),
				ExpectError: regexp.MustCompile("expected status to be one of .*"),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s"
  status = "active"
  tags = [netbox_tag.test.name]
  mark_utilized = false
  is_pool = false
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "mark_utilized", "false"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "is_pool", "false"),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s"
  status = "deprecated"
  tags = [netbox_tag.test.name]
  mark_utilized = true
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "deprecated"),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s"
  status = "container"
  tags = [netbox_tag.test.name]
  mark_utilized = true
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "container"),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s 2"
  status = "active"
  tags = [netbox_tag.test.name]
  mark_utilized = true
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "description", fmt.Sprintf("%s 2", testDesc)),
					resource.TestCheckResourceAttr("netbox_prefix.test", "vrf_id", "0"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tenant_id", "0"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "site_id", "0"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_prefix.test", "mark_utilized", "true"),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s 2"
  status = "active"
  vrf_id = netbox_vrf.test.id
  tenant_id = netbox_tenant.test.id
  tags = [netbox_tag.test.name]
  mark_utilized = true
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "description", fmt.Sprintf("%s 2", testDesc)),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_prefix.test", "mark_utilized", "true"),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s 2"
  status = "active"
  vrf_id = netbox_vrf.test.id
  tenant_id = netbox_tenant.test.id
  site_id = netbox_site.test.id
  vlan_id = netbox_vlan.test.id
  role_id = netbox_ipam_role.test.id
  tags = [netbox_tag.test.name]
  mark_utilized = true
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "description", fmt.Sprintf("%s 2", testDesc)),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "vlan_id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_prefix.test", "mark_utilized", "true"),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s 2"
  status = "active"
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "description", fmt.Sprintf("%s 2", testDesc)),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  status = "active"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
				),
			},
			{
				ResourceName:      "netbox_prefix.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPrefix_cf(t *testing.T) {
	testPrefix := "1.1.2.128/25"
	testSlug := "prefix_cf"
	testVid := "124"
	randomSlug := testAccGetTestName(testSlug)
	testDesc := "test cf prefix"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxPrefixFullDependencies(testName, randomSlug, testVid) + fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name   = "%s"
  type   = "text"
  weight = 100
  content_types = ["ipam.prefix"]
}

resource "netbox_prefix" "test" {
  prefix = "%s"
  description = "%s 2"
  status = "active"
  mark_utilized = true

  custom_fields = {
    "${netbox_custom_field.test.name}" = "test-field"
  }
}`, testSlug, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", fmt.Sprintf("custom_fields.%s", testSlug), "test-field"),
				),
			},
			{
				ResourceName:      "netbox_prefix.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_prefix", &resource.Sweeper{
		Name:         "netbox_prefix",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamPrefixesListParams()
			res, err := api.Ipam.IpamPrefixesList(params, nil)
			if err != nil {
				return err
			}
			for _, prefix := range res.GetPayload().Results {
				if len(prefix.Tags) > 0 && (prefix.Tags[0] == &models.NestedTag{Name: strToPtr("acctest"), Slug: strToPtr("acctest")}) {
					deleteParams := ipam.NewIpamPrefixesDeleteParams().WithID(prefix.ID)
					_, err := api.Ipam.IpamPrefixesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a prefix")
				}
			}
			return nil
		},
	})
}
