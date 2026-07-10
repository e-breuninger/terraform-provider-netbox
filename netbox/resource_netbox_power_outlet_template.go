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

func resourceNetboxPowerOutletTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxPowerOutletTemplateCreate,
		ReadContext:   resourceNetboxPowerOutletTemplateRead,
		UpdateContext: resourceNetboxPowerOutletTemplateUpdate,
		DeleteContext: resourceNetboxPowerOutletTemplateDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/poweroutlettemplate/):

> A template for a power outlet that will be created on all instantiations of the parent device type. See the power outlet documentation for more detail.`,
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
				Description: "One of [iec-60320-c5, iec-60320-c7, iec-60320-c13, iec-60320-c15, iec-60320-c19, iec-60320-c21, iec-60309-p-n-e-4h, iec-60309-p-n-e-6h, iec-60309-p-n-e-9h, iec-60309-2p-e-4h, iec-60309-2p-e-6h, iec-60309-2p-e-9h, iec-60309-3p-e-4h, iec-60309-3p-e-6h, iec-60309-3p-e-9h, iec-60309-3p-n-e-4h, iec-60309-3p-n-e-6h, iec-60309-3p-n-e-9h, nema-1-15r, nema-5-15r, nema-5-20r, nema-5-30r, nema-5-50r, nema-6-15r, nema-6-20r, nema-6-30r, nema-6-50r, nema-10-30r, nema-10-50r, nema-14-20r, nema-14-30r, nema-14-50r, nema-14-60r, nema-15-15r, nema-15-20r, nema-15-30r, nema-15-50r, nema-15-60r, nema-l1-15r, nema-l5-15r, nema-l5-20r, nema-l5-30r, nema-l5-50r, nema-l6-15r, nema-l6-20r, nema-l6-30r, nema-l6-50r, nema-l10-30r, nema-l14-20r, nema-l14-30r, nema-l14-50r, nema-l14-60r, nema-l15-20r, nema-l15-30r, nema-l15-50r, nema-l15-60r, nema-l21-20r, nema-l21-30r, nema-l22-30r, CS6360C, CS6364C, CS8164C, CS8264C, CS8364C, CS8464C, ita-e, ita-f, ita-g, ita-h, ita-i, ita-j, ita-k, ita-l, ita-m, ita-n, ita-o, ita-multistandard, usb-a, usb-micro-b, usb-c, dc-terminal, hdot-cx, saf-d-grid, neutrik-powercon-20a, neutrik-powercon-32a, neutrik-powercon-true1, neutrik-powercon-true1-top, ubiquiti-smartpower, hardwired, other]",
			},
			"power_port_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of a power port template on the same parent device type or module type.",
			},
			"feed_leg": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "One of [A, B, C]",
				ValidateFunc: validation.StringInSlice([]string{"A", "B", "C"}, false),
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

func resourceNetboxPowerOutletTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	data := models.WritablePowerOutletTemplate{
		Name:        &name,
		Description: d.Get("description").(string),
		Label:       d.Get("label").(string),
		Type:        d.Get("type").(string),
		PowerPort:   getOptionalInt(d, "power_port_id"),
		FeedLeg:     d.Get("feed_leg").(string),
	}

	if deviceTypeID, ok := d.Get("device_type_id").(int); ok && deviceTypeID != 0 {
		data.DeviceType = int64ToPtr(int64(deviceTypeID))
	}
	if moduleTypeID, ok := d.Get("module_type_id").(int); ok && moduleTypeID != 0 {
		data.ModuleType = int64ToPtr(int64(moduleTypeID))
	}
	params := dcim.NewDcimPowerOutletTemplatesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimPowerOutletTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxPowerOutletTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimPowerOutletTemplatesReadParams().WithID(id)

	res, err := api.Dcim.DcimPowerOutletTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimPowerOutletTemplatesReadDefault); ok {
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
	if tmpl.PowerPort != nil {
		d.Set("power_port_id", tmpl.PowerPort.ID)
	}
	if tmpl.FeedLeg != nil {
		d.Set("feed_leg", tmpl.FeedLeg.Value)
	}
	if tmpl.DeviceType != nil {
		d.Set("device_type_id", tmpl.DeviceType.ID)
	}
	if tmpl.ModuleType != nil {
		d.Set("module_type_id", tmpl.ModuleType.ID)
	}

	return diags
}

func resourceNetboxPowerOutletTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)

	data := models.WritablePowerOutletTemplate{
		Name:        &name,
		Description: d.Get("description").(string),
		Label:       d.Get("label").(string),
		Type:        d.Get("type").(string),
		PowerPort:   getOptionalInt(d, "power_port_id"),
		FeedLeg:     d.Get("feed_leg").(string),
	}

	if d.HasChange("device_type_id") {
		deviceTypeID := int64(d.Get("device_type_id").(int))
		data.DeviceType = &deviceTypeID
	}

	if d.HasChange("module_type_id") {
		moduleTypeID := int64(d.Get("module_type_id").(int))
		data.ModuleType = &moduleTypeID
	}

	params := dcim.NewDcimPowerOutletTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimPowerOutletTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxPowerOutletTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerOutletTemplatesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimPowerOutletTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimPowerOutletTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
