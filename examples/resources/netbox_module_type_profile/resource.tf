resource "netbox_module_type_profile" "test" {
  name        = "test-module-type-profile"
  description = "test-description"
  schema = jsonencode({
    properties = {
      capacity = { type = "integer" }
    }
  })
}
