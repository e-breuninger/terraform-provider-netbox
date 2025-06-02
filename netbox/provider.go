package netbox

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/status"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type providerState struct {
	*client.NetBoxAPI
	defaultTags *schema.Set

	// concurrent access ok, only populated on provider start
	tagCache map[string]*models.NestedTag
}

// This makes the description contain the default value, particularly useful for the docs
// From https://github.com/hashicorp/terraform-plugin-docs/issues/65#issuecomment-1152842370
func init() {
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		desc = strings.TrimSpace(desc)

		if !bytes.HasSuffix([]byte(s.Description), []byte(".")) && s.Description != "" {
			desc += "."
		}

		if len(s.AtLeastOneOf) > 0 {
			atLeastOneOf := make([]string, len(s.AtLeastOneOf))
			for i, l := range s.AtLeastOneOf {
				atLeastOneOf[i] = fmt.Sprintf("`%s`", l)
			}
			desc += fmt.Sprintf(" At least one of %s must be given.", joinStringWithFinalConjunction(atLeastOneOf, ", ", "or"))
		}

		if len(s.ExactlyOneOf) > 0 {
			exactlyOneOf := make([]string, len(s.ExactlyOneOf))
			for i, l := range s.ExactlyOneOf {
				exactlyOneOf[i] = fmt.Sprintf("`%s`", l)
			}
			desc += fmt.Sprintf(" Exactly one of %s must be given.", joinStringWithFinalConjunction(exactlyOneOf, ", ", "or"))
		}

		if len(s.RequiredWith) > 0 {
			requires := make([]string, len(s.RequiredWith))
			for i, c := range s.RequiredWith {
				requires[i] = fmt.Sprintf("`%s`", c)
			}
			desc += fmt.Sprintf(" Required when %s is set.", joinStringWithFinalConjunction(requires, ", ", "and"))
		}

		if len(s.ConflictsWith) > 0 {
			conflicts := make([]string, len(s.ConflictsWith))
			for i, c := range s.ConflictsWith {
				conflicts[i] = fmt.Sprintf("`%s`", c)
			}
			desc += fmt.Sprintf(" Conflicts with %s.", joinStringWithFinalConjunction(conflicts, ", ", "and"))
		}

		if s.Default != nil {
			if s.Default == "" {
				desc += " Defaults to `\"\"`."
			} else {
				desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
			}
		}

		return strings.TrimSpace(desc)
	}
}

// Provider returns a schema.Provider for Netbox.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"netbox_available_ip_address":       resourceNetboxAvailableIPAddress(),
			"netbox_virtual_machine":            resourceNetboxVirtualMachine(),
			"netbox_cluster_type":               resourceNetboxClusterType(),
			"netbox_cluster":                    resourceNetboxCluster(),
			"netbox_contact":                    resourceNetboxContact(),
			"netbox_contact_group":              resourceNetboxContactGroup(),
			"netbox_contact_assignment":         resourceNetboxContactAssignment(),
			"netbox_contact_role":               resourceNetboxContactRole(),
			"netbox_device":                     resourceNetboxDevice(),
			"netbox_device_interface":           resourceNetboxDeviceInterface(),
			"netbox_device_type":                resourceNetboxDeviceType(),
			"netbox_manufacturer":               resourceNetboxManufacturer(),
			"netbox_tenant":                     resourceNetboxTenant(),
			"netbox_tenant_group":               resourceNetboxTenantGroup(),
			"netbox_vrf":                        resourceNetboxVrf(),
			"netbox_ip_address":                 resourceNetboxIPAddress(),
			"netbox_interface_template":         resourceNetboxInterfaceTemplate(),
			"netbox_interface":                  resourceNetboxInterface(),
			"netbox_service":                    resourceNetboxService(),
			"netbox_platform":                   resourceNetboxPlatform(),
			"netbox_prefix":                     resourceNetboxPrefix(),
			"netbox_available_prefix":           resourceNetboxAvailablePrefix(),
			"netbox_primary_ip":                 resourceNetboxPrimaryIP(),
			"netbox_device_primary_ip":          resourceNetboxDevicePrimaryIP(),
			"netbox_device_role":                resourceNetboxDeviceRole(),
			"netbox_tag":                        resourceNetboxTag(),
			"netbox_cluster_group":              resourceNetboxClusterGroup(),
			"netbox_site":                       resourceNetboxSite(),
			"netbox_vlan":                       resourceNetboxVlan(),
			"netbox_vlan_group":                 resourceNetboxVlanGroup(),
			"netbox_available_vlan":             resourceNetboxAvailableVLAN(),
			"netbox_ipam_role":                  resourceNetboxIpamRole(),
			"netbox_ip_range":                   resourceNetboxIPRange(),
			"netbox_region":                     resourceNetboxRegion(),
			"netbox_aggregate":                  resourceNetboxAggregate(),
			"netbox_rir":                        resourceNetboxRir(),
			"netbox_route_target":               resourceNetboxRouteTarget(),
			"netbox_circuit":                    resourceNetboxCircuit(),
			"netbox_circuit_type":               resourceNetboxCircuitType(),
			"netbox_circuit_provider":           resourceNetboxCircuitProvider(),
			"netbox_circuit_termination":        resourceNetboxCircuitTermination(),
			"netbox_user":                       resourceNetboxUser(),
			"netbox_group":                      resourceNetboxGroup(),
			"netbox_permission":                 resourceNetboxPermission(),
			"netbox_token":                      resourceNetboxToken(),
			"netbox_custom_field":               resourceCustomField(),
			"netbox_asn":                        resourceNetboxAsn(),
			"netbox_location":                   resourceNetboxLocation(),
			"netbox_site_group":                 resourceNetboxSiteGroup(),
			"netbox_rack":                       resourceNetboxRack(),
			"netbox_rack_type":                  resourceNetboxRackType(),
			"netbox_rack_role":                  resourceNetboxRackRole(),
			"netbox_rack_reservation":           resourceNetboxRackReservation(),
			"netbox_cable":                      resourceNetboxCable(),
			"netbox_device_console_port":        resourceNetboxDeviceConsolePort(),
			"netbox_device_console_server_port": resourceNetboxDeviceConsoleServerPort(),
			"netbox_device_power_port":          resourceNetboxDevicePowerPort(),
			"netbox_device_power_outlet":        resourceNetboxDevicePowerOutlet(),
			"netbox_device_front_port":          resourceNetboxDeviceFrontPort(),
			"netbox_device_rear_port":           resourceNetboxDeviceRearPort(),
			"netbox_device_module_bay":          resourceNetboxDeviceModuleBay(),
			"netbox_module":                     resourceNetboxModule(),
			"netbox_module_type":                resourceNetboxModuleType(),
			"netbox_power_feed":                 resourceNetboxPowerFeed(),
			"netbox_power_panel":                resourceNetboxPowerPanel(),
			"netbox_inventory_item_role":        resourceNetboxInventoryItemRole(),
			"netbox_inventory_item":             resourceNetboxInventoryItem(),
			"netbox_webhook":                    resourceNetboxWebhook(),
			"netbox_custom_field_choice_set":    resourceNetboxCustomFieldChoiceSet(),
			"netbox_virtual_chassis":            resourceNetboxVirtualChassis(),
			"netbox_virtual_disk":               resourceNetboxVirtualDisks(),
			"netbox_config_template":            resourceNetboxConfigTemplate(),
			"netbox_event_rule":                 resourceNetboxEventRule(),
			"netbox_vpn_tunnel_group":           resourceNetboxVpnTunnelGroup(),
			"netbox_vpn_tunnel":                 resourceNetboxVpnTunnel(),
			"netbox_vpn_tunnel_termination":     resourceNetboxVpnTunnelTermination(),
			"netbox_config_context":             resourceNetboxConfigContext(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_asn":                dataSourceNetboxAsn(),
			"netbox_asns":               dataSourceNetboxAsns(),
			"netbox_available_prefix":   dataSourceNetboxAvailablePrefix(),
			"netbox_cluster":            dataSourceNetboxCluster(),
			"netbox_cluster_group":      dataSourceNetboxClusterGroup(),
			"netbox_cluster_type":       dataSourceNetboxClusterType(),
			"netbox_contact":            dataSourceNetboxContact(),
			"netbox_contact_role":       dataSourceNetboxContactRole(),
			"netbox_contact_group":      dataSourceNetboxContactGroup(),
			"netbox_tenant":             dataSourceNetboxTenant(),
			"netbox_tenants":            dataSourceNetboxTenants(),
			"netbox_tenant_group":       dataSourceNetboxTenantGroup(),
			"netbox_vrf":                dataSourceNetboxVrf(),
			"netbox_vrfs":               dataSourceNetboxVrfs(),
			"netbox_platform":           dataSourceNetboxPlatform(),
			"netbox_prefix":             dataSourceNetboxPrefix(),
			"netbox_prefixes":           dataSourceNetboxPrefixes(),
			"netbox_devices":            dataSourceNetboxDevices(),
			"netbox_device_role":        dataSourceNetboxDeviceRole(),
			"netbox_device_type":        dataSourceNetboxDeviceType(),
			"netbox_site":               dataSourceNetboxSite(),
			"netbox_location":           dataSourceNetboxLocation(),
			"netbox_locations":          dataSourceNetboxLocations(),
			"netbox_tag":                dataSourceNetboxTag(),
			"netbox_tags":               dataSourceNetboxTags(),
			"netbox_virtual_machines":   dataSourceNetboxVirtualMachine(),
			"netbox_interfaces":         dataSourceNetboxInterfaces(),
			"netbox_device_interfaces":  dataSourceNetboxDeviceInterfaces(),
			"netbox_device_power_ports": dataSourceNetboxDevicePowerPorts(),
			"netbox_ipam_role":          dataSourceNetboxIPAMRole(),
			"netbox_route_target":       dataSourceNetboxRouteTarget(),
			"netbox_ip_addresses":       dataSourceNetboxIPAddresses(),
			"netbox_ip_range":           dataSourceNetboxIPRange(),
			"netbox_ip_ranges":          dataSourceNetboxIPRanges(),
			"netbox_region":             dataSourceNetboxRegion(),
			"netbox_vlan":               dataSourceNetboxVlan(),
			"netbox_vlans":              dataSourceNetboxVlans(),
			"netbox_vlan_group":         dataSourceNetboxVlanGroup(),
			"netbox_site_group":         dataSourceNetboxSiteGroup(),
			"netbox_racks":              dataSourceNetboxRacks(),
			"netbox_rack_role":          dataSourceNetboxRackRole(),
			"netbox_config_context":     dataSourceNetboxConfigContext(),
		},
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_SERVER_URL", nil),
				Description: "Location of Netbox server including scheme (http or https) and optional port. Can be set via the `NETBOX_SERVER_URL` environment variable.",
			},
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_API_TOKEN", nil),
				Description: "Netbox API authentication token. Can be set via the `NETBOX_API_TOKEN` environment variable.",
			},
			"allow_insecure_https": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_ALLOW_INSECURE_HTTPS", false),
				Description: "Flag to set whether to allow https with invalid certificates. Can be set via the `NETBOX_ALLOW_INSECURE_HTTPS` environment variable. Defaults to `false`.",
			},
			"headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_HEADERS", map[string]interface{}{}),
				Description: "Set these header on all requests to Netbox. Can be set via the `NETBOX_HEADERS` environment variable.",
			},
			"strip_trailing_slashes_from_url": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_STRIP_TRAILING_SLASHES_FROM_URL", true),
				Description: "If true, strip trailing slashes from the `server_url` parameter and print a warning when doing so. Note that using trailing slashes in the `server_url` parameter will usually lead to errors. Can be set via the `NETBOX_STRIP_TRAILING_SLASHES_FROM_URL` environment variable. Defaults to `true`.",
			},
			"skip_version_check": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_SKIP_VERSION_CHECK", false),
				Description: "If true, do not try to determine the running Netbox version at provider startup. Disables warnings about possibly unsupported Netbox version. Also useful for local testing on terraform plans. Can be set via the `NETBOX_SKIP_VERSION_CHECK` environment variable. Defaults to `false`.",
			},
			"request_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_REQUEST_TIMEOUT", 10),
				Description: "Netbox API HTTP request timeout in seconds. Can be set via the `NETBOX_REQUEST_TIMEOUT` environment variable.",
			},
			"default_tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Tags to add to every resource managed by this provider",
			},
		},
		ConfigureContextFunc: providerConfigure,
	}

	// all resources that have tags get a custom diff function
	for _, def := range provider.ResourcesMap {
		if _, ok := def.Schema[tagsKey]; ok {
			def.Schema[tagsAllKey] = tagsAllSchema // add computed key for all tags
			if existingDiff := def.CustomizeDiff; existingDiff != nil {
				def.CustomizeDiff = customdiff.Sequence(existingDiff, tagsCustomDiff)
			} else {
				def.CustomizeDiff = tagsCustomDiff
			}
		}
	}

	return provider
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	config := Config{
		APIToken:                    data.Get("api_token").(string),
		AllowInsecureHTTPS:          data.Get("allow_insecure_https").(bool),
		Headers:                     data.Get("headers").(map[string]interface{}),
		RequestTimeout:              data.Get("request_timeout").(int),
		StripTrailingSlashesFromURL: data.Get("strip_trailing_slashes_from_url").(bool),
	}

	serverURL := data.Get("server_url").(string)

	// Unless explicitly switched off, strip trailing slashes from the server url
	// Trailing slashes cause errors as seen in https://github.com/e-breuninger/terraform-provider-netbox/issues/198
	// and https://github.com/e-breuninger/terraform-provider-netbox/issues/300
	stripTrailingSlashesFromURL := data.Get("strip_trailing_slashes_from_url").(bool)

	if stripTrailingSlashesFromURL {
		trimmed := false

		// This is Go's poor man's while loop
		for strings.HasSuffix(serverURL, "/") {
			serverURL = strings.TrimRight(serverURL, "/")
			trimmed = true
		}
		if trimmed {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Stripped trailing slashes from the `server_url` parameter",
				Detail:   "Trailing slashes in the `server_url` parameter lead to problems in most setups, so all trailing slashes were stripped. Use the `strip_trailing_slashes_from_url` parameter to disable this feature or remove all trailing slashes in the `server_url` to disable this warning.",
			})
		}
	}

	config.ServerURL = serverURL

	netboxClient, clientError := config.Client()
	if clientError != nil {
		return nil, diag.FromErr(clientError)
	}

	// Unless explicitly switched off, use the client to retrieve the Netbox version
	// so we can determine compatibility of the provider with the used Netbox
	skipVersionCheck := data.Get("skip_version_check").(bool)

	if !skipVersionCheck {
		req := status.NewStatusListParams()
		res, err := netboxClient.Status.StatusList(req, nil)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		netboxVersion := res.GetPayload().(map[string]interface{})["netbox-version"].(string)

		supportedVersions := []string{"4.1.0", "4.1.1", "4.1.2", "4.1.3", "4.1.4", "4.1.5", "4.1.6", "4.1.7", "4.1.8", "4.1.10", "4.1.11"}

		if !slices.Contains(supportedVersions, netboxVersion) {
			// Currently, there is no way to test these warnings. There is an issue to track this: https://github.com/hashicorp/terraform-plugin-sdk/issues/864
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Possibly unsupported Netbox version",
				Detail:   fmt.Sprintf("Your Netbox version is v%v. The provider was successfully tested against the following versions:\n\n  %v\n\nUnexpected errors may occur.", netboxVersion, strings.Join(supportedVersions, ", ")),
			})
		}
	}

	tags, ok := data.Get("default_tags").(*schema.Set)
	tagCache := make(map[string]*models.NestedTag, tags.Len())
	if ok {
		for _, tag := range tags.List() {
			if tagName, ok := tag.(string); ok {
				nbTag, err := findTag(netboxClient, tagName)
				if err != nil {
					d := diag.FromErr(fmt.Errorf("default tag not found: %w", err))
					d[0].Severity = diag.Warning
					diags = append(diags, d...)
				} else {
					tagCache[tagName] = nbTag
				}
			} else {
				diags = append(diags, diag.Errorf("invalid type for default tag: %T", tag)...)
			}
		}
	}

	state := &providerState{
		NetBoxAPI:   netboxClient,
		defaultTags: schema.CopySet(tags),
		tagCache:    tagCache,
	}
	return state, diags
}

func tagsCustomDiff(ctx context.Context, diff *schema.ResourceDiff, m interface{}) error {
	state := m.(*providerState)

	tagSet := diff.Get(tagsKey).(*schema.Set)
	allTags := tagSet.Union(state.defaultTags)

	// check if tags are already up-to-date
	if diff.Get(tagsAllKey).(*schema.Set).Equal(allTags) {
		return nil // nothing to do, same set
	}

	return diff.SetNew(tagsAllKey, allTags.List())
}
