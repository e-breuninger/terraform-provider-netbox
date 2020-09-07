terraform {
  required_providers {
    netbox = {
      source  = "e-breuninger/netbox"
      version = ">=0.0.3"
    }
  }
}

provider "netbox" {
  server_url = "https://netboxdemo.com/"
  api_token  = " 72830d67beff4ae178b94d8f781842408df8069d"
}

resource "netbox_device_role" "testdevicerole" {
  name      = "my-device-role"
  vm_role   = true
  color_hex = "ff0000" # beautiful red
}

resource "netbox_platform" "testplatform" {
  name = "my-test-platform"
}

resource "netbox_cluster_type" "testclustertype" {
  name = "my-test-cluster-type"
}

resource "netbox_cluster" "testcluster" {
  name            = "my-test-cluster"
  cluster_type_id = netbox_cluster_type.testclustertype.id
}

resource "netbox_tenant" "testtenant" {
  name = "my-test-tenant"
}

resource "netbox_virtual_machine" "testvm" {
  name         = "my-test-vm"
  comments     = "my-test-comment"
  memory_mb    = 1024
  vcpus        = 4
  disk_size_gb = 512
  cluster_id   = netbox_cluster.testcluster.id
  tenant_id    = netbox_tenant.testtenant.id
  platform_id  = netbox_platform.testplatform.id
  role_id      = netbox_device_role.testdevicerole.id
  # NOTE: Custom fields have to be created in the API first
  #  custom_fields = {
  #    "custom_field_name" = "custom field value"
  #  }
}

resource "netbox_interface" "testinterface" {
  virtual_machine_id = netbox_virtual_machine.testvm.id
  name               = "my-test-interface"
  description        = "description"
  type               = "virtual"

  tags = ["my:tag", "bar"]
}

resource "netbox_ip_address" "testip" {
  ip_address   = "1.2.3.4/32"
  interface_id = netbox_interface.testinterface.id
  status       = "active"
}

resource "netbox_primary_ip" "testprimaryip" {
  virtual_machine_id = netbox_virtual_machine.testvm.id
  ip_address_id      = netbox_ip_address.testip.id
}

resource "netbox_service" "testservice" {
  name               = "my-test-service"
  virtual_machine_id = netbox_virtual_machine.testvm.id
  protocol           = "tcp"
  port               = 80
}
