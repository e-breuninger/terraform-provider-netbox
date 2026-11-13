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

resource "netbox_virtual_chassis" "test" {
  name = "test-stack"
}

resource "netbox_device" "member1" {
  name                     = "test-switch-1"
  device_type_id           = netbox_device_type.test.id
  role_id                  = netbox_device_role.test.id
  site_id                  = netbox_site.test.id
  virtual_chassis_id       = netbox_virtual_chassis.test.id
  virtual_chassis_position = 1
}

resource "netbox_device" "member2" {
  name                     = "test-switch-2"
  device_type_id           = netbox_device_type.test.id
  role_id                  = netbox_device_role.test.id
  site_id                  = netbox_site.test.id
  virtual_chassis_id       = netbox_virtual_chassis.test.id
  virtual_chassis_position = 2
}

resource "netbox_virtual_chassis_master" "test" {
  virtual_chassis_id = netbox_virtual_chassis.test.id
  device_id          = netbox_device.member1.id
}
