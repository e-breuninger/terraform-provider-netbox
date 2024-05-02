package netbox

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxConfigContext() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxConfigContextRead,
		Description: `:meta:subcategory:Extras:`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"cluster_types": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"device_types": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"locations": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"platforms": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"regions": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"site_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"sites": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tenant_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tenants": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceNetboxConfigContextRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	params := extras.NewExtrasConfigContextsListParams()
	params.Name = &name
	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Extras.ExtrasConfigContextsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("More than one result. Specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("No result")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("weight", result.Weight)
	if result.Data != nil {
		if jsonArr, err := json.Marshal(result.Data); err == nil {
			d.Set("data", string(jsonArr))
		}
	} else {
		d.Set("data", nil)
	}
	cluster_groups := make([]int64, len(result.ClusterGroups))
	for i, v := range result.ClusterGroups {
		cluster_groups[i] = int64(v.ID)
	}
	d.Set("cluster_groups", cluster_groups)
	cluster_types := make([]int64, len(result.ClusterTypes))
	for i, v := range result.ClusterTypes {
		cluster_types[i] = int64(v.ID)
	}
	d.Set("cluster_types", cluster_types)
	clusters := make([]int64, len(result.Clusters))
	for i, v := range result.Clusters {
		clusters[i] = int64(v.ID)
	}
	d.Set("clusters", clusters)
	device_types := make([]int64, len(result.DeviceTypes))
	for i, v := range result.DeviceTypes {
		device_types[i] = int64(v.ID)
	}
	d.Set("device_types", device_types)
	locations := make([]int64, len(result.Locations))
	for i, v := range result.Locations {
		locations[i] = int64(v.ID)
	}
	d.Set("locations", locations)
	platforms := make([]int64, len(result.Platforms))
	for i, v := range result.Platforms {
		platforms[i] = int64(v.ID)
	}
	d.Set("platforms", platforms)
	regions := make([]int64, len(result.Regions))
	for i, v := range result.Regions {
		regions[i] = int64(v.ID)
	}
	d.Set("regions", regions)
	roles := make([]int64, len(result.Roles))
	for i, v := range result.Roles {
		roles[i] = int64(v.ID)
	}
	d.Set("roles", roles)
	site_groups := make([]int64, len(result.SiteGroups))
	for i, v := range result.SiteGroups {
		site_groups[i] = int64(v.ID)
	}
	d.Set("site_groups", site_groups)
	sites := make([]int64, len(result.Sites))
	for i, v := range result.Sites {
		sites[i] = int64(v.ID)
	}
	d.Set("sites", sites)
	tenant_groups := make([]int64, len(result.TenantGroups))
	for i, v := range result.TenantGroups {
		tenant_groups[i] = int64(v.ID)
	}
	d.Set("tenant_groups", tenant_groups)
	tenants := make([]int64, len(result.Tenants))
	for i, v := range result.Tenants {
		tenants[i] = int64(v.ID)
	}
	d.Set("tenants", tenants)
	tags := make([]string, len(result.Tags))
	for i, v := range result.Tags {
		tags[i] = string(v)
	}
	d.Set("tags", tags)
	return nil
}
