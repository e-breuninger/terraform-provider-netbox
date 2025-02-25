package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxLocation_basic(t *testing.T) {
	testSlug := "location_basic"
	testName := testAccGetTestName(testSlug)
	testNameSub := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	randomSlugSub := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "my-description"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
}

resource "netbox_location" "test-sub" {
  name        = "%[3]s"
  slug        = "%[4]s"
  description = "my-description"
  parent_id   = netbox_location.test.id
  site_id     = netbox_site.test.id
}`, testName, randomSlug, testNameSub, randomSlugSub),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_location.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_location.test", "description", "my-description"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "id", "netbox_location.test-sub", "parent_id"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test_2" {
  name = "%[1]s_b"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name = "%[1]s"
  slug = "%[2]s"
  site_id = netbox_site.test_2.id
  tenant_id = netbox_tenant.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_location.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_location.test", "description", ""),
				),
			},
			{
				ResourceName:      "netbox_location.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxLocation_updateParent(t *testing.T) {
	testSlug := "loc_upd_parent"
	testName := testAccGetTestName(testSlug)
	testNameSub := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	randomSlugSub := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxLocationUpdateParent1(testName, randomSlug, testNameSub, randomSlugSub),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_location.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_location.test", "description", "my-description"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "id", "netbox_location.test_sub", "parent_id"),
				),
			},
			{
				Config: testAccNetboxLocationUpdateParent2(testName, randomSlug, testNameSub, randomSlugSub),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_location.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_location.test", "description", "my-description"),
					resource.TestCheckResourceAttr("netbox_location.test_sub", "parent_id", "0"),
				),
			},
		},
	})
}

func testAccNetboxLocationUpdateParent1(testName string, randomSlug string, testNameSub string, randomSlugSub string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "my-description"
  site_id     = netbox_site.test.id
}

resource "netbox_location" "test_sub" {
  name        = "%[3]s"
  slug        = "%[4]s"
  description = "my-description"
  parent_id   = netbox_location.test.id
  site_id     = netbox_site.test.id
}`, testName, randomSlug, testNameSub, randomSlugSub)
}

func testAccNetboxLocationUpdateParent2(testName string, randomSlug string, testNameSub string, randomSlugSub string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "my-description"
  site_id     = netbox_site.test.id
}

resource "netbox_location" "test_sub" {
  name        = "%[3]s"
  slug        = "%[4]s"
  description = "my-description"
  parent_id   = "0"
  site_id     = netbox_site.test.id
}`, testName, randomSlug, testNameSub, randomSlugSub)
}

func init() {
	resource.AddTestSweepers("netbox_location", &resource.Sweeper{
		Name:         "netbox_location",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimLocationsListParams()
			res, err := api.Dcim.DcimLocationsList(params, nil)
			if err != nil {
				return err
			}
			for _, Location := range res.GetPayload().Results {
				if strings.HasPrefix(*Location.Name, testPrefix) {
					deleteParams := dcim.NewDcimLocationsDeleteParams().WithID(Location.ID)
					_, err := api.Dcim.DcimLocationsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a location")
				}
			}
			return nil
		},
	})
}
