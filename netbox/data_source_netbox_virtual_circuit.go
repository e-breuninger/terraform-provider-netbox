package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxVirtualCircuit() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVirtualCircuitRead,
		Description: `:meta:subcategory:Circuits:`,
		Schema: map[string]*schema.Schema{
			"id": {
				AtLeastOneOf: []string{"id", "cid"},
				Computed:     true,
				Optional:     true,
				Type:         schema.TypeString,
			},
			"cid": {
				AtLeastOneOf: []string{"id", "cid"},
				Optional:     true,
				Type:         schema.TypeString,
			},
			"provider_network_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"provider_account_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"type_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"status": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"tenant_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"description": {
				Computed: true,
				Type:     schema.TypeString,
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

func dataSourceNetboxVirtualCircuitRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := circuits.NewCircuitsVirtualCircuitsListParams()

	if cid, ok := d.Get("cid").(string); ok && cid != "" {
		params.Cid = &cid
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.ID = &id
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Circuits.CircuitsVirtualCircuitsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no virtual circuit found matching filter")
	}
	result := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("cid", result.Cid)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)

	if result.Status != nil {
		d.Set("status", result.Status.Value)
	}

	if result.ProviderNetwork != nil {
		d.Set("provider_network_id", result.ProviderNetwork.ID)
	}

	if result.ProviderAccount != nil {
		d.Set("provider_account_id", result.ProviderAccount.ID)
	}

	if result.Type != nil {
		d.Set("type_id", result.Type.ID)
	}

	if result.Tenant != nil {
		d.Set("tenant_id", result.Tenant.ID)
	}

	if result.CustomFields != nil {
		d.Set("custom_fields", flattenCustomFields(result.CustomFields))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
