package netbox

import (
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxCables() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxCablesRead,
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
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"cables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"tenant_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"label": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"color": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"length": {
							Type:     schema.TypeFloat,
							Computed: true,
							Optional: true,
						},
						"length_unit": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"a_termination": {
							Type:     schema.TypeSet,
							Required: true,
							Elem:     genericObjectSchema,
						},
						"b_termination": {
							Type:     schema.TypeSet,
							Required: true,
							Elem:     genericObjectSchema,
						},
						"description": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(0, 200),
						},
						"comments": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			// tagsKey:         tagsSchema,
			// customFieldsKey: customFieldsSchema,
		},
	}
}

func dataSourceNetboxCablesRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimCablesListParams()

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
			case "termination_a_id":
				params.TerminationaID = &vString
			case "termination_b_id":
				params.TerminationbID = &vString
			case "termination_a_type":
				params.TerminationaType = &vString
			case "termination_b_type":
				params.TerminationbType = &vString
			case "length_unit":
				params.LengthUnit = &vString
			case "length":
				params.Length = &vString
			case "color":
				params.Color = &vString
			case "label":
				params.Label = &vString
			case "status":
				params.Status = &vString
			case "type":
				params.Type = &vString
			case "id":
				params.ID = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Dcim.DcimCablesList(params, nil)
	if err != nil {
		return err
	}

	var s []map[string]interface{}
	for _, cable := range res.GetPayload().Results {
		var mapping = make(map[string]interface{})
		mapping["id"] = cable.ID
		mapping["a_termination"] = getSchemaSetFromGenericObjects(cable.ATerminations)
		mapping["b_termination"] = getSchemaSetFromGenericObjects(cable.BTerminations)

		if cable.Status != nil {
			mapping["status"] = cable.Status.Value
		} else {
			mapping["status"] = nil
		}

		mapping["type"] = cable.Type

		if cable.Tenant != nil {
			mapping["tenant_id"] = cable.Tenant.ID
		} else {
			mapping["tenant_id"] = nil
		}

		mapping["label"] = cable.Label
		if cable.Color != "" {
			mapping["color_hex"] = cable.Color
		}
		if cable.Length != nil {
			mapping["length"] = cable.Length
		}

		if cable.LengthUnit != nil {
			mapping["length_unit"] = cable.LengthUnit.Value
		} else {
			mapping["length_unit"] = nil
		}

		if cable.Tenant != nil {
			mapping["tenant_id"] = cable.Tenant.ID
		}
		mapping["description"] = cable.Description
		mapping["comments"] = cable.Comments

		// cf := getCustomFields(res.GetPayload().CustomFields)
		// if cf != nil {
		// 	mapping[customFieldsKey] = cf
		// }

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("cables", s)
}
