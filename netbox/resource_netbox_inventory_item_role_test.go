package netbox

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxInventoryItemRoleFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]sa"
}
`, testName)
}

func TestAccNetboxInventoryItemRole_basic(t *testing.T) {
	testSlug := "inventory_item_role_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckInventoryItemRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxInventoryItemRoleFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_inventory_item_role" "test" {
  name = "%[1]s"
  slug = "%[1]s_slug"
	color_hex = "123456"

  description = "%[1]s_description"
  tags = [ netbox_tag.test.name ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "slug", testName+"_slug"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "color_hex", "123456"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "tags.0", testName+"a"),
				),
			},
			{
				Config: testAccNetboxInventoryItemRoleFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_inventory_item_role" "test" {
  name = "%[1]s"
  slug = "%[1]s_slug"
	color_hex = "123456"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "slug", testName+"_slug"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "color_hex", "123456"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "tags.#", "0"),
				),
			},
			{
				ResourceName:      "netbox_inventory_item_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckInventoryItemRoleDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each inventory item role
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_inventory_item_role" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimInventoryItemRolesReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimInventoryItemRolesRead(params, nil)

		if err == nil {
			return fmt.Errorf("inventory item role (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimInventoryItemRolesReadDefault); ok {
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
	resource.AddTestSweepers("netbox_inventory_item_role", &resource.Sweeper{
		Name:         "netbox_inventory_item_role",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimInventoryItemRolesListParams()
			res, err := api.Dcim.DcimInventoryItemRolesList(params, nil)
			if err != nil {
				return err
			}
			for _, role := range res.GetPayload().Results {
				if strings.HasPrefix(*role.Name, testPrefix) {
					deleteParams := dcim.NewDcimInventoryItemRolesDeleteParams().WithID(role.ID)
					_, err := api.Dcim.DcimInventoryItemRolesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an inventory item role")
				}
			}
			return nil
		},
	})
}
