data "netbox_prefix" "test" {
  cidr = "10.0.0.0/24"
}

resource "netbox_available_prefix" "test" {
  parent_prefix_id = data.netbox_prefix.test.id
  prefix_length    = 25
  status           = "active"
}
