resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
}

resource "netbox_device_role" "test" {
  name = "test-device-role"
}

resource "netbox_device" "test" {
  name           = "test-device"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_inventory_item_role" "test" {
  name      = "test-inventory-item-role"
  color_hex = "00ff00"
}

resource "netbox_inventory_item" "test" {
  device_id       = netbox_device.test.id
  name            = "test-inventory-item"
  role_id         = netbox_inventory_item_role.test.id
  manufacturer_id = netbox_manufacturer.test.id
  part_id         = "test-part-id"
  serial          = "test-serial"
  asset_tag       = "test-asset-tag"
  description     = "test-description"
}
