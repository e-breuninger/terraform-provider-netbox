package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"testing"
)

func TestAccNetboxTagDatasource_basic(t *testing.T) {
	testSlug := "tag_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Read test
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
  color = "112233"
  description = "This is a test"
}
data "netbox_tag" "test" {
	name = netbox_tag.test.name
}
`, testName, randomSlug),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.netbox_tag.test", tfjsonpath.New("name"), knownvalue.StringExact(testName)),
					statecheck.ExpectKnownValue("data.netbox_tag.test", tfjsonpath.New("slug"), knownvalue.StringExact(randomSlug)),
					statecheck.ExpectKnownValue("data.netbox_tag.test", tfjsonpath.New("color"), knownvalue.StringExact("112233")),
					statecheck.ExpectKnownValue("data.netbox_tag.test", tfjsonpath.New("description"), knownvalue.StringExact("This is a test")),
				},
			},
		},
	})
}
