resource "netbox_circuit_provider" "test" {
  name = "test-circuit-provider"
}

resource "netbox_circuit_provider_network" "test" {
  name                = "test-circuit-provider-network"
  circuit_provider_id = netbox_circuit_provider.test.id
  service_id          = "test-service-id"
  description         = "test-description"
}
