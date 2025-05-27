package netbox

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxDeviceInterfaces() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDeviceInterfaceRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
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
			"interfaces": {
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
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
						},
						"mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
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
						"tagged_vlans": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						// Do as a TypeList due to limitation of TypeMap
						"untagged_vlan": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"device_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxDeviceInterfaceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := dcim.NewDcimInterfacesListParams()

	if limit, ok := d.GetOk("limit"); ok {
		limitInt := int64(limit.(int))
		params.Limit = &limitInt
	}


	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "mac_address":
				params.MacAddress = &vString
			case "name":
				params.Name = &vString
			case "tag":
				params.Tag = []string{vString} //TODO: switch schema to list?
			case "device_id":
				params.DeviceID = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Dcim.DcimInterfacesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	var filteredInterfaces []*models.Interface
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, dcimInterface := range res.GetPayload().Results {
			if r.MatchString(*dcimInterface.Name) {
				filteredInterfaces = append(filteredInterfaces, dcimInterface)
			}
		}
	} else {
		filteredInterfaces = res.GetPayload().Results
	}

	var s []map[string]interface{}
	for _, v := range filteredInterfaces {
		var mapping = make(map[string]interface{})
		mapping["id"] = v.ID
		if v.Description != "" {
			mapping["description"] = v.Description
		}
		mapping["enabled"] = v.Enabled
		if v.MacAddress != nil {
			mapping["mac_address"] = *v.MacAddress
		}
		if v.Mode != nil {
			mapping["mode"] = map[string]string{
				"label": *v.Mode.Label,
				"value": *v.Mode.Value,
			}
		}
		if v.Mtu != nil {
			mapping["mtu"] = *v.Mtu
		}
		if v.Name != nil {
			mapping["name"] = *v.Name
		}
		if v.TaggedVlans != nil {
			mapping["tagged_vlans"] = flattenVlanAttributes(v.TaggedVlans)
		}
		if v.Tags != nil {
			var tags []int64
			for _, t := range v.Tags {
				tags = append(tags, t.ID)
			}
			mapping["tag_ids"] = tags
		}
		if v.UntaggedVlan != nil {
			vlanSlice := []*models.NestedVLAN{v.UntaggedVlan}
			mapping["untagged_vlan"] = flattenVlanAttributes(vlanSlice)
		}

		mapping["device_id"] = v.Device.ID

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("interfaces", s)
}

/*

TODO: maybe separate this function from both VM Interfaces and DCIM Interfaces?

func flattenVlanAttributes(vlans []*models.NestedVLAN) []map[string]interface{} {
	var mappedVlans []map[string]interface{}
	for _, vlan := range vlans {
		v := *vlan
		mappedVlan := map[string]interface{}{
			"id":   v.ID,
			"vid":  *v.Vid,
			"name": *v.Name,
		}
		mappedVlans = append(mappedVlans, mappedVlan)
	}
	return mappedVlans
}
*/
