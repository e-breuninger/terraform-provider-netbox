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

func resourceNetboxModuleBayTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxModuleBayTemplateCreate,
		ReadContext:   resourceNetboxModuleBayTemplateRead,
		UpdateContext: resourceNetboxModuleBayTemplateUpdate,
		DeleteContext: resourceNetboxModuleBayTemplateDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/modulebaytemplate/):

> A template for a module bay that will be created on all instantiations of the parent device type. See the module bay documentation for more detail.`,
		Schema: map[string]*schema.Schema{
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
			"position": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identifier to reference when renaming installed components.",
			},
			"device_type_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxModuleBayTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	data := models.WritableModuleBayTemplate{
		Name:        &name,
		Description: d.Get("description").(string),
		Label:       d.Get("label").(string),
		Position:    d.Get("position").(string),
		DeviceType:  int64ToPtr(int64(d.Get("device_type_id").(int))),
	}

	params := dcim.NewDcimModuleBayTemplatesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimModuleBayTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxModuleBayTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimModuleBayTemplatesReadParams().WithID(id)

	res, err := api.Dcim.DcimModuleBayTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimModuleBayTemplatesReadDefault); ok {
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

	d.Set("name", tmpl.Name)
	d.Set("description", tmpl.Description)
	d.Set("label", tmpl.Label)
	d.Set("position", tmpl.Position)

	if tmpl.DeviceType != nil {
		d.Set("device_type_id", tmpl.DeviceType.ID)
	}

	return diags
}

func resourceNetboxModuleBayTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)

	data := models.WritableModuleBayTemplate{
		Name:        &name,
		Description: d.Get("description").(string),
		Label:       d.Get("label").(string),
		Position:    d.Get("position").(string),
		DeviceType:  int64ToPtr(int64(d.Get("device_type_id").(int))),
	}

	params := dcim.NewDcimModuleBayTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimModuleBayTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxModuleBayTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModuleBayTemplatesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimModuleBayTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimModuleBayTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
