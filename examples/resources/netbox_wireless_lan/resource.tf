resource "netbox_wireless_lan" "guest" {
  ssid   = "guest-wifi"
  status = "active"
}
