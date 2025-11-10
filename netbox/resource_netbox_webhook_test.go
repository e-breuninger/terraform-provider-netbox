package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxWebhook_basic(t *testing.T) {
	testName := testAccGetTestName("webhook_basic")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name         = "%s"
  payload_url  = "https://example.com/webhook"
  http_method  = "POST"
  http_content_type = "application/json"
  body_template = "{\"event\": \"{{ event }}\", \"data\": {{ data | tojson }}}"
  additional_headers = "X-Custom-Header: test-value"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", "https://example.com/webhook"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_method", "POST"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_content_type", "application/json"),
				),
			},
			{
				ResourceName:      "netbox_webhook.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxWebhook_minimal(t *testing.T) {
	testName := testAccGetTestName("webhook_minimal")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name        = "%s"
  payload_url = "https://example.com/hook"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", "https://example.com/hook"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_method", "POST"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_content_type", "application/json"),
				),
			},
		},
	})
}

func TestAccNetboxWebhook_withGETMethod(t *testing.T) {
	testName := testAccGetTestName("webhook_get")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name        = "%s"
  payload_url = "https://example.com/get-hook"
  http_method = "GET"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_method", "GET"),
				),
			},
		},
	})
}
