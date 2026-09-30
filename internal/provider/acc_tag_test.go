//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxTag_basic(t *testing.T) {
	testSlug := "tag_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
  color_hex = "112233"
  description = "This is a test"
  object_types = ["dcim.device"]
  weight      = 50
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tag.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tag.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_tag.test", "color_hex", "112233"),
					resource.TestCheckResourceAttr("netbox_tag.test", "description", "This is a test"),
					resource.TestCheckResourceAttr("netbox_tag.test", "weight", "50"),
				),
			},
			{
				ResourceName:      "netbox_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: description clears; slug is computed and keeps its value, color_hex
				// falls back to its default.
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_tag.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_tag.test", "object_types"),
					resource.TestCheckResourceAttr("netbox_tag.test", "weight", "50"),
					resource.TestCheckResourceAttr("netbox_tag.test", "color_hex", "9e9e9e"),
				),
			},
		},
	})
}

func TestAccNetboxTag_defaultSlug(t *testing.T) {
	testSlug := "tag_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tag.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tag.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttrSet("netbox_tag.test", "color_hex"),
				),
			},
		},
	})
}

func TestAccNetboxTagsDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("tags_ds")
	setUp := fmt.Sprintf(`
resource "netbox_tag" "test_1" {
  name = "%[1]s-1"
}

resource "netbox_tag" "test_2" {
  name = "%[1]s-2"
}

resource "netbox_tag" "test_3" {
  name = "%[1]s-3"
  slug = "%[2]s-weird"
}`, testName, getSlug(testName))
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check:  resource.TestCheckResourceAttr("netbox_tag.test_1", "slug", getSlug(testName+"-1")),
			},
			{
				Config: setUp + fmt.Sprintf(`
data "netbox_tags" "test" {
  filters = [
    { name = "name", value = "%s-1" },
  ]
  depends_on = [netbox_tag.test_1]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tags.test", "tags.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_tags.test", "tags.0.id", "netbox_tag.test_1", "id"),
				),
			},
			{
				Config: setUp + fmt.Sprintf(`
data "netbox_tags" "test" {
  filters = [
    { name = "slug", value = "%s-weird" },
  ]
  depends_on = [netbox_tag.test_3]
}
data "netbox_tag" "one" {
  slug = "%[1]s-weird"
  depends_on = [netbox_tag.test_3]
}`, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tags.test", "tags.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_tags.test", "tags.0.id", "netbox_tag.test_3", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tag.one", "id", "netbox_tag.test_3", "id"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_tag",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasTagsList(extras.NewExtrasTagsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasTagsDestroy(extras.NewExtrasTagsDestroyParams().WithID(id), nil)
			return err
		})
}
