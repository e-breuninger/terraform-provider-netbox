resource "netbox_region" "test" {
  name        = "test-region"
  slug        = "test-region"
  description = "test-description"
}

resource "netbox_region" "test_child" {
  name             = "test-region-child"
  parent_region_id = netbox_region.test.id
}
