package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxAsn() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxAsnRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"asn": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"asn", "tag"},
			},
			"tag": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"asn", "tag"},
				Description:  "Tag to include in the data source filter (must match the tag's slug).",
			},
			"tag__n": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Tag to exclude from the data source filter (must match the tag's slug).
Refer to [Netbox's documentation](https://demo.netbox.dev/static/docs/rest-api/filtering/#lookup-expressions)
for more information on available lookup expressions.`,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchemaRead,
		},
	}
}

func dataSourceNetboxAsnRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := ipam.NewIpamAsnsListParams()

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	if asn, ok := d.Get("asn").(string); ok && asn != "" {
		params.Asn = &asn
	}

	if tag, ok := d.Get("tag").(string); ok && tag != "" {
		params.Tag = []string{tag} //TODO: switch schema to list?
	}
	if tagn, ok := d.Get("tag__n").(string); ok && tagn != "" {
		params.Tagn = &tagn
	}

	res, err := api.Ipam.IpamAsnsList(params, nil)
	if err != nil {
		return err
	}

	if count := *res.GetPayload().Count; count != int64(1) {
		return fmt.Errorf("expected one ASN, but got %d", count)
	}

	result := res.GetPayload().Results[0]
	d.Set("id", result.ID)
	d.Set("asn", strconv.FormatInt(*result.Asn, 10))
	d.Set("description", result.Description)
	d.Set("tags", getTagListFromNestedTagList(result.Tags))
	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}
