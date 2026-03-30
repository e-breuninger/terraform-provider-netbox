package netbox

import (
	"fmt"
	"regexp"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxClusters() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxClustersRead,
		Description: `:meta:subcategory:Virtualization:`,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
			},
			"clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cluster_type_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cluster_group_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tenant_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"site_id": {
							Type:     schema.TypeInt,
							Computed: true,
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
						"scope_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scope_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"comments": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_fields": {
							Type:     schema.TypeMap,
							Computed: true,
						},
						tagsKey: tagsSchemaRead,
					},
				},
			},
		},
	}
}

func dataSourceNetboxClustersRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := virtualization.NewVirtualizationClustersListParams()

	// Get user limit (0 = fetch all)
	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		var tags []string
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "name":
				params.Name = &vString
			case "cluster_type_id":
				params.TypeID = &vString
			case "cluster_group_id":
				params.GroupID = &vString
			case "site_id":
				params.SiteID = &vString
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination (fetch all when name_regex is used)
	paginationHelper := NewPaginationHelper(FetchAll)
	var allClusters []*models.Cluster

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Virtualization.VirtualizationClustersList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch clusters at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allClusters = append(allClusters, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allClusters)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Apply name_regex filter
	var filteredClusters []*models.Cluster
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, cluster := range allClusters {
			if r.MatchString(*cluster.Name) {
				filteredClusters = append(filteredClusters, cluster)
			}
		}
	} else {
		filteredClusters = allClusters
	}

	// Apply user limit to filtered results
	if userLimit > 0 && int64(len(filteredClusters)) > userLimit {
		filteredClusters = filteredClusters[:userLimit]
	}

	var s []map[string]interface{}
	for _, cluster := range filteredClusters {
		var mapping = make(map[string]interface{})

		mapping["cluster_id"] = cluster.ID

		if cluster.Name != nil {
			mapping["name"] = *cluster.Name
		}

		if cluster.Type != nil {
			mapping["cluster_type_id"] = cluster.Type.ID
		}

		if cluster.Group != nil {
			mapping["cluster_group_id"] = cluster.Group.ID
		}

		if cluster.Tenant != nil {
			mapping["tenant_id"] = cluster.Tenant.ID
		}

		if cluster.Description != "" {
			mapping["description"] = cluster.Description
		}

		if cluster.Comments != "" {
			mapping["comments"] = cluster.Comments
		}

		if cluster.ScopeType != nil && cluster.ScopeID != nil {
			mapping["scope_type"] = *cluster.ScopeType
			mapping["scope_id"] = *cluster.ScopeID
			switch *cluster.ScopeType {
			case "dcim.site":
				mapping["site_id"] = *cluster.ScopeID
			case "dcim.sitegroup":
				mapping["site_group_id"] = *cluster.ScopeID
			case "dcim.location":
				mapping["location_id"] = *cluster.ScopeID
			case "dcim.region":
				mapping["region_id"] = *cluster.ScopeID
			}
		}

		if cluster.CustomFields != nil {
			mapping["custom_fields"] = cluster.CustomFields
		}

		if cluster.Tags != nil {
			mapping[tagsKey] = getTagListFromNestedTagList(cluster.Tags)
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("clusters", s)
}
