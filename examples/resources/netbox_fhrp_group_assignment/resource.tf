data "netbox_fhrp_group" "test" {
  protocol    = "vrrp"
  group_id    = 1234
}

data "netbox_devices" "my_device" {
  name = "my-device"
}

data "netbox_device_interfaces" "interfaces" {
  device_id = data.netbox_devices.my_device[0].id
  name = "eth0"
}

resource "netbox_fhrp_group_assignment" "test" {
  group_id = data.netbox_fhrp_group.test.id
  interface_id = data.netbox_device_interfaces.interfaces[0].id
  interface_type = "dcim.interface"
  priority = 150
}