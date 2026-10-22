data "netbox_vpn_tunnels" "test" {
  filters = [
    { name = "name", value = "test-vpn_tunnel" },
  ]
}
