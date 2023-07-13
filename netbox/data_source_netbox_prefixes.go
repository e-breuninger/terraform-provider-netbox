package netbox

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/iancoleman/strcase"
)

func dataSourceNetboxPrefixes() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxPrefixesRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of filters to apply to the API query when requesting prefixes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the field to filter on. Supported fields are: `prefix`, `vlan_vid`, `vrf_id`, `vlan_id`, `status`, `site_id`, & `tag`.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value to pass to the specified filter.",
						},
					},
				},
			},
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
				Description:      "The limit of objects to return from the API lookup.",
			},
			"prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan_vid": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"vrf_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
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

func dataSourceNetboxPrefixesRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := ipam.NewIpamPrefixesListParams()

	if limitValue, ok := d.GetOk("limit"); ok {
		params.Limit = int64ToPtr(int64(limitValue.(int)))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			paramName := strcase.ToCamel(strings.Replace(k.(string), "__n", "n", -1))
			paramName = strings.Replace(paramName, "Id", "ID", -1)

			params_reflect := reflect.ValueOf(params).Elem()
			field := params_reflect.FieldByName(paramName)

			if !(field.IsValid()) {
				return fmt.Errorf("'%s' is not a supported filter parameter.  Netbox go SDK does not have the associated parameter [(%s)]", k, paramName)
			}

			if field.Kind() == reflect.Slice {
				//Param is an array/slice
				field.Set(reflect.Append(field, reflect.ValueOf(vString)))
			} else if (reflect.PtrTo(field.Type().Elem()).Elem().Kind()) == reflect.Float64 {
				// ^ This CANT be the best way to do this, but it works
				vFloat, err := strconv.ParseFloat(vString, 64)
				if field.Set(reflect.ValueOf(&vFloat)); err != nil {
					return fmt.Errorf("Failed to set parameter [(%s)] with error [(%s)]", paramName, err)
				}
			} else {
				//Param is a scalar
				field.Set(reflect.ValueOf(&vString))
			}
		}
	}

	res, err := api.Ipam.IpamPrefixesList(params, nil)
	if err != nil {
		return err
	}

	filteredPrefixes := res.GetPayload().Results

	var s []map[string]interface{}
	for _, v := range filteredPrefixes {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["prefix"] = v.Prefix
		mapping["description"] = v.Description
		if v.Vlan != nil {
			mapping["vlan_vid"] = v.Vlan.Vid
			mapping["vlan_id"] = v.Vlan.ID
		}
		if v.Vrf != nil {
			mapping["vrf_id"] = v.Vrf.ID
		}
		mapping["status"] = v.Status.Value
		mapping["tags"] = getTagListFromNestedTagList(v.Tags)

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("prefixes", s)
}
