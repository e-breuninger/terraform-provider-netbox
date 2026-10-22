resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-module-type"
  part_number     = "test-part-number"
  weight          = 1.2
  weight_unit     = "kg"
  description     = "test-description"
}
