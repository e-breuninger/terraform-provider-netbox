resource "netbox_export_template" "site_names" {
  name           = "site-names"
  content_types  = ["dcim.site"]
  description    = "Exports a plain list of site names"
  template_code  = "{% for obj in queryset %}{{ obj.name }}\n{% endfor %}"
  mime_type      = "text/plain"
  file_extension = "txt"
  as_attachment  = true
}
