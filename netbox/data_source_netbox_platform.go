package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxPlatform() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxPlatformRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataSourceNetboxPlatformRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	params := dcim.NewDcimPlatformsListParams()
	params.Name = &name
	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Dcim.DcimPlatformsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one platform returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no platform found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("slug", result.Slug)
	if result.Manufacturer != nil {
		d.Set("manufacturer_id", result.Manufacturer.ID)
	}
	return nil
}
