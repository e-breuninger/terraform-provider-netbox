//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxModuleType_basic(t *testing.T) {
	testName := testAccGetTestName("modtype")
	deps := fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
  part_number     = "PN-%[1]s"
  weight          = 1.2
  weight_unit     = "kg"
  description     = "%[1]s"
  comments        = "Created by acceptance test."
}
data "netbox_module_type" "test" {
  depends_on = [netbox_module_type.test]
  model      = "%[1]s"
}
data "netbox_module_types" "all" {
  depends_on     = [netbox_module_type.test]
  filters = [
    { name = "model__ic", value = "%[1]s" },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_module_type.test", "part_number", "PN-"+testName),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight", "1.2"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight_unit", "kg"),
					resource.TestCheckResourceAttrPair("netbox_module_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_module_type.test", "id", "netbox_module_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_module_types.all", "module_types.#", "1"),
				),
			},
			{
				// Shrink: part number, weight and unit must clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_module_type.test", "part_number"),
					resource.TestCheckNoResourceAttr("netbox_module_type.test", "weight"),
					resource.TestCheckNoResourceAttr("netbox_module_type.test", "weight_unit"),
					resource.TestCheckNoResourceAttr("netbox_module_type.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_module_type.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_module_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_module_type",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimModuleTypesList(dcim.NewDcimModuleTypesListParams(), nil)
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
			_, err := client.Dcim.DcimModuleTypesDestroy(dcim.NewDcimModuleTypesDestroyParams().WithID(id), nil)
			return err
		})
}
