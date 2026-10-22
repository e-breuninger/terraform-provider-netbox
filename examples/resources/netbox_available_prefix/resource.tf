resource "netbox_prefix" "test" {
  prefix = "192.0.2.0/24"
  status = "container"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

# Allocates the next free /26 inside the parent prefix.
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.test.id
  prefix_length    = 26
  status           = "active"
  tenant_id        = netbox_tenant.test.id
  description      = "test-description"
}
