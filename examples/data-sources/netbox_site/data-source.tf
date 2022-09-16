data "netbox_site" "get_by_name" {
  name = "Example Site 1"
}

data "netbox_site" "get_by_slug" {
  slug = "example-site-1"
}
