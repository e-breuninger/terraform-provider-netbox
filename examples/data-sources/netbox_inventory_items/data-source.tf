data "netbox_inventory_items" "test" {
  filters = [
    { name = "name", value = "test-inventory_item" },
  ]
}
