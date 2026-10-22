//go:build acctest

// The flat dcim objects NetBox 4.6 added: rack groups, cable bundles and module type profiles.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxRackGroup_basic(t *testing.T) {
	testName := testAccGetTestName("rackgrp")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rack_group" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "Acceptance test rack group."
  comments    = "Created by acceptance test."
}
data "netbox_rack_group" "test" {
  name = netbox_rack_group.test.name
}
data "netbox_rack_groups" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_rack_group.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_rack_group.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttrPair("data.netbox_rack_group.test", "id", "netbox_rack_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_rack_groups.test", "rack_groups.#", "1"),
				),
			},
			{
				// Shrink: every optional attribute clears.
				Config: fmt.Sprintf(`
resource "netbox_rack_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_rack_group.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_rack_group.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_rack_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxCableBundle_basic(t *testing.T) {
	testName := testAccGetTestName("cablebundle")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cable_bundle" "test" {
  name        = "%[1]s"
  description = "Acceptance test cable bundle."
  comments    = "Created by acceptance test."
}
data "netbox_cable_bundle" "test" {
  name = netbox_cable_bundle.test.name
}
data "netbox_cable_bundles" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_cable_bundle.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cable_bundle.test", "name", testName),
					resource.TestCheckResourceAttrPair("data.netbox_cable_bundle.test", "id", "netbox_cable_bundle.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_cable_bundles.test", "cable_bundles.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_cable_bundle" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_cable_bundle.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_cable_bundle.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_cable_bundle.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxModuleTypeProfile_basic(t *testing.T) {
	testName := testAccGetTestName("mtprofile")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_module_type_profile" "test" {
  name        = "%[1]s"
  description = "Acceptance test module type profile."
  comments    = "Created by acceptance test."
  schema = jsonencode({
    properties = {
      capacity = { type = "integer" }
    }
  })
}
data "netbox_module_type_profile" "test" {
  name = netbox_module_type_profile.test.name
}
data "netbox_module_type_profiles" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_module_type_profile.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "name", testName),
					resource.TestCheckResourceAttrSet("netbox_module_type_profile.test", "schema"),
					resource.TestCheckResourceAttrPair("data.netbox_module_type_profile.test", "id", "netbox_module_type_profile.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_module_type_profiles.test", "module_type_profiles.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_module_type_profile" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_module_type_profile.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_module_type_profile.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_module_type_profile.test", "schema"),
				),
			},
			{
				ResourceName:      "netbox_module_type_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_rack_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimRackGroupsList(dcim.NewDcimRackGroupsListParams(), nil)
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
			_, err := client.Dcim.DcimRackGroupsDestroy(dcim.NewDcimRackGroupsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_cable_bundle",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimCableBundlesList(dcim.NewDcimCableBundlesListParams(), nil)
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
			_, err := client.Dcim.DcimCableBundlesDestroy(dcim.NewDcimCableBundlesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_module_type_profile",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimModuleTypeProfilesList(dcim.NewDcimModuleTypeProfilesListParams(), nil)
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
			_, err := client.Dcim.DcimModuleTypeProfilesDestroy(dcim.NewDcimModuleTypeProfilesDestroyParams().WithID(id), nil)
			return err
		})
}
