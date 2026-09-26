resource "netbox_circuit_provider" "test" {
  name = "test-circuit-provider"
}

resource "netbox_circuit_type" "test" {
  name = "test-circuit-type"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_circuit" "test" {
  cid                 = "test-cid"
  circuit_provider_id = netbox_circuit_provider.test.id
  circuit_type_id     = netbox_circuit_type.test.id
  tenant_id           = netbox_tenant.test.id
  status              = "active"
  install_date        = "2024-05-01"
  termination_date    = "2030-05-01"
  commit_rate_kbps    = 1000000
  description         = "test-description"
}
