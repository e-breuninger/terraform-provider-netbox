//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxGroup_basic(t *testing.T) {
	testName := testAccGetTestName("group")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_group" "test" {
  name        = "%[1]s"
  description = "Acceptance test group."
}
data "netbox_group" "test" {
  name = netbox_group.test.name
}
data "netbox_groups" "test" {
  filters = [
    { name = "name", value = netbox_group.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_group.test", "description", "Acceptance test group."),
					resource.TestCheckResourceAttr("netbox_group.test", "user_count", "0"),
					resource.TestCheckResourceAttrPair("data.netbox_group.test", "id", "netbox_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_groups.test", "groups.#", "1"),
				),
			},
			{
				// Drop description: it must clear.
				Config: fmt.Sprintf(`
resource "netbox_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_group.test", "description"),
			},
			{
				ResourceName:      "netbox_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxUser_basic(t *testing.T) {
	testName := testAccGetTestName("user")
	deps := fmt.Sprintf(`
resource "netbox_group" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_user" "test" {
  username   = "%[1]s"
  password   = "Acceptance-Test-Pw-9271"
  first_name = "Acceptance"
  last_name  = "Test"
  email      = "%[1]s@example.com"
  is_active  = true
  group_ids  = [netbox_group.test.id]
}
data "netbox_user" "test" {
  username = netbox_user.test.username
}
data "netbox_users" "test" {
  filters = [
    { name = "username", value = netbox_user.test.username },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_user.test", "username", testName),
					// NetBox never returns the password, so state keeps what was configured.
					resource.TestCheckResourceAttr("netbox_user.test", "password", "Acceptance-Test-Pw-9271"),
					resource.TestCheckResourceAttr("netbox_user.test", "first_name", "Acceptance"),
					resource.TestCheckResourceAttr("netbox_user.test", "email", testName+"@example.com"),
					resource.TestCheckResourceAttr("netbox_user.test", "is_active", "true"),
					resource.TestCheckResourceAttr("netbox_user.test", "group_ids.#", "1"),
					resource.TestCheckResourceAttrSet("netbox_user.test", "date_joined"),
					resource.TestCheckNoResourceAttr("netbox_user.test", "last_login"),
					resource.TestCheckResourceAttrPair("data.netbox_user.test", "id", "netbox_user.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_users.test", "users.#", "1"),
				),
			},
			{
				// Shrink: everything optional clears (is_active is computed and keeps NetBox's value).
				Config: deps + fmt.Sprintf(`
resource "netbox_user" "test" {
  username = "%s"
  password = "Acceptance-Test-Pw-9271"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_user.test", "is_active", "true"),
					resource.TestCheckNoResourceAttr("netbox_user.test", "first_name"),
					resource.TestCheckNoResourceAttr("netbox_user.test", "last_name"),
					resource.TestCheckNoResourceAttr("netbox_user.test", "email"),
					resource.TestCheckNoResourceAttr("netbox_user.test", "group_ids"),
				),
			},
			{
				ResourceName:      "netbox_user.test",
				ImportState:       true,
				ImportStateVerify: true,
				// An imported user has no password: NetBox never returns it.
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

func TestAccNetboxToken_basic(t *testing.T) {
	testName := testAccGetTestName("token")
	deps := fmt.Sprintf(`
resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Acceptance-Test-Pw-9271"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_token" "test" {
  user_id       = netbox_user.test.id
  description   = "Acceptance test token."
  write_enabled = false
  enabled       = false
  expires       = "2030-01-01T00:00:00.000Z"
}
resource "netbox_token" "chosen" {
  user_id = netbox_user.test.id
  token   = "0123456789abcdef0123456789abcdef01234567"
}
data "netbox_token" "test" {
  id = netbox_token.test.id
}
data "netbox_tokens" "test" {
  depends_on = [netbox_token.test, netbox_token.chosen]
  filters = [
    { name = "user_id", value = netbox_user.test.id },
  ]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_token.test", "user_id", "netbox_user.test", "id"),
					resource.TestCheckResourceAttr("netbox_token.test", "description", "Acceptance test token."),
					resource.TestCheckResourceAttr("netbox_token.test", "write_enabled", "false"),
					resource.TestCheckResourceAttr("netbox_token.test", "expires", "2030-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr("netbox_token.test", "enabled", "false"),
					// NetBox generates the key (the public prefix) and, when unset, the secret; the secret
					// is in the create response only and stays in state from there.
					resource.TestCheckResourceAttrSet("netbox_token.test", "key"),
					resource.TestCheckResourceAttrWith("netbox_token.test", "token", func(v string) error {
						if len(v) != 40 {
							return fmt.Errorf("generated token %q is not 40 characters", v)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("netbox_token.test", "version", "2"),
					resource.TestCheckResourceAttrSet("netbox_token.test", "pepper_id"),
					resource.TestCheckResourceAttrSet("netbox_token.test", "created"),
					resource.TestCheckResourceAttr("netbox_token.chosen", "token", "0123456789abcdef0123456789abcdef01234567"),
					resource.TestCheckResourceAttr("netbox_token.chosen", "enabled", "true"),
					resource.TestCheckResourceAttrPair("data.netbox_token.test", "id", "netbox_token.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_tokens.test", "tokens.#", "2"),
				),
			},
			{
				// The same config again: the secret survives the refresh that reads it back as null, and
				// the update of the other token (description) does not resend it.
				Config: deps + `
resource "netbox_token" "test" {
  user_id       = netbox_user.test.id
  description   = "Acceptance test token."
  write_enabled = false
  enabled       = false
  expires       = "2030-01-01T00:00:00.000Z"
}
resource "netbox_token" "chosen" {
  user_id     = netbox_user.test.id
  token       = "0123456789abcdef0123456789abcdef01234567"
  description = "relabelled"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("netbox_token.test", "token"),
					resource.TestCheckResourceAttr("netbox_token.chosen", "token", "0123456789abcdef0123456789abcdef01234567"),
					resource.TestCheckResourceAttr("netbox_token.chosen", "description", "relabelled"),
				),
			},
			{
				// Shrink: description clears. write_enabled is computed, so it keeps the value it
				// already has, and expires is crossed as raw, which NetBox never clears.
				Config: deps + `
resource "netbox_token" "test" {
  user_id = netbox_user.test.id
}
resource "netbox_token" "chosen" {
  user_id = netbox_user.test.id
  token   = "0123456789abcdef0123456789abcdef01234567"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_token.test", "write_enabled", "false"),
					resource.TestCheckResourceAttr("netbox_token.test", "expires", "2030-01-01T00:00:00.000Z"),
					resource.TestCheckNoResourceAttr("netbox_token.test", "description"),
					// enabled is optional+computed and keeps its value; the secret stays too.
					resource.TestCheckResourceAttr("netbox_token.test", "enabled", "false"),
					resource.TestCheckResourceAttrSet("netbox_token.test", "token"),
				),
			},
			{
				ResourceName:      "netbox_token.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The secret is in the create response only; an imported token has none.
				ImportStateVerifyIgnore: []string{"token"},
			},
		},
	})
}

func TestAccNetboxToken_v1(t *testing.T) {
	testName := testAccGetTestName("token_v1")
	deps := fmt.Sprintf(`
resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Acceptance-Test-Pw-9271"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// version 1 is the legacy plaintext format: NetBox generates (or accepts) a
				// 40-character secret and stores it recoverably; key and pepper_id stay null.
				Config: deps + `
resource "netbox_token" "test" {
  user_id = netbox_user.test.id
  version = 1
}
resource "netbox_token" "chosen" {
  user_id = netbox_user.test.id
  version = 1
  token   = "76543210fedcba9876543210fedcba9876543210"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_token.test", "version", "1"),
					resource.TestCheckNoResourceAttr("netbox_token.test", "key"),
					resource.TestCheckNoResourceAttr("netbox_token.test", "pepper_id"),
					resource.TestCheckResourceAttrWith("netbox_token.test", "token", func(v string) error {
						if len(v) != 40 {
							return fmt.Errorf("generated token %q is not 40 characters", v)
						}
						return nil
					}),
					resource.TestCheckResourceAttr("netbox_token.chosen", "version", "1"),
					resource.TestCheckResourceAttr("netbox_token.chosen", "token", "76543210fedcba9876543210fedcba9876543210"),
				),
			},
			{
				// version is force_new + create_only: NetBox rejects a version change in place
				// ("enforce_version_dependent_fields"), so this apply succeeding proves the token
				// was replaced with a v2 one.
				Config: deps + `
resource "netbox_token" "test" {
  user_id = netbox_user.test.id
  version = 2
}
resource "netbox_token" "chosen" {
  user_id = netbox_user.test.id
  version = 1
  token   = "76543210fedcba9876543210fedcba9876543210"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_token.test", "version", "2"),
					resource.TestCheckResourceAttrSet("netbox_token.test", "key"),
					resource.TestCheckResourceAttrSet("netbox_token.test", "pepper_id"),
				),
			},
			{
				// An imported v1 token has no secret in state either: even for the recoverable
				// format, NetBox returns the plaintext in the create response only.
				ResourceName:            "netbox_token.chosen",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"token"},
			},
		},
	})
}

func TestAccNetboxPermission_basic(t *testing.T) {
	testName := testAccGetTestName("perm")
	deps := fmt.Sprintf(`
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
resource "netbox_permission" "test" {
  name         = "%[1]s"
  description  = "Acceptance test permission."
  enabled      = true
  object_types = ["dcim.device", "dcim.site"]
  actions      = ["view", "add"]
  group_ids    = [netbox_group.test.id]
  user_ids     = [netbox_user.test.id]
  constraints  = jsonencode({ status = "active" })
}
data "netbox_permission" "test" {
  name = netbox_permission.test.name
}
data "netbox_permissions" "test" {
  filters = [
    { name = "name", value = netbox_permission.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_permission.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_permission.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_permission.test", "object_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_permission.test", "actions.#", "2"),
					resource.TestCheckResourceAttr("netbox_permission.test", "group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test", "user_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test", "constraints", `{"status":"active"}`),
					resource.TestCheckResourceAttrPair("data.netbox_permission.test", "id", "netbox_permission.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_permissions.test", "permissions.#", "1"),
				),
			},
			{
				// Shrink: everything optional clears (enabled is computed and keeps NetBox's value).
				Config: deps + fmt.Sprintf(`
resource "netbox_permission" "test" {
  name         = "%s"
  object_types = ["dcim.device"]
  actions      = ["view"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_permission.test", "object_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test", "enabled", "true"),
					resource.TestCheckNoResourceAttr("netbox_permission.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_permission.test", "group_ids"),
					resource.TestCheckNoResourceAttr("netbox_permission.test", "user_ids"),
					resource.TestCheckNoResourceAttr("netbox_permission.test", "constraints"),
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
