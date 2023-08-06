resource "netbox_manufacturer" "test" {
  name = "Dell"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "Networking"
}
