//go:build acctest

package provider_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// testAccV1Token creates a v1 token of the admin user, the only token version provider v5 sends,
// and deletes it when the test ends.
func testAccV1Token(t *testing.T) string {
	t.Helper()
	client, err := testAccClient()
	if err != nil {
		t.Fatal(err)
	}
	userList, err := client.Users.UsersUsersList(users.NewUsersUsersListParams().WithUsername([]string{"admin"}), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(userList.Payload.Results) == 0 {
		t.Skip("no admin user to own a v1 token")
	}
	version := int64(1)
	userID := userList.Payload.Results[0].ID
	created, err := client.Users.UsersTokensCreate(users.NewUsersTokensCreateParams().WithData(&models.WritableToken{
		User:         &userID,
		Version:      &version,
		WriteEnabled: true,
		Description:  "v5 migration acceptance test",
	}), nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, _ = client.Users.UsersTokensDestroy(users.NewUsersTokensDestroyParams().WithID(created.Payload.ID), nil)
	})
	return created.Payload.Token
}

// TestAccNetboxMigrationFromV5_device creates devices with provider v5.8.0 from the registry, then
// applies the same configuration with this provider: each device keeps its id (no replacement) and
// plans no change. The v5 state of a device without local_context_data does not decode against
// this provider's schema (the unset value is ""); one with it decodes but holds v5's other unset
// strings as "".
func TestAccNetboxMigrationFromV5_device(t *testing.T) {
	// The provider under test answers for the address the v5 state is recorded under.
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "e-breuninger")
	testAccPreCheck(t)
	v1Token := testAccV1Token(t)
	testName := testAccGetTestName("v5migr")
	// Shared by both providers: the same syntax in v5 and here.
	common := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device" "context" {
  name               = "%[1]s-context"
  device_type_id     = netbox_device_type.test.id
  role_id            = netbox_device_role.test.id
  site_id            = netbox_site.test.id
  local_context_data = jsonencode({ a = 1 })
}
`, testName)
	v5Config := fmt.Sprintf(`
provider "netbox" {
  server_url = "%[1]s"
  api_token  = "%[2]s"
}
`, os.Getenv("NETBOX_SERVER_URL"), v1Token) + common
	devices := []string{"netbox_device.test", "netbox_device.context"}
	v5IDs := map[string]string{}
	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"netbox": {Source: "registry.terraform.io/e-breuninger/netbox", VersionConstraint: "5.8.0"},
				},
				Config: v5Config,
				Check: func(state *terraform.State) error {
					for _, name := range devices {
						v5IDs[name] = state.RootModule().Resources[name].Primary.ID
					}
					return nil
				},
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   common,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("netbox_device.test", plancheck.ResourceActionNoop),
						plancheck.ExpectResourceAction("netbox_device.context", plancheck.ResourceActionNoop),
					},
				},
				Check: func(state *terraform.State) error {
					for _, name := range devices {
						if id := state.RootModule().Resources[name].Primary.ID; id != v5IDs[name] {
							return fmt.Errorf("%s: id %q, want %q (replaced instead of migrated)", name, id, v5IDs[name])
						}
					}
					return nil
				},
			},
		},
	})
}
