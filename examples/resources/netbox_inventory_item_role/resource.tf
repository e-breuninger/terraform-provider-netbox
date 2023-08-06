# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  tenant_id      = netbox_tenant.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_inventory_item_role" "test" {
  name      = "Role 1"
  slug      = "role-1-slug"
  color_hex = "123456"
}

resource "netbox_inventory_item" "parent" {
  device_id = netbox_device.test.id
  name      = "Inventory Item 1"
  role_id   = netbox_inventory_item_role.test.id
}
