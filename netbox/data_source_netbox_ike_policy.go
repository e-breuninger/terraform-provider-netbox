package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxIkePolicy() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxIkePolicyRead,
		Description: `:meta:subcategory:VPN Tunnels:`,
		Schema: map[string]*schema.Schema{
			"id": {
				AtLeastOneOf: []string{"id", "name"},
				Computed:     true,
				Optional:     true,
				Type:         schema.TypeString,
			},
			"name": {
				AtLeastOneOf: []string{"id", "name"},
				Optional:     true,
				Type:         schema.TypeString,
			},
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"version": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"mode": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"proposal_ids": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"preshared_key": {
				Computed:  true,
				Sensitive: true,
				Type:      schema.TypeString,
			},
			"comments": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"custom_fields": {
				Computed: true,
				Type:     schema.TypeMap,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxIkePolicyRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := vpn.NewVpnIkePoliciesListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Vpn.VpnIkePoliciesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no IKE policy found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	d.Set("preshared_key", result.PresharedKey)

	if result.Version != nil {
		d.Set("version", result.Version.Value)
	}

	if result.Mode != nil {
		d.Set("mode", result.Mode.Value)
	}

	proposals := result.Proposals
	proposalIDs := make([]int64, len(proposals))
	for i, v := range proposals {
		proposalIDs[i] = v.ID
	}
	d.Set("proposal_ids", proposalIDs)

	if result.CustomFields != nil {
		d.Set("custom_fields", flattenCustomFields(result.CustomFields))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
