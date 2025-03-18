package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNetboxEventRule_basic(t *testing.T) {
	testName := testAccGetTestName("evt_rule_basic")
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetBoxEventRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name        = "%[1]s"
  payload_url = "https://example.com/webhook"
}

resource "netbox_event_rule" "test" {
  name             = "%[1]s"
  description      = "foo description"
  content_types    = ["dcim.site"]
  action_type      = "webhook"
  action_object_id = netbox_webhook.test.id
  event_types      = ["object_created", "object_updated", "object_deleted", "job_started", "job_completed", "job_failed", "job_errored"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_event_rule.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "content_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "content_types.0", "dcim.site"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "action_type", "webhook"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "description", "foo description"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "event_types.#", "7"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name        = "%[1]s"
  payload_url = "https://example.com/webhook"
}

resource "netbox_event_rule" "test" {
  name             = "%[1]s"
  content_types    = ["dcim.site", "virtualization.cluster"]
  action_type      = "webhook"
  action_object_id = netbox_webhook.test.id
  event_types      = ["object_created"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_event_rule.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "content_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "content_types.0", "dcim.site"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "content_types.1", "virtualization.cluster"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "action_type", "webhook"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "event_types.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_event_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckNetBoxEventRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_event_rule" {
			continue
		}

		// Fetch the eventRule by ID
		// Retrieve our interface by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		eventRule, err := client.Extras.ExtrasEventRulesRead(extras.NewExtrasEventRulesReadParams().WithID(stateID), nil)
		if err == nil && eventRule != nil {
			return fmt.Errorf("EventRule %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func init() {
	resource.AddTestSweepers("netbox_event_rule", &resource.Sweeper{
		Name:         "netbox_event_rule",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := extras.NewExtrasEventRulesListParams()
			res, err := api.Extras.ExtrasEventRulesList(params, nil)
			if err != nil {
				return err
			}
			for _, eventRule := range res.GetPayload().Results {
				if strings.HasPrefix(*eventRule.Name, testPrefix) {
					deleteParams := extras.NewExtrasEventRulesDeleteParams().WithID(eventRule.ID)
					_, err := api.Extras.ExtrasEventRulesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a eventRule")
				}
			}
			return nil
		},
	})
}
