//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxRack_basic(t *testing.T) {
	testName := testAccGetTestName("rack")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_rack_role" "test" {
  name      = "%[1]s"
  slug      = "%[2]s"
  color_hex = "112233"
}
resource "netbox_rack_group" "test" {
  name = "%[1]s"
}`, testName, getSlug(testName))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_rack" "test" {
  name           = "%[1]s"
  site_id        = netbox_site.test.id
  location_id    = netbox_location.test.id
  tenant_id      = netbox_tenant.test.id
  role_id        = netbox_rack_role.test.id
  group_id       = netbox_rack_group.test.id
  status         = "active"
  form_factor    = "4-post-cabinet"
  airflow        = "front-to-rear"
  width          = 19
  u_height       = 47
  starting_unit  = 2
  desc_units     = true
  serial         = "SN-%[1]s"
  asset_tag      = "AT-%[1]s"
  facility_id    = "ROW1-R1"
  outer_width    = 600
  outer_height   = 2000
  outer_depth    = 1200
  outer_unit     = "mm"
  mounting_depth = 900
  weight         = 120.5
  max_weight     = 1000
  weight_unit    = "kg"
  description    = "%[1]s"
  comments       = "Installed by acceptance test."
}
data "netbox_rack" "test" {
  depends_on = [netbox_rack.test]
  name       = "%[1]s"
}
data "netbox_racks" "by_site" {
  depends_on = [netbox_rack.test]
  filters = [
    { name = "site_id", value = netbox_site.test.id },
    { name = "status", value = "active" },
    { name = "u_height", value = "47" },
    { name = "desc_units", value = "true" },
    { name = "outer_unit", value = "mm" },
    { name = "serial", value = "SN-%[1]s" },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					// The pre-6.0 rack filters, across the parameter shapes: repeatable ints and
					// strings, single-value bool and string.
					resource.TestCheckResourceAttr("data.netbox_racks.by_site", "racks.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_racks.by_site", "racks.0.id", "netbox_rack.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "group_id", "netbox_rack_group.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "form_factor", "4-post-cabinet"),
					resource.TestCheckResourceAttr("netbox_rack.test", "airflow", "front-to-rear"),
					resource.TestCheckResourceAttr("netbox_rack.test", "starting_unit", "2"),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_height", "2000"),
					resource.TestCheckResourceAttr("netbox_rack.test", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack.test", "u_height", "47"),
					resource.TestCheckResourceAttr("netbox_rack.test", "desc_units", "true"),
					resource.TestCheckResourceAttr("netbox_rack.test", "asset_tag", "AT-"+testName),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_unit", "mm"),
					resource.TestCheckResourceAttr("netbox_rack.test", "weight", "120.5"),
					resource.TestCheckResourceAttr("netbox_rack_role.test", "color_hex", "112233"),
					resource.TestCheckResourceAttrPair("data.netbox_rack.test", "id", "netbox_rack.test", "id"),
				),
			},
			{
				// Minimal update: shrink and drop the asset tag / outer dims.
				Config: deps + fmt.Sprintf(`
resource "netbox_rack" "test" {
  name     = "%[1]s"
  site_id  = netbox_site.test.id
  status   = "planned"
  u_height = 42
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_rack.test", "u_height", "42"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "asset_tag"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "location_id"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "group_id"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "form_factor"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "airflow"),
					// starting_unit is computed: unset keeps the value.
					resource.TestCheckResourceAttr("netbox_rack.test", "starting_unit", "2"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "serial"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "facility_id"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "outer_width"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "outer_height"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "outer_depth"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "outer_unit"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "mounting_depth"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "weight"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "max_weight"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "weight_unit"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_rack.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_rack.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxRack_defaults(t *testing.T) {
	testName := testAccGetTestName("rack_min")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_rack" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					// Server-side defaults must land in state.
					resource.TestCheckResourceAttr("netbox_rack.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_rack.test", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack.test", "u_height", "42"),
					resource.TestCheckResourceAttr("netbox_rack.test", "desc_units", "false"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_rack",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimRacksList(dcim.NewDcimRacksListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimRacksDestroy(dcim.NewDcimRacksDestroyParams().WithID(id), nil)
			return err
		})
}
