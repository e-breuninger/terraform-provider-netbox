// assumes that the referenced console port resources exist
resource "netbox_cable" "test" {
  a_termination {
    object_type = "dcim.consoleserverport"
    object_id   = netbox_device_console_server_port.kvm1.id
  }
  a_termination {
    object_type = "dcim.consoleserverport"
    object_id   = netbox_device_console_server_port.kvm2.id
  }

  b_termination {
    object_type = "dcim.consoleport"
    object_id   = netbox_device_console_port.server1.id
  }
  b_termination {
    object_type = "dcim.consoleport"
    object_id   = netbox_device_console_port.server2.id
  }

  status      = "connected"
  label       = "KVM cable"
  type        = "cat8"
  color_hex   = "123456"
  length      = 10
  length_unit = "m"
}
