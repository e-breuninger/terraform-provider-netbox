// Get all clusters of a specific type
data "netbox_cluster_type" "vmware" {
  name = "VMware ESXi"
}

data "netbox_clusters" "vmware_clusters" {
  filter {
    name  = "cluster_type_id"
    value = data.netbox_cluster_type.vmware.id
  }
}

// Get clusters by name regex
data "netbox_clusters" "prod_clusters" {
  name_regex = "prod-.*"
}

// Get clusters at a specific site
data "netbox_clusters" "site_clusters" {
  filter {
    name  = "site_id"
    value = data.netbox_site.main.id
  }
}
