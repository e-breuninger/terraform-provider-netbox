resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_vrf" "test" {
  name = "test-vrf"
}

resource "netbox_ipam_role" "test" {
  name = "test-ipam-role"
}

resource "netbox_vlan" "test" {
  name = "test-vlan"
  vid  = 100
}

resource "netbox_prefix" "test" {
  prefix        = "192.0.2.0/24"
  status        = "active"
  is_pool       = true
  mark_utilized = false
  description   = "test-description"
  tenant_id     = netbox_tenant.test.id
  vrf_id        = netbox_vrf.test.id
  role_id       = netbox_ipam_role.test.id
  vlan_id       = netbox_vlan.test.id
  site_id       = netbox_site.test.id
}
