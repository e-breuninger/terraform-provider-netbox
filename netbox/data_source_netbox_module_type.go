package netbox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxModuleType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxModuleTypeRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletype/):

> A module type represents a specific make and model of hardware component which is installable within a device's module bay and has its own child components. For example, consider a chassis-based switch or router with a number of field-replaceable line cards. Each line card has its own model number and includes a certain set of components such as interfaces. Each module type may have a manufacturer, model number, and part number assigned to it.`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"weight_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"weight"},
				Description:  "One of [kg, g, lb, oz]",
				ValidateFunc: validation.StringInSlice([]string{"kg", "g", "lb", "oz"}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
	}
}

func dataSourceNetboxModuleTypeRead(d *schema.ResourceData, m interface{}) error {
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
