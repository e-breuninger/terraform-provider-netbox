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
			"u_height": {
				Type:     schema.TypeFloat,
				Computed: true,
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
	d.Set("slug", result.Slug)
	d.Set("u_height", result.UHeight)
	return nil
}
