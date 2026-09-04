# This example attaches a tag to a device that already exists in NetBox but
# is not otherwise managed by Terraform (e.g. it was onboarded automatically).

resource "netbox_tag" "managed_by_terraform" {
  name = "managed-by-terraform"
}

resource "netbox_device_tag" "example" {
  device_id = 123
  tag       = netbox_tag.managed_by_terraform.name
}
