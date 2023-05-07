package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxContactDataSource_basic(t *testing.T) {

	testSlug := "cntct_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact" "test" {
  name = "%[1]s"
}

data "netbox_contact" "by_name" {
  depends_on = [netbox_contact.test]
  name = "%[1]s"
}

data "netbox_contact" "by_slug" {
  depends_on = [netbox_contact.test]
  slug = "%[1]s"
}

data "netbox_contact" "by_description" {
  depends_on = [netbox_contact.test]
  name = "%[1]s"
  description = "%[1]s"
}

data "netbox_contact" "by_both" {
  depends_on = [netbox_contact.test]
  name = "%[1]s"
  slug = "%[1]s"
  description = "%[1]s"
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_contact.by_name", "id", "netbox_contact.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_contact.by_slug", "id", "netbox_contact.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_contact.by_description", "id", "netbox_contact.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_contact.by_both", "id", "netbox_contact.test", "id"),
				),
			},
		},
	})
}
