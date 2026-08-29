resource "netbox_ipsec_proposal" "test" {
  name                      = "my-ipsec-proposal"
  description               = "This is a description."
  encryption_algorithm      = "aes-256-gcm"
  authentication_algorithm  = "hmac-sha256"
  sa_lifetime_seconds       = 3600
  sa_lifetime_data          = 1000000

  comments = "some comments"
}
