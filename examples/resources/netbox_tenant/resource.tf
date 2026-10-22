resource "netbox_tenant_group" "test" {
  name = "test-tenant-group"
}

resource "netbox_tenant" "test" {
  name        = "test-tenant"
  group_id    = netbox_tenant_group.test.id
  description = "test-description"
}
