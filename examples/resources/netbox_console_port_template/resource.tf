resource "netbox_manufacturer" "test" {
  name = "Dell"
}

resource "netbox_device_type" "test" {
  model           = "PowerSwitch S4112F-ON"
  slug            = "powerswitch-s4112f-on"
  part_number     = "S4112F-ON"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_console_port_template" "test" {
  name           = "console"
  description    = "RJ-45 serial console port"
  label          = "Serial console"
  device_type_id = netbox_device_type.test.id
  type           = "rj-45"
}
