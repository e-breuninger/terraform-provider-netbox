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

resource "netbox_device" "test_a" {
  name           = "test-device-a"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device" "test_b" {
  name           = "test-device-b"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_interface" "test_a" {
  device_id = netbox_device.test_a.id
  name      = "test-interface"
  type      = "1000base-t"
}

resource "netbox_device_interface" "test_b" {
  device_id = netbox_device.test_b.id
  name      = "test-interface"
  type      = "1000base-t"
}

resource "netbox_cable" "test" {
  a_side = { device_interface_ids = [netbox_device_interface.test_a.id] }
  b_side = { device_interface_ids = [netbox_device_interface.test_b.id] }

  type        = "cat6"
  status      = "connected"
  label       = "test-label"
  color_hex   = "ff0000"
  length      = 3
  length_unit = "m"
  description = "test-description"
}
