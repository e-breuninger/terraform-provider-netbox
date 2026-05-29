package netbox

import (
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxManufacturer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxManufacturerRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):Fetches the list of manufacturers from netbox. Can optionally be filtered by name, slug or tag.

From the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/manufacturer/):

> A manufacturer represents the "make" of a device; e.g. Cisco or Dell. Each device type must be assigned to a manufacturer. (Inventory items and platforms may also be associated with manufacturers.)`,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      `The name of the field to filter by. Supported values are "name", "slug" and "tag".`,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"name", "slug", "tag"}, false)),
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The value to filter by. For "name" and "slug" this is a string, for "tag" this is the name of the tag.`,
						},
					},
				},
			},
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
				Description:      `The maximum number of items to return. Will return all items if not set`,
			},
			"manufacturers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxManufacturerRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := dcim.NewDcimManufacturersListParams()

	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	if filter, ok := d.GetOk("filter"); ok {
		filterParams := filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "name":
				params.Name = &vString
			case "tag":
				params.Tag = []string{vString}
			case "slug":
				params.Slug = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	paginationHelper := NewPaginationHelper(userLimit)
	var allManufacturers []*models.Manufacturer

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Dcim.DcimManufacturersList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch manufacturers at offset %d: %w", currentOffset, err)
		}

		payload := res.Payload
		allManufacturers = append(allManufacturers, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allManufacturers)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	var s []map[string]interface{}
	for _, v := range allManufacturers {
		mapping := make(map[string]interface{})
		mapping["id"] = v.ID
		if v.Description != "" {
			mapping["description"] = v.Description
		}
		if v.Slug != nil {
			mapping["slug"] = *v.Slug
		}
		if v.Name != nil {
			mapping["name"] = *v.Name
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("manufacturers", s)
}
