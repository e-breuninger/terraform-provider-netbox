package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxAvailablePrefix() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxAvailablePrefixRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"prefix_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"prefixes_available": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"family": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrf_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxAvailablePrefixRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	params := ipam.NewIpamPrefixesAvailablePrefixesListParams()

	if prefixID, ok := d.Get("prefix_id").(int); ok && prefixID != 0 {
		params.ID = int64(prefixID)
	}

	res, err := api.Ipam.IpamPrefixesAvailablePrefixesList(params, nil)
	if err != nil {
		return err
	}

	result := res.GetPayload()

	var s []map[string]interface{}
	for _, v := range result {
		var mapping = make(map[string]interface{})

		mapping["prefix"] = v.Prefix
		mapping["family"] = v.Family
		if v.Vrf != nil {
			mapping["vrf_id"] = v.Vrf.ID
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("prefixes_available", s)
}
