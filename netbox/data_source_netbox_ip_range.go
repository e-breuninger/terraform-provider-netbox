package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxIPRange() *schema.Resource {
	filterAtLeastOneOf := []string{
		"contains",
		"family",
		"vrf_id",
		"tenant_id",
		"status",
		"role_id",
		"description",
		"tag",
	}
	return &schema.Resource{
		Read:        dataSourceNetboxIPRangeRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"contains": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: filterAtLeastOneOf,
				ValidateFunc: validation.IsCIDR,
			},
			"family": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: filterAtLeastOneOf,
				ValidateFunc: validation.IntInSlice([]int{4, 6}),
				Description:  "The IP family of the IP range. One of 4 or 6",
			},
			"vrf_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: filterAtLeastOneOf,
			},
			"tenant_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: filterAtLeastOneOf,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: filterAtLeastOneOf,
			},
			"role_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: filterAtLeastOneOf,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: filterAtLeastOneOf,
				Description:  "Description to include in the data source filter.",
			},
			"tag": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: filterAtLeastOneOf,
				Description:  "Tag to include in the data source filter (must match the tag's slug).",
			},
			"tag__n": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Tag to exclude from the data source filter (must match the tag's slug).
Refer to [Netbox's documentation](https://netboxlabs.com/docs/netbox/reference/filtering/#lookup-expressions)
for more information on available lookup expressions.`,
			},
			"tags": tagsSchemaRead,
		},
	}
}

func dataSourceNetboxIPRangeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamIPRangesListParams()

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	if contains, ok := d.Get("contains").(string); ok && contains != "" {
		params.Contains = &contains
	}

	if family, ok := d.Get("family").(int); ok && family != 0 {
		familyFloat := float64(family)
		params.Family = &familyFloat
	}

	if vrfID, ok := d.Get("vrf_id").(int); ok && vrfID != 0 {
		// Note that vrf_id is a string pointer in the netbox filter, but we use a number in the provider
		params.VrfID = strToPtr(strconv.Itoa(vrfID))
	}

	if tenantID, ok := d.Get("tenant_id").(int); ok && tenantID != 0 {
		// Note that tenant_id is a string pointer in the netbox filter, but we use a number in the provider
		params.TenantID = strToPtr(strconv.Itoa(tenantID))
	}

	if status, ok := d.Get("status").(string); ok && status != "" {
		params.Status = &status
	}

	if roleID, ok := d.Get("role_id").(int); ok && roleID != 0 {
		params.RoleID = strToPtr(strconv.Itoa(roleID))
	}

	if description, ok := d.Get("description").(string); ok && description != "" {
		params.Description = &description
	}

	if tag, ok := d.Get("tag").(string); ok && tag != "" {
		params.Tag = []string{tag} //TODO: switch schema to list
	}
	if tagn, ok := d.Get("tag__n").(string); ok && tagn != "" {
		params.Tagn = &tagn
	}

	res, err := api.Ipam.IpamIPRangesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than IP range returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no IP range found matching filter")
	}

	result := res.GetPayload().Results[0]
	d.Set("id", result.ID)
	d.Set("start_address", result.StartAddress)
	d.Set("end_address", result.EndAddress)
	d.Set("status", result.Status.Value)
	d.Set("description", result.Description)
	d.Set("family", int(*result.Family.Value))
	d.Set("tags", getTagListFromNestedTagList(result.Tags))

	if result.Role != nil {
		d.Set("role_id", result.Role.ID)
	}
	if result.Vrf != nil {
		d.Set("vrf_id", result.Vrf.ID)
	}
	if result.Tenant != nil {
		d.Set("tenant_id", result.Tenant.ID)
	}

	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}
