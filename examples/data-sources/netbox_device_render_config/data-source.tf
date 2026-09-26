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

resource "netbox_config_template" "test" {
  name          = "test-config-template"
  template_code = "hostname {{ device.name }}"
}

resource "netbox_device" "test" {
  name               = "test-device"
  device_type_id     = netbox_device_type.test.id
  role_id            = netbox_device_role.test.id
  site_id            = netbox_site.test.id
  config_template_id = netbox_config_template.test.id
}

data "netbox_device_render_config" "test" {
  device_id = netbox_device.test.id
}
