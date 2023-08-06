# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  tenant_id      = netbox_tenant.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_rear_port" "test" {
  device_id      = netbox_device.test.id
  name           = "rear port"
  type           = "8p8c"
  positions      = 1
  mark_connected = true
}

resource "netbox_inventory_item" "parent" {
  device_id = netbox_device.test.id
  name      = "Parent Item"
}

resource "netbox_inventory_item" "test" {
  device_id = netbox_device.test.id
  name      = "Child Item"
  parent_id = netbox_inventory_item.parent.id

  component_type = "dcim.rearport"
  component_id   = netbox_device_rear_port.test.id
}
