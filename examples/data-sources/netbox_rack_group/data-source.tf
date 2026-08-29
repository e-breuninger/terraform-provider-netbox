resource "netbox_rack_group" "test" {
  name = "test"
}

data "netbox_rack_group" "test" {
  name = netbox_rack_group.test.name
}
