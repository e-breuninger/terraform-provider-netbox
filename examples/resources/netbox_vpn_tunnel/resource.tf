resource "netbox_vpn_tunnel_group" "test" {
  name = "test-vpn-tunnel-group"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_vpn_tunnel" "test" {
  name                = "test-vpn-tunnel"
  encapsulation       = "gre"
  status              = "active"
  vpn_tunnel_group_id = netbox_vpn_tunnel_group.test.id
  tenant_id           = netbox_tenant.test.id
  tunnel_id           = 42
  description         = "test-description"
}
