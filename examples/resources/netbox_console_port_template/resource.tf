resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
}

resource "netbox_console_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "test-console-port-template"
  type           = "rj-45"
  label          = "test-label"
  description    = "test-description"
}
