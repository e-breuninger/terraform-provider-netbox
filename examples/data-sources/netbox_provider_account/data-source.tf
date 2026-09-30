// Get Provider Account by Name
data "netbox_provider_account" "example" {
    name = "AccountA"
}

// Get Provider Account by ID
data "netbox_provider_account" "example_by_id" {
    id = "1"
}
