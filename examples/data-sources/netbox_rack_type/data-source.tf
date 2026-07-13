# Get rack type by model
data "netbox_rack_type" "ex1" {
  model = "APC 4-Post Cabinet"
}

# Get rack type by slug
data "netbox_rack_type" "ex2" {
  slug = "apc-4-post-cabinet"
}
