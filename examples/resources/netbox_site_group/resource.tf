resource "netbox_site_group" "test" {
  name        = "test-site-group"
  slug        = "test-site-group"
  description = "test-description"
}

resource "netbox_site_group" "test_child" {
  name      = "test-site-group-child"
  parent_id = netbox_site_group.test.id
}
