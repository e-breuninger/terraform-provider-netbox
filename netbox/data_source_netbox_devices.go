// Copyright (c) 2022 Cisco Systems, Inc. and its affiliates
// All rights reserved.

package netbox

import (
	"fmt"
	"regexp"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
						"tag_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"tag_slugs": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
			vString := v.(string)
			switch k {
			case "asset_tag":
				params.AssetTag = &vString
			case "cluster_id":
				params.ClusterID = &vString
			case "name":
				params.Name = &vString
			case "region":
				params.Region = &vString
			case "role_id":
				params.RoleID = &vString
			case "site_id":
				params.SiteID = &vString
			case "tenant_id":
				params.TenantID = &vString
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
		if device.Tags != nil {
			var tagsIds []int64
			var tagsSlugs []string
			for _, t := range device.Tags {
				tagsIds = append(tagsIds, t.ID)
			}
			for _, t := range device.Tags {
				tagsSlugs = append(tagsSlugs, *t.Slug)
			}
			mapping["tag_ids"] = tagsIds
			mapping["tag_slugs"] = tagsSlugs
		}
		if device.Tenant != nil {
			mapping["tenant_id"] = device.Tenant.ID
		}
		if device.DeviceRole != nil {
			mapping["role_id"] = device.DeviceRole.ID
		}
		if device.Serial != "" {
			mapping["serial"] = device.Serial
		}
		if device.Status != nil {
			mapping["status"] = *device.Status.Value
		}
		s = append(s, mapping)
	}

	d.SetId(resource.UniqueId())
	return d.Set("devices", s)
}
