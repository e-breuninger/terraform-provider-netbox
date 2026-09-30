resource "netbox_l2vpn" "test" {
  name = "test-l2vpn"
  type = "vxlan"
}

resource "netbox_vlan" "test" {
  name = "test-vlan"
  vid  = 100
}

resource "netbox_l2vpn_termination" "test" {
  l2vpn_id = netbox_l2vpn.test.id
  vlan_id  = netbox_vlan.test.id
}
