terraform {
  required_providers {
    netbox = {
      source  = "e-breuninger/netbox"
      version = "~> 6.0"
    }
  }
}

provider "netbox" {
  server_url = "https://netbox.example.com"
  api_token  = "<your api token>"
}
