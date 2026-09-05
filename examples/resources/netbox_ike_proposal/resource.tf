resource "netbox_ike_proposal" "test" {
  name                      = "my-ike-proposal"
  description               = "This is a description."
  authentication_method     = "preshared-keys"
  encryption_algorithm      = "aes-256-gcm"
  authentication_algorithm  = "hmac-sha256"
  group                     = 14
  sa_lifetime               = 3600

  comments = "some comments"
}
