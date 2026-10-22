resource "netbox_cluster_type" "lab" {
  name = "lab"
}

resource "netbox_cluster" "lab" {
  name            = "lab"
  cluster_type_id = netbox_cluster_type.lab.id
}

resource "netbox_virtual_machine" "web" {
  name       = "web"
  cluster_id = netbox_cluster.lab.id
}

# Was netbox_interface.
resource "netbox_virtual_machine_interface" "eth0" {
  virtual_machine_id = netbox_virtual_machine.web.id
  name               = "eth0"
}

# Was interface_id plus object_type = "virtualization.vminterface".
resource "netbox_ip_address" "web" {
  ip_address                   = "192.0.2.10/24"
  virtual_machine_interface_id = netbox_virtual_machine_interface.eth0.id
}

# virtual_machine_id as before.
resource "netbox_primary_ip" "web" {
  virtual_machine_id = netbox_virtual_machine.web.id
  ip_address_id      = netbox_ip_address.web.id
}
