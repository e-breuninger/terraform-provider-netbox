package netbox

import (
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxWebhookHTTPMethodOptions = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

func resourceNetboxWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxWebhookCreate,
		Read:   resourceNetboxWebhookRead,
		Update: resourceNetboxWebhookUpdate,
		Delete: resourceNetboxWebhookDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/integrations/webhooks/):

> A webhook is a mechanism for conveying to some external system a change that took place in NetBox. For example, you may want to notify a monitoring system whenever the status of a device is updated in NetBox. This can be done by creating a webhook for the device model in NetBox and identifying the webhook receiver. When NetBox detects a change to a device, an HTTP request containing the details of the change and who made it be sent to the specified receiver.`,

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
			"trigger_on_create": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"trigger_on_update": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"trigger_on_delete": {
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
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					equal, _ := jsonSemanticCompare(oldValue, newValue)
					return equal
				},
				DiffSuppressOnRefresh: true,
			},
			"http_method": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxWebhookHTTPMethodOptions, false),
				Description:  buildValidValueDescription(resourceNetboxWebhookHTTPMethodOptions),
				Default:      "POST",
			},
			"http_content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The complete list of official content types is available [here](https://www.iana.org/assignments/media-types/media-types.xhtml).",
				Default:     "application/json",
			},
			"additional_headers": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"conditions": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					equal, _ := jsonSemanticCompare(oldValue, newValue)
					return equal
				},
				DiffSuppressOnRefresh: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxWebhookCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := &models.Webhook{}
	for _, contentType := range d.Get("content_types").(*schema.Set).List() {
		data.ContentTypes = append(data.ContentTypes, contentType.(string))
	}
	name := d.Get("name").(string)
	data.Name = &name
	triggerOnCreate := d.Get("trigger_on_create").(bool)
	data.TypeCreate = triggerOnCreate
	triggerOnUpdate := d.Get("trigger_on_update").(bool)
	data.TypeUpdate = triggerOnUpdate
	triggerOnDelete := d.Get("trigger_on_delete").(bool)
	data.TypeDelete = triggerOnDelete
	enabled := d.Get("enabled").(bool)
	data.Enabled = enabled
	payloadURL := d.Get("payload_url").(string)
	data.PayloadURL = &payloadURL
	bodyTemplate := d.Get("body_template").(string)
	data.BodyTemplate = bodyTemplate
	data.HTTPMethod = getOptionalStr(d, "http_method", false)
	data.HTTPContentType = getOptionalStr(d, "http_content_type", false)
	data.AdditionalHeaders = getOptionalStr(d, "additional_headers", false)

	if conditionsData, ok := d.GetOk("conditions"); ok {
		var conditions any
		err := json.Unmarshal([]byte(conditionsData.(string)), &conditions)
		if err != nil {
			return err
		}
		data.Conditions = conditions
	}

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

	webhook := res.GetPayload()
	d.Set("content_types", webhook.ContentTypes)
	d.Set("name", webhook.Name)
	d.Set("trigger_on_create", webhook.TypeCreate)
	d.Set("trigger_on_update", webhook.TypeUpdate)
	d.Set("trigger_on_delete", webhook.TypeDelete)
	d.Set("enabled", webhook.Enabled)
	d.Set("payload_url", webhook.PayloadURL)
	d.Set("body_template", webhook.BodyTemplate)
	d.Set("http_method", webhook.HTTPMethod)
	d.Set("http_content_type", webhook.HTTPContentType)
	d.Set("additional_headers", webhook.AdditionalHeaders)

	if webhook.Conditions != nil {
		conditions, err := json.Marshal(webhook.Conditions)
		if err != nil {
			return err
		}
		d.Set("conditions", string(conditions))
	}

	return nil
}

func resourceNetboxWebhookUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.Webhook{}

	for _, contentType := range d.Get("content_types").(*schema.Set).List() {
		data.ContentTypes = append(data.ContentTypes, contentType.(string))
	}

	name := d.Get("name").(string)
	triggerOnCreate := d.Get("trigger_on_create").(bool)
	triggerOnUpdate := d.Get("trigger_on_update").(bool)
	triggerOnDelete := d.Get("trigger_on_delete").(bool)
	enabled := d.Get("enabled").(bool)
	payloadURL := d.Get("payload_url").(string)
	bodyTemplate := d.Get("body_template").(string)

	data.Name = &name
	data.TypeCreate = triggerOnCreate
	data.TypeUpdate = triggerOnUpdate
	data.TypeDelete = triggerOnDelete
	data.Enabled = enabled
	data.PayloadURL = &payloadURL
	data.BodyTemplate = bodyTemplate
	data.HTTPMethod = getOptionalStr(d, "http_method", false)
	data.HTTPContentType = getOptionalStr(d, "http_content_type", false)
	data.AdditionalHeaders = getOptionalStr(d, "additional_headers", false)

	if conditionsData, ok := d.GetOk("conditions"); ok {
		var conditions any
		err := json.Unmarshal([]byte(conditionsData.(string)), &conditions)
		if err != nil {
			return err
		}
		data.Conditions = conditions
	}

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
