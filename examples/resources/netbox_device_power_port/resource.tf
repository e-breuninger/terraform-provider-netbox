# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_power_port" "test" {
  device_id      = netbox_device.test.id
  name           = "power port"
  maximum_draw   = 750
  allocated_draw = 500
  type           = "iec-60320-c6"
}
