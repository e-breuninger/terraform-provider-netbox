resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "test"
}

resource "netbox_device_type" "test" {
  model           = "test"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id

  # custom fields - optional, not required
  custom_fields = {
    "test_field_1" = "test_field_value_1",
    "test_field_2" = "test_field_value_2"
  }
}
