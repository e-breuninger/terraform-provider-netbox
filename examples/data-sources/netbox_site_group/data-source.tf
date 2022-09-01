// Assumes the corresponding site groups exist
data "netbox_site_group" "get_by_name" {
  name = "example-sitegroup-1"
}

data "netbox_site_group" "get_by_slug" {
  slug = "sitegrp"
}
