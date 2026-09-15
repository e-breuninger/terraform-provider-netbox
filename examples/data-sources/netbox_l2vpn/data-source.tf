// Get L2VPN by name
data "netbox_l2vpn" "example" {
  name = "example-l2vpn"
}

// Get L2VPN by ID
data "netbox_l2vpn" "by_id" {
  id = "42"
}
