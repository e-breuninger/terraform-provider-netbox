package netbox

import (
	"fmt"
	"log"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxFhrpGroup_basic(t *testing.T) {
	testSlug := "fhrp_group_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_tag" "test" {
  name = "%[1]sa"
}

resource "netbox_fhrp_group" "test" {
  protocol    = "other"
  group_id    = "1234"
  auth_type   = "md5"
  auth_key    = "test"
  name        = "test"
  description = "test"
  comments    = "test"

  tags = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "protocol", "other"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "group_id", "1234"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "auth_type", "md5"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "auth_key", "test"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "name", "test"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "description", "test"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "comments", "test"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "tags.0", testName+"a"),
				),
			},
			{
				ResourceName:      "netbox_fhrp_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_fhrp_group", &resource.Sweeper{
		Name:         "netbox_fhrp_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamFhrpGroupsListParams()
			res, err := api.Ipam.IpamFhrpGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, asn := range res.GetPayload().Results {
				deleteParams := ipam.NewIpamFhrpGroupsDeleteParams().WithID(asn.ID)
				_, err := api.Ipam.IpamFhrpGroupsDelete(deleteParams, nil)
				if err != nil {
					return err
				}
				log.Print("[DEBUG] Deleted an asn")
			}
			return nil
		},
	})
}
