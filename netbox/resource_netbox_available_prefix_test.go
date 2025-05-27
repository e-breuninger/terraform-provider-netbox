package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxAvailablePrefixFullDependencies(testName string, parentPrefix string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_prefix" "parent" {
  prefix = "%[2]s"
  description = "%[1]s"
  status = "container"
  tags = [netbox_tag.test.name]
}
`, testName, parentPrefix)
}

func TestAccNetboxAvailablePrefix_basic(t *testing.T) {
	testParentPrefix := "1.1.0.0/24"
	testPrefixLength := 25
	expectedPrefix := "1.1.0.0/25"
	testSlug := "prefix"
	testDesc := "test prefix"
	testName := testAccGetTestName(testSlug)

	parentResourceName := "netbox_prefix.parent"
	resourceName := "netbox_available_prefix.test"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  description = "%s"
  status = "active"
  tags = [netbox_tag.test.name]
  mark_utilized = true
  is_pool = true
}`, testPrefixLength, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "prefix", expectedPrefix),
					resource.TestCheckResourceAttr(resourceName, "status", "active"),
					resource.TestCheckResourceAttr(resourceName, "description", testDesc),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", testName),
					resource.TestCheckResourceAttr(resourceName, "mark_utilized", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_pool", "true"),
				),
			},
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  description = "%s"
  status = "active"
  tags = [netbox_tag.test.name]
  mark_utilized = false
  is_pool = false
}`, testPrefixLength, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "mark_utilized", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_pool", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					parent, ok := s.RootModule().Resources[parentResourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", parentResourceName)
					}
					resource, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceName)
					}

					return fmt.Sprintf("%s %s %d", parent.Primary.ID, resource.Primary.ID, testPrefixLength), nil
				},
			},
		},
	})
}

func TestAccNetboxAvailablePrefix_multiplePrefixesSerial(t *testing.T) {
	testParentPrefix := "1.1.0.0/24"
	testPrefixLength := 25
	expectedPrefixes := []string{
		"1.1.0.0/25",
		"1.1.0.128/25",
	}
	testSlug := "prefix"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test1" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_available_prefix" "test2" {
  depends_on = [netbox_available_prefix.test1]
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = netbox_available_prefix.test1.prefix_length
  status = "active"
  tags = [netbox_tag.test.name]
}
`, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test1", "prefix", expectedPrefixes[0]),
					resource.TestCheckResourceAttr("netbox_available_prefix.test2", "prefix", expectedPrefixes[1]),
				),
			},
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test1" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_available_prefix" "test2" {
  depends_on = [netbox_available_prefix.test1]
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = netbox_available_prefix.test1.prefix_length
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_available_prefix" "test3" {
  depends_on = [netbox_available_prefix.test2]
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = netbox_available_prefix.test1.prefix_length
  status = "active"
  tags = [netbox_tag.test.name]
}
`, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test1", "prefix", expectedPrefixes[0]),
					resource.TestCheckResourceAttr("netbox_available_prefix.test2", "prefix", expectedPrefixes[1]),
				),
				ExpectError: regexp.MustCompile(".*Insufficient resources are available to satisfy the request.*"),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_available_prefix", &resource.Sweeper{
		Name:         "netbox_available_prefix",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamPrefixesListParams()
			res, err := api.Ipam.IpamPrefixesList(params, nil)
			if err != nil {
				return err
			}
			for _, prefix := range res.GetPayload().Results {
				if len(prefix.Tags) > 0 && (prefix.Tags[0] == &models.NestedTag{Name: strToPtr("acctest"), Slug: strToPtr("acctest")}) {
					deleteParams := ipam.NewIpamPrefixesDeleteParams().WithID(prefix.ID)
					_, err := api.Ipam.IpamPrefixesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a prefix")
				}
			}
			return nil
		},
	})
}
