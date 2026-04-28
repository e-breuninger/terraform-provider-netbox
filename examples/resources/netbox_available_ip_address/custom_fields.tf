# custom_fields can be set, edited, or cleared on netbox_available_ip_address.
# Removing the entire block (or setting it to {}) clears every CF on the IP.
data "netbox_prefix" "leases" {
  cidr = "10.0.0.0/24"
}

resource "netbox_available_ip_address" "leased" {
  prefix_id = data.netbox_prefix.leases.id
  status    = "active"

  custom_fields = {
    purpose      = "edge-gateway"
    last_audited = "2026-04-27"
  }
}
