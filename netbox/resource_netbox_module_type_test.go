package netbox

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxModuleTypeFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test" {
	name = "%[1]sa"
}
`, testName)
}

func TestAccNetboxModuleType_basic(t *testing.T) {
	testSlug := "module_type_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckModuleTypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxModuleTypeFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s"
  part_number = "%[1]s_pn"
  description = "%[1]s_description"
  comments = "%[1]s_comments"

  weight = 1
  weight_unit = "kg"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_module_type.test", "part_number", testName+"_pn"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "comments", testName+"_comments"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight", "1"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight_unit", "kg"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttrPair("netbox_module_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
				),
			},
			{
				Config: testAccNetboxModuleTypeFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_module_type.test", "part_number", ""),
					resource.TestCheckResourceAttr("netbox_module_type.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_module_type.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight", "0"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight_unit", ""),
					resource.TestCheckResourceAttr("netbox_module_type.test", "tags.#", "0"),

					resource.TestCheckResourceAttrPair("netbox_module_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
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

func testAccCheckModuleTypeDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*providerState)

	// loop through the resources in state, verifying each module type
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_module_type" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimModuleTypesReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimModuleTypesRead(params, nil)

		if err == nil {
			return fmt.Errorf("module type (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimModuleTypesReadDefault); ok {
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
	resource.AddTestSweepers("netbox_module_type", &resource.Sweeper{
		Name:         "netbox_module_type",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimModuleTypesListParams()
			res, err := api.Dcim.DcimModuleTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, moduleType := range res.GetPayload().Results {
				if strings.HasPrefix(*moduleType.Model, testPrefix) {
					deleteParams := dcim.NewDcimModuleTypesDeleteParams().WithID(moduleType.ID)
					_, err := api.Dcim.DcimModuleTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a module_type")
				}
			}
			return nil
		},
	})
}
