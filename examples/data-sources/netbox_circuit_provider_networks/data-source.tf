data "netbox_circuit_provider_networks" "test" {
  filters = [
    { name = "name", value = "test-circuit_provider_network" },
  ]
}
