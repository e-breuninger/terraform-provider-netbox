package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxUsersDataSource_basic(t *testing.T) {
	testSlug := "users_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_user" "test_0" {
  username = "%[1]s_0"
	password = "supersecurepassword123!"
	active   = true
  staff    = true
}

resource "netbox_user" "test_1" {
  username = "%[1]s_1"
	password = "supersecurepassword123!"
	active   = true
  staff    = false
}

data "netbox_users" "test_0" {
	filter {
		name  = "username"
		value = "%[1]s_0"
	}
  depends_on = [netbox_user.test_0]
}

data "netbox_users" "test_1" {
	filter {
		name  = "username"
		value = "%[1]s_1"
	}
  depends_on = [netbox_user.test_1]
}

data "netbox_users" "test_id_1" {
	filter {
		name  = "id"
		value = netbox_user.test_1.id
	}
  depends_on = [netbox_user.test_1]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					// 3 created by test config plus one default user
					resource.TestCheckResourceAttr("data.netbox_users.test_0", "users.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_users.test_1", "users.#", "1"),

					// All results needs an offset of one because user 0 is admin
					resource.TestCheckResourceAttrPair("data.netbox_users.test_0", "users.0.username", "netbox_user.test_0", "username"),
					resource.TestCheckResourceAttrPair("data.netbox_users.test_1", "users.0.username", "netbox_user.test_1", "username"),

					resource.TestCheckResourceAttrPair("data.netbox_users.test_0", "users.0.active", "netbox_user.test_0", "active"),
					resource.TestCheckResourceAttrPair("data.netbox_users.test_1", "users.0.active", "netbox_user.test_1", "active"),

					resource.TestCheckResourceAttrPair("data.netbox_users.test_0", "users.0.staff", "netbox_user.test_0", "staff"),
					resource.TestCheckResourceAttrPair("data.netbox_users.test_1", "users.0.staff", "netbox_user.test_1", "staff"),

					resource.TestCheckResourceAttr("data.netbox_users.test_id_1", "users.0.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_users.test_id_1", "users.0.id", "netbox_user.test_1", "id"),
				),
			},
		},
	})
}
