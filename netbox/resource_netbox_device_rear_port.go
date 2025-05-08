package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceRearPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceRearPortCreate,
		Read:   resourceNetboxDeviceRearPortRead,
		Update: resourceNetboxDeviceRearPortUpdate,
		Delete: resourceNetboxDeviceRearPortDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rearport/):

> Like front ports, rear ports are pass-through ports which represent the continuation of a path from one cable to the next. Each rear port is defined with its physical type and a number of positions: Rear ports with more than one position can be mapped to multiple front ports. This can be useful for modeling instances where multiple paths share a common cable (for example, six discrete two-strand fiber connections sharing a 12-strand MPO cable).`,

		Schema: map[string]*schema.Schema{
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
				Required:    true,
				Description: "One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other]",
			},
			"positions": {
				Type:     schema.TypeInt,
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

func resourceNetboxDeviceRearPortCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableRearPort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Name:          strToPtr(d.Get("name").(string)),
		Type:          strToPtr(d.Get("type").(string)),
		Positions:     int64(d.Get("positions").(int)),
		Module:        getOptionalInt(d, "module_id"),
		Label:         getOptionalStr(d, "label", false),
		Color:         getOptionalStr(d, "color_hex", false),
		Description:   getOptionalStr(d, "description", false),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimRearPortsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimRearPortsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceRearPortRead(d, m)
}

func resourceNetboxDeviceRearPortRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRearPortsReadParams().WithID(id)

	res, err := api.Dcim.DcimRearPortsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimRearPortsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	rearPort := res.GetPayload()

	if rearPort.Device != nil {
		d.Set("device_id", rearPort.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", rearPort.Name)

	if rearPort.Type != nil {
		d.Set("type", rearPort.Type.Value)
	} else {
		d.Set("type", nil)
	}

	d.Set("positions", rearPort.Positions)

	if rearPort.Module != nil {
		d.Set("module_id", rearPort.Module.ID)
	} else {
		d.Set("module_id", nil)
	}

	d.Set("label", rearPort.Label)
	d.Set("color_hex", rearPort.Color)
	d.Set("description", rearPort.Description)
	d.Set("mark_connected", rearPort.MarkConnected)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxDeviceRearPortUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableRearPort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Name:          strToPtr(d.Get("name").(string)),
		Type:          strToPtr(d.Get("type").(string)),
		Positions:     int64(d.Get("positions").(int)),
		Module:        getOptionalInt(d, "module_id"),
		Label:         getOptionalStr(d, "label", true),
		Color:         getOptionalStr(d, "color_hex", false),
		Description:   getOptionalStr(d, "description", true),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimRearPortsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimRearPortsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDeviceRearPortRead(d, m)
}

func resourceNetboxDeviceRearPortDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRearPortsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimRearPortsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
