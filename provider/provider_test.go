package provider_test

import (
	"context"
	"testing"

	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/e-breuninger/terraform-provider-netbox/provider"
)

// TestNew checks that the public constructor returns the NetBox provider and
// that it can be served in-process, without the Terraform CLI.
func TestNew(t *testing.T) {
	ctx := context.Background()
	p := provider.New("1.2.3")()

	var meta fwprovider.MetadataResponse
	p.Metadata(ctx, fwprovider.MetadataRequest{}, &meta)
	if meta.TypeName != "netbox" {
		t.Errorf("TypeName: want netbox, got %q", meta.TypeName)
	}
	if meta.Version != "1.2.3" {
		t.Errorf("Version: want 1.2.3, got %q", meta.Version)
	}

	server, err := providerserver.NewProtocol6WithError(p)()
	if err != nil {
		t.Fatalf("NewProtocol6WithError: %v", err)
	}
	schema, err := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %v", err)
	}
	for _, d := range schema.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Errorf("GetProviderSchema: %s: %s", d.Summary, d.Detail)
		}
	}
	if _, ok := schema.ResourceSchemas["netbox_device"]; !ok {
		t.Errorf("GetProviderSchema: netbox_device missing from %d resource schemas", len(schema.ResourceSchemas))
	}
}
