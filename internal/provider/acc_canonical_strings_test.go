//go:build acctest

// Optional strings NetBox stores in a canonical form: an empty string is read
// back as null, and surrounding whitespace is trimmed (DRF CharField). Both
// must converge instead of failing with "Provider produced inconsistent result
// after apply".
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxCanonicalStrings(t *testing.T) {
	testName := testAccGetTestName("canonstr")
	config := func(description, comments string) string {
		return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name        = "%[1]s"
  description = %[2]s
}
resource "netbox_site" "test" {
  name        = "%[1]s"
  description = %[2]s
  comments    = %[3]s
}`, testName, description, comments)
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config(`""`, "<<-EOT\n    trailing newline\n  EOT"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_site.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_site.test", "comments", "trailing newline\n"),
				),
			},
			{
				Config: config(`"set"`, `"  padded  "`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "description", "set"),
					resource.TestCheckResourceAttr("netbox_site.test", "comments", "  padded  "),
				),
			},
			{
				Config: config(`""`, `""`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_site.test", "comments", ""),
				),
			},
		},
	})
}
