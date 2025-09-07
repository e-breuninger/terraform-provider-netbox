package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDevicePowerPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDevicePowerPortCreate,
		Read:   resourceNetboxDevicePowerPortRead,
		Update: resourceNetboxDevicePowerPortUpdate,
		Delete: resourceNetboxDevicePowerPortDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerport/):

> A power port is a device component which draws power from some external source (e.g. an upstream power outlet), and generally represents a power supply internal to a device.`,

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
				Description: "One of [iec-60320-c6, iec-60320-c8, iec-60320-c14, iec-60320-c16, iec-60320-c20, iec-60320-c22, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15p, nema-5-15p, nema-5-20p, nema-5-30p, nema-5-50p, nema-6-15p, nema-6-20p, nema-6-30p, nema-6-50p, nema-10-30p, nema-10-50p, nema-14-20p, nema-14-30p, nema-14-50p, nema-14-60p, nema-15-15p, nema-15-20p, nema-15-30p, nema-15-50p, nema-15-60p, nema-l1-15p, nema-l5-15p, nema-l5-20p, nema-l5-30p, nema-l5-50p, nema-l6-15p, nema-l6-20p, nema-l6-30p, nema-l6-50p, nema-l10-30p, nema-l14-20p, nema-l14-30p, nema-l14-50p, nema-l14-60p, nema-l15-20p, nema-l15-30p, nema-l15-50p, nema-l15-60p, nema-l21-20p, nema-l21-30p, nema-l22-30p, cs6361c, cs6365c, cs8165c, cs8265c, cs8365c, cs8465c, ita-c, ita-e, ita-f, ita-ef, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, usb-3-b, usb-3-micro-b, dc-terminal, saf-d-grid, neutrik-powercon-20, neutrik-powercon-32, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other]",
			},
			"maximum_draw": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"allocated_draw": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mark_connected": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDevicePowerPortCreate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	data := models.WritablePowerPort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", false),
		Type:          getOptionalStr(d, "type", false),
		MaximumDraw:   getOptionalInt(d, "maximum_draw"),
		AllocatedDraw: getOptionalInt(d, "allocated_draw"),
		Description:   getOptionalStr(d, "description", false),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerPortsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimPowerPortsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDevicePowerPortRead(d, m)
}

func resourceNetboxDevicePowerPortRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerPortsReadParams().WithID(id)

	res, err := api.Dcim.DcimPowerPortsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimPowerPortsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	powerPort := res.GetPayload()

	if powerPort.Device != nil {
		d.Set("device_id", powerPort.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", powerPort.Name)

	if powerPort.Module != nil {
		d.Set("module_id", powerPort.Module.ID)
	} else {
		d.Set("module_id", nil)
	}

	d.Set("label", powerPort.Label)

	if powerPort.Type != nil {
		d.Set("type", powerPort.Type.Value)
	} else {
		d.Set("type", nil)
	}

	d.Set("maximum_draw", powerPort.MaximumDraw)
	d.Set("allocated_draw", powerPort.AllocatedDraw)

	d.Set("description", powerPort.Description)
	d.Set("mark_connected", powerPort.MarkConnected)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	state.readTags(d, res.GetPayload().Tags)

	return nil
}

func resourceNetboxDevicePowerPortUpdate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritablePowerPort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", true),
		Type:          getOptionalStr(d, "type", false),
		MaximumDraw:   getOptionalInt(d, "maximum_draw"),
		AllocatedDraw: getOptionalInt(d, "allocated_draw"),
		Description:   getOptionalStr(d, "description", true),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerPortsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimPowerPortsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDevicePowerPortRead(d, m)
}

func resourceNetboxDevicePowerPortDelete(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerPortsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimPowerPortsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
