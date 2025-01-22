resource "netbox_user" "test" {
  username = "johndoe"
  password = "Abcdefghijkl1"
  active   = true
  staff    = true
}
