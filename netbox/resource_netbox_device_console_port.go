package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceConsolePort() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceConsolePortCreate,
		Read:   resourceNetboxDeviceConsolePortRead,
		Update: resourceNetboxDeviceConsolePortUpdate,
		Delete: resourceNetboxDeviceConsolePortDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/consoleport/):

> A console port provides connectivity to the physical console of a device. These are typically used for temporary access by someone who is physically near the device, or for remote out-of-band access provided via a networked console server.`,

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

func resourceNetboxDeviceConsolePortCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableConsolePort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", false),
		Type:          getOptionalStr(d, "type", false),
		Speed:         getOptionalInt(d, "speed"),
		Description:   getOptionalStr(d, "description", false),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimConsolePortsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimConsolePortsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceConsolePortRead(d, m)
}

func resourceNetboxDeviceConsolePortRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimConsolePortsReadParams().WithID(id)

	res, err := api.Dcim.DcimConsolePortsRead(params, nil)

	if err != nil {
		if errIntf, ok := err.(*dcim.DcimConsolePortsReadDefault); ok && errIntf.Code() == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	consolePort := res.GetPayload()

	if consolePort.Device != nil {
		d.Set("device_id", consolePort.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", consolePort.Name)

	if consolePort.Module != nil {
		d.Set("module_id", consolePort.Module.ID)
	} else {
		d.Set("module_id", nil)
	}

	d.Set("label", consolePort.Label)

	if consolePort.Type != nil {
		d.Set("type", consolePort.Type.Value)
	} else {
		d.Set("type", nil)
	}

	if consolePort.Speed != nil {
		d.Set("speed", consolePort.Speed.Value)
	} else {
		d.Set("speed", nil)
	}

	d.Set("description", consolePort.Description)
	d.Set("mark_connected", consolePort.MarkConnected)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxDeviceConsolePortUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableConsolePort{
		Device:        int64ToPtr(int64(d.Get("device_id").(int))),
		Module:        getOptionalInt(d, "module_id"),
		Name:          strToPtr(d.Get("name").(string)),
		Label:         getOptionalStr(d, "label", true),
		Type:          getOptionalStr(d, "type", false),
		Speed:         getOptionalInt(d, "speed"),
		Description:   getOptionalStr(d, "description", true),
		MarkConnected: d.Get("mark_connected").(bool),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimConsolePortsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimConsolePortsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDeviceConsolePortRead(d, m)
}

func resourceNetboxDeviceConsolePortDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimConsolePortsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimConsolePortsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
