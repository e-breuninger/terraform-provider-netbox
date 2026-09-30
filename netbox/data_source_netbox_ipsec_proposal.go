package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxIpsecProposal() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxIpsecProposalRead,
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
			"encryption_algorithm": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"authentication_algorithm": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"sa_lifetime_seconds": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"sa_lifetime_data": {
				Computed: true,
				Type:     schema.TypeInt,
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

func dataSourceNetboxIpsecProposalRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := vpn.NewVpnIpsecProposalsListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Vpn.VpnIpsecProposalsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no IPSec proposal found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	d.Set("sa_lifetime_seconds", result.SaLifetimeSeconds)
	d.Set("sa_lifetime_data", result.SaLifetimeData)

	if result.EncryptionAlgorithm != nil {
		d.Set("encryption_algorithm", result.EncryptionAlgorithm.Value)
	}

	if result.AuthenticationAlgorithm != nil {
		d.Set("authentication_algorithm", result.AuthenticationAlgorithm.Value)
	}

	if result.CustomFields != nil {
		d.Set("custom_fields", flattenCustomFields(result.CustomFields))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
