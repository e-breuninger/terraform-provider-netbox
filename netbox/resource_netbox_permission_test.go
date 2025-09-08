package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPermission_basic(t *testing.T) {
	testName := testAccGetTestName("permission_basic")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_permission" "test" {
  name         = "%s"
  description  = "Test permission"
  enabled      = true
  object_types = ["dcim.device"]
  actions      = ["view", "change"]
  constraints  = "{}"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_permission.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_permission.test", "description", "Test permission"),
					resource.TestCheckResourceAttr("netbox_permission.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_permission.test", "object_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test", "actions.#", "2"),
				),
			},
			{
				ResourceName:      "netbox_permission.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPermission_minimal(t *testing.T) {
	testName := testAccGetTestName("permission_minimal")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_permission" "test" {
  name         = "%s"
  object_types = ["dcim.device", "dcim.interface"]
  actions      = ["view"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_permission.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_permission.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_permission.test", "object_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_permission.test", "actions.#", "1"),
				),
			},
		},
	})
}

func TestAccNetboxPermission_withUsersAndGroups(t *testing.T) {
	testName := testAccGetTestName("permission_users_groups")
	testUser := testAccGetTestName("user")
	testGroup := testAccGetTestName("group")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_user" "test" {
  username = "%s"
}

resource "netbox_group" "test" {
  name = "%s"
}

resource "netbox_permission" "test" {
  name         = "%s"
  object_types = ["dcim.device"]
  actions      = ["view", "add"]
  users        = [netbox_user.test.id]
  groups       = [netbox_group.test.id]
}`, testUser, testGroup, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_permission.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_permission.test", "users.#", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test", "groups.#", "1"),
				),
			},
		},
	})
}

func TestAccNetboxPermission_disabled(t *testing.T) {
	testName := testAccGetTestName("permission_disabled")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_permission" "test" {
  name         = "%s"
  enabled      = false
  object_types = ["dcim.device"]
  actions      = ["view"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_permission.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_permission.test", "enabled", "false"),
				),
			},
		},
	})
}
