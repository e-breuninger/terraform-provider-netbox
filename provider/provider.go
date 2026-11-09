// Package provider exposes the NetBox Terraform provider for use as a Go
// library, e.g. by tools that serve it in-process instead of through the
// Terraform CLI (such as Crossplane providers generated with upjet).
//
// The implementation lives in internal/provider, which Go does not allow other
// modules to import; this package only re-exports its constructor.
package provider

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	internalprovider "github.com/e-breuninger/terraform-provider-netbox/internal/provider"
)

// New returns a constructor for the NetBox provider, as passed to
// providerserver.Serve by main.go.
func New(version string) func() fwprovider.Provider {
	return internalprovider.New(version)
}
