// Get IPSec profile by name
data "netbox_ipsec_profile" "test" {
  name = "my-ipsec-profile"
}

// Get IPSec profile by id
data "netbox_ipsec_profile" "test" {
  id = 1
}
