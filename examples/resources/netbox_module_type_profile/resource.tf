resource "netbox_module_type_profile" "test" {
  name        = "test"
  description = "This is my module type profile"
  comments    = "Some comments"
  schema = jsonencode({
    type = "object"
    properties = {
      port_count = {
        type = "integer"
      }
    }
  })
}
