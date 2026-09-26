data "netbox_circuit_providers" "test" {
  filters = [
    { name = "name", value = "test-circuit_provider" },
  ]
}
