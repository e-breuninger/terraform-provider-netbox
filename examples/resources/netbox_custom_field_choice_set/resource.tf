resource "netbox_custom_field_choice_set" "test" {
  name        = "test-custom-field-choice-set"
  description = "test-description"
  extra_choices = [
    { value = "a", label = "test-a" },
    { value = "b", label = "test-b" },
  ]
}
