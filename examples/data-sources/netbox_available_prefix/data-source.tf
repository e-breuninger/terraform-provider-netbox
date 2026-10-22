resource "netbox_prefix" "test" {
  prefix = "192.0.2.0/24"
  status = "container"
}

data "netbox_available_prefix" "test" {
  prefix_id = netbox_prefix.test.id
}
