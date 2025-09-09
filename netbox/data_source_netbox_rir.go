package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxRir() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxRirRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#regional-internet-registries-rirs):

> Regional Internet registries are responsible for the allocation of globally-routable address space. The five RIRs are ARIN, RIPE, APNIC, LACNIC, and AFRINIC. However, some address space has been set aside for internal use, such as defined in RFCs 1918 and 6598. NetBox considers these RFCs as a sort of RIR as well; that is, an authority which "owns" certain address space.`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"name", "slug"},
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"name", "slug"},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxRirRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamRirsListParams()

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.Slug = &slug
	}

	res, err := api.Ipam.IpamRirsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one rir returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no rir found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.Set("id", result.ID)
	d.Set("name", result.Name)
	d.Set("slug", result.Slug)
	d.Set("description", result.Description)
	d.Set("is_private", result.IsPrivate)
	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}
