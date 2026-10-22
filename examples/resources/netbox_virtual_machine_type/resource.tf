resource "netbox_platform" "test" {
  name = "test-platform"
}

resource "netbox_virtual_machine_type" "test" {
  name                = "test-virtual-machine-type"
  slug                = "test-virtual-machine-type"
  default_platform_id = netbox_platform.test.id
  default_vcpus       = 4
  default_memory      = 8192
  description         = "test-description"
}
