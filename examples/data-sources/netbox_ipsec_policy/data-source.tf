// Get IPSec policy by name
data "netbox_ipsec_policy" "test" {
  name = "my-ipsec-policy"
}

// Get IPSec policy by id
data "netbox_ipsec_policy" "test" {
  id = 1
}
