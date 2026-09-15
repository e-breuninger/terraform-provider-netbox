resource "netbox_site" "test" {
  name   = "test"
  status = "active"
}

resource "netbox_manufacturer" "test" {
  name = "test"
}

resource "netbox_device_type" "test" {
  model           = "test"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device_role" "test" {
  name      = "test"
  color_hex = "123456"
}

resource "netbox_device" "test" {
  name           = "test"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
  status         = "active"
}

resource "netbox_virtual_device_context" "test" {
  name        = "test"
  device_id   = netbox_device.test.id
  identifier  = 1
  status      = "active"
  description = "This is my virtual device context"
}
