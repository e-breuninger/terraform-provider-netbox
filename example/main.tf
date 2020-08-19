terraform {
  required_providers {
    netbox = {
      source = "e-breuninger/netbox"
      version = "0.0.1"
    }
  }
}

provider "netbox" {
    server_url           = "https://netboxdemo.com/"
    api_token            = " 72830d67beff4ae178b94d8f781842408df8069d"
}

resource "netbox_platform" "testplatform" {
    name = "my-test-platform"
}