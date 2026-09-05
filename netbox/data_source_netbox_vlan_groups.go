package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVlanGroups() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVlanGroupsRead,
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
			"vlan_groups": {
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
						"ranges": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"end": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"used": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tag_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxVlanGroupsRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamVlanGroupsListParams()

	// Get user limit (0 = fetch all)
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
			case "name":
				params.Name = []string{vString}
			case "name__empty":
				b, perr := boolFromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.NameEmpty = b
			case "name__ic":
				params.NameIc = []string{vString}
			case "name__ie":
				params.NameIe = []string{vString}
			case "name__iew":
				params.NameIew = []string{vString}
			case "name__isw":
				params.NameIsw = []string{vString}
			case "name__n":
				params.Namen = []string{vString}
			case "name__nic":
				params.NameNic = []string{vString}
			case "name__nie":
				params.NameNie = []string{vString}
			case "name__niew":
				params.NameNiew = []string{vString}
			case "name__nisw":
				params.NameNisw = []string{vString}
			case "slug":
				params.Slug = []string{vString}
			case "slug__empty":
				b, perr := boolFromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.SlugEmpty = b
			case "slug__ic":
				params.SlugIc = []string{vString}
			case "slug__ie":
				params.SlugIe = []string{vString}
			case "slug__iew":
				params.SlugIew = []string{vString}
			case "slug__isw":
				params.SlugIsw = []string{vString}
			case "slug__n":
				params.Slugn = []string{vString}
			case "slug__nic":
				params.SlugNic = []string{vString}
			case "slug__nie":
				params.SlugNie = []string{vString}
			case "slug__niew":
				params.SlugNiew = []string{vString}
			case "slug__nisw":
				params.SlugNisw = []string{vString}
			case "description":
				params.Description = []string{vString}
			case "description__empty":
				b, perr := boolFromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.DescriptionEmpty = b
			case "description__ic":
				params.DescriptionIc = []string{vString}
			case "description__ie":
				params.DescriptionIe = []string{vString}
			case "description__iew":
				params.DescriptionIew = []string{vString}
			case "description__isw":
				params.DescriptionIsw = []string{vString}
			case "description__n":
				params.Descriptionn = []string{vString}
			case "description__nic":
				params.DescriptionNic = []string{vString}
			case "description__nie":
				params.DescriptionNie = []string{vString}
			case "description__niew":
				params.DescriptionNiew = []string{vString}
			case "description__nisw":
				params.DescriptionNisw = []string{vString}
			case "id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ID = n
			case "id__gt":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.IDGt = n
			case "id__gte":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.IDGte = n
			case "id__lt":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.IDLt = n
			case "id__lte":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.IDLte = n
			case "id__n":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.IDn = n
			case "scope_type":
				params.ScopeType = []string{vString}
			case "scope_type__n":
				params.ScopeTypen = []string{vString}
			case "scope_id":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ScopeID = n
			case "scope_id__gt":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ScopeIDGt = n
			case "scope_id__gte":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ScopeIDGte = n
			case "scope_id__lt":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ScopeIDLt = n
			case "scope_id__lte":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ScopeIDLte = n
			case "scope_id__n":
				n, perr := int64FromFilterString(vString)
				if perr != nil {
					return perr
				}
				params.ScopeIDn = n
			case "site_id", "site":
				if n, err := strconv.ParseInt(vString, 10, 64); err == nil {
					params.Site = &n
				}
			case "location_id", "location":
				if n, err := strconv.ParseInt(vString, 10, 64); err == nil {
					params.Location = &n
				}
			case "rack_id", "rack":
				if n, err := strconv.ParseInt(vString, 10, 64); err == nil {
					params.Rack = &n
				}
			case "region_id", "region":
				if n, err := strconv.ParseInt(vString, 10, 64); err == nil {
					params.Region = &n
				}
			case "sitegroup_id", "sitegroup":
				if n, err := strconv.ParseInt(vString, 10, 64); err == nil {
					params.SiteGroup = &n
				}
			case "cluster_id", "cluster":
				if n, err := strconv.ParseInt(vString, 10, 64); err == nil {
					params.Cluster = &n
				}
			case "clustergroup_id", "clustergroup":
				if n, err := strconv.ParseInt(vString, 10, 64); err == nil {
					params.ClusterGroup = &n
				}
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination
	paginationHelper := NewPaginationHelper(userLimit)
	var allVlanGroups []*models.VLANGroup

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Ipam.IpamVlanGroupsList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch VLAN groups at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allVlanGroups = append(allVlanGroups, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allVlanGroups)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Trim to user limit if specified
	trimmedCount := paginationHelper.TrimToLimit(len(allVlanGroups))
	filteredVlanGroups := allVlanGroups[:trimmedCount]

	var s []map[string]interface{}
	for _, vg := range filteredVlanGroups {
		var mapping = make(map[string]interface{})

		mapping["id"] = vg.ID
		mapping["name"] = vg.Name
		mapping["slug"] = vg.Slug
		mapping["description"] = vg.Description

		mapping["ranges"] = []map[string]int64{}
		for _, v := range vg.VidRanges {
			mapping["ranges"] = append(mapping["ranges"].([]map[string]int64), map[string]int64{
				"start": v[0],
				"end":   v[1],
			})
		}
		mapping["used"] = vg.VlanCount

		if vg.Tags != nil {
			var tagIDs []int64
			for _, t := range vg.Tags {
				tagIDs = append(tagIDs, t.ID)
			}
			mapping["tag_ids"] = tagIDs
		}
		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("vlan_groups", s)
}
