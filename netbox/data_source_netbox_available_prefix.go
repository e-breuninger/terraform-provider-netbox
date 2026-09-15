package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxAvailablePrefix() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxAvailablePrefixRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):This data source lists the free prefixes inside a parent prefix.

~> **Note:** Reading this list does not reserve anything. The prefixes are read while the plan is built, so they can be taken by another apply, or by somebody working in the Netbox UI, before your configuration gets around to using them. Use the ` + "`netbox_available_prefix`" + ` resource instead when you want a prefix allocated and held.`,
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
	api := m.(*providerState)

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
