package netbox

import (
	"fmt"
	"regexp"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxCircuitProviderNetworks() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxCircuitProviderNetworksRead,
		Description: `:meta:subcategory:Circuits:`,
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
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
				Default:          0,
			},
			"provider_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_fields": {
							Type:     schema.TypeMap,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provider_network_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"provider_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						tagsKey: tagsSchemaRead,
					},
				},
			},
		},
	}
}

func dataSourceNetboxCircuitProviderNetworksRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := circuits.NewCircuitsProviderNetworksListParams()

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
			case "provider_id":
				params.ProviderID = &vString
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination (fetch all when name_regex is used)
	paginationHelper := NewPaginationHelper(userLimit)
	var allProviderNetworks []*models.ProviderNetwork

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Circuits.CircuitsProviderNetworksList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch Provider Networks at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allProviderNetworks = append(allProviderNetworks, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allProviderNetworks)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Apply name_regex filter
	var filteredProviderNetworks []*models.ProviderNetwork
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, provider_network := range allProviderNetworks {
			if r.MatchString(*provider_network.Name) {
				filteredProviderNetworks = append(filteredProviderNetworks, provider_network)
			}
		}
	} else {
		filteredProviderNetworks = allProviderNetworks
	}

	// Trim to user limit if specified
	trimmedCount := paginationHelper.TrimToLimit(len(filteredProviderNetworks))
	filteredProviderNetworks = filteredProviderNetworks[:trimmedCount]

	var s []map[string]interface{}
	for _, provider_network := range filteredProviderNetworks {
		var mapping = make(map[string]interface{})

		mapping["provider_network_id"] = provider_network.ID

		if provider_network.Comments != "" {
			mapping["comments"] = provider_network.Comments
		}

		if provider_network.CustomFields != nil {
			mapping["custom_fields"] = provider_network.CustomFields
		}

		if provider_network.Description != "" {
			mapping["description"] = provider_network.Description
		}

		if provider_network.Name != nil {
			mapping["name"] = *provider_network.Name
		}

		if provider_network.Provider != nil {
			mapping["provider_id"] = provider_network.Provider.ID
		}

		if provider_network.ServiceID != "" {
			mapping["service_id"] = provider_network.ServiceID
		}

		if provider_network.Tags != nil {
			mapping[tagsKey] = getTagListFromNestedTagList(provider_network.Tags)
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("provider_networks", s)
}
