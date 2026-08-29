package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxCableBundle() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxCableBundleRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
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
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"comments": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"cable_count": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"custom_fields": {
				Computed: true,
				Type:     schema.TypeMap,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxCableBundleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := dcim.NewDcimCableBundlesListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Dcim.DcimCableBundlesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no cable bundle found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	d.Set("cable_count", result.CableCount)

	if result.CustomFields != nil {
		d.Set("custom_fields", flattenCustomFields(result.CustomFields))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
