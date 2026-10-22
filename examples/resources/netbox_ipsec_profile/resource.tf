resource "netbox_ike_proposal" "test" {
  name                  = "test-ike-proposal"
  authentication_method = "preshared-keys"
  encryption_algorithm  = "aes-256-cbc"
  group                 = 14
}

resource "netbox_ike_policy" "test" {
  name             = "test-ike-policy"
  version          = 2
  ike_proposal_ids = [netbox_ike_proposal.test.id]
}

resource "netbox_ipsec_proposal" "test" {
  name                 = "test-ipsec-proposal"
  encryption_algorithm = "aes-256-cbc"
}

resource "netbox_ipsec_policy" "test" {
  name               = "test-ipsec-policy"
  ipsec_proposal_ids = [netbox_ipsec_proposal.test.id]
}

resource "netbox_ipsec_profile" "test" {
  name            = "test-ipsec-profile"
  mode            = "esp"
  ike_policy_id   = netbox_ike_policy.test.id
  ipsec_policy_id = netbox_ipsec_policy.test.id
  description     = "test-description"
}
