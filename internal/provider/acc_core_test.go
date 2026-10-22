//go:build acctest

// Data sources (core) and config context profiles (extras), added by NetBox 4.6.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/core"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("datasource")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_data_source" "test" {
  name          = "%[1]s"
  type          = "git"
  source_url    = "https://github.com/netbox-community/netbox-example-config.git"
  enabled       = false
  ignore_rules  = "*.md"
  sync_interval = 1440
  parameters    = jsonencode({ branch = "main" })
  description   = "Acceptance test data source."
  comments      = "Created by acceptance test."
}
data "netbox_data_source" "test" {
  name = netbox_data_source.test.name
}
data "netbox_data_sources" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_data_source.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_data_source.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_data_source.test", "type", "git"),
					resource.TestCheckResourceAttr("netbox_data_source.test", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_data_source.test", "sync_interval", "1440"),
					resource.TestCheckResourceAttrSet("netbox_data_source.test", "status"),
					resource.TestCheckResourceAttrPair("data.netbox_data_source.test", "id", "netbox_data_source.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_data_sources.test", "data_sources.#", "1"),
				),
			},
			{
				// Shrink: enabled is optional+computed and keeps NetBox's value, the rest clears.
				Config: fmt.Sprintf(`
resource "netbox_data_source" "test" {
  name       = "%[1]s"
  type       = "git"
  source_url = "https://github.com/netbox-community/netbox-example-config.git"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_data_source.test", "ignore_rules"),
					resource.TestCheckNoResourceAttr("netbox_data_source.test", "sync_interval"),
					resource.TestCheckNoResourceAttr("netbox_data_source.test", "parameters"),
					resource.TestCheckNoResourceAttr("netbox_data_source.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_data_source.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_data_source.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxConfigContextProfile_basic(t *testing.T) {
	testName := testAccGetTestName("ccprofile")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_config_context_profile" "test" {
  name        = "%[1]s"
  description = "Acceptance test config context profile."
  comments    = "Created by acceptance test."
  schema = jsonencode({
    properties = {
      ntp_servers = { type = "array" }
    }
  })
}
data "netbox_config_context_profile" "test" {
  name = netbox_config_context_profile.test.name
}
data "netbox_config_context_profiles" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_config_context_profile.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_context_profile.test", "name", testName),
					resource.TestCheckResourceAttrSet("netbox_config_context_profile.test", "schema"),
					resource.TestCheckResourceAttrPair("data.netbox_config_context_profile.test", "id", "netbox_config_context_profile.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_config_context_profiles.test", "config_context_profiles.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_config_context_profile" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_config_context_profile.test", "schema"),
					resource.TestCheckNoResourceAttr("netbox_config_context_profile.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_config_context_profile.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_config_context_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_data_source",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Core.CoreDataSourcesList(core.NewCoreDataSourcesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Core.CoreDataSourcesDestroy(core.NewCoreDataSourcesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_config_context_profile",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasConfigContextProfilesList(extras.NewExtrasConfigContextProfilesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasConfigContextProfilesDestroy(extras.NewExtrasConfigContextProfilesDestroyParams().WithID(id), nil)
			return err
		})
}
