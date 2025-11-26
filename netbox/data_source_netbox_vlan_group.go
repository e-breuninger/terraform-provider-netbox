package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVlanGroup() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVlanGroupRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				AtLeastOneOf: []string{"id", "name", "slug", "scope_type"},
			},
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				AtLeastOneOf: []string{"id", "name", "slug", "scope_type"},
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"id", "name", "slug", "scope_type"},
			},
			"scope_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxVlanGroupScopeTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxVlanGroupScopeTypeOptions),
				AtLeastOneOf: []string{"id", "name", "slug", "scope_type"},
			},
			"scope_id": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"scope_type"},
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
	api := m.(*providerState)
	params := ipam.NewIpamVlanGroupsListParams()

	params.Limit = int64ToPtr(2)

	if id, ok := d.Get("id").(string); ok && id != "" {
		params.ID = strToPtr(id)
	}
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = strToPtr(name)
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.Slug = strToPtr(slug)
	}
	if scopeType, ok := d.Get("scope_type").(string); ok && scopeType != "" {
		params.ScopeType = strToPtr(scopeType)
	}
	if scopeID, ok := d.Get("scope_id").(string); ok && scopeID != "" {
		params.ScopeID = strToPtr(scopeID)
	}

	res, err := api.Ipam.IpamVlanGroupsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one vlan group returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no vlan group found matching filter")
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("id", strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("slug", result.Slug)
	d.Set("vlan_count", result.VlanCount)
	d.Set("description", result.Description)
	d.Set("scope_id", strconv.FormatInt(*result.ScopeID, 10))
	d.Set("scope_type", result.ScopeType)
	return nil
}
