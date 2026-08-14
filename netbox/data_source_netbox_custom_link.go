package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxCustomLink() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxCustomLinkRead,
		Description: `:meta:subcategory:Extras:`,
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
			"content_types": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enabled": {
				Computed: true,
				Type:     schema.TypeBool,
			},
			"link_text": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"link_url": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"weight": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"group_name": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"button_class": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"new_window": {
				Computed: true,
				Type:     schema.TypeBool,
			},
		},
	}
}

func dataSourceNetboxCustomLinkRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := extras.NewExtrasCustomLinksListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Extras.ExtrasCustomLinksList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no custom link found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("content_types", result.ContentTypes)
	d.Set("enabled", result.Enabled)
	d.Set("link_text", result.LinkText)
	d.Set("link_url", result.LinkURL)
	d.Set("weight", result.Weight)
	d.Set("group_name", result.GroupName)
	d.Set("button_class", result.ButtonClass)
	d.Set("new_window", result.NewWindow)

	return nil
}
