package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxCustomFieldChoiceSetBaseChoicesOptions = []string{"IATA", "ISO_3166", "UN_LOCODE"}

func resourceNetboxCustomFieldChoiceSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCustomFieldChoiceSetCreate,
		Read:   resourceNetboxCustomFieldChoiceSetRead,
		Update: resourceNetboxCustomFieldChoiceSetUpdate,
		Delete: resourceNetboxCustomFieldChoiceSetDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/customfieldchoiceset/):

Single- and multi-selection custom fields must define a set of valid choices from which the user may choose when defining the field value. These choices are defined as sets that may be reused among multiple custom fields.

A choice set must define a base choice set and/or a set of arbitrary extra choices.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"base_choices": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxCustomFieldChoiceSetBaseChoicesOptions, false),
				Description:  buildValidValueDescription(resourceNetboxCustomFieldChoiceSetBaseChoicesOptions),
				AtLeastOneOf: []string{"base_choices", "extra_choices"},
			},
			"extra_choices": {
				Type:        schema.TypeList,
				Description: "This length of the inner lists must be exactly two, where the first value is the value of a choice and the second value is the label of the choice.",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				Optional:     true,
				AtLeastOneOf: []string{"base_choices", "extra_choices"},
				//				ValidateFunc: func (i interface{}, k string) (warnings []string, _errors []error) {
				//	// Outer list length must be > 0
				//	extraChoiceListList := i.([][]string)
				//	if len(extraChoiceListList) == 0 {
				//		return nil, []error{errors.New("length of list must be > 0")}
				//	}
				//
				//	// Inner list length must be exactly 2
				//	for _, innerList := range extraChoiceListList {
				//		if len(innerList) != 2 {
				//			return nil, []error{errors.New("length of list must be > 0")}
				//		}
				//	}
				//
				//	return warnings, _errors
				//},
			},
			"order_alphabetically": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "experimental",
				Default:     false,
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCustomFieldChoiceSetCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)

	data := models.CustomFieldChoiceSet{
		Name: &name,
	}

	data.Description = getOptionalStr(d, "description", false)

	var extraChoiceListList [][]string

	extraChoices, ok := d.GetOk("extra_choices")
	if ok {
		for _, innerList := range extraChoices.([]interface{}) {
			tmp := innerList.([]interface{})
			if len(tmp) != 2 {
				return errors.New("length of inner lists must be exactly two for custom field choice sets")
			}
			extraChoiceListList = append(extraChoiceListList, []string{tmp[0].(string), tmp[1].(string)})
		}
		data.ExtraChoices = extraChoiceListList
	}

	params := extras.NewExtrasCustomFieldChoiceSetsCreateParams().WithData(&data)

	res, err := api.Extras.ExtrasCustomFieldChoiceSetsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCustomFieldChoiceSetRead(d, m)
}

func resourceNetboxCustomFieldChoiceSetRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasCustomFieldChoiceSetsReadParams().WithID(id)

	res, err := api.Extras.ExtrasCustomFieldChoiceSetsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*extras.ExtrasCustomFieldChoiceSetsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	choiceSet := res.GetPayload()

	d.Set("name", choiceSet.Name)

	if choiceSet.Description != "" {
		d.Set("description", choiceSet.Description)
	} else {
		d.Set("description", nil)
	}

	return nil
}

func resourceNetboxCustomFieldChoiceSetUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)

	data := models.CustomFieldChoiceSet{
		Name: &name,
	}

	data.Description = getOptionalStr(d, "description", true)

	var extraChoiceListList [][]string

	extraChoices, ok := d.GetOk("extra_choices")
	if ok {
		for _, innerList := range extraChoices.([]interface{}) {
			tmp := innerList.([]interface{})
			if len(tmp) != 2 {
				return errors.New("length of inner lists must be exactly two for custom field choice sets")
			}
			extraChoiceListList = append(extraChoiceListList, []string{tmp[0].(string), tmp[1].(string)})
		}
		data.ExtraChoices = extraChoiceListList
	}

	params := extras.NewExtrasCustomFieldChoiceSetsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasCustomFieldChoiceSetsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCustomFieldChoiceSetRead(d, m)
}

func resourceNetboxCustomFieldChoiceSetDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasCustomFieldChoiceSetsDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasCustomFieldChoiceSetsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasCustomFieldChoiceSetsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
