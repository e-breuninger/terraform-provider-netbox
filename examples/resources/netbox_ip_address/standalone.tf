resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_vrf" "test" {
  name = "test-vrf"
}

resource "netbox_ip_address" "test" {
  ip_address  = "192.0.2.40/24"
  status      = "reserved"
  role        = "vip"
  dns_name    = "test-vip.example.com"
  description = "test-description"
  tenant_id   = netbox_tenant.test.id
  vrf_id      = netbox_vrf.test.id
}
