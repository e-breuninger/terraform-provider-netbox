package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxVirtualCircuitType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVirtualCircuitTypeCreate,
		Read:   resourceNetboxVirtualCircuitTypeRead,
		Update: resourceNetboxVirtualCircuitTypeUpdate,
		Delete: resourceNetboxVirtualCircuitTypeDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/models/circuits/virtualcircuittype/):

> Virtual circuits are classified by functional type, just as physical circuits are. These types are completely customizable, and are typically used to convey the type of service being delivered over a virtual circuit.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"color_hex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceNetboxVirtualCircuitTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableVirtualCircuitType{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Color = d.Get("color_hex").(string)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := circuits.NewCircuitsVirtualCircuitTypesCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsVirtualCircuitTypesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualCircuitTypeRead(d, m)
}

func resourceNetboxVirtualCircuitTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsVirtualCircuitTypesReadParams().WithID(id)

	res, err := api.Circuits.CircuitsVirtualCircuitTypesRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsVirtualCircuitTypesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	virtualCircuitType := res.GetPayload()

	d.Set("name", virtualCircuitType.Name)
	d.Set("slug", virtualCircuitType.Slug)
	d.Set("color_hex", virtualCircuitType.Color)
	d.Set("description", virtualCircuitType.Description)
	d.Set("comments", virtualCircuitType.Comments)

	api.readTags(d, virtualCircuitType.Tags)

	cf := getCustomFields(virtualCircuitType.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxVirtualCircuitTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualCircuitType{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Color = d.Get("color_hex").(string)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := circuits.NewCircuitsVirtualCircuitTypesPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsVirtualCircuitTypesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVirtualCircuitTypeRead(d, m)
}

func resourceNetboxVirtualCircuitTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsVirtualCircuitTypesDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsVirtualCircuitTypesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsVirtualCircuitTypesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
