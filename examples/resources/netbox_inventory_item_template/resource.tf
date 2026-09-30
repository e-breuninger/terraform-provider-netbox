resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
}

resource "netbox_inventory_item_role" "test" {
  name = "test-inventory-item-role"
}

resource "netbox_inventory_item_template" "test" {
  device_type_id  = netbox_device_type.test.id
  name            = "test-inventory-item-template"
  role_id         = netbox_inventory_item_role.test.id
  manufacturer_id = netbox_manufacturer.test.id
  part_id         = "test-part-id"
  label           = "test-label"
  description     = "test-description"
}
