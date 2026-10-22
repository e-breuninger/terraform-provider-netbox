//go:build acctest

package provider_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccNetboxCustomField_basic(t *testing.T) {
	testName := strings.ReplaceAll(testAccGetTestName("cf_basic"), "-", "_")
	// Serial on purpose: this test's custom field carries a default, which NetBox stamps onto every
	// newly created object of its types. A parallel test creating such an object between our create
	// and our delete would then fail its own ImportStateVerify once the field is gone. Serial tests
	// finish before any t.Parallel test starts.
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name         = "%[1]s"
  type         = "integer"
  # Object types no other test creates: the default below would otherwise
  # show up on their objects.
  object_types       = ["dcim.cable", "dcim.powerfeed"]
  label              = "%[1]s label"
  description        = "desc"
  comments           = "Acceptance test comments."
  group_name         = "grp"
  required           = false
  unique             = true
  filter_logic       = "exact"
  weight             = 150
  search_weight      = 500
  ui_visible         = "if-set"
  ui_editable        = "no"
  is_cloneable       = true
  default            = jsonencode(5)
  validation_minimum = 1
  validation_maximum = 100
}
resource "netbox_custom_field" "objref" {
  name                  = "%[1]s_ref"
  type                  = "object"
  object_types          = ["dcim.powerpanel"]
  related_object_type   = "dcim.device"
  related_object_filter = jsonencode({ status = "active" })
}
resource "netbox_custom_field" "txt" {
  name             = "%[1]s_txt"
  type             = "text"
  object_types     = ["dcim.cable"]
  validation_regex = "^[a-z]+$"
}
resource "netbox_custom_field" "json" {
  name              = "%[1]s_json"
  type              = "json"
  object_types      = ["dcim.cable"]
  validation_schema = jsonencode({ type = "object" })
}
data "netbox_custom_field" "test" {
  name = netbox_custom_field.test.name
}
data "netbox_custom_fields" "test" {
  filters = [
    { name = "name", value = netbox_custom_field.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "type", "integer"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "object_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "label", testName+" label"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "unique", "true"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "is_cloneable", "true"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "ui_editable", "no"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "ui_visible", "if-set"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "search_weight", "500"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "required", "false"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "filter_logic", "exact"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "weight", "150"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "default", "5"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "validation_minimum", "1"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "validation_maximum", "100"),
					resource.TestCheckResourceAttr("netbox_custom_field.objref", "related_object_type", "dcim.device"),
					resource.TestCheckResourceAttr("netbox_custom_field.objref", "related_object_filter", `{"status":"active"}`),
					resource.TestCheckResourceAttr("netbox_custom_field.json", "validation_schema", `{"type":"object"}`),
					resource.TestCheckResourceAttr("netbox_custom_field.txt", "validation_regex", "^[a-z]+$"),
					resource.TestCheckResourceAttrPair("data.netbox_custom_field.test", "id", "netbox_custom_field.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_custom_fields.test", "custom_fields.#", "1"),
				),
			},
			{
				// Shrink: the labels and validations clear, the computed flags keep what NetBox holds.
				// objref keeps related_object_type — NetBox requires it while the type is object.
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name         = "%[1]s"
  type         = "integer"
  object_types = ["dcim.cable", "dcim.powerfeed"]
}
resource "netbox_custom_field" "objref" {
  name                = "%[1]s_ref"
  type                = "object"
  object_types        = ["dcim.powerpanel"]
  related_object_type = "dcim.device"
}
resource "netbox_custom_field" "txt" {
  name         = "%[1]s_txt"
  type         = "text"
  object_types = ["dcim.cable"]
}
resource "netbox_custom_field" "json" {
  name         = "%[1]s_json"
  type         = "json"
  object_types = ["dcim.cable"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "label"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "unique", "true"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "is_cloneable", "true"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "ui_editable", "no"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "ui_visible", "if-set"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "search_weight", "500"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "group_name"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "default"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "validation_minimum"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "validation_maximum"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.txt", "validation_regex"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.objref", "related_object_filter"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.json", "validation_schema"),
				),
			},
			{
				ResourceName:      "netbox_custom_field.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxCustomField_selectWithChoiceSet(t *testing.T) {
	testName := strings.ReplaceAll(testAccGetTestName("cf_select"), "-", "_")
	// Serial on purpose: this test's custom field carries a default, which NetBox stamps onto every
	// newly created object of its types. A parallel test creating such an object between our create
	// and our delete would then fail its own ImportStateVerify once the field is gone. Serial tests
	// finish before any t.Parallel test starts.
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name          = "%[1]s"
  description   = "colors"
  extra_choices = [
    { value = "r", label = "Red" },
    { value = "g", label = "Green" },
  ]
}

resource "netbox_custom_field" "test" {
  name          = "%[1]s"
  type          = "select"
  object_types  = ["dcim.cable"]
  choice_set_id = netbox_custom_field_choice_set.test.id
  default       = jsonencode("r")
}

data "netbox_custom_field_choice_set" "test" {
  id = netbox_custom_field_choice_set.test.id
}
data "netbox_custom_field_choice_sets" "test" {
  filters = [
    { name = "id", value = netbox_custom_field_choice_set.test.id },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.#", "2"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.1.label", "Green"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "choices_count", "2"),
					resource.TestCheckResourceAttrPair("netbox_custom_field.test", "choice_set_id", "netbox_custom_field_choice_set.test", "id"),
					resource.TestCheckResourceAttr("netbox_custom_field.test", "default", `"r"`),
					resource.TestCheckResourceAttr("data.netbox_custom_field_choice_set.test", "extra_choices.0.value", "r"),
					resource.TestCheckResourceAttr("data.netbox_custom_field_choice_sets.test", "custom_field_choice_sets.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name          = "%[1]s"
  extra_choices = [
    { value = "r", label = "Red" },
  ]
}

resource "netbox_custom_field" "test" {
  name          = "%[1]s"
  type          = "select"
  object_types  = ["dcim.cable"]
  choice_set_id = netbox_custom_field_choice_set.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.#", "1"),
					resource.TestCheckNoResourceAttr("netbox_custom_field_choice_set.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_custom_field.test", "default"),
				),
			},
		},
	})
}

func TestAccNetboxCustomFieldChoiceSet_baseChoices(t *testing.T) {
	testName := testAccGetTestName("cs_base")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name         = "%[1]s"
  base_choices = "IATA"
  extra_choices = [
    { value = "zzz", label = "Extra" },
  ]
  choice_colors = {
    zzz = "red"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "base_choices", "IATA"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.#", "1"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "choice_colors.%", "1"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "choice_colors.zzz", "red"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "order_alphabetically", "false"),
				),
			},
			{
				// Shrink half one: the extra choices and their colors clear while the base stays.
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name         = "%[1]s"
  base_choices = "IATA"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_custom_field_choice_set.test", "extra_choices"),
					resource.TestCheckNoResourceAttr("netbox_custom_field_choice_set.test", "choice_colors"),
				),
			},
			{
				// Shrink half two: the base clears (raw_null sends an explicit null) while extra
				// choices return — a set must keep one of the two.
				Config: fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name = "%[1]s"
  extra_choices = [
    { value = "zzz", label = "Extra" },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_custom_field_choice_set.test", "base_choices"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_custom_field_choice_set.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_custom_field",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasCustomFieldsList(extras.NewExtrasCustomFieldsListParams(), nil)
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
			_, err := client.Extras.ExtrasCustomFieldsDestroy(extras.NewExtrasCustomFieldsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_custom_field_choice_set",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasCustomFieldChoiceSetsList(extras.NewExtrasCustomFieldChoiceSetsListParams(), nil)
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
			_, err := client.Extras.ExtrasCustomFieldChoiceSetsDestroy(extras.NewExtrasCustomFieldChoiceSetsDestroyParams().WithID(id), nil)
			return err
		})
}

func TestAccNetboxCustomField_typedValues(t *testing.T) {
	base := strings.ReplaceAll(testAccGetTestName("cf_typed"), "-", "_")
	config := func(tier, extra string) string {
		return fmt.Sprintf(`
resource "netbox_custom_field" "tier" {
  name               = "%[1]s_tier"
  type               = "integer"
  object_types       = ["tenancy.tenant"]
  validation_minimum = 1
  validation_maximum = 3
}

resource "netbox_custom_field" "ratio" {
  name         = "%[1]s_ratio"
  type         = "decimal"
  object_types = ["tenancy.tenant"]
}

resource "netbox_custom_field" "managed" {
  name         = "%[1]s_managed"
  type         = "boolean"
  object_types = ["tenancy.tenant"]
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
  custom_fields = {
    (netbox_custom_field.tier.name)  = "%[2]s"
    (netbox_custom_field.ratio.name) = "1.5"
%[3]s
  }
}`, base, tier, extra)
	}
	managedLine := `    (netbox_custom_field.managed.name) = "true"`
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config("2", managedLine),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "custom_fields."+base+"_tier", "2"),
					resource.TestCheckResourceAttr("netbox_tenant.test", "custom_fields."+base+"_ratio", "1.5"),
					resource.TestCheckResourceAttr("netbox_tenant.test", "custom_fields."+base+"_managed", "true"),
				),
			},
			{
				// Change the integer, drop the boolean (must be cleared).
				Config: config("3", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "custom_fields."+base+"_tier", "3"),
					resource.TestCheckNoResourceAttr("netbox_tenant.test", "custom_fields."+base+"_managed"),
				),
			},
		},
	})
}

func TestAccNetboxCustomFieldChoiceSet_orderAlphabetically(t *testing.T) {
	testName := strings.ReplaceAll(testAccGetTestName("cs_order"), "-", "_")
	// Deliberately not in alphabetical order: NetBox returns the choices sorted, and the provider must
	// keep the configured order in state (the implicit post-apply plan check fails otherwise).
	config := fmt.Sprintf(`
resource "netbox_custom_field_choice_set" "test" {
  name                 = "%[1]s"
  order_alphabetically = true
  extra_choices = [
    { value = "prod", label = "Production" },
    { value = "staging", label = "Staging" },
    { value = "dev", label = "Development" },
  ]
}`, testName)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.0.value", "prod"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "extra_choices.2.value", "dev"),
					resource.TestCheckResourceAttr("netbox_custom_field_choice_set.test", "choices_count", "3"),
				),
			},
		},
	})
}

// Typed custom field values keep their configured spelling: NetBox echoes "1.50" as 1.5 and "007"
// as 7, the fold keeps the spelling in state, and the next plan is empty.
func TestAccNetboxCustomField_spelledValues(t *testing.T) {
	base := strings.ReplaceAll(testAccGetTestName("cf_spelled"), "-", "_")
	config := fmt.Sprintf(`
resource "netbox_custom_field" "tier" {
  name         = "%[1]s_tier"
  type         = "integer"
  object_types = ["tenancy.tenant"]
}

resource "netbox_custom_field" "ratio" {
  name         = "%[1]s_ratio"
  type         = "decimal"
  object_types = ["tenancy.tenant"]
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
  custom_fields = {
    (netbox_custom_field.tier.name)  = "007"
    (netbox_custom_field.ratio.name) = "1.50"
  }
}`, base)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "custom_fields."+base+"_tier", "007"),
					resource.TestCheckResourceAttr("netbox_tenant.test", "custom_fields."+base+"_ratio", "1.50"),
				),
			},
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
		},
	})
}
