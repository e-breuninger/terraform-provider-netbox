resource "netbox_data_source" "example" {
  name         = "example-git-source"
  type         = "git"
  source_url   = "https://github.com/example/netbox-data.git"
  enabled      = true
  description  = "Example git-backed data source"
  sync_interval = 1440
  ignore_rules = ".git\n.gitignore"
}
