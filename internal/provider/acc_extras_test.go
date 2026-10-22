//go:build acctest

// Extras: webhooks, event rules, config templates and contexts, custom links, export templates
// and notification groups.
package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccNetboxWebhook_basic(t *testing.T) {
	testName := testAccGetTestName("webhook")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name               = "%[1]s"
  payload_url        = "https://example.com/hooks/%[1]s"
  http_method        = "PUT"
  http_content_type  = "application/x-yaml"
  additional_headers = "X-Test: %[1]s"
  body_template      = "{{ event }}"
  secret             = "hunter2"
  ssl_verification   = false
  description        = "Acceptance test webhook."
}
data "netbox_webhook" "test" {
  name = netbox_webhook.test.name
}
data "netbox_webhooks" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_webhook.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_method", "PUT"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_content_type", "application/x-yaml"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "ssl_verification", "false"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "secret", "hunter2"),
					resource.TestCheckResourceAttrPair("data.netbox_webhook.test", "id", "netbox_webhook.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_webhooks.test", "webhooks.#", "1"),
				),
			},
			{
				// Shrink, and verify against a CA file, which NetBox only allows while verification
				// is on. http_method, http_content_type and ssl_verification are optional+computed and
				// keep their values unless set; the rest clears.
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name             = "%[1]s"
  payload_url      = "https://example.com/hooks/%[1]s"
  ssl_verification = true
  ca_file_path     = "/etc/ssl/certs/ca.pem"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_webhook.test", "http_method", "PUT"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "ssl_verification", "true"),
					resource.TestCheckResourceAttr("netbox_webhook.test", "ca_file_path", "/etc/ssl/certs/ca.pem"),
					resource.TestCheckNoResourceAttr("netbox_webhook.test", "additional_headers"),
					resource.TestCheckNoResourceAttr("netbox_webhook.test", "body_template"),
					resource.TestCheckNoResourceAttr("netbox_webhook.test", "secret"),
					resource.TestCheckNoResourceAttr("netbox_webhook.test", "description"),
				),
			},
			{
				// The CA file clears too while verification stays on.
				Config: fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name             = "%[1]s"
  payload_url      = "https://example.com/hooks/%[1]s"
  ssl_verification = true
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_webhook.test", "ca_file_path"),
			},
			{
				ResourceName:      "netbox_webhook.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxEventRule_basic(t *testing.T) {
	testName := testAccGetTestName("eventrule")
	deps := fmt.Sprintf(`
resource "netbox_webhook" "test" {
  name        = "%[1]s"
  payload_url = "https://example.com/hooks/%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_event_rule" "test" {
  name               = "%[1]s"
  object_types       = ["dcim.site", "dcim.device"]
  event_types        = ["object_created", "object_updated"]
  enabled            = false
  conditions         = jsonencode({ attr = "status", value = "active" })
  action_type        = "webhook"
  action_object_type = "extras.webhook"
  action_object_id   = netbox_webhook.test.id
  description        = "Acceptance test event rule."
}
data "netbox_event_rule" "test" {
  name = netbox_event_rule.test.name
}
data "netbox_event_rules" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_event_rule.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_event_rule.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "object_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "event_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "action_type", "webhook"),
					resource.TestCheckResourceAttrPair("netbox_event_rule.test", "action_object_id", "netbox_webhook.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_event_rule.test", "id", "netbox_event_rule.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_event_rules.test", "event_rules.#", "1"),
				),
			},
			{
				// Shrink: enabled is optional+computed and keeps NetBox's value; the rest clears.
				Config: deps + fmt.Sprintf(`
resource "netbox_event_rule" "test" {
  name               = "%[1]s"
  object_types       = ["dcim.site"]
  event_types        = ["object_deleted"]
  action_type        = "webhook"
  action_object_type = "extras.webhook"
  action_object_id   = netbox_webhook.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_event_rule.test", "object_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_event_rule.test", "event_types.#", "1"),
					resource.TestCheckNoResourceAttr("netbox_event_rule.test", "conditions"),
					resource.TestCheckNoResourceAttr("netbox_event_rule.test", "description"),
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

func TestAccNetboxConfigTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("cfgtemplate")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_config_template" "test" {
  name               = "%[1]s"
  description        = "Acceptance test config template."
  template_code      = "hostname {{ device.name }}"
  environment_params = jsonencode({ trim_blocks = true })
  mime_type          = "text/x-config"
  file_name          = "startup"
  file_extension     = "cfg"
  as_attachment      = true
  debug              = true
}
data "netbox_config_template" "test" {
  name = netbox_config_template.test.name
}
data "netbox_config_templates" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_config_template.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_config_template.test", "template_code", "hostname {{ device.name }}"),
					resource.TestCheckResourceAttr("netbox_config_template.test", "mime_type", "text/x-config"),
					resource.TestCheckResourceAttr("netbox_config_template.test", "as_attachment", "true"),
					resource.TestCheckResourceAttr("netbox_config_template.test", "debug", "true"),
					resource.TestCheckResourceAttrPair("data.netbox_config_template.test", "id", "netbox_config_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_config_templates.test", "config_templates.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_config_template" "test" {
  name          = "%[1]s"
  template_code = "hostname {{ device.name }}"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_config_template.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_config_template.test", "environment_params"),
					resource.TestCheckNoResourceAttr("netbox_config_template.test", "mime_type"),
					resource.TestCheckNoResourceAttr("netbox_config_template.test", "file_name"),
					resource.TestCheckNoResourceAttr("netbox_config_template.test", "file_extension"),
				),
			},
			{
				ResourceName:      "netbox_config_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxConfigContext_basic(t *testing.T) {
	testName := testAccGetTestName("cfgcontext")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_tenant_group" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
resource "netbox_platform" "test" {
  name = "%[1]s"
}
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster_group" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_config_context_profile" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name              = "%[1]s"
  weight            = 500
  description       = "Acceptance test config context."
  is_active         = false
  data              = jsonencode({ ntp_servers = ["10.0.0.1", "10.0.0.2"] })
  profile_id        = netbox_config_context_profile.test.id
  site_ids          = [netbox_site.test.id]
  site_group_ids    = [netbox_site_group.test.id]
  region_ids        = [netbox_region.test.id]
  location_ids      = [netbox_location.test.id]
  tenant_ids        = [netbox_tenant.test.id]
  tenant_group_ids  = [netbox_tenant_group.test.id]
  device_type_ids   = [netbox_device_type.test.id]
  device_role_ids   = [netbox_device_role.test.id]
  platform_ids      = [netbox_platform.test.id]
  cluster_type_ids  = [netbox_cluster_type.test.id]
  cluster_group_ids = [netbox_cluster_group.test.id]
  cluster_ids       = [netbox_cluster.test.id]
  tag_slugs         = [netbox_tag.test.slug]
}
data "netbox_config_context" "test" {
  name = netbox_config_context.test.name
}
data "netbox_config_contexts" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_config_context.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_context.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_config_context.test", "weight", "500"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "is_active", "false"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "site_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "site_group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "region_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "location_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "tenant_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "tenant_group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "device_type_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "device_role_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "platform_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "cluster_type_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "cluster_group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "cluster_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_context.test", "tag_slugs.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_config_context.test", "profile_id", "netbox_config_context_profile.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_config_context.test", "id", "netbox_config_context.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_config_contexts.test", "config_contexts.#", "1"),
				),
			},
			{
				// data is normalize json: the state keeps the configured text even though NetBox
				// echoes the document compact, with sorted keys and "<" escaped.
				Config: deps + fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name = "%[1]s"
  data = <<EOT
%[2]sEOT
}`, testName, prettyJSON),
				Check: resource.TestCheckResourceAttr("netbox_config_context.test", "data", prettyJSON),
			},
			{
				// ... and the next plan is empty: semantic equality, no drift.
				Config: deps + fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name = "%[1]s"
  data = <<EOT
%[2]sEOT
}`, testName, prettyJSON),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{
				// Text that is not JSON fails at plan.
				Config: deps + fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name = "%[1]s"
  data = "{not json"
}`, testName),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("(?i)invalid json"),
			},
			{
				// Shrink: weight and is_active are optional+computed and keep NetBox's values; the
				// scopes and the profile clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name = "%[1]s"
  data = jsonencode({ ntp_servers = ["10.0.0.1"] })
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "profile_id"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "site_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "site_group_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "region_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "location_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "tenant_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "tenant_group_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "device_type_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "device_role_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "platform_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "cluster_type_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "cluster_group_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "cluster_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_config_context.test", "tag_slugs.#"),
				),
			},
			{
				ResourceName:      "netbox_config_context.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxCustomLink_basic(t *testing.T) {
	testName := testAccGetTestName("customlink")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_link" "test" {
  name         = "%[1]s"
  object_types = ["dcim.device", "dcim.site"]
  link_text    = "Open {{ object.name }}"
  link_url     = "https://example.com/%[1]s/{{ object.pk }}"
  enabled      = false
  weight       = 50
  group_name   = "%[1]s"
  button_class = "blue"
  new_window   = true
}
data "netbox_custom_link" "test" {
  name = netbox_custom_link.test.name
}
data "netbox_custom_links" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_custom_link.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_link.test", "object_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "weight", "50"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "button_class", "blue"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "new_window", "true"),
					resource.TestCheckResourceAttrPair("data.netbox_custom_link.test", "id", "netbox_custom_link.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_custom_links.test", "custom_links.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_custom_link.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: enabled, weight and button_class are optional+computed and keep NetBox's
				// values; new_window falls back to its default and group_name clears.
				Config: fmt.Sprintf(`
resource "netbox_custom_link" "test" {
  name         = "%[1]s"
  object_types = ["dcim.device"]
  link_text    = "Open {{ object.name }}"
  link_url     = "https://example.com/%[1]s/{{ object.pk }}"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_link.test", "object_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "weight", "50"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "button_class", "blue"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "new_window", "false"),
					resource.TestCheckNoResourceAttr("netbox_custom_link.test", "group_name"),
				),
			},
		},
	})
}

func TestAccNetboxExportTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("exporttmpl")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_export_template" "test" {
  name               = "%[1]s"
  object_types       = ["dcim.device"]
  template_code      = "{%% for obj in queryset %%}{{ obj.name }}\n{%% endfor %%}"
  mime_type          = "text/csv"
  file_name          = "%[1]s"
  file_extension     = "csv"
  as_attachment      = false
  environment_params = jsonencode({ trim_blocks = true })
  description        = "Acceptance test export template."
}
data "netbox_export_template" "test" {
  name = netbox_export_template.test.name
}
data "netbox_export_templates" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_export_template.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_export_template.test", "mime_type", "text/csv"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "file_extension", "csv"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "as_attachment", "false"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "environment_params", `{"trim_blocks":true}`),
					resource.TestCheckResourceAttrPair("data.netbox_export_template.test", "id", "netbox_export_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_export_templates.test", "export_templates.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_export_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: as_attachment is optional+computed and keeps its value; the rest clears.
				Config: fmt.Sprintf(`
resource "netbox_export_template" "test" {
  name          = "%[1]s"
  object_types  = ["dcim.device"]
  template_code = "{{ queryset | length }}"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_export_template.test", "as_attachment", "false"),
					resource.TestCheckNoResourceAttr("netbox_export_template.test", "mime_type"),
					resource.TestCheckNoResourceAttr("netbox_export_template.test", "file_name"),
					resource.TestCheckNoResourceAttr("netbox_export_template.test", "file_extension"),
					resource.TestCheckNoResourceAttr("netbox_export_template.test", "environment_params"),
					resource.TestCheckNoResourceAttr("netbox_export_template.test", "description"),
				),
			},
		},
	})
}

// TestAccNetboxNotificationGroup_basic: notification groups have no list filters in NetBox, hence
// no data source.
func TestAccNetboxNotificationGroup_basic(t *testing.T) {
	testName := testAccGetTestName("notifgroup")
	deps := fmt.Sprintf(`
resource "netbox_group" "test" {
  name = "%[1]s"
}
resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Acceptance-Test-Pw-9271"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_notification_group" "test" {
  name        = "%[1]s"
  description = "Acceptance test notification group."
  group_ids   = [netbox_group.test.id]
  user_ids    = [netbox_user.test.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_notification_group.test", "group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_notification_group.test", "user_ids.#", "1"),
					resource.TestCheckTypeSetElemAttrPair("netbox_notification_group.test", "user_ids.*", "netbox_user.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_notification_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_notification_group" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_notification_group.test", "group_ids"),
					resource.TestCheckNoResourceAttr("netbox_notification_group.test", "user_ids"),
					resource.TestCheckNoResourceAttr("netbox_notification_group.test", "description"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_custom_link",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasCustomLinksList(extras.NewExtrasCustomLinksListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasCustomLinksDestroy(extras.NewExtrasCustomLinksDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_export_template",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasExportTemplatesList(extras.NewExtrasExportTemplatesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasExportTemplatesDestroy(extras.NewExtrasExportTemplatesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_notification_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasNotificationGroupsList(extras.NewExtrasNotificationGroupsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasNotificationGroupsDestroy(extras.NewExtrasNotificationGroupsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_webhook",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasWebhooksList(extras.NewExtrasWebhooksListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasWebhooksDestroy(extras.NewExtrasWebhooksDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_event_rule",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasEventRulesList(extras.NewExtrasEventRulesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasEventRulesDestroy(extras.NewExtrasEventRulesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_config_template",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasConfigTemplatesList(extras.NewExtrasConfigTemplatesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasConfigTemplatesDestroy(extras.NewExtrasConfigTemplatesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_config_context",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Extras.ExtrasConfigContextsList(extras.NewExtrasConfigContextsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Extras.ExtrasConfigContextsDestroy(extras.NewExtrasConfigContextsDestroyParams().WithID(id), nil)
			return err
		})
}

// prettyJSON is a JSON document as a person writes it: pretty-printed, keys unsorted, a character
// Go's encoder would escape. NetBox echoes it compact, sorted and escaped; normalize json keeps
// this spelling in state. It ends in a newline, as a heredoc does.
const prettyJSON = `{
  "zebra": {"note": "a<b"},
  "alpha": [1, 2]
}
`
