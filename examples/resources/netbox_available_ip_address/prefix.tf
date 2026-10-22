resource "netbox_prefix" "test" {
  prefix = "192.0.2.0/24"
  status = "active"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

# Allocates the next free address of the prefix.
resource "netbox_available_ip_address" "test" {
  prefix_id   = netbox_prefix.test.id
  status      = "reserved"
  role        = "vip"
  dns_name    = "test-host.example.com"
  description = "test-description"
  tenant_id   = netbox_tenant.test.id
}
