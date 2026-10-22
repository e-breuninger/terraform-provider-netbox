data "netbox_circuit_provider_accounts" "test" {
  filters = [
    { name = "name", value = "test-circuit_provider_account" },
  ]
}
