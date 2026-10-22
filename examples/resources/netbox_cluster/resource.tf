resource "netbox_cluster_type" "test" {
  name = "test-cluster-type"
}

resource "netbox_cluster_group" "test" {
  name = "test-cluster-group"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_cluster" "test" {
  name             = "test-cluster"
  cluster_type_id  = netbox_cluster_type.test.id
  cluster_group_id = netbox_cluster_group.test.id
  tenant_id        = netbox_tenant.test.id
  site_id          = netbox_site.test.id
  status           = "active"
  description      = "test-description"
}
