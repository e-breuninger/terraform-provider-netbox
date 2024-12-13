resource "netbox_custom_field_choice_set" "test" {
  name        = "my-custom-field-set"
  description = "Description"
  extra_choices = [
    ["choice1", "label1"], # label and choice are different
    ["choice2", "choice2"] # label and choice are the same
  ]
}
