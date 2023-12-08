package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxDeviceModuleBay() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxDeviceModuleBayRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/modulebay/):

> Module bays represent a space or slot within a device in which a field-replaceable module may be installed. A common example is that of a chassis-based switch such as the Cisco Nexus 9000 or Juniper EX9200. Modules in turn hold additional components that become available to the parent device.`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//"position": {
			//	Type:     schema.TypeInt,
			//	Optional: true,
			//},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
	}
}

func dataSourceNetboxDeviceModuleBayRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimModuleBaysListParams()

	params.Limit = int64ToPtr(2)
	if deviceId, ok := d.Get("device_id").(int); ok && deviceId != 0 {
		deviceId := strconv.Itoa(deviceId)
		params.SetDeviceID(&deviceId)
	}
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.SetName(&name)
	}
	if label, ok := d.Get("label").(string); ok && label != "" {
		params.SetLabel(&label)
	}

	res, err := api.Dcim.DcimModuleBaysList(params, nil)
	if err != nil {
		return err
	}

	if count := *res.GetPayload().Count; count != 1 {
		return fmt.Errorf("expected one `netbox_device_module_bay`, but got %d", count)
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("device_id", result.Device.ID)
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("label", result.Label)
	//d.Set("position", result.Position)
	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))

	return nil
}
