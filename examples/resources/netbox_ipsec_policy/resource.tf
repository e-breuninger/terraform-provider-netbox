resource "netbox_ipsec_proposal" "test" {
  name                  = "my-ipsec-proposal"
  encryption_algorithm  = "aes-256-gcm"
}

resource "netbox_ipsec_policy" "test" {
  name         = "my-ipsec-policy"
  description  = "This is a description."
  proposal_ids = [netbox_ipsec_proposal.test.id]
  pfs_group    = 14

  comments = "some comments"
}
