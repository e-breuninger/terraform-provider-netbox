resource "netbox_user" "test" {
  username = "johndoe"
  password = "abcdefghijkl"
  active   = true
  staff    = true
}
