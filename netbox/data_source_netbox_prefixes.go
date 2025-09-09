package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxPrefixes() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxPrefixesRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of filters to apply to the API query when requesting prefixes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the field to filter on. Supported fields are: `prefix`, `contains`, `vlan_vid`, `vrf_id`, `vlan_id`, `status`, `tenant_id`, `site_id`, `description` & `tag`.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value to pass to the specified filter.",
						},
					},
				},
			},
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
				Description:      "The limit of objects to return from the API lookup.",
			},
			"prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"site_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"site_group_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"location_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"region_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_vid": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"vrf_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": tagsSchemaRead,
					},
				},
			},
		},
	}
}

func dataSourceNetboxPrefixesRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	params := ipam.NewIpamPrefixesListParams()

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
			case "prefix":
				params.Prefix = &vString
			case "vlan_vid":
				float, err := strconv.ParseFloat(vString, 64)
				if err != nil {
					return err
				}
				params.VlanVid = &float
			case "contains":
				params.Contains = &vString
			case "vrf_id":
				params.VrfID = &vString
			case "vlan_id":
				params.VlanID = &vString
			case "status":
				params.Status = &vString
			case "tenant_id":
				params.TenantID = &vString
			case "site_id":
				params.SiteID = &vString
			case "description":
				params.Description = &vString
			case "tag":
				params.Tag = []string{vString}
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Ipam.IpamPrefixesList(params, nil)
	if err != nil {
		return err
	}

	filteredPrefixes := res.GetPayload().Results

	var s []map[string]interface{}
	for _, v := range filteredPrefixes {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["prefix"] = v.Prefix
		mapping["description"] = v.Description
		if v.Vlan != nil {
			mapping["vlan_vid"] = v.Vlan.Vid
			mapping["vlan_id"] = v.Vlan.ID
		}
		if v.Vrf != nil {
			mapping["vrf_id"] = v.Vrf.ID
		}
		if v.Tenant != nil {
			mapping["tenant_id"] = v.Tenant.ID
		}
		if v.ScopeType != nil && v.ScopeID != nil {
			scopeID := v.ScopeID
			switch scopeType := v.ScopeType; *scopeType {
			case "dcim.site":
				mapping["site_id"] = scopeID
			case "dcim.sitegroup":
				mapping["site_group_id"] = scopeID
			case "dcim.location":
				mapping["location_id"] = scopeID
			case "dcim.region":
				mapping["region_id"] = scopeID
			}
		}
		mapping["status"] = v.Status.Value
		mapping["tags"] = getTagListFromNestedTagList(v.Tags)

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("prefixes", s)
}
