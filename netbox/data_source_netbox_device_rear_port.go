package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxDeviceRearPort() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDeviceRearPortRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletype/):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "One of [8p8c, 8p6c, 8p4c, 8p2c, 6p6c, 6p4c, 6p2c, 4p4c, 4p2c, gg45, tera-4p, tera-2p, tera-1p, 110-punch, bnc, f, n, mrj21, fc, lc, lc-pc, lc-upc, lc-apc, lsh, lsh-pc, lsh-upc, lsh-apc, mpo, mtrj, sc, sc-pc, sc-upc, sc-apc, st, cs, sn, sma-905, sma-906, urm-p2, urm-p4, urm-p8, splice, other]",
			},
			"positions": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"module_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"color_hex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mark_connected": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
	}
}

func dataSourceNetboxDeviceRearPortRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimModuleTypesListParams()

	params.Limit = int64ToPtr(2)
	if manufacturerID, ok := d.Get("manufacturer_id").(int); ok && manufacturerID != 0 {
		manufacturerID := strconv.Itoa(manufacturerID)
		params.SetManufacturerID(&manufacturerID)
	}
	if model, ok := d.Get("model").(string); ok && model != "" {
		params.SetModel(&model)
	}

	res, err := api.Dcim.DcimModuleTypesList(params, nil)
	if err != nil {
		return err
	}

	if count := *res.GetPayload().Count; count != 1 {
		return fmt.Errorf("expected one `netbox_module_type`, but got %d", count)
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("manufacturer_id", result.Manufacturer.ID)
	d.Set("model", result.Model)
	d.Set("part_number", result.PartNumber)
	d.Set("weight", result.Weight)

	if result.WeightUnit != nil {
		d.Set("weight_unit", result.WeightUnit.Value)
	} else {
		d.Set("weight_unit", nil)
	}

	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))

	return nil
}
