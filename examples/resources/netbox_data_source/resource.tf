resource "netbox_data_source" "test" {
  name          = "test-data-source"
  type          = "git"
  source_url    = "https://example.com/test.git"
  enabled       = true
  ignore_rules  = "*.md"
  sync_interval = 1440
  parameters    = jsonencode({ branch = "main" })
  description   = "test-description"
}
