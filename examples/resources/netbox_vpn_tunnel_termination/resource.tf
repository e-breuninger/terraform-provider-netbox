resource "netbox_vpn_tunnel_group" "test" {
  name        = "my-tunnel-group"
  description = "description"
}

resource "netbox_vpn_tunnel" "test" {
  name            = "my-tunnel"
  encapsulation   = "ipsec-transport"
  status          = "active"
  tunnel_group_id = netbox_vpn_tunnel_group.test.id
}

resource "netbox_vpn_tunnel_termination" "device" {
  role                = "peer"
  tunnel_id           = netbox_vpn_tunnel.test.id
  device_interface_id = 123
}

resource "netbox_vpn_tunnel_termination" "vm" {
  role                         = "peer"
  tunnel_id                    = netbox_vpn_tunnel.test.id
  virtual_machine_interface_id = 234
}
