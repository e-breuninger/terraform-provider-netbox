package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxContactTagDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}

resource "netbox_tag" "test_b" {
  name = "%[1]sb"
}
`, testName)
}

func TestAccNetboxContact_basic(t *testing.T) {
	testSlug := "contact_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact.test", "name", testName),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_contact" "test" {
  name = "%s"
  email = "test@test.com"
  phone = "123-123123"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_contact.test", "email", "test@test.com"),
					resource.TestCheckResourceAttr("netbox_contact.test", "phone", "123-123123"),
				),
			},
			{
				ResourceName:      "netbox_contact.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxContact_tags(t *testing.T) {
	testSlug := "contact_tags"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxContactTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_contact" "test_tags" {
  name = "%[1]s"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact.test_tags", "name", testName),
					resource.TestCheckResourceAttr("netbox_contact.test_tags", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_contact.test_tags", "tags.0", testName+"a"),
				),
			},
			{
				Config: testAccNetboxContactTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_contact" "test_tags" {
  name = "%[1]s"
  tags = ["%[1]sa", "%[1]sb"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact.test_tags", "tags.#", "2"),
					resource.TestCheckResourceAttr("netbox_contact.test_tags", "tags.0", testName+"a"),
					resource.TestCheckResourceAttr("netbox_contact.test_tags", "tags.1", testName+"b"),
				),
			},
			{
				Config: testAccNetboxContactTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_contact" "test_tags" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact.test_tags", "tags.#", "0"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_contact", &resource.Sweeper{
		Name:         "netbox_contact",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := tenancy.NewTenancyContactsListParams()
			res, err := api.Tenancy.TenancyContactsList(params, nil)
			if err != nil {
				return err
			}
			for _, contact := range res.GetPayload().Results {
				if strings.HasPrefix(*contact.Name, testPrefix) {
					deleteParams := tenancy.NewTenancyContactsDeleteParams().WithID(contact.ID)
					_, err := api.Tenancy.TenancyContactsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a contact")
				}
			}
			return nil
		},
	})
}
