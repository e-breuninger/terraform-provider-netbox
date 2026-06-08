package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxContact() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxContactRead,
		Description: `:meta:subcategory:Tenancy:`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				AtLeastOneOf: []string{"name", "slug"},
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"name", "slug"},
			},
			"group_id": {
				Type:       schema.TypeInt,
				Deprecated: "This field is deprecated. Please use the new \"group_ids\" attribute instead.",
				Computed:   true,
			},
			"group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNetboxContactRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	params := tenancy.NewTenancyContactsListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Tenancy.TenancyContactsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one contact returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no contact found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)

	if result.Groups != nil {
		groups := result.Groups
		groupIDs := make([]int64, len(groups))
		for i, group := range groups {
			groupIDs[i] = group.ID
		}
		d.Set("group_ids", groupIDs)
		if len(groupIDs) > 0 {
			d.Set("group_id", groupIDs[0])
		}
	}

	return nil
}
