package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIPAMRoleDataSource_basic(t *testing.T) {
	testSlug := "ipamrole_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ipam_role" "test" {
	name = "%[1]s"
}
data "netbox_ipam_role" "test" {
	depends_on = [netbox_ipam_role.test]
	name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ipam_role.test", "id", "netbox_ipam_role.test", "id"),
				),
			},
		},
	})
}
