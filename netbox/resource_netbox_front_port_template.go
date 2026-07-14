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

> A template for a front port that will be created on all instantiations of the parent device type. See the front port documentation for more detail.

This resource manages a single rear port mapping per front port template (the common case for patch panels); a template with multiple mappings created outside Terraform is read as its first mapping.`,
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
				Description: "ID of a rear port template on the same parent device type or module type.",
			},
			"rear_port_position": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(1, 1024),
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

func frontPortTemplateMapping(d *schema.ResourceData) []*models.FrontPortTemplateMapping {
	rearPortID := int64(d.Get("rear_port_id").(int))
	return []*models.FrontPortTemplateMapping{{
		Position:         int64ToPtr(1),
		RearPort:         &rearPortID,
		RearPortPosition: int64(d.Get("rear_port_position").(int)),
	}}
}

func resourceNetboxFrontPortTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	name := d.Get("name").(string)
	portType := d.Get("type").(string)

	data := models.WritableFrontPortTemplate{
		Name:        &name,
		Description: d.Get("description").(string),
		Label:       d.Get("label").(string),
		Type:        &portType,
		Color:       d.Get("color_hex").(string),
		// NetBox <= 4.4 requires the legacy singular rear_port field and ignores
		// rear_ports; NetBox >= 4.5 ignores the legacy field and expects the
		// rear_ports mapping array instead. Sending both keeps a single request
		// shape working across all supported NetBox versions.
		RearPort:         int64(d.Get("rear_port_id").(int)),
		RearPortPosition: int64(d.Get("rear_port_position").(int)),
		RearPorts:        frontPortTemplateMapping(d),
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

	return resourceNetboxFrontPortTemplateRead(ctx, d, m)
}

func resourceNetboxFrontPortTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

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
	d.Set("color_hex", tmpl.Color)

	if tmpl.Type != nil {
		d.Set("type", tmpl.Type.Value)
	}

	// NetBox >= 4.5 returns mappings in the rear_ports array and no longer
	// returns the legacy singular rear_port field; NetBox <= 4.4 does the
	// opposite.
	if len(tmpl.RearPorts) > 0 {
		if tmpl.RearPorts[0].RearPort != nil {
			d.Set("rear_port_id", *tmpl.RearPorts[0].RearPort)
		}
		d.Set("rear_port_position", tmpl.RearPorts[0].RearPortPosition)
	} else if tmpl.RearPort != nil {
		d.Set("rear_port_id", tmpl.RearPort.ID)
		d.Set("rear_port_position", tmpl.RearPortPosition)
	} else {
		d.Set("rear_port_id", nil)
		d.Set("rear_port_position", tmpl.RearPortPosition)
	}

	if tmpl.DeviceType != nil {
		d.Set("device_type_id", tmpl.DeviceType.ID)
	}
	if tmpl.ModuleType != nil {
		d.Set("module_type_id", tmpl.ModuleType.ID)
	}

	return nil
}

func resourceNetboxFrontPortTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	portType := d.Get("type").(string)

	data := models.WritableFrontPortTemplate{
		Name:             &name,
		Description:      d.Get("description").(string),
		Label:            d.Get("label").(string),
		Type:             &portType,
		Color:            d.Get("color_hex").(string),
		RearPort:         int64(d.Get("rear_port_id").(int)),
		RearPortPosition: int64(d.Get("rear_port_position").(int)),
	}

	if d.HasChange("device_type_id") {
		deviceTypeID := int64(d.Get("device_type_id").(int))
		data.DeviceType = &deviceTypeID
	}

	if d.HasChange("module_type_id") {
		moduleTypeID := int64(d.Get("module_type_id").(int))
		data.ModuleType = &moduleTypeID
	}

	// NetBox >= 4.5 rejects re-sending a mapping that already exists (its
	// uniqueness validator does not exclude the object's own rows) but
	// replaces the mapping set atomically when it differs, and preserves it
	// when rear_ports is omitted. So the mapping is only sent when it
	// actually changed. NetBox <= 4.4 ignores rear_ports entirely and
	// applies the legacy fields above.
	if d.HasChange("rear_port_id") || d.HasChange("rear_port_position") {
		data.RearPorts = frontPortTemplateMapping(d)
	}
	params := dcim.NewDcimFrontPortTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimFrontPortTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxFrontPortTemplateRead(ctx, d, m)
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
