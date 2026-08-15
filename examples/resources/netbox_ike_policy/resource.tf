resource "netbox_ike_proposal" "test" {
  name                   = "my-ike-proposal"
  authentication_method  = "preshared-keys"
  encryption_algorithm   = "aes-256-gcm"
  group                  = 14
}

resource "netbox_ike_policy" "test" {
  name          = "my-ike-policy"
  description   = "This is a description."
  version       = 1
  mode          = "main"
  proposal_ids  = [netbox_ike_proposal.test.id]
  preshared_key = "supersecret"

  comments = "some comments"
}
