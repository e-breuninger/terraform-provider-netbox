package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxRacks() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxRacksRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
			},
			"racks": {
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
						"site_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"width": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"u_height": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						tagsKey: tagsSchemaRead,
						"tenant_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"facility_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"role_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"serial": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asset_tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"weight": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"max_weight": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"weight_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"desc_units": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"outer_width": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"outer_depth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"outer_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mounting_depth": {
							Type:     schema.TypeInt,
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
						"custom_fields": {
							Type:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxRacksRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := dcim.NewDcimRacksListParams()

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
			case "asset_tag":
				params.AssetTag = &vString
			case "contact":
				params.Contact = &vString
			case "contact_group":
				params.ContactGroup = &vString
			case "contact_role":
				params.ContactRole = &vString
			case "desc_units":
				params.DescUnits = &vString
			case "facility_id":
				params.FacilityID = &vString
			case "id":
				params.ID = &vString
			case "location_id":
				params.LocationID = &vString
			case "max_weight":
				params.MaxWeight = &vString
			case "mounting_depth":
				params.MountingDepth = &vString
			case "name":
				params.Name = &vString
			case "outer_depth":
				params.OuterDepth = &vString
			case "outer_unit":
				params.OuterUnit = &vString
			case "outer_width":
				params.OuterWidth = &vString
			case "region_id":
				params.RegionID = &vString
			case "role_id":
				params.RoleID = &vString
			case "serial":
				params.Serial = &vString
			case "site_id":
				params.SiteID = &vString
			case "status":
				params.Status = &vString
			case "tenant_id":
				params.TenantID = &vString
			case "type":
				params.Type = &vString
			case "u_height":
				params.UHeight = &vString
			case "weight":
				params.Weight = &vString
			case "weight_unit":
				params.WeightUnit = &vString
			case "width":
				params.Width = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Dcim.DcimRacksList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	filteredRacks := res.GetPayload().Results

	var s []map[string]interface{}
	for _, v := range filteredRacks {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["name"] = v.Name
		if v.Site != nil {
			mapping["site_id"] = v.Site.ID
		}
		if v.Status != nil {
			mapping["status"] = v.Status.Value
		}
		if v.Width != nil {
			mapping["width"] = v.Width.Value
		}
		mapping["u_height"] = v.UHeight
		mapping["tags"] = getTagListFromNestedTagList(v.Tags)
		if v.Tenant != nil {
			mapping["tenant_id"] = v.Tenant.ID
		}
		mapping["facility_id"] = v.FacilityID
		if v.Location != nil {
			mapping["location_id"] = v.Location.ID
		}
		if v.Role != nil {
			mapping["role_id"] = v.Role.ID
		}
		mapping["serial"] = v.Serial
		mapping["asset_tag"] = v.AssetTag
		if v.Type != nil {
			mapping["type"] = v.Type.Value
		}
		mapping["weight"] = v.Weight
		mapping["max_weight"] = v.MaxWeight
		mapping["desc_units"] = v.DescUnits
		mapping["outer_width"] = v.OuterWidth
		mapping["outer_depth"] = v.OuterDepth
		if v.OuterUnit != nil {
			mapping["outer_unit"] = v.OuterUnit.Value
		}
		mapping["mounting_depth"] = v.MountingDepth
		mapping["description"] = v.Description
		mapping["comments"] = v.Comments
		mapping["custom_fields"] = getCustomFields(v.CustomFields)

		s = append(s, mapping)
	}

	d.SetId(resource.UniqueId())
	return d.Set("racks", s)
}
