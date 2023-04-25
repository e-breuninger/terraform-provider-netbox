package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCustomField() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCustomFieldCreate,
		Read:   resourceNetboxCustomFieldRead,
		Update: resourceNetboxCustomFieldUpdate,
		Delete: resourceNetboxCustomFieldDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/customization/custom-fields/#custom-fields):

> Each model in NetBox is represented in the database as a discrete table, and each attribute of a model exists as a column within its table. For example, sites are stored in the dcim_site table, which has columns named name, facility, physical_address, and so on. As new attributes are added to objects throughout the development of NetBox, tables are expanded to include new rows.
>
> However, some users might want to store additional object attributes that are somewhat esoteric in nature, and that would not make sense to include in the core NetBox database schema. For instance, suppose your organization needs to associate each device with a ticket number correlating it with an internal support system record. This is certainly a legitimate use for NetBox, but it's not a common enough need to warrant including a field for every NetBox installation. Instead, you can create a custom field to hold this data.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					models.CustomFieldTypeValueText,
					models.CustomFieldTypeValueInteger,
					models.CustomFieldTypeValueBoolean,
					models.CustomFieldTypeValueDate,
					models.CustomFieldTypeValueURL,
					models.CustomFieldTypeValueSelect,
					models.CustomFieldTypeValueMultiselect,
				}, false),
			},
			"content_types": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"weight": {
				Type:     schema.TypeInt,
				Required: true,
				DefaultFunc: func() (interface{}, error) {
					return 100, nil
				},
			},
			"choices": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Default:  nil,
			},
			"default": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"validation_maximum": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"validation_minimum": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"validation_regex": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCustomFieldUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := &models.WritableCustomField{
		Name:            strToPtr(d.Get("name").(string)),
		Type:            d.Get("type").(string),
		Description:     d.Get("description").(string),
		GroupName:       d.Get("group_name").(string),
		Label:           d.Get("label").(string),
		Required:        d.Get("required").(bool),
		ValidationRegex: d.Get("validation_regex").(string),
		Weight:          int64ToPtr(int64(d.Get("weight").(int))),
	}

	choices, ok := d.GetOk("choices")
	if ok {
		if data.Type != "select" && data.Type != "multiselect" {
			return fmt.Errorf("choices may be set only for custom selection fields")
		}
		for _, choice := range choices.(*schema.Set).List() {
			data.Choices = append(data.Choices, choice.(string))
		}
	}

	ctypes, ok := d.GetOk("content_types")
	if ok {
		for _, t := range ctypes.(*schema.Set).List() {
			data.ContentTypes = append(data.ContentTypes, t.(string))
		}
	}

	vmax, ok := d.GetOk("validation_maximum")
	if ok {
		data.ValidationMaximum = int64ToPtr(int64(vmax.(int)))
	}
	vmin, ok := d.GetOk("validation_minimum")
	if ok {
		data.ValidationMinimum = int64ToPtr(int64(vmin.(int)))
	}

	params := extras.NewExtrasCustomFieldsUpdateParams().WithID(id).WithData(data)
	res, err := api.Extras.ExtrasCustomFieldsUpdate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCustomFieldRead(d, m)
}

func resourceNetboxCustomFieldCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := &models.WritableCustomField{
		Name:            strToPtr(d.Get("name").(string)),
		Type:            d.Get("type").(string),
		Description:     d.Get("description").(string),
		GroupName:       d.Get("group_name").(string),
		Label:           d.Get("label").(string),
		Required:        d.Get("required").(bool),
		ValidationRegex: d.Get("validation_regex").(string),
		Weight:          int64ToPtr(int64(d.Get("weight").(int))),
	}

	choices, ok := d.GetOk("choices")
	if ok {
		if data.Type != "select" && data.Type != "multiselect" {
			return fmt.Errorf("choices may be set only for custom selection fields")
		}
		for _, choice := range choices.(*schema.Set).List() {
			data.Choices = append(data.Choices, choice.(string))
		}
	}

	ctypes, ok := d.GetOk("content_types")
	if ok {
		for _, t := range ctypes.(*schema.Set).List() {
			data.ContentTypes = append(data.ContentTypes, t.(string))
		}
	}

	vmax, ok := d.GetOk("validation_maximum")
	if ok {
		data.ValidationMaximum = int64ToPtr(int64(vmax.(int)))
	}
	vmin, ok := d.GetOk("validation_minimum")
	if ok {
		data.ValidationMinimum = int64ToPtr(int64(vmin.(int)))
	}

	params := extras.NewExtrasCustomFieldsCreateParams().WithData(data)

	res, err := api.Extras.ExtrasCustomFieldsCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCustomFieldRead(d, m)
}

func resourceNetboxCustomFieldRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasCustomFieldsReadParams().WithID(id)
	res, err := api.Extras.ExtrasCustomFieldsRead(params, nil)
	if err != nil {
		errapi, ok := err.(*extras.ExtrasCustomFieldsReadDefault)
		if !ok {
			return err
		}
		errorcode := errapi.Code()
		if errorcode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}
	d.Set("name", res.GetPayload().Name)
	d.Set("type", *res.GetPayload().Type.Value)

	d.Set("content_types", res.GetPayload().ContentTypes)

	choices := res.GetPayload().Choices
	if choices != nil {
		d.Set("choices", res.GetPayload().Choices)
	}

	d.Set("weight", res.GetPayload().Weight)
	if res.GetPayload().Default != nil {
		d.Set("default", res.GetPayload().Default)
	}

	d.Set("description", res.GetPayload().Description)
	d.Set("group_name", res.GetPayload().GroupName)
	d.Set("label", res.GetPayload().Label)
	d.Set("required", res.GetPayload().Required)

	d.Set("validation_maximum", res.GetPayload().ValidationMaximum)
	d.Set("validation_minimum", res.GetPayload().ValidationMinimum)
	d.Set("validation_regex", res.GetPayload().ValidationRegex)

	return nil
}

func resourceNetboxCustomFieldDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasCustomFieldsDeleteParams().WithID(id)
	_, err := api.Extras.ExtrasCustomFieldsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasCustomFieldsDeleteDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
			}
		}
		return err
	}
	return nil
}
