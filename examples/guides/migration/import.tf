# Re-imports a VLAN group whose 5.x state was removed with `terraform state rm`.
resource "netbox_vlan_group" "office" {
  name = "office"
  vid_ranges = [
    { start = 100, end = 199 },
  ]
}

import {
  to = netbox_vlan_group.office
  id = "42"
}
