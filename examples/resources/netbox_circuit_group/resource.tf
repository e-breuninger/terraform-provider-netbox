resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_circuit_group" "test" {
  name        = "test-circuit-group"
  slug        = "test-circuit-group"
  tenant_id   = netbox_tenant.test.id
  description = "test-description"
}
