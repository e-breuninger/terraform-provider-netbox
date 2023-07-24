package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxWebhookCreate,
		Read:   resourceNetboxWebhookRead,
		Update: resourceNetboxWebhookUpdate,
		Delete: resourceNetboxWebhookDelete,

		Description: `:meta:subcategory:Webhook:From the [official documentation](https://docs.netbox.dev/en/stable/integrations/webhooks/):

> Webhooks are used to send HTTP POST requests to a specified URL in response to events within NetBox.
> You can configure different types of events`,

		Schema: map[string]*schema.Schema{
			"content_types": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type_create": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type_update": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type_delete": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"payload_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"body_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxWebhookCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := &models.WritableWebhook{}
	for _, contentType := range d.Get("content_types").(*schema.Set).List() {
		data.ContentTypes = append(data.ContentTypes, contentType.(string))
	}
	name := d.Get("name").(string)
	data.Name = &name
	type_create := d.Get("type_create").(bool)
	data.TypeCreate = type_create
	type_update := d.Get("type_update").(bool)
	data.TypeUpdate = type_update
	type_delete := d.Get("type_delete").(bool)
	data.TypeDelete = type_delete
	enabled := d.Get("enabled").(bool)
	data.Enabled = enabled
	payload_url := d.Get("payload_url").(string)
	data.PayloadURL = &payload_url
	body_template := d.Get("body_template").(string)
	data.BodyTemplate = body_template

	params := extras.NewExtrasWebhooksCreateParams().WithData(data)

	res, err := api.Extras.ExtrasWebhooksCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxWebhookRead(d, m)
}

func resourceNetboxWebhookRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasWebhooksReadParams().WithID(id)

	res, err := api.Extras.ExtrasWebhooksRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasWebhooksReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.Set("content_types", res.GetPayload().ContentTypes)
	d.Set("name", res.GetPayload().Name)
	d.Set("type_create", res.GetPayload().TypeCreate)
	d.Set("type_update", res.GetPayload().TypeUpdate)
	d.Set("type_delete", res.GetPayload().TypeDelete)
	d.Set("enabled", res.GetPayload().Enabled)
	d.Set("payload_url", res.GetPayload().PayloadURL)
	d.Set("body_template", res.GetPayload().BodyTemplate)

	return nil
}

func resourceNetboxWebhookUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableWebhook{}

	for _, contentType := range d.Get("content_types").(*schema.Set).List() {
		data.ContentTypes = append(data.ContentTypes, contentType.(string))
	}

	name := d.Get("name").(string)
	type_create := d.Get("type_create").(bool)
	type_update := d.Get("type_update").(bool)
	type_delete := d.Get("type_delete").(bool)
	enabled := d.Get("enabled").(bool)
	payload_url := d.Get("payload_url").(string)
	body_template := d.Get("body_template").(string)

	data.Name = &name
	data.TypeCreate = type_create
	data.TypeUpdate = type_update
	data.TypeDelete = type_delete
	data.Enabled = enabled
	data.PayloadURL = &payload_url
	data.BodyTemplate = body_template

	params := extras.NewExtrasWebhooksUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasWebhooksUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxWebhookRead(d, m)
}

func resourceNetboxWebhookDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasWebhooksDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasWebhooksDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasWebhooksDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
