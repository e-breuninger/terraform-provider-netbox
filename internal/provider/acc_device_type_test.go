//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxDeviceType_basic(t *testing.T) {
	testName := testAccGetTestName("devtype")
	deps := fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}
resource "netbox_platform" "test" {
  name = "%[1]s"
}`, testName, getSlug(testName))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_type" "test" {
  manufacturer_id          = netbox_manufacturer.test.id
  model                    = "%[1]s"
  slug                     = "%[2]s"
  part_number              = "PN-%[1]s"
  default_platform_id      = netbox_platform.test.id
  u_height                 = 1.5
  is_full_depth            = false
  exclude_from_utilization = true
  subdevice_role           = "parent"
  airflow                  = "front-to-rear"
  weight                   = 12.5
  weight_unit              = "kg"
  description              = "%[1]s"
  comments                 = "Created by acceptance test."
}
data "netbox_device_type" "test" {
  depends_on = [netbox_device_type.test]
  model      = "%[1]s"
}
data "netbox_device_types" "all" {
  depends_on     = [netbox_device_type.test]
  filters = [
    { name = "model__ic", value = "%[1]s" },
  ]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_device_type.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_device_type.test", "u_height", "1.5"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "is_full_depth", "false"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "exclude_from_utilization", "true"),
					resource.TestCheckResourceAttrPair("netbox_device_type.test", "default_platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "subdevice_role", "parent"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "airflow", "front-to-rear"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "weight", "12.5"),
					resource.TestCheckResourceAttrPair("netbox_device_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_device_type.test", "id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_types.all", "device_types.#", "1"),
				),
			},
			{
				// Shrink: choice fields and part number must clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "subdevice_role"),
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "airflow"),
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "part_number"),
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "default_platform_id"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "exclude_from_utilization", "false"),
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "weight"),
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "weight_unit"),
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_device_type.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_device_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxDeviceType_defaults(t *testing.T) {
	testName := testAccGetTestName("devtype_min")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					// is_full_depth defaults to true (schema default matching NetBox); u_height to 1 server-side.
					resource.TestCheckResourceAttr("netbox_device_type.test", "is_full_depth", "true"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "u_height", "1"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_device_type",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimDeviceTypesList(dcim.NewDcimDeviceTypesListParams(), nil)
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
			_, err := client.Dcim.DcimDeviceTypesDestroy(dcim.NewDcimDeviceTypesDestroyParams().WithID(id), nil)
			return err
		})
}
