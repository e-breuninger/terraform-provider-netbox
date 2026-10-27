resource "netbox_ip_range" "test" {
  start_address = "10.0.0.1/24"
  end_address   = "10.0.0.50/24"
}

resource "netbox_ip_address" "test" {
  ip_range_id = netbox_ip_range.test.id
  status      = "active"
}
