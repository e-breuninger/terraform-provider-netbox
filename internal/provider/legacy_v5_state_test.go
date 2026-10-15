package provider

import "testing"

func TestIsNetboxProviderAddress(t *testing.T) {
	for address, want := range map[string]bool{
		"registry.terraform.io/e-breuninger/netbox": true, // Terraform
		"netbox":                               true, // OpenTofu
		"registry.terraform.io/hashicorp/null": false,
		"netboxx":                              false,
		"":                                     false,
	} {
		if got := isNetboxProviderAddress(address); got != want {
			t.Errorf("isNetboxProviderAddress(%q) = %v, want %v", address, got, want)
		}
	}
}
