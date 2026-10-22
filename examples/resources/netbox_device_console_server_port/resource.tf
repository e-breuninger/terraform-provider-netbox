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

resource "netbox_device_console_server_port" "test" {
  device_id      = netbox_device.test.id
  name           = "test-console-server-port"
  type           = "rj-45"
  speed          = 115200
  label          = "test-label"
  mark_connected = true
  description    = "test-description"
}
