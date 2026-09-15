resource "netbox_manufacturer" "test" {
  name = "Cisco"
}

resource "netbox_device_type" "test" {
  model           = "Test Device Type"
  slug            = "test-device-type"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_inventory_item_role" "test" {
  name      = "Power Supply"
  slug      = "power-supply"
  color_hex = "f44336"
}

resource "netbox_inventory_item_template" "test" {
  name            = "PSU-1"
  label           = "Power Supply 1"
  device_type_id  = netbox_device_type.test.id
  role_id         = netbox_inventory_item_role.test.id
  manufacturer_id = netbox_manufacturer.test.id
  part_id         = "PWR-750"
  description     = "Redundant power supply"
}
