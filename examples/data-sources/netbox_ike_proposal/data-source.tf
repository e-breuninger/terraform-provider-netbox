// Get IKE proposal by name
data "netbox_ike_proposal" "test" {
  name = "my-ike-proposal"
}

// Get IKE proposal by id
data "netbox_ike_proposal" "test" {
  id = 1
}
