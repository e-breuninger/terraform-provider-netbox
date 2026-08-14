package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxProviderAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxProviderAccountCreate,
		Read:   resourceNetboxProviderAccountRead,
		Update: resourceNetboxProviderAccountUpdate,
		Delete: resourceNetboxProviderAccountDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://netboxlabs.com/docs/netbox/models/circuits/provideraccount/):

> This model represents an account associated with a given provider with which a NetBox user has some kind of authorization to access, such as a login username, account number, or contract number.`,

		Schema: map[string]*schema.Schema{
			"account": {
				Type:     schema.TypeString,
				Required: true,
			},
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
				Optional: true,
			},
			"provider_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxProviderAccountCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableProviderAccount{}

	account := d.Get("account").(string)
	data.Account = &account
	data.Comments = d.Get("comments").(string)
	data.Description = d.Get("description").(string)
	data.Name = d.Get("name").(string)
	provider := int64(d.Get("provider_id").(int))
	data.Provider = &provider

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := circuits.NewCircuitsProviderAccountsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsProviderAccountsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxProviderAccountRead(d, m)
}

func resourceNetboxProviderAccountRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsProviderAccountsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsProviderAccountsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsProviderAccountsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("account", res.GetPayload().Account)
	d.Set("comments", res.GetPayload().Comments)
	d.Set("description", res.GetPayload().Description)
	d.Set("name", res.GetPayload().Name)

	if res.GetPayload().Provider != nil {
		d.Set("provider_id", res.GetPayload().Provider.ID)
	}

	api.readTags(d, res.GetPayload().Tags)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxProviderAccountUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableProviderAccount{}

	account := d.Get("account").(string)
	data.Account = &account
	data.Comments = d.Get("comments").(string)
	data.Description = d.Get("description").(string)
	data.Name = d.Get("name").(string)
	provider := int64(d.Get("provider_id").(int))
	data.Provider = &provider

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := circuits.NewCircuitsProviderAccountsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsProviderAccountsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxProviderAccountRead(d, m)
}

func resourceNetboxProviderAccountDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsProviderAccountsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsProviderAccountsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsProviderAccountsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
