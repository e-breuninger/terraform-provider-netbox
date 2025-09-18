package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxTagDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("tag_ds_basic")
	setUp := fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
  description = "Test tag"
}`, testName, testName)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tag.test", "name", testName),
				),
			},
			{
				Config: setUp + fmt.Sprintf(`
data "netbox_tag" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tag.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_tag.test", "slug", testName),
					resource.TestCheckResourceAttr("data.netbox_tag.test", "description", "Test tag"),
				),
			},
		},
	})
}

func TestAccNetboxTagDataSource_noResults(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "netbox_tag" "test" {
		name = "nonexistent-tag-%s"
}`, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)),
				ExpectError: regexp.MustCompile("no tag found matching filter"),
			},
		},
	})
}
