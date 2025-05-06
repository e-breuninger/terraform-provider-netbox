package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxPlatformDataSource_basic(t *testing.T) {
	testSlug := "pltf_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_platform" "test" {
  name = "%[1]s"
}
data "netbox_platform" "test" {
  depends_on = [netbox_platform.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_platform.test", "id", "netbox_platform.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxPlatformDataSource_manufacturer(t *testing.T) {
	testSlug := "pltf_ds_manufacturer"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_platform" "test" {
  name = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}
data "netbox_platform" "test" {
  depends_on = [netbox_platform.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_platform.test", "id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_platform.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
				),
			},
		},
	})
}
