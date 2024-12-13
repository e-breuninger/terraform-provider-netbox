# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name      = "SFP"
}

resource "netbox_manufacturer" "test" {
  name = "Dell"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "Networking"
}

resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status         = "active"

  description = "SFP card"
}
