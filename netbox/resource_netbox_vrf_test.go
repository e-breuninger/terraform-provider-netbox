package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVrfTagDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}

resource "netbox_tag" "test_b" {
  name = "%[1]sb"
}
`, testName)
}

func testAccNetboxVrfTenantDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test_tenant_a" {
  name = "%[1]sa"
}
resource "netbox_tenant" "test_tenant_b" {
  name = "%[1]sb"
}
`, testName)
}

func TestAccNetboxVrf_basic(t *testing.T) {

	testSlug := "vrf_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name        = "%s"
  description = "my-description"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vrf.test", "description", "my-description"),
				),
			},
			{
				ResourceName:      "netbox_vrf.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVrf_tags(t *testing.T) {

	testSlug := "vrf_tag"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVrfTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_vrf" "test_tags" {
  name = "%[1]s"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test_tags", "name", testName),
					resource.TestCheckResourceAttr("netbox_vrf.test_tags", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_vrf.test_tags", "tags.0", testName+"a"),
				),
			},
			{
				Config: testAccNetboxVrfTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_vrf" "test_tags" {
  name = "%[1]s"
  tags = ["%[1]sa", "%[1]sb"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test_tags", "tags.#", "2"),
					resource.TestCheckResourceAttr("netbox_vrf.test_tags", "tags.0", testName+"a"),
					resource.TestCheckResourceAttr("netbox_vrf.test_tags", "tags.1", testName+"b"),
				),
			},
			{
				Config: testAccNetboxVrfTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_vrf" "test_tags" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test_tags", "tags.#", "0"),
				),
			},
		},
	})
}

func TestAccNetboxVrf_tenant(t *testing.T) {

	testSlug := "vrf_tenant"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVrfTenantDependencies(testName) + fmt.Sprintf(`
resource "netbox_vrf" "test_tenant" {
  name = "%[1]s"
  tenant_id = netbox_tenant.test_tenant_a.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test_tenant", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_vrf.test_tenant", "tenant_id", "netbox_tenant.test_tenant_a", "id"),
				),
			},
			{
				Config: testAccNetboxVrfTenantDependencies(testName) + fmt.Sprintf(`
resource "netbox_vrf" "test_tenant" {
  name = "%[1]s"
  tenant_id = netbox_tenant.test_tenant_b.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test_tenant", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_vrf.test_tenant", "tenant_id", "netbox_tenant.test_tenant_b", "id"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_vrf", &resource.Sweeper{
		Name:         "netbox_vrf",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamVrfsListParams()
			res, err := api.Ipam.IpamVrfsList(params, nil)
			if err != nil {
				return err
			}
			for _, vrf := range res.GetPayload().Results {
				if strings.HasPrefix(*vrf.Name, testPrefix) {
					deleteParams := ipam.NewIpamVrfsDeleteParams().WithID(vrf.ID)
					_, err := api.Ipam.IpamVrfsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a vrf")
				}
			}
			return nil
		},
	})
}
