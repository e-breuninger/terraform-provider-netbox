resource "netbox_manufacturer" "test" {
  name = "Cisco"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "PWR-C1-715WAC-P"
  part_number     = "PWR-C1-715WAC-P"
}

resource "netbox_power_port_template" "test" {
  name           = "PSU-1"
  label          = "715W AC power supply"
  module_type_id = netbox_module_type.test.id
  type           = "iec-60320-c16"
  maximum_draw   = 715
}
