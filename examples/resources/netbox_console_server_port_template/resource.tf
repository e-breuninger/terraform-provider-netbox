resource "netbox_manufacturer" "test" {
  name = "Opengear"
}

resource "netbox_device_type" "test" {
  model           = "CM8148"
  slug            = "cm8148"
  part_number     = "CM8148"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_console_server_port_template" "test" {
  name           = "port1"
  description    = "Console server port 1"
  label          = "Port 1"
  device_type_id = netbox_device_type.test.id
  type           = "rj-45"
}
