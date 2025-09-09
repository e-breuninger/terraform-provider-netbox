package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVrfs() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVrfsRead,
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
				Default:          0,
			},
			"vrfs": {
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rd": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxVrfsRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	params := ipam.NewIpamVrfsListParams()

	if limitValue, ok := d.GetOk("limit"); ok {
		params.Limit = int64ToPtr(int64(limitValue.(int)))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		var tags []string
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "id":
				params.ID = &vString
			case "name":
				params.Name = &vString
			case "description":
				params.Description = &vString
			case "rd":
				params.Rd = &vString
			case "tenant":
				params.Tenant = &vString
			case "tenant__n":
				params.Tenantn = &vString
			case "tenant_group":
				params.TenantGroup = &vString
			case "tenant_group__n":
				params.TenantGroupn = &vString
			case "tenant_group_id":
				params.TenantGroupID = &vString
			case "tenant_group_id__n":
				params.TenantGroupIDn = &vString
			case "tenant_id":
				params.TenantID = &vString
			case "tenant_id__n":
				params.TenantIDn = &vString
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	res, err := api.Ipam.IpamVrfsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}

	filteredVrfs := res.GetPayload().Results

	var s []map[string]interface{}
	for _, v := range filteredVrfs {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["name"] = v.Name
		mapping["description"] = v.Description
		if v.Rd != nil {
			mapping["rd"] = v.Rd
		}
		if v.Tenant != nil {
			mapping["tenant"] = v.Tenant.ID
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("vrfs", s)
}
