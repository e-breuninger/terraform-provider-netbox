data "netbox_circuit_groups" "test" {
  filters = [
    { name = "name", value = "test-circuit_group" },
  ]
}
