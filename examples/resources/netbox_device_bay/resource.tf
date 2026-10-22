resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
  subdevice_role  = "parent"
}

resource "netbox_device_type" "test_child" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type-child"
  subdevice_role  = "child"
  u_height        = 0
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

resource "netbox_device" "test_child" {
  name           = "test-device-child"
  device_type_id = netbox_device_type.test_child.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_bay" "test" {
  device_id           = netbox_device.test.id
  name                = "test-device-bay"
  installed_device_id = netbox_device.test_child.id
  label               = "test-label"
  description         = "test-description"
}
