resource "netbox_cluster_type" "test" {
  name = "test-cluster-type"
}

resource "netbox_cluster" "test" {
  name            = "test-cluster"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_device_role" "test" {
  name    = "test-device-role"
  vm_role = true
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_platform" "test" {
  name = "test-platform"
}

resource "netbox_virtual_machine" "test" {
  name               = "test-virtual-machine"
  cluster_id         = netbox_cluster.test.id
  status             = "active"
  role_id            = netbox_device_role.test.id
  tenant_id          = netbox_tenant.test.id
  platform_id        = netbox_platform.test.id
  vcpus              = 2
  memory_mb          = 2048
  disk_size_mb       = 10240
  local_context_data = jsonencode({ key = "test-value" })
  description        = "test-description"
}
