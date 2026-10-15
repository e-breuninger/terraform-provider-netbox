//go:build acctest

// State written by provider v5 holds "" for every optional string it did not
// set, and [] for unset tags once it has refreshed them. The first v6 plan over such state must be
// empty, not an update of every object from "" or [] to null.
package provider_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccNetboxLegacyBlanks(t *testing.T) {
	// The provider under test must answer for the address the v5 state is
	// recorded under, registry.terraform.io/e-breuninger/netbox.
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "e-breuninger")
	testAccPreCheck(t)
	testName := testAccGetTestName("blank")
	config := fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
  tags = [netbox_tag.test.slug]
}
resource "netbox_contact" "test" {
  name  = "%[1]s"
  email = "%[2]s@example.com"
}
`, testName, getSlug(testName))
	// v5 stores a site's unset tags as [].
	bare := func(extra string) string {
		return fmt.Sprintf(`
resource "netbox_site" "bare" {
  name = "%[1]s-bare"
  %[2]s
}
`, testName, extra)
	}
	v5 := map[string]resource.ExternalProvider{
		"netbox": {Source: "registry.terraform.io/e-breuninger/netbox", VersionConstraint: "5.8.0"},
	}
	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: v5,
				Config: fmt.Sprintf(`
provider "netbox" {
  server_url = "%s"
  api_token  = "%s"
}
`, os.Getenv("NETBOX_SERVER_URL"), os.Getenv("NETBOX_API_TOKEN")) + config + bare(""),
			},
			{ // v5 writes the unset tags as [] when it refreshes
				ExternalProviders:  v5,
				RefreshState:       true,
				ExpectNonEmptyPlan: true, // v5's own plan against this NetBox
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   config + bare(""),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{ // an explicit "" still round-trips once v6 has written the state
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   config + bare(`description = ""`),
			},
		},
	})
}
