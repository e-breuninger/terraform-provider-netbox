# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_ip_address" "test_v4" {
  ip_address          = "1.1.1.1/32"
  status              = "active"
  device_interface_id = netbox_device_interface.test.id
}

resource "netbox_device_primary_ip" "test_v4" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.test.id
}
