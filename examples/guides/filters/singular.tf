data "netbox_site" "berlin" {
  name = "Berlin"
}

# The one device named "core-01" at the site; zero or several matches fail the read.
data "netbox_device" "core_01" {
  name    = "core-01"
  site_id = data.netbox_site.berlin.id
}
