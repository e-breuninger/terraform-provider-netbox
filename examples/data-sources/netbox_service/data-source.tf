data "netbox_service" "example" {
  name               = "ssh"
  virtual_machine_id = netbox_virtual_machine.example.id
}
