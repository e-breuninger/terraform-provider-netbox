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

resource "netbox_device_power_port" "test" {
  device_id = netbox_device.test.id
  name      = "test-power-port"
  type      = "iec-60320-c14"
}

resource "netbox_device_power_outlet" "test" {
  device_id     = netbox_device.test.id
  name          = "test-power-outlet"
  type          = "iec-60320-c13"
  power_port_id = netbox_device_power_port.test.id
  feed_leg      = "A"
  label         = "test-label"
  description   = "test-description"
}
