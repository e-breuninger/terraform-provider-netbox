resource "netbox_route_target" "import" {
  name = "65000:1"
}

resource "netbox_route_target" "export" {
  name = "65000:2"
}

resource "netbox_l2vpn" "example" {
  name           = "example-l2vpn"
  slug           = "example-l2vpn"
  type           = "vxlan"
  identifier     = 1234
  import_targets = [netbox_route_target.import.id]
  export_targets = [netbox_route_target.export.id]
}
