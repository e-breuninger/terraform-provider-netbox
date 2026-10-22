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
  mgmt_only = true
}

resource "netbox_ip_address" "test" {
  ip_address          = "192.0.2.10/24"
  device_interface_id = netbox_device_interface.test.id
}

resource "netbox_device_oob_ip" "test" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.test.id
}
