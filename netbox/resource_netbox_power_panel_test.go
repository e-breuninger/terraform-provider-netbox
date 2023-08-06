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

func testAccNetboxPowerPanelFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_location" "test" {
  name = "%[1]s"
  site_id =netbox_site.test.id
}

resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}`, testName)
}

func TestAccNetboxPowerPanel_basic(t *testing.T) {
	testSlug := "power_panel_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckPowerPanelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxPowerPanelFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_power_panel" "test" {
  name = "%[1]s"
  description = "%[1]sdescription"
  comments = "%[1]scomments"
  
  site_id = netbox_site.test.id
  location_id = netbox_location.test.id
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_panel.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "description", testName+"description"),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "comments", testName+"comments"),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttrPair("netbox_power_panel.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_power_panel.test", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				Config: testAccNetboxPowerPanelFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_power_panel" "test" {
  name = "%[1]s" 
  site_id = netbox_site.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_panel.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_power_panel.test", "location_id", "0"),

					resource.TestCheckResourceAttrPair("netbox_power_panel.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_power_panel.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPowerPanelDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each power panel
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_power_panel" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimPowerPanelsReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimPowerPanelsRead(params, nil)

		if err == nil {
			return fmt.Errorf("power panel (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimPowerPanelsReadDefault); ok {
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
	resource.AddTestSweepers("netbox_power_panel", &resource.Sweeper{
		Name:         "netbox_power_panel",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimPowerPanelsListParams()
			res, err := api.Dcim.DcimPowerPanelsList(params, nil)
			if err != nil {
				return err
			}
			for _, powerPanel := range res.GetPayload().Results {
				if strings.HasPrefix(*powerPanel.Name, testPrefix) {
					deleteParams := dcim.NewDcimPowerPanelsDeleteParams().WithID(powerPanel.ID)
					_, err := api.Dcim.DcimPowerPanelsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a power_panel")
				}
			}
			return nil
		},
	})
}
