package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxPrefixes() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxPrefixesRead,
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
	api := m.(*client.NetBoxAPI)

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
			case "vrf_id":
				params.VrfID = &vString
			case "vlan_id":
				params.VlanID = &vString
			case "status":
				params.Status = &vString
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
		mapping["status"] = v.Status.Value
		mapping["tags"] = getTagListFromNestedTagList(v.Tags)

		s = append(s, mapping)
	}

	d.SetId(resource.UniqueId())
	return d.Set("prefixes", s)
}
