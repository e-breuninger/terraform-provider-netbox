package netbox

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCustomFieldChoiceSet_basic(t *testing.T) {
	testSlug := "cfields_choiceset"
	testName := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name        = "%s"
  description = "foo"
  extra_choices = [
    ["choice1", "label1"], # label and choice are different
    ["choice2", "choice2"]  # label and choice are the same
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "description", "foo"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.0.0", "choice1"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.0.1", "label1"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.1.0", "choice2"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.1.1", "choice2"),
				),
			},
		},
	})
}

func TestAccNetboxCustomFieldChoiceSet_listlength(t *testing.T) {
	testSlug := "cfields_choiceset_length"
	testName := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name        = "%s"
  description = "foo"
  extra_choices = [
    ["choice1", "label1", "toolong"]
  ]
}`, testName),
				ExpectError: regexp.MustCompile("length of inner lists must be exactly two for custom field choice sets"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name        = "%s"
  description = "foo"
  extra_choices = [
    ["choice1"]
  ]
}`, testName),
				ExpectError: regexp.MustCompile("length of inner lists must be exactly two for custom field choice sets"),
			},
		},
	})
}
