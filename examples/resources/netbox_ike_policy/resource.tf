resource "netbox_ike_proposal" "test" {
  name                  = "test-ike-proposal"
  authentication_method = "preshared-keys"
  encryption_algorithm  = "aes-256-cbc"
  group                 = 14
}

resource "netbox_ike_policy" "test" {
  name             = "test-ike-policy"
  version          = 1
  mode             = "main"
  ike_proposal_ids = [netbox_ike_proposal.test.id]
  preshared_key    = "test-preshared-key"
  description      = "test-description"
}
