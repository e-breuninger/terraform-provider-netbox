// Assumes the 'dc-west' cluster group already exists
data "netbox_cluster_group" "dc-west" {
    name = "dc-west"
}

resource "netbox_cluster_type" "vmw-vsphere" {
    name = "VMware vSphere 6"
}

resource "netbox_cluster" "vmw-cluster-01" {
    cluster_type_id = netbox_cluster_type.vmw-vsphere.id
    name = "vmw-cluster-01"
    cluster_group_id = data.netbox_cluster_group.dc-west.id
}
