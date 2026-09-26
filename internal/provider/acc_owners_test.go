//go:build acctest

// Object ownership (users.owner and users.ownergroup), added by NetBox 4.6.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxOwnerGroup_basic(t *testing.T) {
	testName := testAccGetTestName("ownergrp")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_owner_group" "test" {
  name        = "%[1]s"
  description = "Acceptance test owner group."
}
data "netbox_owner_group" "test" {
  name = netbox_owner_group.test.name
}
data "netbox_owner_groups" "test" {
  filters = [
    { name = "name", value = "%[1]s" },
  ]
  depends_on = [netbox_owner_group.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_owner_group.test", "name", testName),
					resource.TestCheckResourceAttrPair("data.netbox_owner_group.test", "id", "netbox_owner_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_owner_groups.test", "owner_groups.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_owner_group" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_owner_group.test", "description"),
			},
			{
				ResourceName:      "netbox_owner_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxOwner_basic(t *testing.T) {
	testName := testAccGetTestName("owner")
	deps := fmt.Sprintf(`
resource "netbox_owner_group" "test" {
  name = "%[1]s"
}
resource "netbox_group" "test" {
  name = "%[1]s"
}
resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Acceptance-test-passw0rd"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_owner" "test" {
  name           = "%[1]s"
  owner_group_id = netbox_owner_group.test.id
  user_ids       = [netbox_user.test.id]
  user_group_ids = [netbox_group.test.id]
  description    = "Acceptance test owner."
}
data "netbox_owner" "test" {
  name = netbox_owner.test.name
}
data "netbox_owners" "test" {
  filters = [
    { name = "group_id", value = netbox_owner_group.test.id },
  ]
  depends_on     = [netbox_owner.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_owner.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_owner.test", "user_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_owner.test", "user_group_ids.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_owner.test", "owner_group_id", "netbox_owner_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_owner.test", "id", "netbox_owner.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_owners.test", "owners.#", "1"),
				),
			},
			{
				// Shrink: the member lists and the description clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_owner" "test" {
  name           = "%[1]s"
  owner_group_id = netbox_owner_group.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_owner.test", "user_ids"),
					resource.TestCheckNoResourceAttr("netbox_owner.test", "user_group_ids"),
					resource.TestCheckNoResourceAttr("netbox_owner.test", "description"),
				),
			},
			{
				ResourceName:      "netbox_owner.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// owner_id is the object-ownership reference NetBox 4.6 puts on nearly every object. netbox_site
// has always had a Writable request model; netbox_device_role is one of the sixteen that only got
// one when go-netbox 45aad9ef added the missing twins, so it is the one to watch for a regression
// there. Both set and clear the owner.
func TestAccNetboxOwnerID_basic(t *testing.T) {
	testName := testAccGetTestName("ownerid")
	deps := fmt.Sprintf(`
resource "netbox_owner_group" "test" {
  name = "%[1]s"
}
resource "netbox_owner" "test" {
  name           = "%[1]s"
  owner_group_id = netbox_owner_group.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_site" "test" {
  name     = "%[1]s"
  owner_id = netbox_owner.test.id
}
resource "netbox_device_role" "test" {
  name     = "%[1]s"
  owner_id = netbox_owner.test.id
}
data "netbox_sites" "by_owner" {
  filters = [
    { name = "owner_id", value = netbox_owner.test.id },
  ]
  depends_on = [netbox_site.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_site.test", "owner_id", "netbox_owner.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_role.test", "owner_id", "netbox_owner.test", "id"),
					// The owner is this test's alone, so the lookup finds exactly our site.
					resource.TestCheckResourceAttrPair("data.netbox_sites.by_owner", "sites.0.id", "netbox_site.test", "id"),
				),
			},
			{
				// Unsetting clears the owner in NetBox on both request shapes.
				Config: deps + fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_site.test", "owner_id"),
					resource.TestCheckNoResourceAttr("netbox_device_role.test", "owner_id"),
				),
			},
			{
				ResourceName:      "netbox_device_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_owner",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Users.UsersOwnersList(users.NewUsersOwnersListParams(), nil)
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
			_, err := client.Users.UsersOwnersDestroy(users.NewUsersOwnersDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_owner_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Users.UsersOwnerGroupsList(users.NewUsersOwnerGroupsListParams(), nil)
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
			_, err := client.Users.UsersOwnerGroupsDestroy(users.NewUsersOwnerGroupsDestroyParams().WithID(id), nil)
			return err
		})
}
