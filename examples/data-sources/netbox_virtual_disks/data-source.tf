data "netbox_virtual_disks" "test" {
  filters = [
    { name = "name", value = "test-virtual_disk" },
  ]
}
