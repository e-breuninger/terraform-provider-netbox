package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testAccNetboxRackTypeFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}`, testName)
}

func TestAccNetboxRackType_basic(t *testing.T) {
	testSlug := "racktype_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRackTypeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxRackTypeFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_rack_type" "test" {
  model             = "%[1]s"
  manufacturer_id   = netbox_manufacturer.test.id
  width             = 19
  u_height          = 48
  starting_unit     = 1
  form_factor       = "2-post-frame"
  tags              = ["%[1]sa"]
  description       = "%[1]s"
  outer_width       = 10
  outer_depth       = 15
  outer_unit        = "mm"
  weight            = 15
  max_weight        = 20
  weight_unit       = "kg"
  mounting_depth_mm = 21
  comments          = "%[1]scomments"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "description", testName),
					resource.TestCheckResourceAttrPair("netbox_rack_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "u_height", "48"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "tags.0", testName+"a"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "outer_width", "10"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "outer_depth", "15"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "outer_unit", "mm"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "weight", "15"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "max_weight", "20"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "weight_unit", "kg"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "comments", testName+"comments"),
					resource.TestCheckResourceAttr("netbox_rack_type.test", "mounting_depth_mm", "21"),
				),
			},
			{
				ResourceName:      "netbox_rack_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckRackTypeDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each rack
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_rack_type" {
			continue
		}

		// Retrieve our rack by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimRackTypesReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimRackTypesRead(params, nil)

		if err == nil {
			return fmt.Errorf("rack type (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimRackTypesReadDefault); ok {
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
	resource.AddTestSweepers("netbox_rack_type", &resource.Sweeper{
		Name:         "netbox_rack_type",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimRackTypesListParams()
			res, err := api.Dcim.DcimRackTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, RackType := range res.GetPayload().Results {
				if strings.HasPrefix(*RackType.Model, testPrefix) {
					deleteParams := dcim.NewDcimRackTypesDeleteParams().WithID(RackType.ID)
					_, err := api.Dcim.DcimRackTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a rack type")
				}
			}
			return nil
		},
	})
}
