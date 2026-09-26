//go:build acctest

// Rack types, and the rack that inherits from one.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxRackType_basic(t *testing.T) {
	testName := testAccGetTestName("racktype")
	deps := fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_rack_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
  slug            = "%[2]s"
  description     = "Acceptance test rack type."
  form_factor     = "4-post-cabinet"
  width           = 23
  u_height        = 47
  starting_unit   = 2
  desc_units      = true
  outer_width     = 600
  outer_height    = 2200
  outer_depth     = 1000
  outer_unit      = "mm"
  mounting_depth  = 900
  weight          = 120.5
  max_weight      = 1500
  weight_unit     = "kg"
  comments        = "Created by acceptance test."
}
data "netbox_rack_type" "test" {
  model = netbox_rack_type.test.model
}
data "netbox_rack_types" "test" {
  filters = [
    { name = "model__ic", value = "%[1]s" },
  ]
  depends_on     = [netbox_rack_type.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "form_factor", "4-post-cabinet"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "width", "23"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "u_height", "47"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "starting_unit", "2"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "desc_units", "true"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "outer_unit", "mm"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "weight", "120.5"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "weight_unit", "kg"),
					resource.TestCheckResourceAttrPair("data.netbox_rack_type.test", "id", "netbox_rack_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_rack_types.test", "rack_types.#", "1"),
				),
			},
			{
				// Shrink the type: width, u_height, starting_unit and desc_units are optional+computed
				// and keep their values, as on netbox_rack; the rest clears. A rack built from the type
				// inherits its physical attributes, so the rack declares the same form factor to keep
				// its plan clean.
				Config: deps + fmt.Sprintf(`
resource "netbox_rack_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
  form_factor     = "2-post-frame"
}
resource "netbox_rack" "test" {
  name         = "%[1]s"
  site_id      = netbox_site.test.id
  rack_type_id = netbox_rack_type.test.id
  form_factor  = "2-post-frame"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_type.test", "form_factor", "2-post-frame"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "width", "23"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "u_height", "47"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "starting_unit", "2"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "desc_units", "true"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "outer_width"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "outer_height"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "outer_depth"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "outer_unit"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "mounting_depth"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "weight"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "max_weight"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "weight_unit"),
					resource.TestCheckNoResourceAttr("netbox_rack_type.test", "comments"),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "rack_type_id", "netbox_rack_type.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "form_factor", "2-post-frame"),
				),
			},
			{
				// The rack detaches from its type in place.
				Config: deps + fmt.Sprintf(`
resource "netbox_rack_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
  form_factor     = "2-post-frame"
}
resource "netbox_rack" "test" {
  name        = "%[1]s"
  site_id     = netbox_site.test.id
  form_factor = "2-post-frame"
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_rack.test", "rack_type_id"),
			},
			{
				ResourceName:      "netbox_rack_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_rack_type",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimRackTypesList(dcim.NewDcimRackTypesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Model)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimRackTypesDestroy(dcim.NewDcimRackTypesDestroyParams().WithID(id), nil)
			return err
		})
}
