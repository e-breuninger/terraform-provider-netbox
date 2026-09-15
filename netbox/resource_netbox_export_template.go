package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxExportTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxExportTemplateCreate,
		Read:   resourceNetboxExportTemplateRead,
		Update: resourceNetboxExportTemplateUpdate,
		Delete: resourceNetboxExportTemplateDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/customization/export-templates/):

> NetBox allows users to define custom templates that can be used when exporting objects. To create an export template, connect to the Django shell and create a new ExportTemplate instance. Each export template is associated with a certain type of object and has a name. Optionally, a MIME type and file extension can be defined; these will be used when the template is rendered for a client for the file download. If a MIME type is not defined, NetBox will fall back to using text/plain.

> Export templates must be written in Jinja2. The list of objects returned from the database when the template is rendered is passed in as a context variable named queryset.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"content_types": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"template_code": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Jinja2 template code. The list of objects being exported is passed as a context variable named `queryset`.",
			},
			"mime_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
				Description:  "Defaults to `text/plain`",
			},
			"file_extension": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Description:  "Extension to append to the rendered filename",
			},
			"as_attachment": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Download file as attachment",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxExportTemplateCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := &models.ExportTemplate{}

	name := d.Get("name").(string)
	data.Name = &name
	data.ContentTypes = toStringList(d.Get("content_types"))
	data.Description = getOptionalStr(d, "description", false)
	templateCode := d.Get("template_code").(string)
	data.TemplateCode = &templateCode
	data.MimeType = getOptionalStr(d, "mime_type", false)
	data.FileExtension = getOptionalStr(d, "file_extension", false)
	data.AsAttachment = d.Get("as_attachment").(bool)

	params := extras.NewExtrasExportTemplatesCreateParams().WithData(data)

	res, err := api.Extras.ExtrasExportTemplatesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxExportTemplateRead(d, m)
}

func resourceNetboxExportTemplateRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasExportTemplatesReadParams().WithID(id)

	res, err := api.Extras.ExtrasExportTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasExportTemplatesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	exportTemplate := res.GetPayload()
	d.Set("name", exportTemplate.Name)
	d.Set("content_types", exportTemplate.ContentTypes)
	d.Set("description", exportTemplate.Description)
	d.Set("template_code", exportTemplate.TemplateCode)
	d.Set("mime_type", exportTemplate.MimeType)
	d.Set("file_extension", exportTemplate.FileExtension)
	d.Set("as_attachment", exportTemplate.AsAttachment)

	return nil
}

func resourceNetboxExportTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.ExportTemplate{}

	name := d.Get("name").(string)
	data.Name = &name
	data.ContentTypes = toStringList(d.Get("content_types"))
	data.Description = getOptionalStr(d, "description", true)
	templateCode := d.Get("template_code").(string)
	data.TemplateCode = &templateCode
	data.MimeType = getOptionalStr(d, "mime_type", true)
	data.FileExtension = getOptionalStr(d, "file_extension", true)
	data.AsAttachment = d.Get("as_attachment").(bool)

	params := extras.NewExtrasExportTemplatesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasExportTemplatesUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxExportTemplateRead(d, m)
}

func resourceNetboxExportTemplateDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasExportTemplatesDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasExportTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasExportTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
