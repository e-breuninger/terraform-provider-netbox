// Get service template by name
data "netbox_service_template" "test" {
  name = "test"
}

// Get service template by id
data "netbox_service_template" "test_by_id" {
  id = "1"
}
