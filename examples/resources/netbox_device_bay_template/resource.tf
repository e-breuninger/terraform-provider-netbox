resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
	subdevice_role = "parent"
}

resource "netbox_device_bay_template" "test" {
	name = "%[1]s"
	device_type_id = netbox_device_type.test.id
}
