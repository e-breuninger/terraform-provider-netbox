package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxConfigContext_basic(t *testing.T) {
	testSlug := "config_context_assignments"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name = "%s"
  description = "test description"
  data = jsonencode({"testkey" = "testval"})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_context.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_config_context.test", "data", "{\"testkey\":\"testval\"}"),
				),
			},
			{
				ResourceName:      "netbox_config_context.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxConfigContext_defaultWeight(t *testing.T) {
	testSlug := "config_context_assignments"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name = "%s"
  data = jsonencode({"testkey" = "testval"})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_context.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_config_context.test", "weight", "1000"),
				),
			},
		},
	})
}

func TestAccNetboxConfigContext_assignments(t *testing.T) {
	testSlug := "config_context_assignments"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_tenant_group" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
  group_id = netbox_tenant_group.test.id
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_site_group" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
  group_id = netbox_site_group.test.id
}

resource "netbox_region" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_group" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id = netbox_site.test.id
  cluster_group_id = netbox_cluster_group.test.id
}

resource "netbox_location" "test" {
    name = "%[1]s"
    site_id =netbox_site.test.id
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}
# Untested: cluster_groups, regions, site_groups, tenant_groups
resource "netbox_config_context" "test" {
  name = "%[1]s"
  data = jsonencode({"testkey" = "testval"})
  regions = [netbox_region.test.id]
  tenant_groups = [netbox_tenant_group.test.id]
  tenants = [netbox_tenant.test.id]
  platforms = [netbox_platform.test.id]
  sites = [netbox_site.test.id]
  site_groups = [netbox_site_group.test.id]
  cluster_types = [netbox_cluster_type.test.id]
  cluster_groups = [netbox_cluster_group.test.id]
  clusters = [netbox_cluster.test.id]
  locations = [netbox_location.test.id]
  roles = [netbox_device_role.test.id]
  tags = [netbox_tag.test.name]
  device_types = [netbox_device_type.test.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_context.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "regions.0", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "tenant_groups.0", "netbox_tenant_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "tenants.0", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "platforms.0", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "sites.0", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "site_groups.0", "netbox_site_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "cluster_types.0", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "cluster_groups.0", "netbox_cluster_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "clusters.0", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "locations.0", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "roles.0", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "tags.0", "netbox_tag.test", "name"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "device_types.0", "netbox_device_type.test", "id"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_config_context", &resource.Sweeper{
		Name:         "netbox_config_context",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			state := m.(*providerState)
			api := state.legacyAPI
			params := extras.NewExtrasConfigContextsListParams()
			res, err := api.Extras.ExtrasConfigContextsList(params, nil)
			if err != nil {
				return err
			}
			for _, configContext := range res.GetPayload().Results {
				if strings.HasPrefix(*configContext.Name, testPrefix) {
					deleteParams := extras.NewExtrasConfigContextsDeleteParams().WithID(configContext.ID)
					_, err := api.Extras.ExtrasConfigContextsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a config context")
				}
			}
			return nil
		},
	})
}
