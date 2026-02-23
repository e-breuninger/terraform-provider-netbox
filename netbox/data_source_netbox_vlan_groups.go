package netbox

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
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

	if limitValue, ok := d.GetOk("limit"); ok {
		params.Limit = int64ToPtr(int64(limitValue.(int)))
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
				params.Name = strToPtr(vString)
			case "name__empty":
				params.NameEmpty = strToPtr(vString)
			case "name__ic":
				params.NameIc = strToPtr(vString)
			case "name__ie":
				params.NameIe = strToPtr(vString)
			case "name__iew":
				params.NameIew = strToPtr(vString)
			case "name__isw":
				params.NameIsw = strToPtr(vString)
			case "name__n":
				params.Namen = strToPtr(vString)
			case "name__nic":
				params.NameNic = strToPtr(vString)
			case "name__nie":
				params.NameNie = strToPtr(vString)
			case "name__niew":
				params.NameNiew = strToPtr(vString)
			case "name__nisw":
				params.NameNisw = strToPtr(vString)
			case "slug":
				params.Slug = strToPtr(vString)
			case "slug__empty":
				params.SlugEmpty = strToPtr(vString)
			case "slug__ic":
				params.SlugIc = strToPtr(vString)
			case "slug__ie":
				params.SlugIe = strToPtr(vString)
			case "slug__iew":
				params.SlugIew = strToPtr(vString)
			case "slug__isw":
				params.SlugIsw = strToPtr(vString)
			case "slug__n":
				params.Slugn = strToPtr(vString)
			case "slug__nic":
				params.SlugNic = strToPtr(vString)
			case "slug__nie":
				params.SlugNie = strToPtr(vString)
			case "slug__niew":
				params.SlugNiew = strToPtr(vString)
			case "slug__nisw":
				params.SlugNisw = strToPtr(vString)
			case "description":
				params.Description = strToPtr(vString)
			case "description__empty":
				params.DescriptionEmpty = strToPtr(vString)
			case "description__ic":
				params.DescriptionIc = strToPtr(vString)
			case "description__ie":
				params.DescriptionIe = strToPtr(vString)
			case "description__iew":
				params.DescriptionIew = strToPtr(vString)
			case "description__isw":
				params.DescriptionIsw = strToPtr(vString)
			case "description__n":
				params.Descriptionn = strToPtr(vString)
			case "description__nic":
				params.DescriptionNic = strToPtr(vString)
			case "description__nie":
				params.DescriptionNie = strToPtr(vString)
			case "description__niew":
				params.DescriptionNiew = strToPtr(vString)
			case "description__nisw":
				params.DescriptionNisw = strToPtr(vString)
			case "id":
				params.ID = strToPtr(vString)
			case "id__gt":
				params.IDGt = strToPtr(vString)
			case "id__gte":
				params.IDGte = strToPtr(vString)
			case "id__lt":
				params.IDLt = strToPtr(vString)
			case "id__lte":
				params.IDLte = strToPtr(vString)
			case "id__n":
				params.IDn = strToPtr(vString)
			case "scope_type":
				params.ScopeType = strToPtr(vString)
			case "scope_type__n":
				params.ScopeTypen = strToPtr(vString)
			case "scope_id":
				params.ScopeID = strToPtr(vString)
			case "scope_id__gt":
				params.ScopeIDGt = strToPtr(vString)
			case "scope_id__gte":
				params.ScopeIDGte = strToPtr(vString)
			case "scope_id__lt":
				params.ScopeIDLt = strToPtr(vString)
			case "scope_id__lte":
				params.ScopeIDLte = strToPtr(vString)
			case "scope_id__n":
				params.ScopeIDn = strToPtr(vString)
			case "site_id", "site":
				if f, err := strconv.ParseFloat(vString, 64); err == nil {
					params.Site = float64ToPtr(f)
				}
			case "location_id", "location":
				if f, err := strconv.ParseFloat(vString, 64); err == nil {
					params.Location = float64ToPtr(f)
				}
			case "rack_id", "rack":
				if f, err := strconv.ParseFloat(vString, 64); err == nil {
					params.Rack = float64ToPtr(f)
				}
			case "region_id", "region":
				if f, err := strconv.ParseFloat(vString, 64); err == nil {
					params.Region = float64ToPtr(f)
				}
			case "sitegroup_id", "sitegroup":
				if f, err := strconv.ParseFloat(vString, 64); err == nil {
					params.Sitegroup = float64ToPtr(f)
				}
			case "cluster_id", "cluster":
				if f, err := strconv.ParseFloat(vString, 64); err == nil {
					params.Cluster = float64ToPtr(f)
				}
			case "clustergroup_id", "clustergroup":
				if f, err := strconv.ParseFloat(vString, 64); err == nil {
					params.Clustergroup = float64ToPtr(f)
				}
			case "minvid":
				params.MinVid = strToPtr(vString)
			case "minvid__gt":
				params.MinVidGt = strToPtr(vString)
			case "minvid__gte":
				params.MinVidGte = strToPtr(vString)
			case "minvid__lt":
				params.MinVidLt = strToPtr(vString)
			case "minvid__lte":
				params.MinVidLte = strToPtr(vString)
			case "minvid__n":
				params.MinVidn = strToPtr(vString)
			case "maxvid":
				params.MaxVid = strToPtr(vString)
			case "maxvid__gt":
				params.MaxVidGt = strToPtr(vString)
			case "maxvid__gte":
				params.MaxVidGte = strToPtr(vString)
			case "maxvid__lt":
				params.MaxVidLt = strToPtr(vString)
			case "maxvid__lte":
				params.MaxVidLte = strToPtr(vString)
			case "maxvid__n":
				params.MaxVidn = strToPtr(vString)
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Ipam.IpamVlanGroupsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	filteredVlanGroups := res.GetPayload().Results

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
