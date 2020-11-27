package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a schema.Provider for Netbox.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"netbox_virtual_machine": resourceNetboxVirtualMachine(),
			"netbox_cluster_type":    resourceNetboxClusterType(),
			"netbox_cluster":         resourceNetboxCluster(),
			"netbox_tenant":          resourceNetboxTenant(),
			"netbox_vrf":             resourceNetboxVrf(),
			"netbox_ip_address":      resourceNetboxIPAddress(),
			"netbox_interface":       resourceNetboxInterface(),
			"netbox_service":         resourceNetboxService(),
			"netbox_platform":        resourceNetboxPlatform(),
			"netbox_primary_ip":      resourceNetboxPrimaryIP(),
			"netbox_device_role":     resourceNetboxDeviceRole(),
			"netbox_tag":             resourceNetboxTag(),
			"netbox_cluster_group":   resourceNetboxClusterGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_cluster":          dataSourceNetboxCluster(),
			"netbox_tenant":           dataSourceNetboxTenant(),
			"netbox_vrf":              dataSourceNetboxVrf(),
			"netbox_platform":         dataSourceNetboxPlatform(),
			"netbox_device_role":      dataSourceNetboxDeviceRole(),
			"netbox_tag":              dataSourceNetboxTag(),
			"netbox_virtual_machines": dataSourceNetboxVirtualMachine(),
			"netbox_interfaces":       dataSourceNetboxInterfaces(),
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
		},
		ConfigureFunc: providerConfigure,
	}
	return provider
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {

	config := Config{
		ServerURL:          data.Get("server_url").(string),
		APIToken:           data.Get("api_token").(string),
		AllowInsecureHttps: data.Get("allow_insecure_https").(bool),
	}

	return config.Client()
}
