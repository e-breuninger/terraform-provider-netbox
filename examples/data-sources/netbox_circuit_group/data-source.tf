// Get Circuit Group by Name
data "netbox_circuit_group" "example" {
    name = "GroupA"
}

// Get Circuit Group by ID
data "netbox_circuit_group" "example_by_id" {
    id = "1"
}
