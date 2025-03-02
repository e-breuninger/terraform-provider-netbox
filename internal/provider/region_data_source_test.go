package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"testing"
)

func TestAccNetboxRegionDataSource_basic(t *testing.T) {
	testPrefix := "region_datasource"
	regionNameParent := testAccGetTestName(testPrefix)
	regionSlugParent := testAccGetTestName(testPrefix)
	tagName := testAccGetTestName(testPrefix)
	tagSlug := testAccGetTestName(testPrefix)
	customFieldName := testAccGetTestCustomFieldName(testPrefix)
	regionName := testAccGetTestName(testPrefix)
	regionSlug := testAccGetTestName(testPrefix)
	regionDescription := testAccGetTestName(testPrefix)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				//Test every field at once since it's a data source.
				Config: fmt.Sprintf(`
resource "netbox_region" "parent" {
	name = "%s"
	slug = "%s"
}
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
  color = "112233"
  description = "This is a test"
}
resource "netbox_custom_field" "test" {
  name = "%s"
  type = "text"
  content_types = ["dcim.region"]
}
resource "netbox_region" "test" {
	name = "%s"
	slug = "%s"
	description = "%s"
	parent = netbox_region.parent.id
	tags = [netbox_tag.test.id]
	custom_fields = {
	"${netbox_custom_field.test.name}" : "testcustomfield"
	}
}
data "netbox_region" "test" {
	id = netbox_region.test.id
}

`, regionNameParent, regionSlugParent, tagName, tagSlug, customFieldName, regionName, regionSlug, regionDescription),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.netbox_region.test", tfjsonpath.New("name"), knownvalue.StringExact(regionName)),
					statecheck.ExpectKnownValue("data.netbox_region.test", tfjsonpath.New("slug"), knownvalue.StringExact(regionSlug)),
					statecheck.ExpectKnownValue("data.netbox_region.test", tfjsonpath.New("description"), knownvalue.StringExact(regionDescription)),
					statecheck.ExpectKnownValue("data.netbox_region.test", tfjsonpath.New("custom_fields").AtMapKey(customFieldName), knownvalue.StringExact("testcustomfield")),
					statecheck.CompareValueCollection("data.netbox_region.test", []tfjsonpath.Path{
						tfjsonpath.New("tags"),
					}, "netbox_tag.test", tfjsonpath.New("id"), compare.ValuesSame()),
					statecheck.CompareValuePairs("netbox_region.parent", tfjsonpath.New("id"), "data.netbox_region.test", tfjsonpath.New("parent"), compare.ValuesSame()),
				},
			},
		},
	})
}
