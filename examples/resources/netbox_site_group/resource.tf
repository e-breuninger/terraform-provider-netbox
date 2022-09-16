resource "netbox_site_group" "parent" {
  name        = "parent"
  description = "sample description"
}

resource "netbox_site_group" "child" {
  name        = "child"
  description = "sample description"

  parent_id = netbox_site_group.parent.id
}
