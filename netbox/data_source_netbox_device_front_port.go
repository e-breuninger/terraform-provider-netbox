package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxDeviceFrontPort() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDeviceFrontPortRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletype/):`,
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
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other]",
			},
			"rear_port_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rear_port_position": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"module_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"color_hex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mark_connected": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxDeviceFrontPortRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimFrontPortsListParams()

	params.Limit = int64ToPtr(2)
	if deviceID, ok := d.Get("device_id").(int); ok && deviceID != 0 {
		deviceID := strconv.Itoa(deviceID)
		params.SetDeviceID(&deviceID)
	}
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.SetName(&name)
	}

	res, err := api.Dcim.DcimFrontPortsList(params, nil)
	if err != nil {
		return err
	}

	if count := *res.GetPayload().Count; count != 1 {
		return fmt.Errorf("expected one `netbox_device_front_port`, but got %d", count)
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("device_id", result.Device.ID)
	d.Set("name", result.Name)
	d.Set("type", result.Type.Value)
	d.Set("rear_port_id", result.RearPort.ID)
	d.Set("rear_port_position", result.RearPortPosition)

	if result.Module != nil {
		d.Set("module_id", result.Module.ID)
	} else {
		d.Set("module_id", nil)
	}

	d.Set("label", result.Label)
	d.Set("color_hex", result.Color)
	d.Set("description", result.Description)
	d.Set("mark_connected", result.MarkConnected)

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))

	return nil
}
