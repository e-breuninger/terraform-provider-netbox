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
}

resource "netbox_power_outlet_template" "test" {
  device_type_id         = netbox_device_type.test.id
  name                   = "test-power-outlet-template"
  type                   = "iec-60320-c13"
  power_port_template_id = netbox_power_port_template.test.id
  feed_leg               = "A"
  label                  = "test-label"
  description            = "test-description"
}
