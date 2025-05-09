package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceConsoleServerPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceConsoleServerPortCreate,
		Read:   resourceNetboxDeviceConsoleServerPortRead,
		Update: resourceNetboxDeviceConsoleServerPortUpdate,
		Delete: resourceNetboxDeviceConsoleServerPortDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/consoleserverport/):

> A console server is a device which provides remote access to the local consoles of connected devices. They are typically used to provide remote out-of-band access to network devices, and generally connect to console ports.`,

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
				Description: "One of [de-9, db-25, rj-11, rj-12, rj-45, mini-din-8, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, other]",
			},
			"speed": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "One of [1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200]",
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

func resourceNetboxDeviceConsoleServerPortCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableConsoleServerPort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", true),
		Type:          getOptionalStr(d, "type", false),
		Speed:         getOptionalInt(d, "speed"),
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

	params := dcim.NewDcimConsoleServerPortsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimConsoleServerPortsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceConsoleServerPortRead(d, m)
}

func resourceNetboxDeviceConsoleServerPortRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimConsoleServerPortsReadParams().WithID(id)

	res, err := api.Dcim.DcimConsoleServerPortsRead(params, nil)

	if err != nil {
		if errIntf, ok := err.(*dcim.DcimConsoleServerPortsReadDefault); ok && errIntf.Code() == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	consoleServerPort := res.GetPayload()

	if consoleServerPort.Device != nil {
		d.Set("device_id", consoleServerPort.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", consoleServerPort.Name)

	if consoleServerPort.Module != nil {
		d.Set("module_id", consoleServerPort.Module.ID)
	} else {
		d.Set("module_id", nil)
	}

	d.Set("label", consoleServerPort.Label)

	if consoleServerPort.Type != nil {
		d.Set("type", consoleServerPort.Type.Value)
	} else {
		d.Set("type", nil)
	}

	if consoleServerPort.Speed != nil {
		d.Set("speed", consoleServerPort.Speed.Value)
	} else {
		d.Set("speed", nil)
	}

	d.Set("description", consoleServerPort.Description)
	d.Set("mark_connected", consoleServerPort.MarkConnected)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, res.GetPayload().Tags)

	return nil
}

func resourceNetboxDeviceConsoleServerPortUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableConsoleServerPort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", true),
		Type:          getOptionalStr(d, "type", false),
		Speed:         getOptionalInt(d, "speed"),
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

	params := dcim.NewDcimConsoleServerPortsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimConsoleServerPortsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDeviceConsoleServerPortRead(d, m)
}

func resourceNetboxDeviceConsoleServerPortDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimConsoleServerPortsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimConsoleServerPortsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
