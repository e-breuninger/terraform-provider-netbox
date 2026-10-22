resource "netbox_vrf" "test" {
  name = "test-vrf"
}

resource "netbox_ip_range" "test" {
  start_address = "192.0.2.129/24"
  end_address   = "192.0.2.200/24"
  vrf_id        = netbox_vrf.test.id
  description   = "test-description"
}
