// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider satisfies various provider interfaces.
var _ provider.Provider = (*netboxProvider)(nil)

// New returns the provider factory used by main.go and acceptance tests.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &netboxProvider{version: version}
	}
}

type netboxProvider struct {
	version string
}

type netboxProviderModel struct {
	ServerURL                   types.String `tfsdk:"server_url"`
	APIToken                    types.String `tfsdk:"api_token"`
	AllowInsecureHTTPS          types.Bool   `tfsdk:"allow_insecure_https"`
	Headers                     types.Map    `tfsdk:"headers"`
	StripTrailingSlashesFromURL types.Bool   `tfsdk:"strip_trailing_slashes_from_url"`
	RequestTimeout              types.Int64  `tfsdk:"request_timeout"`
	CACertFile                  types.String `tfsdk:"ca_cert_file"`
	DefaultTags                 types.Set    `tfsdk:"default_tags"`
	SkipVersionCheck            types.Bool   `tfsdk:"skip_version_check"`
}

func (p *netboxProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netbox"
	resp.Version = p.version
}

func (p *netboxProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				Optional:    true,
				Description: "URL of the NetBox instance (e.g. https://netbox.example.com). The /api path is added automatically). Falls back to the NETBOX_SERVER_URL environment variable.",
			},
			"api_token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Netbox API authentication token. Supports both v1 tokens and v2 tokens. Can be set via the NETBOX_API_TOKEN environment variable.",
			},
			"allow_insecure_https": schema.BoolAttribute{
				Optional:    true,
				Description: "Flag to set whether to allow https with invalid certificates. Can be set via the NETBOX_ALLOW_INSECURE_HTTPS environment variable. Defaults to false.",
			},
			"headers": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Set these header on all requests to Netbox. Can be set via the NETBOX_HEADERS environment variable (a JSON object).",
			},
			"strip_trailing_slashes_from_url": schema.BoolAttribute{
				Optional:    true,
				Description: "If true, strip trailing slashes from the server_url parameter and print a warning when doing so. Note that using trailing slashes in the server_url parameter will usually lead to errors. Can be set via the NETBOX_STRIP_TRAILING_SLASHES_FROM_URL environment variable. Defaults to true.",
			},
			"request_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "Netbox API HTTP request timeout in seconds. Can be set via the NETBOX_REQUEST_TIMEOUT environment variable. Defaults to 10.",
			},
			"ca_cert_file": schema.StringAttribute{
				Optional:    true,
				Description: "Path to a PEM-encoded CA certificate for verifying the Netbox server certificate. Can be set via the NETBOX_CA_CERT_FILE environment variable.",
			},
			"default_tags": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Slugs of tags to add to every resource managed by this provider.",
			},
			"skip_version_check": schema.BoolAttribute{
				Optional:    true,
				Description: "If true, do not try to determine the running Netbox version at provider startup. Can be set via the NETBOX_SKIP_VERSION_CHECK environment variable. Defaults to false.",
			},
		},
	}
}

func (p *netboxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config netboxProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ServerURL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("server_url"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"server_url\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.APIToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"api_token\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.AllowInsecureHTTPS.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("allow_insecure_https"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"allow_insecure_https\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.Headers.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("headers"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"headers\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.StripTrailingSlashesFromURL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("strip_trailing_slashes_from_url"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"strip_trailing_slashes_from_url\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.RequestTimeout.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("request_timeout"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"request_timeout\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.CACertFile.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("ca_cert_file"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"ca_cert_file\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.DefaultTags.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("default_tags"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"default_tags\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if config.SkipVersionCheck.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("skip_version_check"),
			"Unknown provider configuration value",
			"The provider cannot be configured with an unknown value for \"skip_version_check\". Set it statically in the configuration, or use an environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	serverURL := config.ServerURL.ValueString()
	if serverURL == "" {
		serverURL = os.Getenv("NETBOX_SERVER_URL")
	}
	stripTrailingSlashes := true
	if !config.StripTrailingSlashesFromURL.IsNull() {
		stripTrailingSlashes = config.StripTrailingSlashesFromURL.ValueBool()
	} else if v := os.Getenv("NETBOX_STRIP_TRAILING_SLASHES_FROM_URL"); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("strip_trailing_slashes_from_url"),
				"Invalid NETBOX_STRIP_TRAILING_SLASHES_FROM_URL",
				"The NETBOX_STRIP_TRAILING_SLASHES_FROM_URL environment variable must be a boolean (true/false), got "+strconv.Quote(v)+".",
			)
		}
		stripTrailingSlashes = b
	}
	if stripTrailingSlashes && strings.HasSuffix(serverURL, "/") {
		serverURL = strings.TrimRight(serverURL, "/")
		resp.Diagnostics.AddAttributeWarning(
			path.Root("server_url"),
			"Stripped trailing slashes from the `server_url` parameter",
			"Trailing slashes in the `server_url` parameter lead to problems in most setups, so all trailing slashes were stripped.",
		)
	}
	allowInsecureHTTPS := false
	if !config.AllowInsecureHTTPS.IsNull() {
		allowInsecureHTTPS = config.AllowInsecureHTTPS.ValueBool()
	} else if v := os.Getenv("NETBOX_ALLOW_INSECURE_HTTPS"); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("allow_insecure_https"),
				"Invalid NETBOX_ALLOW_INSECURE_HTTPS",
				"The NETBOX_ALLOW_INSECURE_HTTPS environment variable must be a boolean (true/false), got "+strconv.Quote(v)+".",
			)
		}
		allowInsecureHTTPS = b
	}
	skipVersionCheck := false
	if !config.SkipVersionCheck.IsNull() {
		skipVersionCheck = config.SkipVersionCheck.ValueBool()
	} else if v := os.Getenv("NETBOX_SKIP_VERSION_CHECK"); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("skip_version_check"),
				"Invalid NETBOX_SKIP_VERSION_CHECK",
				"The NETBOX_SKIP_VERSION_CHECK environment variable must be a boolean (true/false), got "+strconv.Quote(v)+".",
			)
		}
		skipVersionCheck = b
	}
	if serverURL == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("server_url"),
			"Missing NetBox URL",
			"Set the server_url provider attribute or the NETBOX_SERVER_URL environment variable.",
		)
	}
	apiToken := config.APIToken.ValueString()
	if apiToken == "" {
		apiToken = os.Getenv("NETBOX_API_TOKEN")
	}
	if apiToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing NetBox API token",
			"Set the api_token provider attribute or the NETBOX_API_TOKEN environment variable.",
		)
	}
	headers := map[string]string{}
	if !config.Headers.IsNull() {
		resp.Diagnostics.Append(config.Headers.ElementsAs(ctx, &headers, false)...)
	} else if v := os.Getenv("NETBOX_HEADERS"); v != "" {
		if err := json.Unmarshal([]byte(v), &headers); err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("headers"),
				"Invalid NETBOX_HEADERS",
				"The NETBOX_HEADERS environment variable must be a JSON object of header names to values: "+err.Error(),
			)
		}
	}
	requestTimeout := int64(10)
	if !config.RequestTimeout.IsNull() {
		requestTimeout = config.RequestTimeout.ValueInt64()
	} else if v := os.Getenv("NETBOX_REQUEST_TIMEOUT"); v != "" {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("request_timeout"),
				"Invalid NETBOX_REQUEST_TIMEOUT",
				"The NETBOX_REQUEST_TIMEOUT environment variable must be a number of seconds, got "+strconv.Quote(v)+".",
			)
		}
		requestTimeout = n
	}
	caCertFile := config.CACertFile.ValueString()
	if caCertFile == "" {
		caCertFile = os.Getenv("NETBOX_CA_CERT_FILE")
	}
	clientOptions := netboxapi.Options{
		AllowInsecureHTTPS: allowInsecureHTTPS,
		CACertFile:         caCertFile,
		Headers:            headers,
		RequestTimeout:     time.Duration(requestTimeout) * time.Second,
	}
	if resp.Diagnostics.HasError() {
		return
	}
	c, err := netboxapi.New(serverURL, apiToken, clientOptions)
	if err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("server_url"), "Invalid NetBox URL", err.Error())
		return
	}
	if !skipVersionCheck {
		reported, version, err := netboxapi.Version(ctx, c)
		if err != nil {
			resp.Diagnostics.AddError("Failed to determine Netbox version", netboxapi.VersionErrorDetail(serverURL, err))
			return
		}
		supported := supportedVersions
		if !slices.Contains(supported, version) {
			resp.Diagnostics.AddWarning(
				"Possibly unsupported Netbox version",
				fmt.Sprintf("Your Netbox reports version %v. From that, the provider extracted Netbox version %v.\nThe provider was successfully tested against the following versions:\n\n  %v\n\nUnexpected errors may occur.", reported, version, strings.Join(supported, ", ")),
			)
		}
	}
	defaultTags := []string{}
	if !config.DefaultTags.IsNull() {
		resp.Diagnostics.Append(config.DefaultTags.ElementsAs(ctx, &defaultTags, false)...)
	}
	if err := netboxapi.CheckTags(ctx, c, defaultTags); err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("default_tags"), "Invalid default tag", err.Error())
		return
	}
	client := &netboxapi.Client{NetBoxAPI: c, DefaultTags: defaultTags}
	resp.ResourceData = client
	resp.DataSourceData = client
}

// companionResources and companionDataSources are hand-written resources
// and data sources that companion files register from an init() function:
//
//	func init() { companionResources = append(companionResources, NewXResource) }
//
// They are served next to the generated ones.
var (
	companionResources   []func() resource.Resource
	companionDataSources []func() datasource.DataSource
)

func (p *netboxProvider) Resources(ctx context.Context) []func() resource.Resource {
	return append([]func() resource.Resource{
		NewSiteResource,
		NewRegionResource,
		NewSiteGroupResource,
		NewTenantGroupResource,
		NewTenantResource,
		NewTagResource,
		NewIPAddressResource,
		NewAvailableIPAddressResource,
		NewAvailablePrefixResource,
		NewAvailableAsnResource,
		NewAvailableVlanResource,
		NewCustomFieldResource,
		NewCustomFieldChoiceSetResource,
		NewLocationResource,
		NewManufacturerResource,
		NewPlatformResource,
		NewDeviceRoleResource,
		NewRackRoleResource,
		NewRackResource,
		NewRackReservationResource,
		NewDeviceTypeResource,
		NewModuleTypeResource,
		NewMacAddressResource,
		NewRirResource,
		NewIpamRoleResource,
		NewRouteTargetResource,
		NewVlanResource,
		NewPrefixResource,
		NewAsnResource,
		NewVlanGroupResource,
		NewVrfResource,
		NewContactRoleResource,
		NewContactGroupResource,
		NewContactResource,
		NewClusterTypeResource,
		NewClusterGroupResource,
		NewClusterResource,
		NewCircuitProviderResource,
		NewCircuitTypeResource,
		NewVpnTunnelGroupResource,
		NewVpnTunnelResource,
		NewVpnTunnelTerminationResource,
		NewVirtualMachineResource,
		NewVirtualMachineInterfaceResource,
		NewVirtualDiskResource,
		NewIkeProposalResource,
		NewL2vpnResource,
		NewL2vpnTerminationResource,
		NewDeviceResource,
		NewDeviceInterfaceResource,
		NewDeviceBayResource,
		NewDeviceConsolePortResource,
		NewDeviceConsoleServerPortResource,
		NewDevicePowerPortResource,
		NewDevicePowerOutletResource,
		NewDeviceFrontPortResource,
		NewDeviceRearPortResource,
		NewDeviceModuleBayResource,
		NewModuleResource,
		NewInventoryItemRoleResource,
		NewInventoryItemResource,
		NewVirtualChassisResource,
		NewVirtualDeviceContextResource,
		NewAggregateResource,
		NewIPRangeResource,
		NewServiceResource,
		NewServiceTemplateResource,
		NewFhrpGroupResource,
		NewFhrpGroupAssignmentResource,
		NewCircuitResource,
		NewCircuitTerminationResource,
		NewCircuitProviderNetworkResource,
		NewPowerPanelResource,
		NewPowerFeedResource,
		NewGroupResource,
		NewUserResource,
		NewTokenResource,
		NewPermissionResource,
		NewConsolePortTemplateResource,
		NewConsoleServerPortTemplateResource,
		NewDeviceBayTemplateResource,
		NewInterfaceTemplateResource,
		NewModuleBayTemplateResource,
		NewPowerPortTemplateResource,
		NewPowerOutletTemplateResource,
		NewRearPortTemplateResource,
		NewFrontPortTemplateResource,
		NewInventoryItemTemplateResource,
		NewRackGroupResource,
		NewCableResource,
		NewCableBundleResource,
		NewModuleTypeProfileResource,
		NewCircuitProviderAccountResource,
		NewCircuitGroupResource,
		NewCircuitGroupAssignmentResource,
		NewVirtualCircuitTypeResource,
		NewVirtualCircuitResource,
		NewVirtualCircuitTerminationResource,
		NewAsnRangeResource,
		NewVlanTranslationPolicyResource,
		NewVlanTranslationRuleResource,
		NewVirtualMachineTypeResource,
		NewIkePolicyResource,
		NewIpsecProposalResource,
		NewIpsecPolicyResource,
		NewIpsecProfileResource,
		NewDataSourceResource,
		NewConfigContextProfileResource,
		NewOwnerGroupResource,
		NewOwnerResource,
		NewRackTypeResource,
		NewWirelessLanGroupResource,
		NewWirelessLanResource,
		NewWirelessLinkResource,
		NewContactAssignmentResource,
		NewWebhookResource,
		NewCustomLinkResource,
		NewExportTemplateResource,
		NewNotificationGroupResource,
		NewEventRuleResource,
		NewConfigTemplateResource,
		NewConfigContextResource,
	}, companionResources...)
}

func (p *netboxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return append([]func() datasource.DataSource{
		NewSiteDataSource,
		NewRegionDataSource,
		NewSiteGroupDataSource,
		NewTenantGroupDataSource,
		NewTenantDataSource,
		NewTagDataSource,
		NewIPAddressDataSource,
		NewCustomFieldDataSource,
		NewCustomFieldChoiceSetDataSource,
		NewLocationDataSource,
		NewManufacturerDataSource,
		NewPlatformDataSource,
		NewDeviceRoleDataSource,
		NewRackRoleDataSource,
		NewRackDataSource,
		NewRackReservationDataSource,
		NewDeviceTypeDataSource,
		NewModuleTypeDataSource,
		NewMacAddressDataSource,
		NewRirDataSource,
		NewIpamRoleDataSource,
		NewRouteTargetDataSource,
		NewVlanDataSource,
		NewPrefixDataSource,
		NewAsnDataSource,
		NewVlanGroupDataSource,
		NewVrfDataSource,
		NewContactRoleDataSource,
		NewContactGroupDataSource,
		NewContactDataSource,
		NewClusterTypeDataSource,
		NewClusterGroupDataSource,
		NewClusterDataSource,
		NewCircuitProviderDataSource,
		NewCircuitTypeDataSource,
		NewVpnTunnelGroupDataSource,
		NewVpnTunnelDataSource,
		NewVpnTunnelTerminationDataSource,
		NewVirtualMachineDataSource,
		NewVirtualMachineInterfaceDataSource,
		NewVirtualDiskDataSource,
		NewIkeProposalDataSource,
		NewL2vpnDataSource,
		NewL2vpnTerminationDataSource,
		NewDeviceDataSource,
		NewDeviceInterfaceDataSource,
		NewDeviceBayDataSource,
		NewDeviceConsolePortDataSource,
		NewDeviceConsoleServerPortDataSource,
		NewDevicePowerPortDataSource,
		NewDevicePowerOutletDataSource,
		NewDeviceFrontPortDataSource,
		NewDeviceRearPortDataSource,
		NewDeviceModuleBayDataSource,
		NewModuleDataSource,
		NewInventoryItemRoleDataSource,
		NewInventoryItemDataSource,
		NewVirtualChassisDataSource,
		NewVirtualDeviceContextDataSource,
		NewAggregateDataSource,
		NewIPRangeDataSource,
		NewServiceDataSource,
		NewServiceTemplateDataSource,
		NewFhrpGroupDataSource,
		NewFhrpGroupAssignmentDataSource,
		NewCircuitDataSource,
		NewCircuitTerminationDataSource,
		NewCircuitProviderNetworkDataSource,
		NewPowerPanelDataSource,
		NewPowerFeedDataSource,
		NewGroupDataSource,
		NewUserDataSource,
		NewTokenDataSource,
		NewPermissionDataSource,
		NewConsolePortTemplateDataSource,
		NewConsoleServerPortTemplateDataSource,
		NewDeviceBayTemplateDataSource,
		NewInterfaceTemplateDataSource,
		NewModuleBayTemplateDataSource,
		NewPowerPortTemplateDataSource,
		NewPowerOutletTemplateDataSource,
		NewRearPortTemplateDataSource,
		NewFrontPortTemplateDataSource,
		NewInventoryItemTemplateDataSource,
		NewRackGroupDataSource,
		NewCableDataSource,
		NewCableBundleDataSource,
		NewModuleTypeProfileDataSource,
		NewCircuitProviderAccountDataSource,
		NewCircuitGroupDataSource,
		NewCircuitGroupAssignmentDataSource,
		NewVirtualCircuitTypeDataSource,
		NewVirtualCircuitDataSource,
		NewVirtualCircuitTerminationDataSource,
		NewAsnRangeDataSource,
		NewVlanTranslationPolicyDataSource,
		NewVlanTranslationRuleDataSource,
		NewVirtualMachineTypeDataSource,
		NewIkePolicyDataSource,
		NewIpsecProposalDataSource,
		NewIpsecPolicyDataSource,
		NewIpsecProfileDataSource,
		NewDataSourceDataSource,
		NewConfigContextProfileDataSource,
		NewOwnerGroupDataSource,
		NewOwnerDataSource,
		NewRackTypeDataSource,
		NewWirelessLanGroupDataSource,
		NewWirelessLanDataSource,
		NewWirelessLinkDataSource,
		NewContactAssignmentDataSource,
		NewWebhookDataSource,
		NewCustomLinkDataSource,
		NewExportTemplateDataSource,
		NewEventRuleDataSource,
		NewConfigTemplateDataSource,
		NewConfigContextDataSource,
		NewSitesDataSource,
		NewRegionsDataSource,
		NewSiteGroupsDataSource,
		NewTenantGroupsDataSource,
		NewTenantsDataSource,
		NewTagsDataSource,
		NewIPAddressesDataSource,
		NewCustomFieldsDataSource,
		NewCustomFieldChoiceSetsDataSource,
		NewLocationsDataSource,
		NewManufacturersDataSource,
		NewPlatformsDataSource,
		NewDeviceRolesDataSource,
		NewRackRolesDataSource,
		NewRacksDataSource,
		NewRackReservationsDataSource,
		NewDeviceTypesDataSource,
		NewModuleTypesDataSource,
		NewMacAddressesDataSource,
		NewRirsDataSource,
		NewIpamRolesDataSource,
		NewRouteTargetsDataSource,
		NewVlansDataSource,
		NewPrefixesDataSource,
		NewAsnsDataSource,
		NewVlanGroupsDataSource,
		NewVrfsDataSource,
		NewContactRolesDataSource,
		NewContactGroupsDataSource,
		NewContactsDataSource,
		NewClusterTypesDataSource,
		NewClusterGroupsDataSource,
		NewClustersDataSource,
		NewCircuitProvidersDataSource,
		NewCircuitTypesDataSource,
		NewVpnTunnelGroupsDataSource,
		NewVpnTunnelsDataSource,
		NewVpnTunnelTerminationsDataSource,
		NewVirtualMachinesDataSource,
		NewVirtualMachineInterfacesDataSource,
		NewVirtualDisksDataSource,
		NewIkeProposalsDataSource,
		NewL2vpnsDataSource,
		NewL2vpnTerminationsDataSource,
		NewDevicesDataSource,
		NewDeviceInterfacesDataSource,
		NewDeviceBaysDataSource,
		NewDeviceConsolePortsDataSource,
		NewDeviceConsoleServerPortsDataSource,
		NewDevicePowerPortsDataSource,
		NewDevicePowerOutletsDataSource,
		NewDeviceFrontPortsDataSource,
		NewDeviceRearPortsDataSource,
		NewDeviceModuleBaysDataSource,
		NewModulesDataSource,
		NewInventoryItemRolesDataSource,
		NewInventoryItemsDataSource,
		NewVirtualDeviceContextsDataSource,
		NewAggregatesDataSource,
		NewIPRangesDataSource,
		NewServicesDataSource,
		NewServiceTemplatesDataSource,
		NewFhrpGroupsDataSource,
		NewFhrpGroupAssignmentsDataSource,
		NewCircuitsDataSource,
		NewCircuitTerminationsDataSource,
		NewCircuitProviderNetworksDataSource,
		NewPowerPanelsDataSource,
		NewPowerFeedsDataSource,
		NewGroupsDataSource,
		NewUsersDataSource,
		NewTokensDataSource,
		NewPermissionsDataSource,
		NewConsolePortTemplatesDataSource,
		NewConsoleServerPortTemplatesDataSource,
		NewDeviceBayTemplatesDataSource,
		NewInterfaceTemplatesDataSource,
		NewModuleBayTemplatesDataSource,
		NewPowerPortTemplatesDataSource,
		NewPowerOutletTemplatesDataSource,
		NewRearPortTemplatesDataSource,
		NewFrontPortTemplatesDataSource,
		NewInventoryItemTemplatesDataSource,
		NewRackGroupsDataSource,
		NewCablesDataSource,
		NewCableBundlesDataSource,
		NewModuleTypeProfilesDataSource,
		NewCircuitProviderAccountsDataSource,
		NewCircuitGroupsDataSource,
		NewCircuitGroupAssignmentsDataSource,
		NewVirtualCircuitTypesDataSource,
		NewVirtualCircuitsDataSource,
		NewVirtualCircuitTerminationsDataSource,
		NewAsnRangesDataSource,
		NewVlanTranslationPoliciesDataSource,
		NewVlanTranslationRulesDataSource,
		NewVirtualMachineTypesDataSource,
		NewIkePoliciesDataSource,
		NewIpsecProposalsDataSource,
		NewIpsecPoliciesDataSource,
		NewIpsecProfilesDataSource,
		NewDataSourcesDataSource,
		NewConfigContextProfilesDataSource,
		NewOwnerGroupsDataSource,
		NewOwnersDataSource,
		NewRackTypesDataSource,
		NewWirelessLanGroupsDataSource,
		NewWirelessLansDataSource,
		NewWirelessLinksDataSource,
		NewContactAssignmentsDataSource,
		NewWebhooksDataSource,
		NewCustomLinksDataSource,
		NewExportTemplatesDataSource,
		NewEventRulesDataSource,
		NewConfigTemplatesDataSource,
		NewConfigContextsDataSource,
	}, companionDataSources...)
}
