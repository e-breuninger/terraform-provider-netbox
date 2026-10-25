resource "netbox_group" "network_admins" {
  name = "network-admins"
}

resource "netbox_user" "jdoe" {
  username = "jdoe"
  password = "correct-horse-battery-staple"
  active   = true
}

resource "netbox_notification_group" "network_alerts" {
  name        = "network-alerts"
  description = "Notified about changes to network infrastructure"
  group_ids   = [netbox_group.network_admins.id]
  user_ids    = [netbox_user.jdoe.id]
}
