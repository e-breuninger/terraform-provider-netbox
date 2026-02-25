package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxDeviceModuleBay() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceModuleBayCreate,
		Read:   resourceNetboxDeviceModuleBayRead,
		Update: resourceNetboxDeviceModuleBayUpdate,
		Delete: resourceNetboxDeviceModuleBayDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/modulebay/):

> Module bays represent a space or slot within a device in which a field-replaceable module may be installed. A common example is that of a chassis-based switch such as the Cisco Nexus 9000 or Juniper EX9200. Modules in turn hold additional components that become available to the parent device.`,

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
			"position": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
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

func resourceNetboxDeviceModuleBayCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableModuleBay{
		Device:      int64ToPtr(int64(d.Get("device_id").(int))),
		Name:        strToPtr(d.Get("name").(string)),
		Label:       getOptionalStr(d, "label", false),
		Position:    getOptionalStr(d, "position", false),
		Description: getOptionalStr(d, "description", false),
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

	params := dcim.NewDcimModuleBaysCreateParams().WithData(&data)

	res, err := api.Dcim.DcimModuleBaysCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceModuleBayRead(d, m)
}

func resourceNetboxDeviceModuleBayRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModuleBaysReadParams().WithID(id)

	res, err := api.Dcim.DcimModuleBaysRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimModuleBaysReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	moduleBay := res.GetPayload()

	if moduleBay.Device != nil {
		d.Set("device_id", moduleBay.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", moduleBay.Name)
	d.Set("label", moduleBay.Label)
	d.Set("position", moduleBay.Position)
	d.Set("description", moduleBay.Description)

	cf := flattenCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, res.GetPayload().Tags)

	return nil
}

func resourceNetboxDeviceModuleBayUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableModuleBay{
		Device:      int64ToPtr(int64(d.Get("device_id").(int))),
		Name:        strToPtr(d.Get("name").(string)),
		Label:       getOptionalStr(d, "label", true),
		Position:    getOptionalStr(d, "position", true),
		Description: getOptionalStr(d, "description", true),
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

	params := dcim.NewDcimModuleBaysPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimModuleBaysPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDeviceModuleBayRead(d, m)
}

func resourceNetboxDeviceModuleBayDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModuleBaysDeleteParams().WithID(id)

	_, err := api.Dcim.DcimModuleBaysDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
