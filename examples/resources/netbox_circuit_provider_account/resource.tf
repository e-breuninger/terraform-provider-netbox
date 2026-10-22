resource "netbox_circuit_provider" "test" {
  name = "test-circuit-provider"
}

resource "netbox_circuit_provider_account" "test" {
  account             = "test-account"
  name                = "test-circuit-provider-account"
  circuit_provider_id = netbox_circuit_provider.test.id
  description         = "test-description"
}
