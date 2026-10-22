resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_route_target" "test" {
  name = "65000:100"
}

resource "netbox_l2vpn" "test" {
  name              = "test-l2vpn"
  slug              = "test-l2vpn"
  type              = "vxlan-evpn"
  identifier        = 100
  tenant_id         = netbox_tenant.test.id
  import_target_ids = [netbox_route_target.test.id]
  export_target_ids = [netbox_route_target.test.id]
  description       = "test-description"
}
