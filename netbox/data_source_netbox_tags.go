package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxTags() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxTagsRead,
		Description: `:meta:subcategory:Extras:`,
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
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"color": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxTagsRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := extras.NewExtrasTagsListParams()

	// Get user limit
	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	if filter, ok := d.GetOk("filter"); ok {
		filterParams := filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "id":
				params.ID = &vString
			case "id__gt":
				params.IDGt = &vString
			case "id__gte":
				params.IDGte = &vString
			case "id__lt":
				params.IDLt = &vString
			case "id__lte":
				params.IDLte = &vString
			case "name":
				params.Name = &vString
			case "name__ic":
				params.NameIc = &vString
			case "name__niew":
				params.NameNiew = &vString
			case "name__iew":
				params.NameIew = &vString
			case "name__nisw":
				params.NameNisw = &vString
			case "name__isw":
				params.NameIsw = &vString
			case "slug":
				params.Slug = &vString
			case "slug__ic":
				params.SlugIc = &vString
			case "slug__niew":
				params.SlugNiew = &vString
			case "slug__iew":
				params.SlugIew = &vString
			case "slug__nisw":
				params.SlugNisw = &vString
			case "slug__isw":
				params.SlugIsw = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination
	paginationHelper := NewPaginationHelper(userLimit)
	var allTags []*models.Tag

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Extras.ExtrasTagsList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch tags at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allTags = append(allTags, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allTags)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Trim to user limit if specified
	trimmedCount := paginationHelper.TrimToLimit(len(allTags))
	results := allTags[:trimmedCount]

	if len(results) == 0 {
		return errors.New("no result")
	}

	var s []map[string]interface{}
	for _, v := range results {
		mapping := make(map[string]interface{})

		mapping["tag_id"] = v.ID
		mapping["name"] = v.Name
		mapping["slug"] = v.Slug
		mapping["description"] = v.Description
		mapping["color"] = v.Color

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("tags", s)
}
