package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxDevicePowerOutlet() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDevicePowerOutletCreate,
		Read:   resourceNetboxDevicePowerOutletRead,
		Update: resourceNetboxDevicePowerOutletUpdate,
		Delete: resourceNetboxDevicePowerOutletDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/poweroutlet/):

> Power outlets represent the outlets on a power distribution unit (PDU) or other device that supplies power to dependent devices. Each power port may be assigned a physical type, and may be associated with a specific feed leg (where three-phase power is used) and/or a specific upstream power port. This association can be used to model the distribution of power within a device.

For example, imagine a PDU with one power port which draws from a three-phase feed and 48 power outlets arranged into three banks of 16 outlets each. Outlets 1-16 would be associated with leg A on the port, and outlets 17-32 and 33-48 would be associated with legs B and C, respectively.`,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"module_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other]",
			},
			"power_port_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"feed_leg": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "One of [A, B, C]",
				ValidateFunc: validation.StringInSlice([]string{"A", "B", "C"}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mark_connected": {
				Type:    schema.TypeBool,
				Default: false,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDevicePowerOutletCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritablePowerOutlet{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", false),
		Type:          getOptionalStr(d, "type", false),
		PowerPort:     getOptionalInt(d, "power_port_id"),
		FeedLeg:       getOptionalStr(d, "feed_leg", false),
		Description:   getOptionalStr(d, "description", false),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerOutletsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimPowerOutletsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDevicePowerOutletRead(d, m)
}

func resourceNetboxDevicePowerOutletRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerOutletsReadParams().WithID(id)

	res, err := api.Dcim.DcimPowerOutletsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimPowerOutletsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	powerOutlet := res.GetPayload()

	if powerOutlet.Device != nil {
		d.Set("device_id", powerOutlet.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", powerOutlet.Name)

	if powerOutlet.Module != nil {
		d.Set("module_id", powerOutlet.Module.ID)
	} else {
		d.Set("module_id", nil)
	}

	d.Set("label", powerOutlet.Label)

	if powerOutlet.Type != nil {
		d.Set("type", powerOutlet.Type.Value)
	} else {
		d.Set("type", nil)
	}

	if powerOutlet.PowerPort != nil {
		d.Set("power_port_id", powerOutlet.PowerPort.ID)
	} else {
		d.Set("power_port_id", nil)
	}

	if powerOutlet.FeedLeg != nil {
		d.Set("feed_leg", powerOutlet.FeedLeg.Value)
	} else {
		d.Set("feed_leg", nil)
	}

	d.Set("description", powerOutlet.Description)
	d.Set("mark_connected", powerOutlet.MarkConnected)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxDevicePowerOutletUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritablePowerOutlet{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", false),
		Type:          getOptionalStr(d, "type", false),
		PowerPort:     getOptionalInt(d, "power_port_id"),
		FeedLeg:       getOptionalStr(d, "feed_leg", false),
		Description:   getOptionalStr(d, "description", false),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerOutletsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimPowerOutletsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDevicePowerOutletRead(d, m)
}

func resourceNetboxDevicePowerOutletDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerOutletsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimPowerOutletsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
