data "netbox_inventory_item_roles" "test" {
  filters = [
    { name = "name", value = "test-inventory_item_role" },
  ]
}
