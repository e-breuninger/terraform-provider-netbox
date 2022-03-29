terraform {
  required_providers {
    netbox = {
      source = "e-breuninger/netbox"
    }
  }
}

# example provider configuration for a local netbox deployment
# e.g. https://github.com/netbox-community/netbox-docker
provider "netbox" {
  server_url = "http://localhost:8000"
  api_token  = "0123456789abcdef0123456789abcdef01234567"
}

# example provider configuration for https://netboxdemo.om
#provider "netbox" {
#  server_url = "https://netboxdemo.com/"
#  api_token  = "72830d67beff4ae178b94d8f781842408df8069d"
#}

resource "netbox_tag" "foo" {
  name      = "foo"
  color_hex = "00ff00" # green
}

resource "netbox_tag" "bar" {
  name = "bar"
}

resource "netbox_custom_field" "issue" {
  name = "issue"
  type = "url"
  content_types = ["virtualization.virtualmachine"]
}

resource "netbox_device_role" "testdevicerole" {
  name      = "my-device-role"
  vm_role   = true
  color_hex = "ff0000" # beautiful red
}

resource "netbox_site" "testsite" {
  name   = "my-test-site"
  status = "active"
}

resource "netbox_platform" "testplatform" {
  name = "my-test-platform"
}

resource "netbox_cluster_type" "testclustertype" {
  name = "my-test-cluster-type"
}

resource "netbox_cluster_group" "testclustergroup" {
  name        = "my-test-cluster-group"
  description = "test cluster group description"
}

resource "netbox_vrf" "testvrf" {
  name = "my-test-vrf"
}

resource "netbox_cluster" "testcluster" {
  name             = "my-test-cluster"
  cluster_type_id  = netbox_cluster_type.testclustertype.id
  cluster_group_id = netbox_cluster_group.testclustergroup.id
  site_id          = netbox_site.testsite.id

  # tags can be referenced by name but have to be created first ..
  tags = ["foo"]
  # .. or explicitly depended upon, unless created separately
  depends_on = [netbox_tag.foo]
}

resource "netbox_tenant" "testtenant" {
  name = "my-test-tenant"
}

resource "netbox_virtual_machine" "testvm" {
  name          = "my-test-vm"
  comments      = "my-test-comment"
  memory_mb     = 1024
  vcpus         = 4
  disk_size_gb  = 512
  cluster_id    = netbox_cluster.testcluster.id
  tenant_id     = netbox_tenant.testtenant.id
  platform_id   = netbox_platform.testplatform.id
  role_id       = netbox_device_role.testdevicerole.id
  tags          = [netbox_tag.foo.name, netbox_tag.bar.name]
  custom_fields = {
    "${netbox_custom_field.issue.name}" = "https://github.com/e-breuninger/terraform-provider-netbox/issues/76"
  }
}

resource "netbox_interface" "testinterface" {
  virtual_machine_id = netbox_virtual_machine.testvm.id
  name               = "my-test-interface"
  description        = "description"

  tags = [netbox_tag.foo.name]
}

resource "netbox_ip_address" "testip" {
  ip_address   = "1.2.3.4/32"
  dns_name     = "test.example.com"
  interface_id = netbox_interface.testinterface.id
  status       = "active"
  tenant_id    = netbox_tenant.testtenant.id
  vrf_id       = netbox_vrf.testvrf.id

  tags = [netbox_tag.foo.name]
}

resource "netbox_primary_ip" "testprimaryip" {
  virtual_machine_id = netbox_virtual_machine.testvm.id
  ip_address_id      = netbox_ip_address.testip.id
}

resource "netbox_service" "testservice" {
  name               = "my-test-service"
  virtual_machine_id = netbox_virtual_machine.testvm.id
  protocol           = "tcp"
  ports              = [80]
}
