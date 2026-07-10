resource "netbox_manufacturer" "test" {
  name = "APC"
}

resource "netbox_device_type" "test" {
  model           = "APDU11150ME"
  slug            = "apdu11150me"
  part_number     = "APDU11150ME"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_power_port_template" "test" {
  name           = "inlet"
  device_type_id = netbox_device_type.test.id
  type           = "iec-60309-p-n-e-6h"
}

resource "netbox_power_outlet_template" "test" {
  name           = "outlet-01"
  label          = "C13/C19 combo outlet 01"
  device_type_id = netbox_device_type.test.id
  type           = "iec-60320-c13"
  power_port_id  = netbox_power_port_template.test.id
  feed_leg       = "A"
}
