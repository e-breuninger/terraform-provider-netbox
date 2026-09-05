package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
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
						"type_id": {
							Type:     schema.TypeInt,
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
	api := m.(*providerState)

	params := dcim.NewDcimRacksListParams()

	// Get user limit
	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "asset_tag":
				params.AssetTag = []string{vString}
			case "contact":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.Contact = n
			case "contact_group":
				params.ContactGroup = []string{vString}
			case "contact_role":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ContactRole = n
			case "desc_units":
				b, perr := boolFromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.DescUnits = b
			case "facility_id":
				params.FacilityID = []string{vString}
			case "id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ID = n
			case "location_id":
				params.LocationID = []string{vString}
			case "max_weight":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.MaxWeight = n
			case "mounting_depth":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.MountingDepth = n
			case "name":
				params.Name = []string{vString}
			case "outer_depth":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.OuterDepth = n
			case "outer_unit":
				params.OuterUnit = &vString
			case "outer_width":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.OuterWidth = n
			case "region_id":
				params.RegionID = []string{vString}
			case "role_id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.RoleID = n
			case "serial":
				params.Serial = []string{vString}
			case "site_id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.SiteID = n
			case "status":
				params.Status = []string{vString}
			case "tenant_id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.TenantID = n
			case "u_height":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.UHeight = n
			case "weight":
				f, perr := strconv.ParseFloat(vString, 64)
				if perr != nil {
					return perr
				}
				params.Weight = []float64{f}
			case "weight_unit":
				params.WeightUnit = &vString
			case "width":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.Width = n
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination
	paginationHelper := NewPaginationHelper(userLimit)
	var allRacks []*models.Rack

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Dcim.DcimRacksList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch racks at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allRacks = append(allRacks, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allRacks)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Trim to user limit if specified
	trimmedCount := paginationHelper.TrimToLimit(len(allRacks))
	filteredRacks := allRacks[:trimmedCount]

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
		if v.RackType != nil {
			mapping["type_id"] = v.RackType.ID
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

	d.SetId(id.UniqueId())
	return d.Set("racks", s)
}
