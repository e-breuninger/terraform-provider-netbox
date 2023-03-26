package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxIpAddresses() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxIpAddressesRead,
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
			"ip_addresses": {
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
						"ip_address": {
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
						"dns_name": {
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
					},
				},
			},
		},
	}
}

func dataSourceNetboxIpAddressesRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := ipam.NewIpamIPAddressesListParams()

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
			case "dns_name":
				params.DNSName = &vString
			case "interface_id":
				params.InterfaceID = &vString
			case "device_id":
				params.DeviceID = &vString
			case "ip_address":
				params.Address = &vString
			case "vm_interface_id":
				params.VminterfaceID = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Ipam.IpamIPAddressesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	filteredIpAddresses := res.GetPayload().Results

	var s []map[string]interface{}
	for _, v := range filteredIpAddresses {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["description"] = v.Description
		mapping["created"] = v.Created.String()
		mapping["last_updated"] = v.LastUpdated.String()
		mapping["custom_fields"] = v.CustomFields

		mapping["ip_address"] = v.Address
		mapping["address_family"] = v.Family.Label
		mapping["status"] = v.Status.Value
		mapping["dns_name"] = v.DNSName
		mapping["tenant"] = flattenTenant(v.Tenant)

		if v.Role != nil {
			mapping["role"] = v.Role.Value
		}

		s = append(s, mapping)
	}

	d.SetId(resource.UniqueId())
	return d.Set("ip_addresses", s)

}

func flattenTenant(tenant *models.NestedTenant) []map[string]interface{} {
	var s []map[string]interface{}
	if tenant != nil {
		var mapping = make(map[string]interface{})
		mapping["id"] = tenant.ID
		mapping["name"] = tenant.Name
		mapping["slug"] = tenant.Slug
		s = append(s, mapping)
	}
	return s
}
