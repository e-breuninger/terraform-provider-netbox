package netbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVirtualMachineRead,
		Description: `:meta:subcategory:Virtualization:`,
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
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
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
						"description": {
							Type:     schema.TypeString,
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
							Type:     schema.TypeFloat,
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

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
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
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	if limit, ok := d.GetOk("limit"); ok {
		limitInt := int64(limit.(int))
		params.Limit = &limitInt
	}

	res, err := api.Virtualization.VirtualizationVirtualMachinesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	var filteredVms []*models.VirtualMachineWithConfigContext
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, vm := range res.GetPayload().Results {
			if r.MatchString(*vm.Name) {
				filteredVms = append(filteredVms, vm)
			}
		}
	} else {
		filteredVms = res.GetPayload().Results
	}

	var s []map[string]interface{}
	for _, v := range filteredVms {
		var mapping = make(map[string]interface{})
		if v.Cluster != nil {
			mapping["cluster_id"] = v.Cluster.ID
		}
		if v.Comments != "" {
			mapping["comments"] = v.Comments
		}
		if v.Description != "" {
			mapping["description"] = v.Description
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

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("vms", s)
}
