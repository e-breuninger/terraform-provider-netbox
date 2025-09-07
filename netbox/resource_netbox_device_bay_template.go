package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxDeviceBayTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxDeviceBayTemplateCreate,
		ReadContext:   resourceNetboxDeviceBayTemplateRead,
		UpdateContext: resourceNetboxDeviceBayTemplateUpdate,
		DeleteContext: resourceNetboxDeviceBayTemplateDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/devicebaytemplate/):

> A template for a device bay that will be created on all instantiations of the parent device type.`,
		Schema: map[string]*schema.Schema{
			"device_type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceBayTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	label := d.Get("label").(string)

	data := models.WritableDeviceBayTemplate{
		Name:        &name,
		Description: description,
		Label:       label,
	}

	if deviceTypeID, ok := d.Get("device_type_id").(int); ok && deviceTypeID != 0 {
		data.DeviceType = int64ToPtr(int64(deviceTypeID))
	}
	params := dcim.NewDcimDeviceBayTemplatesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDeviceBayTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxDeviceBayTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimDeviceBayTemplatesReadParams().WithID(id)

	res, err := api.Dcim.DcimDeviceBayTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDeviceBayTemplatesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	tmpl := res.GetPayload()

	d.Set("device_type_id", tmpl.DeviceType.ID)
	d.Set("name", tmpl.Name)
	d.Set("description", tmpl.Description)
	d.Set("label", tmpl.Label)

	return diags
}

func resourceNetboxDeviceBayTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	label := d.Get("label").(string)

	data := models.WritableDeviceBayTemplate{
		Name:        &name,
		Description: description,
		Label:       label,
	}

	if d.HasChange("device_type_id") {
		deviceTypeID := int64(d.Get("device_type_id").(int))
		data.DeviceType = &deviceTypeID
	}

	params := dcim.NewDcimDeviceBayTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimDeviceBayTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxDeviceBayTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDeviceBayTemplatesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimDeviceBayTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDeviceBayTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
