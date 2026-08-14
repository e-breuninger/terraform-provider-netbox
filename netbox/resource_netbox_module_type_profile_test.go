package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxModuleTypeProfile_basic(t *testing.T) {
	testSlug := "module_type_profile_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_module_type_profile" "test" {
  name        = "%[1]s"
  description = "This is my module type profile"
  comments    = "Module type profile comments"
  schema      = jsonencode({
    type = "object"
    properties = {
      port_count = {
        type = "integer"
      }
    }
  })
  tags = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "description", "This is my module type profile"),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "comments", "Module type profile comments"),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_module_type_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_module_type_profile", &resource.Sweeper{
		Name:         "netbox_module_type_profile",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimModuleTypeProfilesListParams()
			res, err := api.Dcim.DcimModuleTypeProfilesList(params, nil)
			if err != nil {
				return err
			}
			for _, moduleTypeProfile := range res.GetPayload().Results {
				if strings.HasPrefix(*moduleTypeProfile.Name, testPrefix) {
					deleteParams := dcim.NewDcimModuleTypeProfilesDeleteParams().WithID(moduleTypeProfile.ID)
					_, err := api.Dcim.DcimModuleTypeProfilesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a module_type_profile")
				}
			}
			return nil
		},
	})
}
