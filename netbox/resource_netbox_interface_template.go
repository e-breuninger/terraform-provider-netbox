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

func resourceNetboxInterfaceTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxInterfaceTemplateCreate,
		ReadContext:   resourceNetboxInterfaceTemplateRead,
		UpdateContext: resourceNetboxInterfaceTemplateUpdate,
		DeleteContext: resourceNetboxInterfaceTemplateDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/interfacetemplate/):

> A template for a network interface that will be created on all instantiations of the parent device type. See the interface documentation for more detail.`,
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
				Type:     schema.TypeString,
				Required: true,
			},
			"mgmt_only": {
				Type:     schema.TypeBool,
				Optional: true,
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

func resourceNetboxInterfaceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	label := d.Get("label").(string)
	interfaceType := d.Get("type").(string)
	mgmtOnly := d.Get("mgmt_only").(bool)

	data := models.WritableInterfaceTemplate{
		Name:        &name,
		Description: description,
		Label:       label,
		Type:        &interfaceType,
		MgmtOnly:    mgmtOnly,
	}

	if deviceTypeID, ok := d.Get("device_type_id").(int); ok && deviceTypeID != 0 {
		data.DeviceType = int64ToPtr(int64(deviceTypeID))
	}
	if moduleTypeID, ok := d.Get("module_type_id").(int); ok && moduleTypeID != 0 {
		data.ModuleType = int64ToPtr(int64(moduleTypeID))
	}
	params := dcim.NewDcimInterfaceTemplatesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimInterfaceTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxInterfaceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimInterfaceTemplatesReadParams().WithID(id)

	res, err := api.Dcim.DcimInterfaceTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimInterfaceTemplatesReadDefault); ok {
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
	d.Set("type", tmpl.Type.Value)
	d.Set("mgmt_only", tmpl.MgmtOnly)

	if tmpl.DeviceType != nil {
		d.Set("device_type_id", tmpl.DeviceType.ID)
	}
	if tmpl.ModuleType != nil {
		d.Set("module_type_id", tmpl.ModuleType.ID)
	}

	return diags
}

func resourceNetboxInterfaceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	label := d.Get("label").(string)
	interfaceType := d.Get("type").(string)
	mgmtOnly := d.Get("mgmt_only").(bool)

	data := models.WritableInterfaceTemplate{
		Name:        &name,
		Description: description,
		Label:       label,
		Type:        &interfaceType,
		MgmtOnly:    mgmtOnly,
	}

	if d.HasChange("device_type_id") {
		deviceTypeID := int64(d.Get("device_type_id").(int))
		data.DeviceType = &deviceTypeID
	}

	if d.HasChange("module_type_id") {
		moduleTypeID := int64(d.Get("module_type_id").(int))
		data.ModuleType = &moduleTypeID
	}

	params := dcim.NewDcimInterfaceTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimInterfaceTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxInterfaceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimInterfaceTemplatesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimInterfaceTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimInterfaceTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
