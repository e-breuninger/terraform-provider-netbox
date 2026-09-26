data "netbox_virtual_circuit_types" "test" {
  filters = [
    { name = "name", value = "test-virtual_circuit_type" },
  ]
}
