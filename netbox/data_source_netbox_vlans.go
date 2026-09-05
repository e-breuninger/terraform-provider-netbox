package netbox

import (
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVlans() *schema.Resource {
	customFieldsFilterSchema := *customFieldsSchema
	return &schema.Resource{
		Read:        dataSourceNetboxVlansRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
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
			customFieldsKey: &customFieldsFilterSchema,
			"vlans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"role": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"site": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"tenant": {
							Type:     schema.TypeInt,
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
				},
			},
		},
	}
}

func dataSourceNetboxVlansRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamVlansListParams()
	var opts []ipam.ClientOption

	// Get user limit
	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		var tags []string
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "vid":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.Vid = n
			case "vid__gt":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.VidGt = n
			case "vid__gte":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.VidGte = n
			case "vid__lt":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.VidLt = n
			case "vid__lte":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.VidLte = n
			case "vid__n":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.Vidn = n
			case "group":
				params.Group = []string{vString}
			case "group__n":
				params.Groupn = []string{vString}
			case "group_id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.GroupID = n
			case "group_id__n":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.GroupIDn = n
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			case "tenant":
				params.Tenant = []string{vString}
			case "tenant__n":
				params.Tenantn = []string{vString}
			case "tenant_group":
				params.TenantGroup = []string{vString}
			case "tenant_group__n":
				params.TenantGroupn = []string{vString}
			case "tenant_group_id":
				params.TenantGroupID = []string{vString}
			case "tenant_group_id__n":
				params.TenantGroupIDn = []string{vString}
			case "tenant_id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.TenantID = n
			case "tenant_id__n":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.TenantIDn = n
			case "site_id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.SiteID = n
			case "site_id__n":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.SiteIDn = n
			case "status":
				params.Status = []string{vString}
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}
	if cfm, ok := d.Get(customFieldsKey).(map[string]interface{}); ok {
		opts = append(opts, WithCustomFieldParamsOption(cfm))
	}

	// Fetch all pages with pagination
	paginationHelper := NewPaginationHelper(userLimit)
	var allVlans []*models.VLAN

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Ipam.IpamVlansList(params, nil, opts...)
		if err != nil {
			return fmt.Errorf("failed to fetch VLANs at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allVlans = append(allVlans, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allVlans)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Trim to user limit if specified
	trimmedCount := paginationHelper.TrimToLimit(len(allVlans))
	filteredVlans := allVlans[:trimmedCount]

	var s []map[string]interface{}
	for _, v := range filteredVlans {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["vid"] = v.Vid
		mapping["name"] = v.Name
		mapping["description"] = v.Description
		if v.Group != nil {
			mapping["group_id"] = v.Group.ID
		}
		mapping["vid"] = v.Vid
		if v.Role != nil {
			mapping["role"] = v.Role.ID
		}
		if v.Site != nil {
			mapping["site"] = v.Site.ID
		}
		mapping["status"] = v.Status.Value
		if v.Tenant != nil {
			mapping["tenant"] = v.Tenant.ID
		}
		cf := flattenCustomFields(v.CustomFields)
		if cf != nil {
			mapping[customFieldsKey] = cf
		}
		if v.Tags != nil {
			var tagIDs []int64
			for _, t := range v.Tags {
				tagIDs = append(tagIDs, t.ID)
			}
			mapping["tag_ids"] = tagIDs
		}
		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("vlans", s)
}
