data "netbox_inventory_item_templates" "test" {
  filters = [
    { name = "name", value = "test-inventory_item_template" },
  ]
}
