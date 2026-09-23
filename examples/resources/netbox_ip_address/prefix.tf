resource "netbox_prefix" "test" {
  prefix = "10.0.0.0/24"
}

resource "netbox_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
  status    = "active"
}
