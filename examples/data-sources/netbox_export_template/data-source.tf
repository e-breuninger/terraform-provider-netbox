// Get an export template by name
data "netbox_export_template" "by_name" {
  name = "site-names"
}

// Get an export template by ID
data "netbox_export_template" "by_id" {
  id = "1"
}
