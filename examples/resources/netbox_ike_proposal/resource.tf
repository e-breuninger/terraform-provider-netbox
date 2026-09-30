resource "netbox_ike_proposal" "test" {
  name                     = "test-ike-proposal"
  authentication_method    = "preshared-keys"
  encryption_algorithm     = "aes-256-cbc"
  authentication_algorithm = "hmac-sha256"
  group                    = 14
  sa_lifetime              = 28800
  description              = "test-description"
}
