resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_rack_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-rack-type"
  slug            = "test-rack-type"
  form_factor     = "4-post-cabinet"
  width           = 19
  u_height        = 42
  starting_unit   = 1
  desc_units      = false
  outer_width     = 600
  outer_height    = 2000
  outer_depth     = 1000
  outer_unit      = "mm"
  mounting_depth  = 900
  weight          = 120
  max_weight      = 1500
  weight_unit     = "kg"
  description     = "test-description"
}
