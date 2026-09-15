package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxL2VPN() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxL2VPNRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
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
			"slug": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"type": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"identifier": {
				Computed: true,
				Type:     schema.TypeInt,
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
			"import_targets": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"export_targets": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			customFieldsKey: {
				Computed: true,
				Type:     schema.TypeMap,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxL2VPNRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamL2vpnsListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "" && id != "0" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Ipam.IpamL2vpnsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no L2VPN found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("slug", result.Slug)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)

	if result.Type != nil {
		d.Set("type", result.Type.Value)
	}

	if result.Identifier != nil {
		d.Set("identifier", result.Identifier)
	}

	if result.Tenant != nil {
		d.Set("tenant_id", result.Tenant.ID)
	}

	d.Set("import_targets", getIDsFromNestedRouteTargetList(result.ImportTargets))
	d.Set("export_targets", getIDsFromNestedRouteTargetList(result.ExportTargets))

	if result.CustomFields != nil {
		d.Set(customFieldsKey, flattenCustomFields(result.CustomFields))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
