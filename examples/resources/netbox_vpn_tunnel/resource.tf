resource "netbox_vpn_tunnel_group" "test" {
  name = "my-tunnel-group"
}

resource "netbox_vpn_tunnel" "test" {
  name            = "my-tunnel"
  encapsulation   = "ipsec-transport"
  status          = "active"
  tunnel_group_id = netbox_vpn_tunnel_group.test.id

  description = "This is a description."
  tunnel_id   = 3
  tenant_id   = 2
}
