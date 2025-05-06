package netbox

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxInventoryItemFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_tag" "test" {
  name = "%[1]sa"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s"
}

resource "netbox_module" "test" {
  device_id = netbox_device.test.id
  module_bay_id = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status = "active"
}

resource "netbox_device_rear_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "8p8c"
  positions = 1
  mark_connected = true
}

resource "netbox_inventory_item_role" "test" {
	name = "%[1]s"
  slug = "%[1]s_slug"
	color_hex = "123456"
}
`, testName)
}

func TestAccNetboxInventoryItem_basic(t *testing.T) {
	testSlug := "inventory_item_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckInventoryItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxInventoryItemFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_inventory_item" "parent" {
	device_id = netbox_device.test.id
	name = "%[1]s_parent"
}

resource "netbox_inventory_item" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"

	parent_id = netbox_inventory_item.parent.id
  label = "%[1]s_label"
	role_id = netbox_inventory_item_role.test.id
	manufacturer_id = netbox_manufacturer.test.id
	part_id = "%[1]s_part"
	serial = "%[1]s_serial"
	asset_tag = "%[1]s_asset"
	discovered = true
	description = "%[1]s_description"
	component_type = "dcim.rearport"
	component_id = netbox_device_rear_port.test.id
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_inventory_item.parent", "name", testName+"_parent"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.parent", "device_id", "netbox_device.test", "id"),

					resource.TestCheckResourceAttr("netbox_inventory_item.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "label", testName+"_label"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "part_id", testName+"_part"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "serial", testName+"_serial"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "asset_tag", testName+"_asset"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "discovered", "true"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "component_type", "dcim.rearport"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "parent_id", "netbox_inventory_item.parent", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "role_id", "netbox_inventory_item_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "component_id", "netbox_device_rear_port.test", "id"),
				),
			},
			{
				Config: testAccNetboxInventoryItemFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_inventory_item" "parent" {
	device_id = netbox_device.test.id
	name = "%[1]s_parent"
}

resource "netbox_inventory_item" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "label", ""),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "part_id", ""),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "serial", ""),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "asset_tag", ""),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "discovered", "false"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "component_type", ""),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "parent_id", "0"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "role_id", "0"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "manufacturer_id", "0"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "component_id", "0"),

					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "device_id", "netbox_device.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_inventory_item.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckInventoryItemDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each inventory item
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_inventory_item" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimInventoryItemsReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimInventoryItemsRead(params, nil)

		if err == nil {
			return fmt.Errorf("inventory_item (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimInventoryItemsReadDefault); ok {
				errorcode := errresp.Code()
				if errorcode == 404 {
					return nil
				}
			}
			return err
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("netbox_inventory_item", &resource.Sweeper{
		Name:         "netbox_inventory_item",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimInventoryItemsListParams()
			res, err := api.Dcim.DcimInventoryItemsList(params, nil)
			if err != nil {
				return err
			}
			for _, rearPort := range res.GetPayload().Results {
				if strings.HasPrefix(*rearPort.Name, testPrefix) {
					deleteParams := dcim.NewDcimInventoryItemsDeleteParams().WithID(rearPort.ID)
					_, err := api.Dcim.DcimInventoryItemsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an inventory_item")
				}
			}
			return nil
		},
	})
}
