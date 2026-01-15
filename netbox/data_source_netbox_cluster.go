package netbox

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxCluster() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxClusterRead,
		Description: `:meta:subcategory:Virtualization:`,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"name", "site_id", "id"},
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_id": {
				Type:         schema.TypeInt,
				Computed:     true,
				Optional:     true,
				AtLeastOneOf: []string{"name", "site_id", "id"},
			},
			"site_group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"location_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"name", "site_id", "id"},
			},
			"cluster_type_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cluster_group_id": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"scope_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scope_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"custom_fields": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxClusterRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := virtualization.NewVirtualizationClustersListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if siteID, ok := d.Get("site_id").(int); ok && siteID != 0 {
		params.SiteID = strToPtr(strconv.FormatInt(int64(siteID), 10))
	}

	if id, ok := d.Get("id").(string); ok && id != "0" {
		params.SetID(&id)
	}

	if clustergroupID, ok := d.Get("cluster_group_id").(int); ok && clustergroupID != 0 {
		clustGroupStr := fmt.Sprintf("%d", clustergroupID)
		params.GroupID = &clustGroupStr
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Virtualization.VirtualizationClustersList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}
	result := res.GetPayload().Results[0]
	d.Set("cluster_id", result.ID)
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("cluster_type_id", result.Type.ID)

	if result.Group != nil {
		d.Set("cluster_group_id", result.Group.ID)
	} else {
		d.Set("cluster_group_id", nil)
	}
	d.Set("comments", result.Comments)
	d.Set("description", result.Description)

	if result.ScopeType != nil && result.ScopeID != nil {
		d.Set("scope_type", result.ScopeType)
		d.Set("scope_id", result.ScopeID)
		scopeID := result.ScopeID
		switch scopeType := result.ScopeType; *scopeType {
		case "dcim.site":
			d.Set("site_id", scopeID)
		case "dcim.sitegroup":
			d.Set("site_group_id", scopeID)
		case "dcim.location":
			d.Set("location_id", scopeID)
		case "dcim.region":
			d.Set("region_id", scopeID)
		}
	}

	if result.CustomFields != nil {
		d.Set("custom_fields", result.CustomFields)
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
