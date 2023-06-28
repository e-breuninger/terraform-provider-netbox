resource "netbox_ip_address" "this" {
  ip_address = "10.0.0.50/24"
  status     = "reserved"
}
