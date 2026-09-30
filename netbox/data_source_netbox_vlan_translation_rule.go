package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// The VLANTranslationRule model has no unique "name" field, unlike most other
// NetBox objects. Instead, a rule is uniquely identified either by its `id`,
// or by the combination of the policy it belongs to (`policy_id`) and its
// `local_vid`.
func dataSourceNetboxVlanTranslationRule() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVlanTranslationRuleRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				AtLeastOneOf: []string{"id", "policy_id"},
				Computed:     true,
				Optional:     true,
				Type:         schema.TypeString,
			},
			"policy_id": {
				AtLeastOneOf: []string{"id", "policy_id"},
				Optional:     true,
				Type:         schema.TypeInt,
			},
			"local_vid": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Used together with `policy_id` to narrow down the result when not looking up by `id`.",
			},
			"remote_vid": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func dataSourceNetboxVlanTranslationRuleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamVlanTranslationRulesListParams()

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	} else {
		if policyID, ok := d.GetOk("policy_id"); ok {
			policyIDStr := strconv.Itoa(policyID.(int))
			params.PolicyID = &policyIDStr
		}
		if localVid, ok := d.GetOk("local_vid"); ok {
			localVidFloat := float64(localVid.(int))
			params.LocalVid = &localVidFloat
		}
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Ipam.IpamVlanTranslationRulesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no VLAN translation rule found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))

	if result.Policy != nil {
		d.Set("policy_id", *result.Policy)
	}
	if result.LocalVid != nil {
		d.Set("local_vid", *result.LocalVid)
	}
	if result.RemoteVid != nil {
		d.Set("remote_vid", *result.RemoteVid)
	}
	d.Set("description", result.Description)

	return nil
}
