package netbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceNetboxVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxVirtualMachineRead,
		Schema: map[string]*schema.Schema{
			"search": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
				},
			},
			"vms": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"comments": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"config_context": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_fields": {
							Type:     schema.TypeMap,
							Computed: true,
						},
						"disk_size_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_context_data": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_mb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"primary_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_ip4": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_ip6": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"site_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"tenant_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vcpus": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vm_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxVirtualMachineRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := virtualization.NewVirtualizationVirtualMachinesListParams()

	if search, ok := d.GetOk("search"); ok {
		var searchParams = search.(map[string]interface{})
		for k, v := range searchParams {
			switch k {
			case "cluster_id":
				var clusterString = v.(string)
				params.ClusterID = &clusterString
			case "cluster_group":
				var clusterGroupString = v.(string)
				params.ClusterGroup = &clusterGroupString
			case "name":
				var nameString = v.(string)
				params.Name = &nameString
			case "region":
				var regionString = v.(string)
				params.Region = &regionString
			case "role":
				var roleString = v.(string)
				params.Role = &roleString
			case "site":
				var siteString = v.(string)
				params.Site = &siteString
			default:
				return fmt.Errorf("'%s' is not a supported search parameter", k)
			}
		}
	}

	res, err := api.Virtualization.VirtualizationVirtualMachinesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	var s []map[string]interface{}
	for _, v := range res.GetPayload().Results {
		var mapping = make(map[string]interface{})
		if v.Cluster != nil {
			mapping["cluster_id"] = v.Cluster.ID
		}
		if v.Comments != "" {
			mapping["comments"] = v.Comments
		}
		if v.ConfigContext != nil {
			if configContext, err := json.Marshal(v.ConfigContext); err == nil {
				mapping["config_context"] = string(configContext)
			}
		}
		if v.CustomFields != nil {
			mapping["custom_fields"] = v.CustomFields
		}
		if v.Disk != nil {
			mapping["disk_size_gb"] = *v.Disk
		}
		if v.LocalContextData != nil {
			if localContextData, err := json.Marshal(v.LocalContextData); err == nil {
				mapping["local_context_data"] = string(localContextData)
			}
		}
		if v.Memory != nil {
			mapping["memory_mb"] = *v.Memory
		}
		if v.Name != nil {
			mapping["name"] = *v.Name
		}
		if v.Platform != nil {
			mapping["platform_id"] = v.Platform.ID
		}
		if v.PrimaryIP != nil {
			mapping["primary_ip"] = v.PrimaryIP.Address
		}
		if v.PrimaryIp4 != nil {
			mapping["primary_ip4"] = v.PrimaryIp4.Address
		}
		if v.PrimaryIp6 != nil {
			mapping["primary_ip6"] = v.PrimaryIp6.Address
		}
		if v.Role != nil {
			mapping["role_id"] = v.Role.ID
		}
		if v.Site != nil {
			mapping["site_id"] = v.Site.ID
		}
		if v.Status != nil {
			mapping["status"] = v.Status.Value
		}
		if v.Tags != nil {
			var tags []int64
			for _, t := range v.Tags {
				tags = append(tags, t.ID)
			}
			mapping["tag_ids"] = tags
		}
		if v.Tenant != nil {
			mapping["tenant_id"] = v.Tenant.ID
		}
		if v.Vcpus != nil {
			mapping["vcpus"] = *v.Vcpus
		}

		mapping["vm_id"] = v.ID

		log.Printf("Map %#v", mapping)
		s = append(s, mapping)
	}

	d.SetId(resource.UniqueId())
	return d.Set("vms", s)
}
