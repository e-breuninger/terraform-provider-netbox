resource "netbox_tag" "monitored" {
  name = "Monitored"
  slug = "monitored"
}

# Tags are referenced by slug.
resource "netbox_tenant" "acme" {
  name = "acme"
  tags = [netbox_tag.monitored.slug]
}
