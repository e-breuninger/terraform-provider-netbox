resource "netbox_wireless_link" "example" {
  interface_a_id = netbox_device_interface.a.id
  interface_b_id = netbox_device_interface.b.id
  ssid           = "example-wireless-link"
  status         = "connected"
}
