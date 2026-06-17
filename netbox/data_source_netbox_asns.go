package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxAsns() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxAsnsRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
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
			"asns": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"asn": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rir_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tags": tagsSchemaRead,
					},
				},
			},
		},
	}
}

func dataSourceNetboxAsnsRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamAsnsListParams()

	// Get user limit
	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "asn":
				params.Asn = &vString
			case "asn__gte":
				params.AsnGte = &vString
			case "asn__lte":
				params.AsnLte = &vString
			case "asn__n":
				params.Asnn = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination
	paginationHelper := NewPaginationHelper(userLimit)
	var allAsns []*models.ASN

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Ipam.IpamAsnsList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch ASNs at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allAsns = append(allAsns, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allAsns)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Trim to user limit if specified
	trimmedCount := paginationHelper.TrimToLimit(len(allAsns))
	filteredAsns := allAsns[:trimmedCount]

	if len(filteredAsns) == 0 {
		return errors.New("no result")
	}

	var s []map[string]interface{}
	for _, v := range filteredAsns {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["asn"] = v.Asn
		mapping["rir_id"] = v.Rir.ID
		mapping["tags"] = getTagListFromNestedTagList(v.Tags)

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("asns", s)
}
