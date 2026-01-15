package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxTagsSetUp() string {
	return `
resource "netbox_tag" "test_1" {
  name = "Tag1234"
  slug = "tag1234"
}

resource "netbox_tag" "test_2" {
  name = "Tag1235"
  slug = "tag1235"
}

resource "netbox_tag" "test_3" {
  name = "Tag2345"
  slug = "weird"
}`
}

func testAccNetboxTagsByName() string {
	return `
data "netbox_tags" "test" {
  filter {
    name  = "name"
    value = "Tag1234"
  }
}`
}

func testAccNetboxTagsBySlug() string {
	return `
data "netbox_tags" "test" {
  filter {
    name = "slug"
    value = "weird"
  }
}`
}

// func testAccNetboxTagsAll() string {
// 	return `
// data "netbox_tags" "test" {
// }`
// }

func TestAccNetboxTagsDataSource_basic(t *testing.T) {
	setUp := testAccNetboxTagsSetUp()
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tag.test_1", "slug", "tag1234"),
				),
			},
			{
				Config: setUp + testAccNetboxTagsByName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tags.test", "tags.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_tags.test", "tags.0.tag_id", "netbox_tag.test_1", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxTagsBySlug(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tags.test", "tags.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_tags.test", "tags.0.tag_id", "netbox_tag.test_3", "id"),
				),
			},
			// {
			// 	Config: setUp + testAccNetboxTagsAll(),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr("data.netbox_tags.test", "tags.#", "3"),
			// 		resource.TestCheckResourceAttrPair("data.netbox_tags.test", "tags.0.tag_id", "netbox_tag.test_1", "id"),
			// 		resource.TestCheckResourceAttrPair("data.netbox_tags.test", "tags.1.tag_id", "netbox_tag.test_2", "id"),
			// 		resource.TestCheckResourceAttrPair("data.netbox_tags.test", "tags.2.tag_id", "netbox_tag.test_3", "id"),
			// 	),
			// },
		},
	})
}
