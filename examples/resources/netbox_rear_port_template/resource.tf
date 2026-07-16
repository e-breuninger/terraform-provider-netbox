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
  label          = "MTP-12 trunk port"
  module_type_id = netbox_module_type.test.id
  type           = "mpo"
  positions      = 12
  color_hex      = "f44336"
}
