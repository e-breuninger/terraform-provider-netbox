data "netbox_circuit_types" "test" {
  filters = [
    { name = "name", value = "test-circuit_type" },
  ]
}
