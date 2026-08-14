package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxProviderAccount() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxProviderAccountRead,
		Description: `:meta:subcategory:Circuits:`,
		Schema: map[string]*schema.Schema{
			"account": {
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
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"name": {
				AtLeastOneOf: []string{"id", "name"},
				Optional:     true,
				Type:         schema.TypeString,
			},
			"id": {
				AtLeastOneOf: []string{"id", "name"},
				Computed:     true,
				Optional:     true,
				Type:         schema.TypeString,
			},
			"provider_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxProviderAccountRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := circuits.NewCircuitsProviderAccountsListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Circuits.CircuitsProviderAccountsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no Provider Account found")
	}
	result := res.GetPayload().Results[0]
	d.Set("account", result.Account)
	d.Set("comments", result.Comments)
	d.Set("description", result.Description)
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)

	if result.Provider != nil {
		d.Set("provider_id", result.Provider.ID)
	}

	if result.CustomFields != nil {
		d.Set("custom_fields", flattenCustomFields((result.CustomFields)))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
