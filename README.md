# terraform-provider-netbox

The Terraform Netbox provider is a plugin for Terraform that allows for the full lifecycle management of [Netbox](https://netbox.readthedocs.io/en/stable/) resources.
This provider is maintained by E. Breuninger.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
- [Go](https://golang.org/doc/install) >= 1.14

## Supported netbox versions
Netbox often makes API-breaking changes even in non-major releases. We aim to always support the latest minor version of Netbox. Check the table below to see which version a provider was tested against. It is generally recommended to use the provider version matching your netbox version.

Provider version | Netbox version
--- | ---
v1.1.x and up | v3.1.3
v1.0.x | v3.0.9
v0.3.x | v2.11.12
v0.2.x | v2.10.10
v0.1.x | v2.9
v0.0.x | v2.9

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

In order to run the full suite of Acceptance tests, run `make docker-up testacc`.

_Note:_ Acceptance tests create a docker compose stack on port 8001.

```sh
$ make testacc
```
If you notice a failed test, it might be due to a stale netbox data volume.  Before concluding there is a problem, 
refresh the docker containers and try again:
```shell
docker rm -f docker_netbox_1 ; docker rm -f  docker_postgres_1 ; docker rm -f docker_redis_1
make testacc
```

## Contribution

We focus on virtual machine management and IPAM. If you want to contribute more resources to this provider feel free to make a PR.
