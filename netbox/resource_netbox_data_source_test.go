package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDataSource_basic(t *testing.T) {
	testSlug := "data_source_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_data_source" "test" {
  name         = "%[1]s"
  type         = "local"
  source_url   = "file:///opt/netbox/data"
  enabled      = true
  description  = "This is my data source"
  comments     = "Some comments"
  ignore_rules = ".git\n.gitignore"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_data_source.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_data_source.test", "type", "local"),
					resource.TestCheckResourceAttr("netbox_data_source.test", "source_url", "file:///opt/netbox/data"),
					resource.TestCheckResourceAttr("netbox_data_source.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_data_source.test", "description", "This is my data source"),
					resource.TestCheckResourceAttr("netbox_data_source.test", "comments", "Some comments"),
				),
			},
			{
				ResourceName:      "netbox_data_source.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_data_source", &resource.Sweeper{
		Name:         "netbox_data_source",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := core.NewCoreDataSourcesListParams()
			res, err := api.Core.CoreDataSourcesList(params, nil)
			if err != nil {
				return err
			}
			for _, dataSource := range res.GetPayload().Results {
				if strings.HasPrefix(*dataSource.Name, testPrefix) {
					deleteParams := core.NewCoreDataSourcesDeleteParams().WithID(dataSource.ID)
					_, err := api.Core.CoreDataSourcesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a data source")
				}
			}
			return nil
		},
	})
}
