resource "netbox_circuit_provider" "test" {
  name = "test-circuit-provider"
}

resource "netbox_circuit_provider_network" "test" {
  name                = "test-circuit-provider-network"
  circuit_provider_id = netbox_circuit_provider.test.id
}

resource "netbox_circuit_provider_account" "test" {
  account             = "test-account"
  circuit_provider_id = netbox_circuit_provider.test.id
}

resource "netbox_virtual_circuit_type" "test" {
  name = "test-virtual-circuit-type"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_virtual_circuit" "test" {
  cid                         = "test-cid"
  circuit_provider_network_id = netbox_circuit_provider_network.test.id
  virtual_circuit_type_id     = netbox_virtual_circuit_type.test.id
  circuit_provider_account_id = netbox_circuit_provider_account.test.id
  tenant_id                   = netbox_tenant.test.id
  status                      = "active"
  description                 = "test-description"
}
