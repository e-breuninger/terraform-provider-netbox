resource "netbox_tenant" "test" {
  name = "test"
}
resource "netbox_route_target" "test" {
  name        = "test"
  description = "my description"
  tenant_id   = netbox_tenant.test.id
}
