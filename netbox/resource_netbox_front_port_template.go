package netbox

import (
	"context"
	"regexp"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxFrontPortTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxFrontPortTemplateCreate,
		ReadContext:   resourceNetboxFrontPortTemplateRead,
		UpdateContext: resourceNetboxFrontPortTemplateUpdate,
		DeleteContext: resourceNetboxFrontPortTemplateDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/frontporttemplate/):

> A template for a front port that will be created on all instantiations of the parent device type. See the front port documentation for more detail.`,
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
				Required:    true,
				Description: "One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other]",
			},
			"rear_port_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The rear port template this front port template is mapped to.",
			},
			"rear_port_position": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntBetween(1, 1024),
				Description:  "Position on the rear port this front port template maps to.",
			},
			"color_hex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[0-9a-f]{6}$"), "Must be hex color string"),
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

func resourceNetboxFrontPortTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	portType := d.Get("type").(string)
	rearPortID := int64(d.Get("rear_port_id").(int))

	data := models.WritableFrontPortTemplate{
		Name:             &name,
		Description:      d.Get("description").(string),
		Label:            d.Get("label").(string),
		Type:             &portType,
		RearPort:         &rearPortID,
		RearPortPosition: int64(d.Get("rear_port_position").(int)),
		Color:            d.Get("color_hex").(string),
	}

	if deviceTypeID, ok := d.Get("device_type_id").(int); ok && deviceTypeID != 0 {
		data.DeviceType = int64ToPtr(int64(deviceTypeID))
	}
	if moduleTypeID, ok := d.Get("module_type_id").(int); ok && moduleTypeID != 0 {
		data.ModuleType = int64ToPtr(int64(moduleTypeID))
	}
	params := dcim.NewDcimFrontPortTemplatesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimFrontPortTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxFrontPortTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimFrontPortTemplatesReadParams().WithID(id)

	res, err := api.Dcim.DcimFrontPortTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimFrontPortTemplatesReadDefault); ok {
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
	d.Set("rear_port_position", tmpl.RearPortPosition)
	d.Set("color_hex", tmpl.Color)

	if tmpl.Type != nil {
		d.Set("type", tmpl.Type.Value)
	}
	if tmpl.RearPort != nil {
		d.Set("rear_port_id", tmpl.RearPort.ID)
	}
	if tmpl.DeviceType != nil {
		d.Set("device_type_id", tmpl.DeviceType.ID)
	}
	if tmpl.ModuleType != nil {
		d.Set("module_type_id", tmpl.ModuleType.ID)
	}

	return diags
}

func resourceNetboxFrontPortTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	portType := d.Get("type").(string)
	rearPortID := int64(d.Get("rear_port_id").(int))

	data := models.WritableFrontPortTemplate{
		Name:             &name,
		Description:      d.Get("description").(string),
		Label:            d.Get("label").(string),
		Type:             &portType,
		RearPort:         &rearPortID,
		RearPortPosition: int64(d.Get("rear_port_position").(int)),
		Color:            d.Get("color_hex").(string),
	}

	if d.HasChange("device_type_id") {
		deviceTypeID := int64(d.Get("device_type_id").(int))
		data.DeviceType = &deviceTypeID
	}

	if d.HasChange("module_type_id") {
		moduleTypeID := int64(d.Get("module_type_id").(int))
		data.ModuleType = &moduleTypeID
	}

	params := dcim.NewDcimFrontPortTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimFrontPortTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxFrontPortTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimFrontPortTemplatesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimFrontPortTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimFrontPortTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
