resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

# A group scoped to the site; the allocated VLAN is placed at the site too.
resource "netbox_vlan_group" "test" {
  name    = "test-vlan-group"
  site_id = netbox_site.test.id
  vid_ranges = [
    { start = 100, end = 199 },
  ]
}

resource "netbox_available_vlan" "test" {
  name      = "test-vlan"
  group_id  = netbox_vlan_group.test.id
  site_id   = netbox_site.test.id
  tenant_id = netbox_tenant.test.id
  status    = "active"
}
