package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"testing"
)

func TestAccNetboxWebhookDatasource_basic(t *testing.T) {
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
			//Read test
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name               = "%s"
  payload_url        = "%s"
  body_template      = "%s"
  additional_headers = "%s"
}
data "netbox_webhook" "test" {
id = netbox_webhook.test.id
}
`, testName, testPayloadURL, testBodyTemplate, testAdditionalHeaders),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.netbox_webhook.test", tfjsonpath.New("name"), knownvalue.StringExact(testName)),
					statecheck.ExpectKnownValue("data.netbox_webhook.test", tfjsonpath.New("payload_url"), knownvalue.StringExact(testPayloadURL)),
					statecheck.ExpectKnownValue("data.netbox_webhook.test", tfjsonpath.New("body_template"), knownvalue.StringExact(testBodyTemplate)),
					statecheck.ExpectKnownValue("data.netbox_webhook.test", tfjsonpath.New("additional_headers"), knownvalue.StringExact(testAdditionalHeaders)),
					statecheck.ExpectKnownValue("data.netbox_webhook.test", tfjsonpath.New("http_content_type"), knownvalue.StringExact("application/json")),
					statecheck.ExpectKnownValue("data.netbox_webhook.test", tfjsonpath.New("http_method"), knownvalue.StringExact("POST")),
				},
			},
		},
	})
}
