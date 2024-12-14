resource "netbox_manufacturer" "test" {
  name = "my-manufacturer"
}

resource "netbox_device_type" "test" {
  model           = "test-model"
  slug            = "test-model"
  part_number     = "test-part-number"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_interface_template" "test" {
  name           = "eth0"
  description    = "eth0 description"
  label          = "eth0 label"
  device_type_id = netbox_device_type.test.id
  type           = "100base-tx"
  mgmt_only      = true
}
