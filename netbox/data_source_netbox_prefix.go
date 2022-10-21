package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxPrefix() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxPrefixRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cidr": {
				Type:          schema.TypeString,
				Optional:      true,
				Deprecated:    "The `cidr` parameter is deprecated in favor of the canonical `prefix` attribute.",
				ConflictsWith: []string{"prefix"},
				ValidateFunc:  validation.IsCIDR,
				AtLeastOneOf:  []string{"prefix", "vlan_vid", "vrf_id", "vlan_id", "cidr"},
			},
			"prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validation.IsCIDR,
				ConflictsWith: []string{"cidr"},
				AtLeastOneOf:  []string{"prefix", "vlan_vid", "vrf_id", "vlan_id", "cidr"},
			},
			"vlan_vid": {
				Type:         schema.TypeFloat,
				Optional:     true,
				AtLeastOneOf: []string{"prefix", "vlan_vid", "vrf_id", "vlan_id", "cidr"},
				ValidateFunc: validation.FloatBetween(1, 4094),
			},
			"vrf_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"prefix", "vlan_vid", "vrf_id", "vlan_id", "cidr"},
			},
			"vlan_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"prefix", "vlan_vid", "vrf_id", "vlan_id", "cidr"},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxPrefixRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := ipam.NewIpamPrefixesListParams()

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	// note: cidr is deprecated in favor of prefix
	if cidr, ok := d.Get("cidr").(string); ok && cidr != "" {
		params.Prefix = &cidr
	}

	if prefix, ok := d.Get("prefix").(string); ok && prefix != "" {
		params.Prefix = &prefix
	}

	if vrfId, ok := d.Get("vrf_id").(int); ok && vrfId != 0 {
		// Note that vrf_id is a string pointer in the netbox filter, but we use a number in the provider
		params.VrfID = strToPtr(strconv.Itoa(vrfId))
	}

	if vlanId, ok := d.Get("vlan_id").(int); ok && vlanId != 0 {
		// Note that vlan_id is a string pointer in the netbox filter, but we use a number in the provider
		params.VlanID = strToPtr(strconv.Itoa(vlanId))
	}

	if vlanVid, ok := d.Get("vlan_vid").(float64); ok && vlanVid != 0 {
		params.VlanVid = &vlanVid
	}

	res, err := api.Ipam.IpamPrefixesList(params, nil)
	if err != nil {
		return err
	}

	if count := *res.GetPayload().Count; count != int64(1) {
		return fmt.Errorf("expected one prefix, but got %d", count)
	}

	result := res.GetPayload().Results[0]
	d.Set("id", result.ID)
	d.Set("cidr", result.Prefix)
	d.Set("prefix", result.Prefix)
	d.Set("status", result.Status.Value)

	if result.Vrf != nil {
		d.Set("vrf_id", result.Vrf.ID)
	}
	if result.Vlan != nil {
		d.Set("vlan_vid", result.Vlan.Vid)
		d.Set("vlan_id", result.Vlan.ID)
	}

	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}
