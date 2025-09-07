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

func testAccNetboxDevicePowerFeedFullDependencies(testName string) string {
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

resource "netbox_location" "test" {
	name = "%[1]s"
	site_id =netbox_site.test.id
}

resource "netbox_power_panel" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
  location_id = netbox_location.test.id
}

resource "netbox_rack" "test" {
  name = "%[1]s"
	site_id = netbox_site.test.id
	status = "reserved"
	width = 19
	u_height = 48
	tenant_id = netbox_tenant.test.id
	location_id = netbox_location.test.id
}
`, testName)
}

func TestAccNetboxDevicePowerFeed_basic(t *testing.T) {
	testSlug := "device_power_feed_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDevicePowerFeedDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePowerFeedFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_power_feed" "test" {
	power_panel_id = netbox_power_panel.test.id
	name = "%[1]s"
	status = "active"
	type = "primary"
	supply = "ac"
	phase = "single-phase"
	voltage = 250
	amperage = 100
  max_percent_utilization = 80

	rack_id = netbox_rack.test.id
	mark_connected = true
	description = "%[1]s_description"
	comments = "%[1]s_comments"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_feed.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "type", "primary"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "supply", "ac"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "phase", "single-phase"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "voltage", "250"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "amperage", "100"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "max_percent_utilization", "80"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "comments", testName+"_comments"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttrPair("netbox_power_feed.test", "power_panel_id", "netbox_power_panel.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_power_feed.test", "rack_id", "netbox_rack.test", "id"),
				),
			},
			{
				Config: testAccNetboxDevicePowerFeedFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_power_feed" "test" {
	power_panel_id = netbox_power_panel.test.id
	name = "%[1]s"
	status = "active"
	type = "primary"
	supply = "ac"
	phase = "single-phase"
	voltage = 250
	amperage = 100
  max_percent_utilization = 80
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_feed.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "type", "primary"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "supply", "ac"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "phase", "single-phase"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "voltage", "250"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "amperage", "100"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "max_percent_utilization", "80"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "mark_connected", "false"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "rack_id", "0"),

					resource.TestCheckResourceAttrPair("netbox_power_feed.test", "power_panel_id", "netbox_power_panel.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_power_feed.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDevicePowerFeedDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	state := testAccProvider.Meta().(*providerState)
	api := state.legacyAPI

	// loop through the resources in state, verifying each power feed
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_power_feed" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimPowerFeedsReadParams().WithID(stateID)
		_, err := api.Dcim.DcimPowerFeedsRead(params, nil)

		if err == nil {
			return fmt.Errorf("device_power_feed (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimPowerFeedsReadDefault); ok {
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
	resource.AddTestSweepers("netbox_power_feed", &resource.Sweeper{
		Name:         "netbox_power_feed",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			state := m.(*providerState)
			api := state.legacyAPI
			params := dcim.NewDcimPowerFeedsListParams()
			res, err := api.Dcim.DcimPowerFeedsList(params, nil)
			if err != nil {
				return err
			}
			for _, powerFeed := range res.GetPayload().Results {
				if strings.HasPrefix(*powerFeed.Name, testPrefix) {
					deleteParams := dcim.NewDcimPowerFeedsDeleteParams().WithID(powerFeed.ID)
					_, err := api.Dcim.DcimPowerFeedsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a power_feed")
				}
			}
			return nil
		},
	})
}
