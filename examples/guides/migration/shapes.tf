# Was vid_ranges = [[100, 199], [300, 399]].
resource "netbox_vlan_group" "office" {
  name = "office"
  vid_ranges = [
    { start = 100, end = 199 },
    { start = 300, end = 399 },
  ]
}

# Was extra_choices = [["prod", "Production"], ["dev", "Development"]].
resource "netbox_custom_field_choice_set" "environment" {
  name = "environment"
  extra_choices = [
    { value = "prod", label = "Production" },
    { value = "dev", label = "Development" },
  ]
}

# Was content_types.
resource "netbox_custom_field" "environment" {
  name          = "environment"
  type          = "select"
  choice_set_id = netbox_custom_field_choice_set.environment.id
  object_types  = ["dcim.device"]
}
