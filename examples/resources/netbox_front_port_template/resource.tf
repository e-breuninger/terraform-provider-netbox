resource "netbox_manufacturer" "test" {
  name = "FS.COM"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "FHD MTP-12 OM4 Cassette"
  part_number     = "FHD-MTP12-OM4-LC12"
}

resource "netbox_rear_port_template" "test" {
  name           = "MTP-1"
  module_type_id = netbox_module_type.test.id
  type           = "mpo"
  positions      = 12
}

resource "netbox_front_port_template" "test" {
  name               = "LC-1"
  label              = "LC duplex port 1"
  module_type_id     = netbox_module_type.test.id
  type               = "lc-upc"
  rear_port_id       = netbox_rear_port_template.test.id
  rear_port_position = 1
  color_hex          = "f44336"
}
