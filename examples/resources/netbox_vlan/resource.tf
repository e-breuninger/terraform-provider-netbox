resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_ipam_role" "test" {
  name = "test-ipam-role"
}

resource "netbox_vlan_group" "test" {
  name = "test-vlan-group"
  vid_ranges = [
    { start = 1, end = 4094 },
  ]
}

resource "netbox_vlan" "test" {
  name        = "test-vlan"
  vid         = 100
  status      = "active"
  description = "test-description"
  tenant_id   = netbox_tenant.test.id
  site_id     = netbox_site.test.id
  group_id    = netbox_vlan_group.test.id
  role_id     = netbox_ipam_role.test.id
}
