resource "netbox_rir" "test" {
  name = "test-rir"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_aggregate" "test" {
  prefix      = "192.0.2.0/24"
  rir_id      = netbox_rir.test.id
  tenant_id   = netbox_tenant.test.id
  date_added  = "2026-01-15"
  description = "test-description"
}
