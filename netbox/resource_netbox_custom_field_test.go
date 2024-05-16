package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCustomField_basic(t *testing.T) {
	testSlug := "custom_fields_basic"
	testName := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name = "%s"
  type = "text"
  content_types = ["virtualization.vminterface"]
  weight = 100
  validation_regex = "^.*$"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "type", "text"),
					resource.TestCheckTypeSetElemAttr("netbox_custom_field.test", "content_types.*", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "weight", "100"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "validation_regex", "^.*$"),
				),
			},
		},
	})
}

func TestAccNetboxCustomField_json(t *testing.T) {
	testSlug := "custom_fields_json"
	testName := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name = "%s"
  type = "json"
  content_types = ["virtualization.vminterface"]
  group_name = "mygroup"
  weight = 100
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "type", "json"),
					resource.TestCheckTypeSetElemAttr("netbox_custom_field.test", "content_types.*", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "group_name", "mygroup"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "weight", "100"),
				),
			},
		},
	})
}

func TestAccNetboxCustomField_object(t *testing.T) {
	testSlug := "custom_fields_object"
	testName := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name = "%s"
  type = "object"
  object_type = "ipam.prefix"
  content_types = ["virtualization.vminterface"]
  group_name = "mygroup"
  weight = 100
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "type", "object"),
					resource.TestCheckTypeSetElemAttr("netbox_custom_field.test", "content_types.*", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "group_name", "mygroup"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "weight", "100"),
				),
			},
		},
	})
}

func TestAccNetboxCustomField_integer(t *testing.T) {
	testSlug := "custom_fields_integer"
	testName := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name = "%s"
  type = "integer"
  content_types = ["virtualization.vminterface"]
  group_name = "mygroup"
  weight = 100
  validation_maximum = 1000
  validation_minimum = 10
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "type", "integer"),
					resource.TestCheckTypeSetElemAttr("netbox_custom_field.test", "content_types.*", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "group_name", "mygroup"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "weight", "100"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "validation_maximum", "1000"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "validation_minimum", "10"),
				),
			},
		},
	})
}

func TestAccNetboxCustomField_select(t *testing.T) {
	testSlug := "custom_fields_select"
	testName := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name = "%[1]s"
  extra_choices = [
    ["red", "red"],
    ["blue", "blue"]
  ]
}

resource "netbox_custom_field" "test" {
  name = "%[1]s"
  type = "select"
  content_types = ["virtualization.vminterface"]
  weight = 101
  default = "red"
  choice_set_id = netbox_custom_field_choice_set.test.id
  description = "select field"
  label = "external"
  required = false
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "type", "select"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "default", "red"),
					resource.TestCheckTypeSetElemAttr("netbox_custom_field.test", "content_types.*", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_custom_field.test", "choice_set_id", "netbox_custom_field_choice_set.test", "id"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "weight", "101"),

					resource.TestCheckResourceAttr("netbox_custom_field.test", "default", "red"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "description", "select field"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "label", "external"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "required", "false"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name = "%[1]s"
  extra_choices = [
    ["red", "red"],
    ["blue", "blue"]
  ]
}

resource "netbox_custom_field" "test" {
  name = "%[1]s"
  type = "select"
  content_types = ["virtualization.vminterface"]
  weight = 102
  default = "red"
  choice_set_id = netbox_custom_field_choice_set.test.id
  description = "select field"
  label = "external"
  required = true
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "type", "select"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "default", "red"),
					resource.TestCheckTypeSetElemAttr("netbox_custom_field.test", "content_types.*", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_custom_field.test", "choice_set_id", "netbox_custom_field_choice_set.test", "id"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "weight", "102"),

					resource.TestCheckResourceAttr("netbox_custom_field.test", "default", "red"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "description", "select field"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "label", "external"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "required", "true"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name = "%[1]s"
  extra_choices = [
    ["red", "red"],
    ["blue", "blue"]
  ]
}

resource "netbox_custom_field" "test" {
  name = "%[1]s"
  type = "select"
  content_types = ["virtualization.vminterface"]
  weight = 102
  default = "red"
  choice_set_id = netbox_custom_field_choice_set.test.id
  description = "select field"
  label = "external"
  required = false
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "required", "false"),
				),
			},
		},
	})
}
