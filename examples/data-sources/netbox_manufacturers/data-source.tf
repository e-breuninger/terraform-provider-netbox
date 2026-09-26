data "netbox_manufacturers" "test" {
  filters = [
    { name = "name", value = "test-manufacturer" },
  ]
}
