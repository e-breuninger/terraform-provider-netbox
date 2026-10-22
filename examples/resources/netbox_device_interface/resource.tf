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

resource "netbox_vlan" "test" {
  name = "test-vlan"
  vid  = 100
}

resource "netbox_vrf" "test" {
  name = "test-vrf"
}

resource "netbox_device_interface" "test_lag" {
  device_id = netbox_device.test.id
  name      = "test-lag"
  type      = "lag"
}

resource "netbox_device_interface" "test" {
  device_id               = netbox_device.test.id
  name                    = "test-interface"
  type                    = "1000base-t"
  label                   = "test-label"
  enabled                 = true
  mgmt_only               = true
  mtu                     = 9000
  speed                   = 1000000
  duplex                  = "full"
  mode                    = "tagged"
  untagged_vlan_id        = netbox_vlan.test.id
  tagged_vlan_ids         = [netbox_vlan.test.id]
  lag_device_interface_id = netbox_device_interface.test_lag.id
  vrf_id                  = netbox_vrf.test.id
  description             = "test-description"
}
