resource "netbox_user" "test" {
  username = "johndoe"
  password = "abcdefghijkl"
  active   = true
  staff    = true
}

resource "netbox_permission" "test" {
  name         = "test"
  description  = "my description"
  enabled      = true
  object_types = ["ipam.prefix"]
  actions      = ["add", "change"]
  users        = [netbox_user.test.id]
  constraints = jsonencode([{
    "status" = "active"
  }])
}
