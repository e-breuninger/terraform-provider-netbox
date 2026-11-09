# Public provider package

This package makes the NetBox provider importable as a Go library.

## Why this folder exists

The provider is implemented in `internal/provider`, and Go doesn't let other modules import `internal` packages.
Other tools that want to serve the provider in-process instead of launching it as a plugin through the Terraform CLI need this.

## Usage

```go
import (
 "github.com/hashicorp/terraform-plugin-framework/providerserver"

 "github.com/e-breuninger/terraform-provider-netbox/provider"
)

// version is reported as the provider version.
server, err := providerserver.NewProtocol6WithError(provider.New(version)())()
```
