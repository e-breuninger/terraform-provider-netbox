resource "netbox_ipsec_proposal" "test" {
  name                     = "test-ipsec-proposal"
  encryption_algorithm     = "aes-256-cbc"
  authentication_algorithm = "hmac-sha256"
  sa_lifetime_seconds      = 3600
  sa_lifetime_data         = 102400
  description              = "test-description"
}
