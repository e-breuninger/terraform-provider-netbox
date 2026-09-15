resource "netbox_l2vpn" "example" {
  name = "example-l2vpn"
  slug = "example-l2vpn"
  type = "vxlan"
}

resource "netbox_l2vpn_termination" "example" {
  l2vpn_id             = netbox_l2vpn.example.id
  assigned_object_type = "dcim.interface"
  assigned_object_id   = netbox_device_interface.example.id
}
