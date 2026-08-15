package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxExportTemplate() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxExportTemplateRead,
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
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"template_code": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"mime_type": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"file_extension": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"as_attachment": {
				Computed: true,
				Type:     schema.TypeBool,
			},
		},
	}
}

func dataSourceNetboxExportTemplateRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := extras.NewExtrasExportTemplatesListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Extras.ExtrasExportTemplatesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no export template found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("content_types", result.ContentTypes)
	d.Set("description", result.Description)
	d.Set("template_code", result.TemplateCode)
	d.Set("mime_type", result.MimeType)
	d.Set("file_extension", result.FileExtension)
	d.Set("as_attachment", result.AsAttachment)

	return nil
}
