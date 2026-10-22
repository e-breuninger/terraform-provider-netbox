resource "netbox_export_template" "test" {
  name           = "test-export-template"
  object_types   = ["dcim.device"]
  template_code  = "{% for obj in queryset %}{{ obj.name }}\n{% endfor %}"
  mime_type      = "text/csv"
  file_name      = "test-file"
  file_extension = "csv"
  as_attachment  = true
  description    = "test-description"
}
