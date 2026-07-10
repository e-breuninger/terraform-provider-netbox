resource "netbox_manufacturer" "test" {
  name = "Cisco"
}

resource "netbox_device_type" "test" {
  model           = "Catalyst 9300X-24Y"
  slug            = "catalyst-9300x-24y"
  part_number     = "C9300X-24Y"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_module_bay_template" "test" {
  name           = "TwentyFiveGigE1/0/1"
  label          = "SFP28 cage"
  position       = "TwentyFiveGigE1/0/1"
  device_type_id = netbox_device_type.test.id
}
