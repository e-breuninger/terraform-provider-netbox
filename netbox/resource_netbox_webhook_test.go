package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetboxWebhook_basic(t *testing.T) {
	testName := testAccGetTestName("webhook_basic")
	testPayloadURL := "https://example.com/webhook"
	testBodyTemplate := "Sample Body"
	testAdditionalHeaders := "Authentication: Bearer abcdef123456"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetBoxWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name               = "%s"
  payload_url        = "%s"
  body_template      = "%s"
  additional_headers = "%s"
}`, testName, testPayloadURL, testBodyTemplate, testAdditionalHeaders),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "body_template", testBodyTemplate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "additional_headers", testAdditionalHeaders),
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

func TestAccNetboxWebhook_update(t *testing.T) {
	testName := testAccGetTestName("webhook_update")
	testPayloadURL := "https://example.com/webhookupdate"
	testBodyTemplate := `{"text": "This is a sample json"}`
	testHTTPMethod := "PUT"
	testHTTPContentType := "application/xml"

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
	name              = "%s"
	payload_url       = "%s"
	body_template        = <<-EOT
	{"text": "This is a sample json"}
	EOT
        http_method       = "%s"
        http_content_type = "%s"
  }`, testName, testPayloadURL, testHTTPMethod, testHTTPContentType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "body_template", testBodyTemplate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_method", testHTTPMethod),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_content_type", testHTTPContentType),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name                 = "%s_updated"
  payload_url          = "%s"
  body_template        = <<-EOT
  {"text": "This is a sample json"}
  EOT
}`, testName, testPayloadURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName+"_updated"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "body_template", testBodyTemplate),
				),
			},
		},
	})
}

func TestAccNetboxWebhook_import(t *testing.T) {
	testName := testAccGetTestName("webhook_import")
	testPayloadURL := "https://test2.com/webhook"

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name                   = "%s"
  payload_url            = "%s"
}`, testName, testPayloadURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
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

func testAccCheckNetBoxWebhookDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_webhook" {
			continue
		}

		// Fetch the webhook by ID
		// Retrieve our interface by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		webhook, err := client.Extras.ExtrasWebhooksRead(extras.NewExtrasWebhooksReadParams().WithID(stateID), nil)
		if err == nil && webhook != nil {
			return fmt.Errorf("Webhook %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func init() {
	resource.AddTestSweepers("netbox_webhook", &resource.Sweeper{
		Name:         "netbox_webhook",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := extras.NewExtrasWebhooksListParams()
			res, err := api.Extras.ExtrasWebhooksList(params, nil)
			if err != nil {
				return err
			}
			for _, webhook := range res.GetPayload().Results {
				if strings.HasPrefix(*webhook.Name, testPrefix) {
					deleteParams := extras.NewExtrasWebhooksDeleteParams().WithID(webhook.ID)
					_, err := api.Extras.ExtrasWebhooksDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a webhook")
				}
			}
			return nil
		},
	})
}
