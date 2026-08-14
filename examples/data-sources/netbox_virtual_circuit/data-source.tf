// Get Virtual Circuit by CID
data "netbox_virtual_circuit" "example" {
  cid = "VC-0001"
}

// Get Virtual Circuit by ID
data "netbox_virtual_circuit" "example_by_id" {
  id = 1
}
