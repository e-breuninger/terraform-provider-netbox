resource "netbox_manufacturer" "test" {
  name = "my-manufacturer"
}

resource "netbox_rack_type" "test" {
  model             = "mymodel"
  manufacturer_id   = netbox_manufacturer.test.id
  width             = 19
  u_height          = 48
  starting_unit     = 1
  form_factor       = "2-post-frame"
  description       = "My description"
  outer_width       = 10
  outer_depth       = 15
  outer_unit        = "mm"
  weight            = 15
  max_weight        = 20
  weight_unit       = "kg"
  mounting_depth_mm = 21
  comments          = "My comments"
}
