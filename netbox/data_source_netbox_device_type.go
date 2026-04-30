package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxDeviceType() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDeviceTypeRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"is_full_depth": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"manufacturer": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"model": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subdevice_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"u_height": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"airflow": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"weight_unit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_platform_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"exclude_from_utilization": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			customFieldsKey: {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceNetboxDeviceTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	params := dcim.NewDcimDeviceTypesListParams()

	params.Limit = int64ToPtr(2)
	if manufacturer, ok := d.Get("manufacturer").(string); ok && manufacturer != "" {
		params.Manufacturer = &manufacturer
	}
	if model, ok := d.Get("model").(string); ok && model != "" {
		params.Model = &model
	}
	if part, ok := d.Get("part_number").(string); ok && part != "" {
		params.PartNumber = &part
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.Slug = &slug
	}

	res, err := api.Dcim.DcimDeviceTypesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one device type returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no device type found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("is_full_depth", result.IsFullDepth)
	d.Set("manufacturer_id", result.Manufacturer.ID)
	d.Set("model", result.Model)
	d.Set("part_number", result.PartNumber)
	if result.SubdeviceRole != nil && result.SubdeviceRole.Value != nil {
		d.Set("subdevice_role", *result.SubdeviceRole.Value)
	} else {
		d.Set("subdevice_role", "")
	}
	d.Set("slug", result.Slug)
	d.Set("u_height", result.UHeight)

	if result.Airflow != nil && result.Airflow.Value != nil {
		d.Set("airflow", *result.Airflow.Value)
	} else {
		d.Set("airflow", "")
	}
	d.Set("weight", result.Weight)
	if result.WeightUnit != nil && result.WeightUnit.Value != nil {
		d.Set("weight_unit", *result.WeightUnit.Value)
	} else {
		d.Set("weight_unit", "")
	}
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	if result.DefaultPlatform != nil {
		d.Set("default_platform_id", result.DefaultPlatform.ID)
	} else {
		d.Set("default_platform_id", nil)
	}
	d.Set("exclude_from_utilization", result.ExcludeFromUtilization)
	d.Set(customFieldsKey, getCustomFields(result.CustomFields))

	return nil
}
