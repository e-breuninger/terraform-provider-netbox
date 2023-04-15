package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVlanGroup() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVlanGroupRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				AtLeastOneOf: []string{"name", "slug"},
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"name", "slug"},
			},
			"scope_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"dcim.location", "dcim.site", "dcim.sitegroup", "dcim.region", "dcim.rack", "virtualization.cluster", "virtualization.clustergroup"}, false),
			},
			"scope_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"scope_type"},
			},
			"min_vid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_vid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlan_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxVlanGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := ipam.NewIpamVlanGroupsListParams()

	params.Limit = int64ToPtr(2)
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.Slug = &slug
	}
	if scopeType, ok := d.Get("scope_type").(string); ok && scopeType != "" {
		params.SetScopeType(&scopeType)
	}
	if scopeID, ok := d.Get("scope_id").(string); ok && scopeID != "" {
		params.SetScopeID(params.ScopeID)
	}

	res, err := api.Ipam.IpamVlanGroupsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("slug", result.Slug)
	d.Set("min_vid", result.MinVid)
	d.Set("max_vid", result.MaxVid)
	d.Set("vlan_count", result.VlanCount)
	d.Set("description", result.Description)
	return nil
}
