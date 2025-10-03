package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceBay() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceBayCreate,
		Read:   resourceNetboxDeviceBayRead,
		Update: resourceNetboxDeviceBayUpdate,
		Delete: resourceNetboxDeviceBayDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/devicebay/):

> Device bays represent a space or slot within a device in which a field-replaceable device may be installed. A common example is that of a chassis-based server that holds a number of blades which may contain device components that don't fit the module pattern. Devices in turn hold additional components that become available to the parent device.`,

		Schema: map[string]*schema.Schema{
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
			"installed_device_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
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

func resourceNetboxDeviceBayCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableDeviceBay{
		Device:          int64ToPtr(int64(d.Get("device_id").(int))),
		Name:            strToPtr(d.Get("name").(string)),
		Label:           getOptionalStr(d, "label", false),
		InstalledDevice: getOptionalInt(d, "installed_device_id"),
		Description:     getOptionalStr(d, "description", false),
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

	params := dcim.NewDcimDeviceBaysCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDeviceBaysCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceBayRead(d, m)
}

func resourceNetboxDeviceBayRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDeviceBaysReadParams().WithID(id)

	res, err := api.Dcim.DcimDeviceBaysRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimDeviceBaysReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	deviceBay := res.GetPayload()

	if deviceBay.Device != nil {
		d.Set("device_id", deviceBay.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", deviceBay.Name)
	d.Set("label", deviceBay.Label)
	if deviceBay.InstalledDevice != nil {
		d.Set("installed_device_id", deviceBay.InstalledDevice.ID)
	}
	d.Set("description", deviceBay.Description)

	cf := flattenCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, res.GetPayload().Tags)

	return nil
}

func resourceNetboxDeviceBayUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableDeviceBay{
		Device:          int64ToPtr(int64(d.Get("device_id").(int))),
		Name:            strToPtr(d.Get("name").(string)),
		Label:           getOptionalStr(d, "label", true),
		InstalledDevice: getOptionalInt(d, "installed_device_id"),
		Description:     getOptionalStr(d, "description", true),
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

	params := dcim.NewDcimDeviceBaysPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimDeviceBaysPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDeviceBayRead(d, m)
}

func resourceNetboxDeviceBayDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDeviceBaysDeleteParams().WithID(id)

	_, err := api.Dcim.DcimDeviceBaysDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
