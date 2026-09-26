//go:build acctest

package provider_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// The netbox_supported_versions data source serves the spec's backend.options.supported_versions.
// Besides the shape, the test checks that the NetBox the suite runs against is in the list — on
// purpose a tripwire: running the suite against a NetBox the spec does not list fails here until
// the list is updated, so the list stays truthful.
func TestAccNetboxSupportedVersionsDataSource_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: `data "netbox_supported_versions" "test" {}`,
				Check:  testAccCheckSupportedVersions("data.netbox_supported_versions.test"),
			},
		},
	})
}

var semanticVersion = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

// testAccCheckSupportedVersions checks that versions holds at least one MAJOR.MINOR.PATCH entry
// and that the running NetBox's version is one of them.
func testAccCheckSupportedVersions(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resource, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("%s not found in state", name)
		}
		attrs := resource.Primary.Attributes
		n, err := strconv.Atoi(attrs["versions.#"])
		if err != nil {
			return fmt.Errorf("%s: versions.# is %q, not a count", name, attrs["versions.#"])
		}
		if n == 0 {
			return fmt.Errorf("%s: versions is empty", name)
		}
		versions := make([]string, n)
		for i := range versions {
			v := attrs["versions."+strconv.Itoa(i)]
			if !semanticVersion.MatchString(v) {
				return fmt.Errorf("%s: versions[%d] = %q is not a MAJOR.MINOR.PATCH version", name, i, v)
			}
			versions[i] = v
		}

		client, err := testAccClient()
		if err != nil {
			return fmt.Errorf("error getting client: %w", err)
		}
		_, running, err := netboxapi.Version(context.Background(), client)
		if err != nil {
			return fmt.Errorf("error reading the NetBox version: %w", err)
		}
		for _, version := range versions {
			if version == running {
				return nil
			}
		}
		return fmt.Errorf("%s: the NetBox under test runs %s, which versions %v does not list; add it to the spec's supported_versions", name, running, versions)
	}
}
