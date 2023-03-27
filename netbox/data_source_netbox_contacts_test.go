package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxContactsDataSource_basic(t *testing.T) {

	testSlug := "tnt_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact" "test_list_0" {
  name = "%[1]s_0"
}
resource "netbox_contact" "test_list_1" {
  name = "%[1]s_1"
}
data "netbox_contacts" "test" {
  depends_on = [netbox_contact.test_list_0, netbox_contact.test_list_1]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_contacts.test", "contacts.0.name", "netbox_contact.test_list_0", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_contacts.test", "contacts.1.name", "netbox_contact.test_list_1", "name"),
				),
			},
		},
	})
}

func testAccNetboxContactsDataSource_manyContacts(testName string) string {
	return fmt.Sprintf(`resource "netbox_contact" "test" {
  count = 51
  name = "%s-${count.index}"
}
`, testName)
}

func TestAccNetboxContactsDataSource_many(t *testing.T) {

	testSlug := "tnt_ds_many"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxContactsDataSource_manyContacts(testName) + `data "netbox_contacts" "test" {
  depends_on = [netbox_contact.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_contacts.test", "contacts.#", "51"),
				),
			},
			{
				Config: testAccNetboxContactsDataSource_manyContacts(testName) + `data "netbox_contacts" "test" {
  depends_on = [netbox_contact.test]
  limit = 2
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_contacts.test", "contacts.#", "2"),
				),
			},
		},
	})
}

func TestAccNetboxContactsDataSource_filter(t *testing.T) {

	testSlug := "tnt_ds_filter"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact" "test_list_0" {
  name = "%[1]s_0"
}
resource "netbox_contact" "test_list_1" {
  name = "%[1]s_1"
}
data "netbox_contacts" "test" {
  depends_on = [netbox_contact.test_list_0, netbox_contact.test_list_1]

  filter {
    name = "name"
    value = "%[1]s_0"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_contacts.test", "contacts.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_contacts.test", "contacts.0.name", "netbox_contact.test_list_0", "name"),
				),
			},
		},
	})
}

func TestAccNetboxContactsDataSource_contactgroups(t *testing.T) {

	testSlug := "tnt_ds_contact_group_filter"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact_group" "group_0" {
  name = "group_%[1]s_1"
}

resource "netbox_contact" "contact_0" {
  name = "contact_%[1]s_0"
  group_id = netbox_contact_group.group_0.id
}

data "netbox_contacts" "test" {
  depends_on = [netbox_contact.contact_0, netbox_contact_group.group_0]

  filter {
    name = "name"
    value = "contact_%[1]s_0"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_contacts.test", "contacts.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_contacts.test", "contacts.0.contact_group.0.name", "netbox_contact_group.group_0", "name"),
				),
			},
		},
	})
}
