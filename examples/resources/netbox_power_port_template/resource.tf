resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
}

resource "netbox_power_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "test-power-port-template"
  type           = "iec-60320-c14"
  maximum_draw   = 600
  allocated_draw = 300
  label          = "test-label"
  description    = "test-description"
}
