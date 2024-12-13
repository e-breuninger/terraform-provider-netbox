package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxTenants() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxTenantsRead,
		Description: `:meta:subcategory:Tenancy:`,
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
			"tenants": {
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
						"comments": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_fields": {
							Type:     schema.TypeMap,
							Computed: true,
						},

						"site_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rack_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"device_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vrf_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip_address_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vm_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"circuit_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cluster_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tenant_group": {
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
									"tenant_count": {
										Type:     schema.TypeInt,
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

func dataSourceNetboxTenantsRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := tenancy.NewTenancyTenantsListParams()

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
			case "name":
				params.Name = &vString
			case "slug":
				params.Slug = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Tenancy.TenancyTenantsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	filteredTenants := res.GetPayload().Results

	var s []map[string]interface{}
	for _, v := range filteredTenants {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["name"] = v.Name
		mapping["slug"] = v.Slug
		mapping["description"] = v.Description
		mapping["created"] = v.Created.String()
		mapping["last_updated"] = v.LastUpdated.String()
		mapping["comments"] = v.Comments
		mapping["custom_fields"] = v.CustomFields

		mapping["site_count"] = v.SiteCount
		mapping["rack_count"] = v.RackCount
		mapping["device_count"] = v.DeviceCount
		mapping["vrf_count"] = v.VrfCount
		mapping["prefix_count"] = v.PrefixCount
		mapping["ip_address_count"] = v.IpaddressCount
		mapping["vlan_count"] = v.VlanCount
		mapping["vm_count"] = v.VirtualmachineCount
		mapping["circuit_count"] = v.CircuitCount
		mapping["cluster_count"] = v.ClusterCount

		mapping["tenant_group"] = flattenTenantGroup(v.Group)
		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("tenants", s)
}

func flattenTenantGroup(group *models.NestedTenantGroup) []map[string]interface{} {
	var s []map[string]interface{}
	if group != nil {
		var mapping = make(map[string]interface{})
		mapping["id"] = group.ID
		mapping["name"] = group.Name
		mapping["slug"] = group.Slug
		mapping["tenant_count"] = group.TenantCount
		s = append(s, mapping)
	}
	return s
}
