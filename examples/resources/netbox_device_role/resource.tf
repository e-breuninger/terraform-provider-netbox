# Device role for physical devices only
resource "netbox_device_role" "core_sw" {
  color_hex = "ff00ff"
  name      = "core-sw"
  vm_role   = false
}

# Device role that can be used for both devices and virtual machines
resource "netbox_device_role" "web_server" {
  color_hex   = "00ff00"
  name        = "web-server"
  description = "Web server role"
  vm_role     = true
}
