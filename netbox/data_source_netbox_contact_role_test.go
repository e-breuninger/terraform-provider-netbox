package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxContactRoleDataSource_basic(t *testing.T) {
	testSlug := "cntctrl_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact_role" "test" {
  name = "%[1]s"
}

data "netbox_contact_role" "by_name" {
  name = netbox_contact_role.test.name
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_contact_role.by_name", "id", "netbox_contact_role.test", "id"),
				),
			},
		},
	})
}
