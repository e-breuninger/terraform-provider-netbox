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

func TestAccNetBoxWebhook_basic(t *testing.T) {
	testName := testAccGetTestName("webhook_basic")
	testEnabled := "true"
	testCreate := "true"
	testPayloadURL := "https://example.com/webhook"
	testContentType := "dcim.site"
	testBodyTemplate := "Sample Body"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name           = "%s"
  enabled           = "%s"
  type_create    = "%s"
  payload_url    = "%s"
  content_types  = ["%s"]
  body_template = "%s"
}`, testName, testEnabled, testCreate, testPayloadURL, testContentType, testBodyTemplate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_create", testCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "content_types.#", "1"),
					resource.TestCheckTypeSetElemAttr("netbox_webhook.test", "content_types.*", testContentType),
					resource.TestCheckResourceAttr("netbox_webhook.test", "body_template", testBodyTemplate),
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
	testCreate := "true"
	testUpdate := "true"
	testDelete := "true"
	testContentType := "dcim.site"
	testContentType1 := "dcim.cable"
	testBodyTemplate := `{"text": "This is a sample json"}`

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
	name           = "%s"
	enabled        = "%s"
	type_create    = "%s"
	type_update    = "%s"
	type_delete    = "%s"
	payload_url    = "%s"
	content_types  = ["%s"]
	body_template = "{\"text\": \"This is a sample json\"}"
  }`, testName, testEnabled, testCreate, testUpdate, testDelete, testPayloadURL, testContentType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_create", testCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_update", testUpdate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_delete", testDelete),
					resource.TestCheckResourceAttr("netbox_webhook.test", "payload_url", testPayloadURL),
					resource.TestCheckResourceAttr("netbox_webhook.test", "content_types.#", "1"),
					resource.TestCheckTypeSetElemAttr("netbox_webhook.test", "content_types.*", testContentType),
					resource.TestCheckResourceAttr("netbox_webhook.test", "body_template", testBodyTemplate),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name           = "%s_updated"
  enabled        = "%s"
  type_create    = "%s"
  type_update    = "%s"
  type_delete    = "%s"
  payload_url    = "%s"
  content_types  = ["%s", "%s"]
  body_template = "{\"text\": \"This is a sample json\"}"
}`, testName, testEnabled, testCreate, testUpdate, testDelete, testPayloadURL, testContentType, testContentType1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName+"_updated"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_create", testCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_update", testUpdate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_delete", testDelete),
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
	testCreate := "true"
	testUpdate := "false"
	testDelete := "false"
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
  name           = "%s"
  enabled        = "%s"
  type_create 	 = "%s"
  type_update 	 = "%s"
  type_delete 	 = "%s"
  payload_url    = "%s"
  content_types  = ["%s"]
}`, testName, testEnabled, testCreate, testUpdate, testDelete, testPayloadURL, testContentType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "enabled", testEnabled),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_create", testCreate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_update", testUpdate),
					resource.TestCheckResourceAttr("netbox_webhook.test", "type_delete", testDelete),
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
