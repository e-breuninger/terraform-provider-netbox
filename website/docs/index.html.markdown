---
layout: "netbox"
page_title: "Provider: Netbox"
sidebar_current: "docs-netbox-index"
description: |-
  Terraform provider for managing Netbox resources.
---

# Netbox Provider

The Terraform Netbox provider is a plugin for Terraform that allows for the full lifecycle management of [Netbox](https://netbox.readthedocs.io/en/stable/) resources.

## Example Usage

```hcl
provider "netbox" {
    server_url           = var.netbox_server
    api_token            = var.netbox_api_token
    allow_insecure_https = false
}

resource "netbox_platform" "testplatform" {
    name = "my-test-platform"
}

resource "netbox_cluster_type" "testclustertype" {
    name = "my-test-cluster-type"
}
```

## Argument Reference

The following arguments are required:

- `server_url` - (Required) The URL to the Netbox endpoint.
- `api_token` - (Required) The API token to authenticate with Netbox.
