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

func dataSourceNetboxDevicePowerPorts() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDevicePowerPortRead,
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
			"power_ports": {
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
						"device_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"module_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_draw": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allocated_draw": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxDevicePowerPortRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	params := dcim.NewDcimPowerPortsListParams()

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

	res, err := api.Dcim.DcimPowerPortsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	var filteredInterfaces []*models.PowerPort
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r, err := regexp.Compile(nameRegex.(string))
		if err != nil {
			return fmt.Errorf("failed to compile name regex: %w", err)
		}
		for _, port := range res.GetPayload().Results {
			if r.MatchString(*port.Name) {
				filteredInterfaces = append(filteredInterfaces, port)
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
		if v.Name != nil {
			mapping["name"] = *v.Name
		}
		if v.Tags != nil {
			var tags []int64
			for _, t := range v.Tags {
				tags = append(tags, t.ID)
			}
			mapping["tag_ids"] = tags
		}

		mapping["device_id"] = v.Device.ID
		if v.Module != nil {
			mapping["module_id"] = v.Module.ID
		}

		if v.Type != nil {
			mapping["type"] = v.Type.Value
		}

		if v.MaximumDraw != nil {
			mapping["maximum_draw"] = *v.MaximumDraw
		}
		if v.AllocatedDraw != nil {
			mapping["allocated_draw"] = *v.AllocatedDraw
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("power_ports", s)
}
