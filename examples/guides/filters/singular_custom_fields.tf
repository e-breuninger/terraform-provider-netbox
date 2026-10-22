# The one device whose custom field "asset_owner" matches "network-team".
data "netbox_device" "network_team" {
  custom_field_filters = {
    asset_owner = "network-team"
  }
}
