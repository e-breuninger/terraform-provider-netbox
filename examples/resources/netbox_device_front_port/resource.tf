# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_rear_port" "test" {
  device_id      = netbox_device.test.id
  name           = "rear port 1"
  type           = "8p8c"
  positions      = 2
  mark_connected = true
}

resource "netbox_device_front_port" "test" {
  device_id          = netbox_device.test.id
  name               = "front port 1"
  type               = "8p8c"
  rear_port_id       = netbox_device_rear_port.test.id
  rear_port_position = 2
}
