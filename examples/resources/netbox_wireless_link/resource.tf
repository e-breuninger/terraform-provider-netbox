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

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_device_interface" "test_a" {
  device_id = netbox_device.test.id
  name      = "test-device-interface-a"
  type      = "ieee802.11ax"
}

resource "netbox_device_interface" "test_b" {
  device_id = netbox_device.test.id
  name      = "test-device-interface-b"
  type      = "ieee802.11ax"
}

resource "netbox_wireless_link" "test" {
  interface_a_id = netbox_device_interface.test_a.id
  interface_b_id = netbox_device_interface.test_b.id
  ssid           = "test-ssid"
  status         = "planned"
  tenant_id      = netbox_tenant.test.id
  auth_type      = "wpa-personal"
  auth_cipher    = "aes"
  auth_psk       = "test-psk"
  distance       = 1.5
  distance_unit  = "km"
  description    = "test-description"
}
