//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxPowerPanel_basic(t *testing.T) {
	testName := testAccGetTestName("powerpanel")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_power_panel" "test" {
  name        = "%[1]s"
  site_id     = netbox_site.test.id
  location_id = netbox_location.test.id
  description = "Acceptance test panel."
  comments    = "Created by acceptance test."
}
data "netbox_power_panel" "test" {
  name = netbox_power_panel.test.name
}
data "netbox_power_panels" "by_site" {
  depends_on = [netbox_power_panel.test]
  filters = [
    { name = "site_id", value = netbox_site.test.id },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_panel.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_power_panel.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_power_panel.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_power_panel.test", "id", "netbox_power_panel.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_power_panels.by_site", "power_panels.#", "1"),
				),
			},
			{
				// Drop the optionals: they must clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_power_panel" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_power_panel.test", "location_id"),
					resource.TestCheckNoResourceAttr("netbox_power_panel.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_power_panel.test", "comments"),
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

func TestAccNetboxPowerFeed_basic(t *testing.T) {
	testName := testAccGetTestName("powerfeed")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_rack" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_power_panel" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_power_feed" "test" {
  name                    = "%[1]s"
  power_panel_id          = netbox_power_panel.test.id
  rack_id                 = netbox_rack.test.id
  status                  = "active"
  type                    = "primary"
  supply                  = "ac"
  phase                   = "three-phase"
  voltage                 = 230
  amperage                = 32
  max_utilization_percent = 75
  mark_connected          = true
  description             = "Acceptance test feed."
  comments                = "Created by acceptance test."
}
data "netbox_power_feed" "test" {
  name = netbox_power_feed.test.name
}
data "netbox_power_feeds" "test" {
  filters = [
    { name = "power_panel_id", value = netbox_power_panel.test.id },
  ]
  depends_on     = [netbox_power_feed.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_feed.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "type", "primary"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "supply", "ac"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "phase", "three-phase"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "voltage", "230"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "amperage", "32"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "max_utilization_percent", "75"),
					resource.TestCheckResourceAttrPair("netbox_power_feed.test", "rack_id", "netbox_rack.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_power_feed.test", "id", "netbox_power_feed.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_power_feeds.test", "power_feeds.#", "1"),
				),
			},
			{
				// Shrink: rack, description and comments clear. The electrical attributes are
				// computed, so they keep the values NetBox holds.
				Config: deps + fmt.Sprintf(`
resource "netbox_power_feed" "test" {
  name           = "%[1]s"
  power_panel_id = netbox_power_panel.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_feed.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "voltage", "230"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "amperage", "32"),
					resource.TestCheckResourceAttr("netbox_power_feed.test", "mark_connected", "true"),
					resource.TestCheckNoResourceAttr("netbox_power_feed.test", "rack_id"),
					resource.TestCheckNoResourceAttr("netbox_power_feed.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_power_feed.test", "comments"),
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
