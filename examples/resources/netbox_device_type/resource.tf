resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
  slug            = "test-device-type"
  part_number     = "test-part-number"
  u_height        = 1
  is_full_depth   = false
  airflow         = "front-to-rear"
  weight          = 12.5
  weight_unit     = "kg"
  description     = "test-description"
}
