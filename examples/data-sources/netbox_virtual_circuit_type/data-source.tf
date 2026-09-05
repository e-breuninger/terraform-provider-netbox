// Get Virtual Circuit Type by Name
data "netbox_virtual_circuit_type" "example" {
  name = "EVPL"
}

// Get Virtual Circuit Type by ID
data "netbox_virtual_circuit_type" "example_by_id" {
  id = 1
}
