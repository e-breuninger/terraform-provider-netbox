data "netbox_prefix" "test" {
  cidr = "10.0.0.0/24"
}

resource "netbox_available_ip_address" "test" {
  prefix_id = data.netbox_prefix.test.id
}
