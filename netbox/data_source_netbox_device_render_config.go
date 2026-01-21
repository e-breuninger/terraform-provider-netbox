package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxDeviceRenderConfig() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDeviceRenderConfigRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):Render the configuration template assigned to a device using the device's config context.`,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the device to render configuration for.",
			},
			"content": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rendered configuration content.",
			},
			"config_template_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the config template that was used for rendering.",
			},
			"config_template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the config template that was used for rendering.",
			},
		},
	}
}

func dataSourceNetboxDeviceRenderConfigRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	deviceID := int64(d.Get("device_id").(int))

	params := dcim.NewDcimDevicesRenderConfigParams().WithID(deviceID)

	res, err := api.Dcim.DcimDevicesRenderConfig(params, nil)
	if err != nil {
		return err
	}

	result := res.GetPayload()

	d.SetId(strconv.FormatInt(deviceID, 10))
	d.Set("content", result.Content)

	if result.Configtemplate != nil {
		d.Set("config_template_id", result.Configtemplate.ID)
		d.Set("config_template_name", result.Configtemplate.Name)
	}

	return nil
}
