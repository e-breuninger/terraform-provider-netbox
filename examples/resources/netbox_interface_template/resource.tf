resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
}

resource "netbox_interface_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "test-interface-template"
  type           = "1000base-t"
  mgmt_only      = true
  poe_mode       = "pse"
  poe_type       = "type1-ieee802.3af"
  label          = "test-label"
  description    = "test-description"
}
