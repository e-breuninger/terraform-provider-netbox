package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccNetboxRackFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_location" "test" {
	name = "%[1]s"
	site_id =netbox_site.test.id
}

resource "netbox_rack_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}`, testName)
}

func TestAccNetboxRack_basic(t *testing.T) {

	testSlug := "rack_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRackDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxRackFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_rack" "test" {
  name = "%[1]s"
	site_id = netbox_site.test.id
	status = "reserved"
	width = 19
	u_height = 48
	tags = ["%[1]sa"]
	tenant_id = netbox_tenant.test.id
	facility_id = "%[1]sfacility"
	location_id = netbox_location.test.id
	role_id = netbox_rack_role.test.id
	serial = "%[1]sserial"
	asset_tag = "%[1]sasset_tag"
	type = "4-post-frame"
	desc_units = true
	outer_width = 10
	outer_depth = 15
	outer_unit = "mm"
	comments = "%[1]scomments"
}
resource "netbox_rack" "test2" {
  name = "%[1]s2"
	site_id = netbox_site.test.id
	location_id = netbox_location.test.id
	status = "reserved"
	width = 19
	u_height = 48
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_rack.test", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack.test", "u_height", "48"),
					resource.TestCheckResourceAttr("netbox_rack.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_rack.test", "tags.0", testName+"a"),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "facility_id", testName+"facility"),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "role_id", "netbox_rack_role.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "serial", testName+"serial"),
					resource.TestCheckResourceAttr("netbox_rack.test", "asset_tag", testName+"asset_tag"),
					resource.TestCheckResourceAttr("netbox_rack.test", "type", "4-post-frame"),
					resource.TestCheckResourceAttr("netbox_rack.test", "desc_units", "true"),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_width", "10"),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_depth", "15"),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_unit", "mm"),
					resource.TestCheckResourceAttr("netbox_rack.test", "comments", testName+"comments"),

					resource.TestCheckResourceAttr("netbox_rack.test2", "name", testName+"2"),
					resource.TestCheckResourceAttrPair("netbox_rack.test2", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_rack.test2", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test2", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_rack.test2", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack.test2", "u_height", "48"),
				),
			},
			{
				Config: testAccNetboxRackFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_rack" "test" {
  name = "%[1]s"
	site_id = netbox_site.test.id
	status = "reserved"
	width = 19
	u_height = 48
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_rack.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_rack.test", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack.test", "u_height", "48"),
					resource.TestCheckResourceAttr("netbox_rack.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "tenant_id", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "facility_id", ""),
					resource.TestCheckResourceAttr("netbox_rack.test", "location_id", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "role_id", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "serial", ""),
					resource.TestCheckResourceAttr("netbox_rack.test", "asset_tag", ""),
					resource.TestCheckResourceAttr("netbox_rack.test", "type", ""),
					resource.TestCheckResourceAttr("netbox_rack.test", "weight", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "max_weight", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "weight_unit", ""),
					resource.TestCheckResourceAttr("netbox_rack.test", "desc_units", "false"),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_width", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_depth", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "outer_unit", ""),
					resource.TestCheckResourceAttr("netbox_rack.test", "mounting_depth", "0"),
					resource.TestCheckResourceAttr("netbox_rack.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_rack.test", "comments", ""),
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

func testAccCheckRackDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each rack
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_rack" {
			continue
		}

		// Retrieve our rack by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimRacksReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimRacksRead(params, nil)

		if err == nil {
			return fmt.Errorf("rack (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimRacksReadDefault); ok {
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
	resource.AddTestSweepers("netbox_rack", &resource.Sweeper{
		Name:         "netbox_rack",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimRacksListParams()
			res, err := api.Dcim.DcimRacksList(params, nil)
			if err != nil {
				return err
			}
			for _, Rack := range res.GetPayload().Results {
				if strings.HasPrefix(*Rack.Name, testPrefix) {
					deleteParams := dcim.NewDcimRacksDeleteParams().WithID(Rack.ID)
					_, err := api.Dcim.DcimRacksDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a rack")
				}
			}
			return nil
		},
	})
}
