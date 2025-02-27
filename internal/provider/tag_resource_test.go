package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"testing"
)

// TODO: Destroy in the TestCase setup
func TestAccNetboxTagResource_basic(t *testing.T) {
	testSlug := "tag_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Test creating basic object.
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
  color = "112233"
  description = "This is a test"
}`, testName, randomSlug),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("name"), knownvalue.StringExact(testName)),
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("slug"), knownvalue.StringExact(randomSlug)),
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("color"), knownvalue.StringExact("112233")),
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("description"), knownvalue.StringExact("This is a test")),
				},
			},
			//Test importing
			{
				ResourceName:      "netbox_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			//Test updating
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s_updated"
  slug = "%s_updated"
  color = "112234"
  description = "This is a test_updated"
}`, testName, randomSlug),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("name"), knownvalue.StringExact(testName+"_updated")),
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("slug"), knownvalue.StringExact(randomSlug+"_updated")),
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("color"), knownvalue.StringExact("112234")),
					statecheck.ExpectKnownValue("netbox_tag.test", tfjsonpath.New("description"), knownvalue.StringExact("This is a test_updated")),
				},
			},
		},
	})
}
