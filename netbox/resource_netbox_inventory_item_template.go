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

func resourceNetboxInventoryItemTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxInventoryItemTemplateCreate,
		ReadContext:   resourceNetboxInventoryItemTemplateRead,
		UpdateContext: resourceNetboxInventoryItemTemplateUpdate,
		DeleteContext: resourceNetboxInventoryItemTemplateDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/inventoryitemtemplate/):

> A template for an inventory item that will be created on all instantiations of the parent device type. See the inventory item documentation for more detail.`,
		Schema: map[string]*schema.Schema{
			"device_type_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"part_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxInventoryItemTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	data := models.WritableInventoryItemTemplate{
		DeviceType:   int64ToPtr(int64(d.Get("device_type_id").(int))),
		Name:         strToPtr(d.Get("name").(string)),
		Parent:       getOptionalInt(d, "parent_id"),
		Label:        getOptionalStr(d, "label", false),
		Role:         getOptionalInt(d, "role_id"),
		Manufacturer: getOptionalInt(d, "manufacturer_id"),
		PartID:       getOptionalStr(d, "part_id", false),
		Description:  getOptionalStr(d, "description", false),
	}

	params := dcim.NewDcimInventoryItemTemplatesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimInventoryItemTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxInventoryItemTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimInventoryItemTemplatesReadParams().WithID(id)

	res, err := api.Dcim.DcimInventoryItemTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimInventoryItemTemplatesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	item := res.GetPayload()

	if item.DeviceType != nil {
		d.Set("device_type_id", item.DeviceType.ID)
	}

	d.Set("name", item.Name)
	d.Set("parent_id", item.Parent)
	d.Set("label", item.Label)

	if item.Role != nil {
		d.Set("role_id", item.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	if item.Manufacturer != nil {
		d.Set("manufacturer_id", item.Manufacturer.ID)
	} else {
		d.Set("manufacturer_id", nil)
	}

	d.Set("part_id", item.PartID)
	d.Set("description", item.Description)

	return diags
}

func resourceNetboxInventoryItemTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableInventoryItemTemplate{
		DeviceType:   int64ToPtr(int64(d.Get("device_type_id").(int))),
		Name:         strToPtr(d.Get("name").(string)),
		Parent:       getOptionalInt(d, "parent_id"),
		Label:        getOptionalStr(d, "label", true),
		Role:         getOptionalInt(d, "role_id"),
		Manufacturer: getOptionalInt(d, "manufacturer_id"),
		PartID:       getOptionalStr(d, "part_id", true),
		Description:  getOptionalStr(d, "description", true),
	}

	params := dcim.NewDcimInventoryItemTemplatesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimInventoryItemTemplatesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxInventoryItemTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimInventoryItemTemplatesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimInventoryItemTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimInventoryItemTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
