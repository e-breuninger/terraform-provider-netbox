package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDataSourceDataSource_basic(t *testing.T) {
	testSlug := "data_source_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_data_source" "test" {
  name         = "%[1]s"
  type         = "local"
  source_url   = "file:///opt/netbox/data"
  enabled      = true
  description  = "This is my data source"
  comments     = "Some comments"
  ignore_rules = ".git\n.gitignore"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_data_source" "test" {
  depends_on = []

  name = netbox_data_source.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_data_source.test", "id", "netbox_data_source.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_data_source.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_data_source.test", "type", "local"),
					resource.TestCheckResourceAttr("data.netbox_data_source.test", "source_url", "file:///opt/netbox/data"),
					resource.TestCheckResourceAttr("data.netbox_data_source.test", "enabled", "true"),
					resource.TestCheckResourceAttr("data.netbox_data_source.test", "description", "This is my data source"),
					resource.TestCheckResourceAttr("data.netbox_data_source.test", "comments", "Some comments"),
				),
			},
		},
	})
}
