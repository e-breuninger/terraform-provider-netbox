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

func TestAccNetboxRegion_basic(t *testing.T) {
	testSlug := "region_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	randomDescription := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
	name = "%s"
	slug = "%s"
	description = "%s"
}
`, testName, randomSlug, randomDescription),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("netbox_region.test", tfjsonpath.New("name"), knownvalue.StringExact(testName)),
					statecheck.ExpectKnownValue("netbox_region.test", tfjsonpath.New("slug"), knownvalue.StringExact(randomSlug)),
					statecheck.ExpectKnownValue("netbox_region.test", tfjsonpath.New("description"), knownvalue.StringExact(randomDescription)),
				},
			},
			{
				ResourceName:      "netbox_region.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
	name = "%s_updated"
	slug = "%s_updated"
	description = "%s_updated"
}
`, testName, randomSlug, randomDescription),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("netbox_region.test", tfjsonpath.New("name"), knownvalue.StringExact(testName+"_updated")),
					statecheck.ExpectKnownValue("netbox_region.test", tfjsonpath.New("slug"), knownvalue.StringExact(randomSlug+"_updated")),
					statecheck.ExpectKnownValue("netbox_region.test", tfjsonpath.New("description"), knownvalue.StringExact(randomDescription+"_updated")),
				},
			},
		},
	})
}

func TestAccNetboxRegion_parent(t *testing.T) {
	testSlug := "region_parent"
	testNameParent := testAccGetTestName(testSlug)
	testSlugParent := testAccGetTestName(testSlug)
	testNameChild := testAccGetTestName(testSlug)
	testSlugChild := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "parent" {
	name = "%s"
	slug = "%s"
}
resource "netbox_region" "child" {
	name = "%s"
	slug = "%s"
	parent = netbox_region.parent.id
}
`, testNameParent, testSlugParent, testNameChild, testSlugChild),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.CompareValuePairs("netbox_region.parent", tfjsonpath.New("id"), "netbox_region.child", tfjsonpath.New("parent"), compare.ValuesSame()),
				},
			},
		},
	})
}

func TestAccNetboxRegionTag(t *testing.T) {
	testPrefix := "region_tag"
	testNameRegion := testAccGetTestName(testPrefix)
	testSlugRegion := testAccGetTestName(testPrefix)
	testNameTag := testAccGetTestName(testPrefix)
	testSlugTag := testAccGetTestName(testPrefix)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
  color_hex = "112233"
  description = "This is a test"
}
resource "netbox_region" "test" {
	name = "%s"
	slug = "%s"
	tags = [netbox_tag.test.id]
}
`, testNameTag, testSlugTag, testNameRegion, testSlugRegion),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.CompareValueCollection("netbox_region.test", []tfjsonpath.Path{
						tfjsonpath.New("tags"),
					}, "netbox_tag.test", tfjsonpath.New("id"), compare.ValuesSame()),
				},
			},
		},
	})
}

func TestAccNetboxRegionCustomField(t *testing.T) {
	testPrefix := "region_customfield"
	customFieldName := testAccGetTestCustomFieldName(testPrefix)
	regionName := testAccGetTestName(testPrefix)
	regionSlug := testAccGetTestName(testPrefix)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name = "%s"
  type = "text"
  content_types = ["dcim.region"]
}
resource "netbox_region" "test" {
	name = "%s"
	slug = "%s"
	custom_fields = {
	"${netbox_custom_field.test.name}" : "testcustomfield"
	}
}
`, customFieldName, regionName, regionSlug),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"netbox_region.test",
						tfjsonpath.New("custom_fields").AtMapKey(customFieldName), knownvalue.StringExact("testcustomfield")),
				},
			},
		},
	})
}
