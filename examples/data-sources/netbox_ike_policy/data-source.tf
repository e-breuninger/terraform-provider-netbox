// Get IKE policy by name
data "netbox_ike_policy" "test" {
  name = "my-ike-policy"
}

// Get IKE policy by id
data "netbox_ike_policy" "test" {
  id = 1
}
