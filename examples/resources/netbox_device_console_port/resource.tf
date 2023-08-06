# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_console_port" "test" {
  device_id      = netbox_device.test.id
  name           = "console port"
  type           = "de-9"
  speed          = 1200
  mark_connected = true
}
