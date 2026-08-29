package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxCircuitGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitGroupCreate,
		Read:   resourceNetboxCircuitGroupRead,
		Update: resourceNetboxCircuitGroupUpdate,
		Delete: resourceNetboxCircuitGroupDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://netboxlabs.com/docs/netbox/models/circuits/circuitgroup/):

> Circuit groups can be used to organize circuits into administrative groups, for instance to indicate that multiple circuits form a redundant pair or are otherwise related.`,

		Schema: map[string]*schema.Schema{
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"tenant_id": {
				Type:     schema.TypeInt,
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

func resourceNetboxCircuitGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableCircuitGroup{}

	data.Comments = d.Get("comments").(string)
	data.Description = d.Get("description").(string)
	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
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

	params := circuits.NewCircuitsCircuitGroupsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsCircuitGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitGroupRead(d, m)
}

func resourceNetboxCircuitGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitGroupsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitGroupsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsCircuitGroupsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("comments", res.GetPayload().Comments)
	d.Set("description", res.GetPayload().Description)
	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	api.readTags(d, res.GetPayload().Tags)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxCircuitGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableCircuitGroup{}

	data.Comments = d.Get("comments").(string)
	data.Description = d.Get("description").(string)
	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := circuits.NewCircuitsCircuitGroupsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsCircuitGroupsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCircuitGroupRead(d, m)
}

func resourceNetboxCircuitGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitGroupsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsCircuitGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsCircuitGroupsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
