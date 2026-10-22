resource "netbox_ipsec_proposal" "test" {
  name                 = "test-ipsec-proposal"
  encryption_algorithm = "aes-256-cbc"
}

resource "netbox_ipsec_policy" "test" {
  name               = "test-ipsec-policy"
  ipsec_proposal_ids = [netbox_ipsec_proposal.test.id]
  pfs_group          = 14
  description        = "test-description"
}
