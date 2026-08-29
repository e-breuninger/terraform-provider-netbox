resource "netbox_ike_proposal" "test" {
  name                    = "my-ike-proposal"
  authentication_method   = "preshared-keys"
  encryption_algorithm    = "aes-256-gcm"
  group                   = 14
}

resource "netbox_ike_policy" "test" {
  name         = "my-ike-policy"
  version      = 2
  proposal_ids = [netbox_ike_proposal.test.id]
}

resource "netbox_ipsec_proposal" "test" {
  name                  = "my-ipsec-proposal"
  encryption_algorithm  = "aes-256-gcm"
}

resource "netbox_ipsec_policy" "test" {
  name         = "my-ipsec-policy"
  proposal_ids = [netbox_ipsec_proposal.test.id]
}

resource "netbox_ipsec_profile" "test" {
  name            = "my-ipsec-profile"
  description     = "This is a description."
  mode            = "esp"
  ike_policy_id   = netbox_ike_policy.test.id
  ipsec_policy_id = netbox_ipsec_policy.test.id

  comments = "some comments"
}
