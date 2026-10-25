// Get a notification group by name
data "netbox_notification_group" "by_name" {
  name = "network-alerts"
}

// Get a notification group by ID
data "netbox_notification_group" "by_id" {
  id = "1"
}
