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

func resourceNetboxPowerPortTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxPowerPortTemplateCreate,
		ReadContext:   resourceNetboxPowerPortTemplateRead,
		UpdateContext: resourceNetboxPowerPortTemplateUpdate,
		DeleteContext: resourceNetboxPowerPortTemplateDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerporttemplate/):

> A template for a power port that will be created on all instantiations of the parent device type. See the power port documentation for more detail.`,
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
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "One of [iec-60320-c6, iec-60320-c8, iec-60320-c14, iec-60320-c16, iec-60320-c20, iec-60320-c22, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15p, nema-5-15p, nema-5-20p, nema-5-30p, nema-5-50p, nema-6-15p, nema-6-20p, nema-6-30p, nema-6-50p, nema-10-30p, nema-10-50p, nema-14-20p, nema-14-30p, nema-14-50p, nema-14-60p, nema-15-15p, nema-15-20p, nema-15-30p, nema-15-50p, nema-15-60p, nema-l1-15p, nema-l5-15p, nema-l5-20p, nema-l5-30p, nema-l5-50p, nema-l6-15p, nema-l6-20p, nema-l6-30p, nema-l6-50p, nema-l10-30p, nema-l14-20p, nema-l14-30p, nema-l14-50p, nema-l14-60p, nema-l15-20p, nema-l15-30p, nema-l15-50p, nema-l15-60p, nema-l21-20p, nema-l21-30p, nema-l22-30p, cs6361c, cs6365c, cs8165c, cs8265c, cs8365c, cs8465c, ita-c, ita-e, ita-f, ita-ef, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, usb-a, usb-b, usb-c, usb-mini-a, usb-mini-b, usb-micro-a, usb-micro-b, usb-micro-ab, usb-3-b, usb-3-micro-b, dc-terminal, saf-d-grid, neutrik-powercon-20, neutrik-powercon-32, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other]",
			},
			"maximum_draw": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum power draw in watts.",
			},
			"allocated_draw": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Allocated power draw in watts.",
			},
			"device_type_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"device_type_id", "module_type_id"},
				ForceNew:     true,
			},
			"module_type_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"device_type_id", "module_type_id"},
				ForceNew:     true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxPowerPortTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	data := models.WritablePowerPortTemplate{
		Name:          &name,
		Description:   d.Get("description").(string),
		Label:         d.Get("label").(string),
		Type:          d.Get("type").(string),
		MaximumDraw:   getOptionalInt(d, "maximum_draw"),
		AllocatedDraw: getOptionalInt(d, "allocated_draw"),
	}

	if deviceTypeID, ok := d.Get("device_type_id").(int); ok && deviceTypeID != 0 {
		data.DeviceType = int64ToPtr(int64(deviceTypeID))
	}
	if moduleTypeID, ok := d.Get("module_type_id").(int); ok && moduleTypeID != 0 {
		data.ModuleType = int64ToPtr(int64(moduleTypeID))
	}
	params := dcim.NewDcimPowerPortTemplatesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimPowerPortTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxPowerPortTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimPowerPortTemplatesReadParams().WithID(id)

	res, err := api.Dcim.DcimPowerPortTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimPowerPortTemplatesReadDefault); ok {
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

	if tmpl.Type != nil {
		d.Set("type", tmpl.Type.Value)
	}
	if tmpl.MaximumDraw != nil {
		d.Set("maximum_draw", tmpl.MaximumDraw)
	}
	if tmpl.AllocatedDraw != nil {
		d.Set("allocated_draw", tmpl.AllocatedDraw)
	}
	if tmpl.DeviceType != nil {
		d.Set("device_type_id", tmpl.DeviceType.ID)
	}
	if tmpl.ModuleType != nil {
		d.Set("module_type_id", tmpl.ModuleType.ID)
	}

	return diags
}

func resourceNetboxPowerPortTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)

	data := models.WritablePowerPortTemplate{
		Name:          &name,
		Description:   d.Get("description").(string),
		Label:         d.Get("label").(string),
		Type:          d.Get("type").(string),
		MaximumDraw:   getOptionalInt(d, "maximum_draw"),
		AllocatedDraw: getOptionalInt(d, "allocated_draw"),
	}

	if d.HasChange("device_type_id") {
		deviceTypeID := int64(d.Get("device_type_id").(int))
		data.DeviceType = &deviceTypeID
	}

	if d.HasChange("module_type_id") {
		moduleTypeID := int64(d.Get("module_type_id").(int))
		data.ModuleType = &moduleTypeID
	}

	params := dcim.NewDcimPowerPortTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimPowerPortTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxPowerPortTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerPortTemplatesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimPowerPortTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimPowerPortTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
