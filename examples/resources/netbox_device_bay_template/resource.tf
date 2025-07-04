resource "netbox_manufacturer" "example" {
	name = "example_manufacturer"
}

resource "netbox_device_type" "example" {
	model = "example_device_model"
	slug = "example_device_slug"
	part_number = "example_part_number"
	manufacturer_id = netbox_manufacturer.example.id
	subdevice_role = "parent"
}

resource "netbox_device_bay_template" "example" {
	name = "example_device_bay_template"
	device_type_id = netbox_device_type.example.id
}
