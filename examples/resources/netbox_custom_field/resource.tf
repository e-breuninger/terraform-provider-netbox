resource "netbox_custom_field" "test" {
  name               = "test_custom_field"
  type               = "integer"
  object_types       = ["dcim.device", "dcim.site"]
  label              = "test-label"
  description        = "test-description"
  group_name         = "test-group"
  required           = false
  filter_logic       = "exact"
  weight             = 100
  default            = jsonencode(5)
  validation_minimum = 1
  validation_maximum = 100
}
