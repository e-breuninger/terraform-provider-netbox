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

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-module-type"
}

resource "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name      = "test-module-bay"
  position  = "1"
}

resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status         = "active"
  serial         = "test-serial"
  asset_tag      = "test-asset-tag"
  description    = "test-description"
}
