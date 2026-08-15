// Get IPSec proposal by name
data "netbox_ipsec_proposal" "test" {
  name = "my-ipsec-proposal"
}

// Get IPSec proposal by id
data "netbox_ipsec_proposal" "test" {
  id = 1
}
