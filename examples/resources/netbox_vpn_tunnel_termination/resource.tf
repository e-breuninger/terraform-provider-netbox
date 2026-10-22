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

resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "test-device-interface"
  type      = "1000base-t"
}

resource "netbox_ip_address" "test" {
  ip_address          = "192.0.2.1/24"
  status              = "active"
  device_interface_id = netbox_device_interface.test.id
}

resource "netbox_vpn_tunnel" "test" {
  name          = "test-vpn-tunnel"
  encapsulation = "ipsec-tunnel"
}

resource "netbox_vpn_tunnel_termination" "test" {
  vpn_tunnel_id         = netbox_vpn_tunnel.test.id
  role                  = "hub"
  device_interface_id   = netbox_device_interface.test.id
  outside_ip_address_id = netbox_ip_address.test.id
}
