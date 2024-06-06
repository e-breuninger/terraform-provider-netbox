resource "netbox_config_template" "test" {
	name = "test"
	description = "test description"
	template_code = "hostname {{ name }}"
	environment_params = jsonencode({"name" = "my-hostname"})
}
