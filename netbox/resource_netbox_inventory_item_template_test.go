package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxInventoryItemTemplate_basic(t *testing.T) {
	testSlug := "inventory_item_template"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_inventory_item_template" "test" {
	name = "%[1]s"
	device_type_id = netbox_device_type.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_inventory_item_template.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_inventory_item_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_inventory_item_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxInventoryItemTemplate_opts(t *testing.T) {
	testSlug := "inventory_item_template"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_inventory_item_role" "test" {
	name = "%[1]s"
	slug = "%[2]s"
	color_hex = "f44336"
}

resource "netbox_inventory_item_template" "parent" {
	name = "%[1]s_parent"
	device_type_id = netbox_device_type.test.id
}

resource "netbox_inventory_item_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	device_type_id = netbox_device_type.test.id
	parent_id = netbox_inventory_item_template.parent.id
	role_id = netbox_inventory_item_role.test.id
	manufacturer_id = netbox_manufacturer.test.id
	part_id = "%[1]s-part"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_inventory_item_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_inventory_item_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_inventory_item_template.test", "label", fmt.Sprintf("%s label", testName)),
					resource.TestCheckResourceAttr("netbox_inventory_item_template.test", "part_id", fmt.Sprintf("%s-part", testName)),
					resource.TestCheckResourceAttrPair("netbox_inventory_item_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item_template.test", "parent_id", "netbox_inventory_item_template.parent", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item_template.test", "role_id", "netbox_inventory_item_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item_template.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_inventory_item_template", &resource.Sweeper{
		Name:         "netbox_inventory_item_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimInventoryItemTemplatesListParams()
			res, err := api.Dcim.DcimInventoryItemTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, item := range res.GetPayload().Results {
				if strings.HasPrefix(*item.Name, testPrefix) {
					deleteParams := dcim.NewDcimInventoryItemTemplatesDeleteParams().WithID(item.ID)
					_, err := api.Dcim.DcimInventoryItemTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an inventory item template")
				}
			}
			return nil
		},
	})
}
