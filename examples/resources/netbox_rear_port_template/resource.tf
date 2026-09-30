resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
}

resource "netbox_rear_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "test-rear-port-template"
  type           = "8p8c"
  positions      = 4
  color_hex      = "ff0000"
  label          = "test-label"
  description    = "test-description"
}
