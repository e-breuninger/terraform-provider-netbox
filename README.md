# terraform-provider-netbox

[![Dependabot Status](https://api.dependabot.com/badges/status?host=github&repo=e-breuninger/terraform-provider-netbox)](https://dependabot.com)


The Terraform Netbox provider is a plugin for Terraform that allows for the full lifecycle management of [Netbox](https://netbox.readthedocs.io/en/stable/) resources.
This provider is maintained by E. Breuninger.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
- [Go](https://golang.org/doc/install) >= 1.14

## Supported netbox versions
Netbox often makes API-breaking changes even in non-major releases. We aim to always support the latest minor version of Netbox.

Version 0.0.4 of this provider is tested against Netbox 2.9.7.

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```sh
$ go install
```

## Installation

When using Terraform 0.13 you can use the provider from the Terraform registry.

For further information on how to use third party providers, see https://www.terraform.io/docs/configuration/providers.html

Releases for all major plattform are available on the release page.

## Using the provider

Here is a short example on how to use this provider:

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

For a more complex example, see the `example` folder.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

In order to run the suite of unit tests, run `make test`.

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create a docker compose stack on port 8001.

```sh
$ make testacc
```

## Contribution

We focus on virtual machine management and IPAM. If you want to contribute more resources to this provider feel free to make a PR.
