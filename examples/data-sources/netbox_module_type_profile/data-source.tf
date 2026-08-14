resource "netbox_module_type_profile" "test" {
  name = "test"
}

data "netbox_module_type_profile" "test" {
  name = netbox_module_type_profile.test.name
}
