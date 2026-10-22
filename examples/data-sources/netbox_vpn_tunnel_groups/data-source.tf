data "netbox_vpn_tunnel_groups" "test" {
  filters = [
    { name = "name", value = "test-vpn_tunnel_group" },
  ]
}
