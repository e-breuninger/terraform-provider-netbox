resource "netbox_device" "mydevice" {
  name           = "mydevice"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_interface" "mydevice_eth0" {
  name      = "eth0"
  device_id = netbox_device.mydevice.id
  type      = "1000base-t"
}

resource "netbox_mac_address" "mydevice_mac" {
  mac_address         = "00:1A:2B:3C:4D:5E"
  device_interface_id = netbox_device_interface.mydevice_eth0.id
}

resource "netbox_device_interface_primary_mac_address" "mydevice_primary_mac" {
  interface_id   = netbox_device_interface.mydevice_eth0.id
  mac_address_id = netbox_mac_address.mydevice_mac.id
}
