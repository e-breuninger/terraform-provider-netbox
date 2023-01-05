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
				AtLeastOneOf:  []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
				Description:  "Description to include in the data source filter.",
			},
			"prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validation.IsCIDR,
				ConflictsWith: []string{"cidr"},
				AtLeastOneOf:  []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
			},
			"vlan_vid": {
				Type:         schema.TypeFloat,
				Optional:     true,
				AtLeastOneOf: []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
				ValidateFunc: validation.FloatBetween(1, 4094),
			},
			"vrf_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
			},
			"vlan_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
			},
			"site_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
			},
			"tag": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"description", "prefix", "vlan_vid", "vrf_id", "vlan_id", "site_id", "cidr", "tag"},
				Description:  "Tag to include in the data source filter (must match the tag's slug).",
			},
			"tag__n": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Tag to exclude from the data source filter (must match the tag's slug).
Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
for more information on available lookup expressions.`,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchemaRead,
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

	if description, ok := d.Get("description").(string); ok && description != "" {
		params.Description = &description
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

	if siteId, ok := d.Get("site_id").(int); ok && siteId != 0 {
		// Note that site_id is a string pointer in the netbox filter, but we use a number in the provider
		params.SiteID = strToPtr(strconv.Itoa(siteId))
	}

	if tag, ok := d.Get("tag").(string); ok && tag != "" {
		params.Tag = &tag
	}
	if tagn, ok := d.Get("tag__n").(string); ok && tagn != "" {
		params.Tagn = &tagn
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
	d.Set("description", result.Description)
	d.Set("tags", getTagListFromNestedTagList(result.Tags))

	if result.Vrf != nil {
		d.Set("vrf_id", result.Vrf.ID)
	}
	if result.Vlan != nil {
		d.Set("vlan_vid", result.Vlan.Vid)
		d.Set("vlan_id", result.Vlan.ID)
	}
	if result.Site != nil {
		d.Set("site_id", result.Site.ID)
	}
	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}
