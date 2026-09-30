resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_vlan" "test" {
  name = "test-vlan"
  vid  = 100
}

resource "netbox_wireless_lan_group" "test" {
  name = "test-wireless-lan-group"
}

resource "netbox_wireless_lan" "test" {
  ssid        = "test-ssid"
  group_id    = netbox_wireless_lan_group.test.id
  status      = "active"
  vlan_id     = netbox_vlan.test.id
  tenant_id   = netbox_tenant.test.id
  site_id     = netbox_site.test.id
  auth_type   = "wpa-personal"
  auth_cipher = "aes"
  auth_psk    = "test-psk"
  description = "test-description"
}
