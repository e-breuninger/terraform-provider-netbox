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

resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "test-interface"
  type      = "1000base-t"
}

resource "netbox_mac_address" "test" {
  mac_address         = "00:11:22:33:44:55"
  device_interface_id = netbox_device_interface.test.id
}
