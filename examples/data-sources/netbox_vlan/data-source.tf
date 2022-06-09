# Get VLAN by name
data "netbox_vlan" "vlan1" {
  name = "vlan-1"
}

# Get VLAN by VLAN ID
data "netbox_vlan" "vlan2" {
  vid = 1234
}
