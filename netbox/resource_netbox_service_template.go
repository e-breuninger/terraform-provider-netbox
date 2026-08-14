package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxServiceTemplateProtocolOptions = []string{"tcp", "udp", "sctp"}

func resourceNetboxServiceTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxServiceTemplateCreate,
		Read:   resourceNetboxServiceTemplateRead,
		Update: resourceNetboxServiceTemplateUpdate,
		Delete: resourceNetboxServiceTemplateDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/servicetemplate/):

> A service template can be used to represent a layer four TCP or UDP service which is commonly used within your organization. This can then be quickly applied to devices and virtual machines using it as a template for the service.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"protocol": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(resourceNetboxServiceTemplateProtocolOptions, false)),
				Description:      buildValidValueDescription(resourceNetboxServiceTemplateProtocolOptions),
			},
			"ports": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxServiceTemplateCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.WritableServiceTemplate{}

	name := d.Get("name").(string)
	data.Name = &name

	protocol := d.Get("protocol").(string)
	data.Protocol = &protocol

	var ports []int64
	for _, v := range d.Get("ports").(*schema.Set).List() {
		ports = append(ports, int64(v.(int)))
	}
	data.Ports = ports

	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamServiceTemplatesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamServiceTemplatesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxServiceTemplateRead(d, m)
}

func resourceNetboxServiceTemplateRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamServiceTemplatesReadParams().WithID(id)

	res, err := api.Ipam.IpamServiceTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamServiceTemplatesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	serviceTemplate := res.GetPayload()

	d.Set("name", serviceTemplate.Name)
	if serviceTemplate.Protocol != nil {
		d.Set("protocol", serviceTemplate.Protocol.Value)
	}
	d.Set("ports", serviceTemplate.Ports)
	d.Set("description", serviceTemplate.Description)
	d.Set("comments", serviceTemplate.Comments)

	api.readTags(d, serviceTemplate.Tags)

	cf := getCustomFields(serviceTemplate.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxServiceTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableServiceTemplate{}

	name := d.Get("name").(string)
	data.Name = &name

	protocol := d.Get("protocol").(string)
	data.Protocol = &protocol

	var ports []int64
	for _, v := range d.Get("ports").(*schema.Set).List() {
		ports = append(ports, int64(v.(int)))
	}
	data.Ports = ports

	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamServiceTemplatesUpdateParams().WithID(id).WithData(&data)
	_, err = api.Ipam.IpamServiceTemplatesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxServiceTemplateRead(d, m)
}

func resourceNetboxServiceTemplateDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamServiceTemplatesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamServiceTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamServiceTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.SetId("")
	return nil
}
