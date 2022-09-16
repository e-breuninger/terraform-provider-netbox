resource "netbox_manufacturer" "test" {
  name = "test"
}

resource "netbox_device_type" "test" {
  model           = "test"
  part_number     = "123"
  manufacturer_id = netbox_manufacturer.test.id
}
