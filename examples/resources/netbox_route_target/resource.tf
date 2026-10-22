resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_route_target" "test" {
  name        = "65000:1"
  description = "test-description"
  tenant_id   = netbox_tenant.test.id
}
