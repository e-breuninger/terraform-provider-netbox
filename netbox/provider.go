package netbox

import (
	"context"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/status"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a schema.Provider for Netbox.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"netbox_available_ip_address": resourceNetboxAvailableIPAddress(),
			"netbox_virtual_machine":      resourceNetboxVirtualMachine(),
			"netbox_cluster_type":         resourceNetboxClusterType(),
			"netbox_cluster":              resourceNetboxCluster(),
			"netbox_tenant":               resourceNetboxTenant(),
			"netbox_tenant_group":         resourceNetboxTenantGroup(),
			"netbox_vrf":                  resourceNetboxVrf(),
			"netbox_ip_address":           resourceNetboxIPAddress(),
			"netbox_interface":            resourceNetboxInterface(),
			"netbox_service":              resourceNetboxService(),
			"netbox_platform":             resourceNetboxPlatform(),
			"netbox_prefix":               resourceNetboxPrefix(),
			"netbox_available_prefix":     resourceNetboxAvailablePrefix(),
			"netbox_primary_ip":           resourceNetboxPrimaryIP(),
			"netbox_device_role":          resourceNetboxDeviceRole(),
			"netbox_tag":                  resourceNetboxTag(),
			"netbox_cluster_group":        resourceNetboxClusterGroup(),
			"netbox_site":                 resourceNetboxSite(),
			"netbox_vlan":                 resourceNetboxVlan(),
			"netbox_ipam_role":            resourceNetboxIpamRole(),
			"netbox_ip_range":             resourceNetboxIpRange(),
			"netbox_region":               resourceNetboxRegion(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_cluster":          dataSourceNetboxCluster(),
			"netbox_cluster_group":    dataSourceNetboxClusterGroup(),
			"netbox_cluster_type":     dataSourceNetboxClusterType(),
			"netbox_tenant":           dataSourceNetboxTenant(),
			"netbox_tenant_group":     dataSourceNetboxTenantGroup(),
			"netbox_vrf":              dataSourceNetboxVrf(),
			"netbox_platform":         dataSourceNetboxPlatform(),
			"netbox_prefix":           dataSourceNetboxPrefix(),
			"netbox_device_role":      dataSourceNetboxDeviceRole(),
			"netbox_site":             dataSourceNetboxSite(),
			"netbox_tag":              dataSourceNetboxTag(),
			"netbox_virtual_machines": dataSourceNetboxVirtualMachine(),
			"netbox_interfaces":       dataSourceNetboxInterfaces(),
			"netbox_ip_range":         dataSourceNetboxIpRange(),
			"netbox_region":           dataSourceNetboxRegion(),
		},
		Schema: map[string]*schema.Schema{
			"server_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_SERVER_URL", nil),
				Description: "Location of Netbox server including scheme and optional port",
			},
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_API_TOKEN", nil),
				Description: "Netbox API authentication token",
			},
			"allow_insecure_https": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_ALLOW_INSECURE_HTTPS", false),
				Description: "Flag to set whether to allow https with invalid certificates",
			},
			"headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_HEADERS", map[string]interface{}{}),
				Description: "Set these header on all requests to Netbox",
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
	return provider
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics

	config := Config{
		ServerURL:          data.Get("server_url").(string),
		APIToken:           data.Get("api_token").(string),
		AllowInsecureHttps: data.Get("allow_insecure_https").(bool),
		Headers:            data.Get("headers").(map[string]interface{}),
	}

	netboxClient, clientError := config.Client()
	if clientError != nil {
		return nil, diag.FromErr(clientError)
	}

	req := status.NewStatusListParams()
	res, _ := netboxClient.(*client.NetBoxAPI).Status.StatusList(req, nil)
	netboxVersion := res.GetPayload().(map[string]interface{})["netbox-version"]

	supportedVersion := "3.1.3"

	if netboxVersion != supportedVersion {

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Possibly unsupported Netbox version",
			Detail:   fmt.Sprintf("This provider was tested against Netbox v%s. Your Netbox version is v%v. Unexpected errors may occur.", supportedVersion, netboxVersion),
		})
	}

	return netboxClient, diags
}
