resource "netbox_config_context" "test" {
  name = "%s"
  data = jsonencode({"testkey" = "testval"})
}
