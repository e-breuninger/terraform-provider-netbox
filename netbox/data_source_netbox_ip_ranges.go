package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxIPRanges() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxIPRangesRead,
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
				Default:          1000,
			},
			"ip_ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_fields": {
							Type:     schema.TypeMap,
							Computed: true,
						},
						"start_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"tenant": {
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
								},
							},
						},
						"tags": {
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
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"slug": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxIPRangesRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := ipam.NewIpamIPRangesListParams()

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
			case "contains":
				params.Contains = &vString
			case "start_address":
				params.StartAddress = &vString
			case "end_address":
				params.EndAddress = &vString
			case "role":
				params.Role = &vString
			case "status":
				params.Status = &vString
			case "vrf":
				params.Vrf = &vString
			case "tenant":
				params.Tenant = &vString
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Ipam.IpamIPRangesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	filteredIPRanges := res.GetPayload().Results

	var s []map[string]interface{}
	for _, v := range filteredIPRanges {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["description"] = v.Description
		mapping["created"] = v.Created.String()
		mapping["last_updated"] = v.LastUpdated.String()
		mapping["custom_fields"] = v.CustomFields

		mapping["start_address"] = v.StartAddress
		mapping["end_address"] = v.EndAddress
		mapping["address_family"] = v.Family.Label
		mapping["status"] = v.Status.Value
		mapping["tenant"] = flattenTenant(v.Tenant)
		if v.Vrf != nil {
			mapping["vrf_id"] = v.Vrf.ID
		}
		if v.Role != nil {
			mapping["role_id"] = v.Role.ID
		}
		mapping["description"] = v.Description
		var stags []map[string]interface{}
		for _, t := range v.Tags {
			var tagmapping = make(map[string]interface{})
			tagmapping["name"] = t.Name
			tagmapping["display"] = t.Display
			tagmapping["slug"] = t.Slug
			tagmapping["id"] = t.ID
			stags = append(stags, tagmapping)
		}
		mapping["tags"] = stags

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("ip_ranges", s)
}
