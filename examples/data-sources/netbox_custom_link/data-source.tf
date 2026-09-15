// Get a custom link by name
data "netbox_custom_link" "by_name" {
  name = "device-docs"
}

// Get a custom link by ID
data "netbox_custom_link" "by_id" {
  id = "1"
}
