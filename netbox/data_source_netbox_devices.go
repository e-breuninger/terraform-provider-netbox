// Copyright (c) 2022 Cisco Systems, Inc. and its affiliates
// All rights reserved.

package netbox

import (
	"encoding/json"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"net"
	"regexp"
	"strings"
)

func dataSourceNetboxDevices() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDevicesRead,
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):",
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
			"devices": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asset_tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
						"local_context_data": {
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
						"device_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"device_type_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"location_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"manufacturer_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"model": {
							Type:     schema.TypeString,
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
						"site_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tenant_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"role_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"serial": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rack_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rack_face": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rack_position": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"primary_ipv4": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_ipv6": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": tagsSchemaRead,
					},
				},
			},
		},
	}
}

func dataSourceNetboxDevicesRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := dcim.NewDcimDevicesListParams()

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			switch k {
			case "asset_tag":
				var assetTagString = v.(string)
				params.AssetTag = &assetTagString
			case "cluster_id":
				var clusterString = v.(string)
				params.ClusterID = &clusterString
			case "device_type_id":
				var deviceTypeIDString = v.(string)
				params.DeviceTypeID = &deviceTypeIDString
			case "name":
				var nameString = v.(string)
				params.Name = &nameString
			case "region":
				var regionString = v.(string)
				params.Region = &regionString
			case "role_id":
				var roleIDString = v.(string)
				params.RoleID = &roleIDString
			case "site_id":
				var siteIDString = v.(string)
				params.SiteID = &siteIDString
			case "location_id":
				var locationIDString = v.(string)
				params.LocationID = &locationIDString
			case "rack_id":
				var rackIDString = v.(string)
				params.RackID = &rackIDString
			case "tenant_id":
				var tenantIDString = v.(string)
				params.TenantID = &tenantIDString
			case "tags":
				var tagsString = v.(string)
				params.Tag = strings.Split(tagsString, ",")
			case "status":
				var statusString = v.(string)
				params.Status = &statusString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	if limit, ok := d.GetOk("limit"); ok {
		limitInt := int64(limit.(int))
		params.Limit = &limitInt
	}

	res, err := api.Dcim.DcimDevicesList(params, nil)
	if err != nil {
		return err
	}

	var filteredDevices []*models.DeviceWithConfigContext
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, device := range res.GetPayload().Results {
			if r.MatchString(*device.Name) {
				filteredDevices = append(filteredDevices, device)
			}
		}
	} else {
		filteredDevices = res.GetPayload().Results
	}

	var s []map[string]interface{}
	for _, device := range filteredDevices {
		var mapping = make(map[string]interface{})
		if device.AssetTag != nil {
			mapping["asset_tag"] = *device.AssetTag
		}
		if device.Cluster != nil {
			mapping["cluster_id"] = device.Cluster.ID
		}
		if device.Comments != "" {
			mapping["comments"] = device.Comments
		}
		if device.Description != "" {
			mapping["description"] = device.Description
		}
		if device.ConfigContext != nil {
			if configContext, err := json.Marshal(device.ConfigContext); err == nil {
				mapping["config_context"] = string(configContext)
			}
		}
		if device.LocalContextData != nil {
			if localContextData, err := json.Marshal(device.LocalContextData); err == nil {
				mapping["local_context_data"] = string(localContextData)
			}
		}
		mapping["device_id"] = device.ID
		if device.DeviceType != nil {
			mapping["device_type_id"] = device.DeviceType.ID
		}
		if device.DeviceType.Manufacturer != nil {
			mapping["manufacturer_id"] = device.DeviceType.Manufacturer.ID
		}
		if device.DeviceType.Model != nil {
			mapping["model"] = *device.DeviceType.Model
		}
		if device.Name != nil {
			mapping["name"] = *device.Name
		}
		if device.Location != nil {
			mapping["location_id"] = device.Location.ID
		}
		if device.Platform != nil {
			mapping["platform_id"] = device.Platform.ID
		}
		if device.Site != nil {
			mapping["site_id"] = device.Site.ID
		}
		if device.Tenant != nil {
			mapping["tenant_id"] = device.Tenant.ID
		}
		if device.Role != nil {
			mapping["role_id"] = device.Role.ID
		}
		if device.Serial != "" {
			mapping["serial"] = device.Serial
		}
		if device.Status != nil {
			mapping["status"] = *device.Status.Value
		}
		if device.CustomFields != nil {
			mapping["custom_fields"] = device.CustomFields
		}
		if device.Rack != nil {
			mapping["rack_id"] = device.Rack.ID
		}
		if device.Position != nil {
			mapping["rack_position"] = device.Position
		}
		if device.Face != nil {
			mapping["rack_face"] = device.Face.Value
		}
		if device.Tags != nil {
			mapping["tags"] = getTagListFromNestedTagList(device.Tags)
		}
		if device.PrimaryIp4 != nil {
			ip, _, err := net.ParseCIDR(*device.PrimaryIp4.Address)
			if err == nil {
				primaryIPv4 := ip.String()
				mapping["primary_ipv4"] = &primaryIPv4
			}
		}
		if device.PrimaryIp6 != nil {
			ip, _, err := net.ParseCIDR(*device.PrimaryIp6.Address)
			if err == nil {
				primaryIPv6 := ip.String()
				mapping["primary_ipv6"] = &primaryIPv6
			}
		}
		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("devices", s)
}
