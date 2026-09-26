resource "netbox_ip_range" "test" {
  start_address = "192.0.2.129/24"
  end_address   = "192.0.2.200/24"
}

# Allocates the next free address of the range.
resource "netbox_available_ip_address" "test" {
  ip_range_id = netbox_ip_range.test.id
  status      = "active"
  description = "test-description"
}
