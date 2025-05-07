package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxLocations() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxLocationsRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of filter to apply to the API query when requesting locations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the field to filter on. Supported fields are: .",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value to pass to the specified filter.",
						},
					},
				},
			},
			"tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "A list of tags to filter on.",
			},
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
				Description:      "The limit of objects to return from the API lookup.",
			},
			"locations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"facility": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"parent_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxLocationsRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimLocationsListParams()

	if limitValue, ok := d.GetOk("limit"); ok {
		params.Limit = int64ToPtr(int64(limitValue.(int)))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "name":
				params.Name = &vString
			case "slug":
				params.Slug = &vString
			case "site":
				params.Site = &vString
			case "site_id":
				params.SiteID = &vString
			case "parent_id":
				params.ParentID = &vString
			case "tenant":
				params.Tenant = &vString
			case "tenant_id":
				params.TenantID = &vString
			case "status":
				params.Status = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}
	if tags, ok := d.GetOk("tags"); ok {
		tagSet := tags.(*schema.Set)
		for _, tag := range tagSet.List() {
			tagV := tag.(string)
			params.Tag = append(params.Tag, tagV)
		}
	}
	res, err := api.Dcim.DcimLocationsList(params, nil)

	if err != nil {
		return err
	}

	filteredLocations := res.GetPayload().Results

	var s []map[string]any
	for _, v := range filteredLocations {
		var mapping = make(map[string]any)

		mapping["id"] = strconv.FormatInt(v.ID, 10)
		mapping["name"] = v.Name
		mapping["slug"] = v.Slug
		mapping["site_id"] = v.Site.ID
		mapping["description"] = v.Description
		mapping["facility"] = v.Facility

		if v.Parent != nil {
			mapping["parent_id"] = v.Parent.ID
		}

		if v.Status != nil {
			mapping["status"] = v.Status.Value
		}

		if v.Tenant != nil {
			mapping["tenant_id"] = v.Tenant.ID
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("locations", s)
}
