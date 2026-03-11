package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVlans() *schema.Resource {
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
			"vlans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
		},
	}
}

func dataSourceNetboxVlansRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamVlansListParams()

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
				params.Vid = &vString
			case "vid__gt":
				params.VidGt = &vString
			case "vid__gte":
				params.VidGte = &vString
			case "vid__lt":
				params.VidLt = &vString
			case "vid__lte":
				params.VidLte = &vString
			case "vid__n":
				params.Vidn = &vString
			case "group":
				params.Group = &vString
			case "group__n":
				params.Groupn = &vString
			case "group_id":
				params.GroupID = &vString
			case "group_id__n":
				params.GroupIDn = &vString
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			case "tenant":
				params.Tenant = &vString
			case "tenant__n":
				params.Tenantn = &vString
			case "tenant_group":
				params.TenantGroup = &vString
			case "tenant_group__n":
				params.TenantGroupn = &vString
			case "tenant_group_id":
				params.TenantGroupID = &vString
			case "tenant_group_id__n":
				params.TenantGroupIDn = &vString
			case "tenant_id":
				params.TenantID = &vString
			case "tenant_id__n":
				params.TenantIDn = &vString
			case "site_id":
				params.SiteID = &vString
			case "site_id__n":
				params.SiteIDn = &vString
			case "status":
				params.Status = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination
	paginationHelper := NewPaginationHelper(userLimit)
	var allVlans []*models.VLAN

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Ipam.IpamVlansList(params, nil)
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

	if len(filteredVlans) == 0 {
		return errors.New("no result")
	}

	var s []map[string]interface{}
	for _, v := range filteredVlans {
		var mapping = make(map[string]interface{})

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
