package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxVirtualDeviceContextStatusOptions = []string{"active", "planned", "offline"}

func resourceNetboxVirtualDeviceContext() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVirtualDeviceContextCreate,
		Read:   resourceNetboxVirtualDeviceContextRead,
		Update: resourceNetboxVirtualDeviceContextUpdate,
		Delete: resourceNetboxVirtualDeviceContextDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/virtualdevicecontext/):

> A virtual device context (VDC) represents a logical partition of a physical device. This is a common feature of high-end network and storage equipment, allowing a single chassis to be divided up into multiple logically independent devices.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"identifier": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 32767),
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  buildValidValueDescription(resourceNetboxVirtualDeviceContextStatusOptions),
				ValidateFunc: validation.StringInSlice(resourceNetboxVirtualDeviceContextStatusOptions, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
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

func resourceNetboxVirtualDeviceContextCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	status := d.Get("status").(string)

	data := &models.WritableVirtualDeviceContext{
		Name:        strToPtr(d.Get("name").(string)),
		Device:      int64ToPtr(int64(d.Get("device_id").(int))),
		Identifier:  getOptionalInt(d, "identifier"),
		Tenant:      getOptionalInt(d, "tenant_id"),
		Status:      &status,
		Description: getOptionalStr(d, "description", false),
		Comments:    getOptionalStr(d, "comments", false),
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

	params := dcim.NewDcimVirtualDeviceContextsCreateParams().WithData(data)

	res, err := api.Dcim.DcimVirtualDeviceContextsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualDeviceContextRead(d, m)
}

func resourceNetboxVirtualDeviceContextRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimVirtualDeviceContextsReadParams().WithID(id)

	res, err := api.Dcim.DcimVirtualDeviceContextsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimVirtualDeviceContextsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	vdc := res.GetPayload()
	d.Set("name", vdc.Name)

	if vdc.Device != nil {
		d.Set("device_id", vdc.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	if vdc.Identifier != nil {
		d.Set("identifier", vdc.Identifier)
	} else {
		d.Set("identifier", nil)
	}

	if vdc.Tenant != nil {
		d.Set("tenant_id", vdc.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if vdc.Status != nil {
		d.Set("status", vdc.Status.Value)
	} else {
		d.Set("status", nil)
	}
	d.Set("description", vdc.Description)
	d.Set("comments", vdc.Comments)

	cf := getCustomFields(vdc.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, vdc.Tags)

	return nil
}

func resourceNetboxVirtualDeviceContextUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	status := d.Get("status").(string)

	data := models.WritableVirtualDeviceContext{
		Name:        strToPtr(d.Get("name").(string)),
		Device:      int64ToPtr(int64(d.Get("device_id").(int))),
		Identifier:  getOptionalInt(d, "identifier"),
		Tenant:      getOptionalInt(d, "tenant_id"),
		Status:      &status,
		Description: getOptionalStr(d, "description", true),
		Comments:    getOptionalStr(d, "comments", true),
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

	params := dcim.NewDcimVirtualDeviceContextsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimVirtualDeviceContextsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVirtualDeviceContextRead(d, m)
}

func resourceNetboxVirtualDeviceContextDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimVirtualDeviceContextsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimVirtualDeviceContextsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimVirtualDeviceContextsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
