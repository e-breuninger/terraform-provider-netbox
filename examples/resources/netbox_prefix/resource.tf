resource "netbox_prefix" "my_prefix" {
  prefix      = "10.0.0.0/24"
  status      = "active"
  description = "test prefix"
}
