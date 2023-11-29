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
	testEnabled := "true"
	triggerOnCreate := "true"
	testPayloadURL := "https://example.com/webhook"
	testContentType := "dcim.site"
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
  enabled            = "%s"
  trigger_on_create  = "%s"
  payload_url        = "%s"
  content_types      = ["%s"]
  body_template      = "%s"
  additional_headers = "%s"
}`, testName, testEnabled, triggerOnCreate, testPayloadURL, testContentType, testBodyTemplate, testAdditionalHeaders),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_create", triggerOnCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "content_types.#", "1"),
					resource.TestCheckTypeSetElemAttr("netbox_webhook.test", "content_types.*", testContentType),
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
	testEnabled := "true"
	testPayloadURL := "https://example.com/webhookupdate"
	triggerOnCreate := "true"
	triggerOnUpdate := "true"
	triggerOnDelete := "true"
	testContentType := "dcim.site"
	testContentType1 := "dcim.cable"
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
	enabled           = "%s"
	trigger_on_create = "%s"
	trigger_on_update = "%s"
	trigger_on_delete = "%s"
	payload_url       = "%s"
	content_types     = ["%s"]
	body_template     = "{\"text\": \"This is a sample json\"}"
        http_method       = "%s"
        http_content_type = "%s"
  }`, testName, testEnabled, triggerOnCreate, triggerOnUpdate, triggerOnDelete, testPayloadURL, testContentType, testHTTPMethod, testHTTPContentType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_create", triggerOnCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_update", triggerOnUpdate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_delete", triggerOnDelete),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "content_types.#", "1"),
					resource.TestCheckTypeSetElemAttr("netbox_webhook.test", "content_types.*", testContentType),
					resource.TestCheckResourceAttr("netbox_webhook.test", "body_template", testBodyTemplate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_method", testHTTPMethod),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_content_type", testHTTPContentType),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name                 = "%s_updated"
  enabled              = "%s"
  trigger_on_create    = "%s"
  trigger_on_update    = "%s"
  trigger_on_delete    = "%s"
  payload_url          = "%s"
  content_types        = ["%s", "%s"]
  body_template        = "{\"text\": \"This is a sample json\"}"
}`, testName, testEnabled, triggerOnCreate, triggerOnUpdate, triggerOnDelete, testPayloadURL, testContentType, testContentType1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName+"_updated"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_create", triggerOnCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_update", triggerOnUpdate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_delete", triggerOnDelete),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "content_types.#", "2"),
					resource.TestCheckTypeSetElemAttr("netbox_webhook.test", "content_types.*", testContentType),
					resource.TestCheckTypeSetElemAttr("netbox_webhook.test", "content_types.*", testContentType1),
					resource.TestCheckResourceAttr("netbox_webhook.test", "body_template", testBodyTemplate),
				),
			},
		},
	})
}

func TestAccNetboxWebhook_import(t *testing.T) {
	testName := testAccGetTestName("webhook_import")
	triggerOnCreate := "true"
	triggerOnUpdate := "false"
	triggerOnDelete := "false"
	testEnabled := "true"
	testPayloadURL := "https://test2.com/webhook"
	testContentType := "dcim.site"

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name                   = "%s"
  enabled                = "%s"
  trigger_on_create 	 = "%s"
  trigger_on_update 	 = "%s"
  trigger_on_delete 	 = "%s"
  payload_url            = "%s"
  content_types          = ["%s"]
}`, testName, testEnabled, triggerOnCreate, triggerOnUpdate, triggerOnDelete, testPayloadURL, testContentType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_create", triggerOnCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_update", triggerOnUpdate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "trigger_on_delete", triggerOnDelete),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "content_types.#", "1"),
					resource.TestCheckTypeSetElemAttr("netbox_webhook.test", "content_types.*", testContentType),
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
