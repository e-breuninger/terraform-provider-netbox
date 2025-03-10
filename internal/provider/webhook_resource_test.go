package provider

import (
	"context"
	"fmt"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"testing"
)

func TestWebhookResource_Schema(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	// Instantiate the resource.Resource and call its Schema method
	NewWebhookResource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	// Validate the schema
	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

// TODO: Destroy in the TestCase setup
// TODO: Test All fields
func TestAccNetboxWebhook_basic(t *testing.T) {
	testName := testAccGetTestName("webhook_basic")
	testPayloadURL := "https://example.com/webhook"
	testBodyTemplate := "Sample Body"
	testAdditionalHeaders := "Authentication: Bearer abcdef123456"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Test creating basic object.
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name               = "%s"
  payload_url        = "%s"
  body_template      = "%s"
  additional_headers = "%s"
}`, testName, testPayloadURL, testBodyTemplate, testAdditionalHeaders),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("name"), knownvalue.StringExact(testName)),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("payload_url"), knownvalue.StringExact(testPayloadURL)),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("body_template"), knownvalue.StringExact(testBodyTemplate)),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("additional_headers"), knownvalue.StringExact(testAdditionalHeaders)),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("http_content_type"), knownvalue.StringExact("application/json")),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("http_method"), knownvalue.StringExact("POST")),
				},
			},
			//Test importing
			{
				ResourceName:      "netbox_webhook.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			//Test updating
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name               = "%s_updated"
  payload_url        = "%s_updated"
  body_template      = "%s_updated"
  additional_headers = "%s_updated"
}`, testName, testPayloadURL, testBodyTemplate, testAdditionalHeaders),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("name"), knownvalue.StringExact(testName+"_updated")),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("payload_url"), knownvalue.StringExact(testPayloadURL+"_updated")),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("body_template"), knownvalue.StringExact(testBodyTemplate+"_updated")),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("additional_headers"), knownvalue.StringExact(testAdditionalHeaders+"_updated")),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("http_content_type"), knownvalue.StringExact("application/json")),
					statecheck.ExpectKnownValue("netbox_webhook.test", tfjsonpath.New("http_method"), knownvalue.StringExact("POST")),
				},
			},
		},
	})
}
