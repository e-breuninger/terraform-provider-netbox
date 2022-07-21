data "netbox_ip_range" "test" {
  start_address = "10.0.0.1/24"
  end_address = "10.0.0.50/24"
}

resource "netbox_available_ip_address" "test" {
  ip_range_id = data.netbox_ip_range.test.id
}
