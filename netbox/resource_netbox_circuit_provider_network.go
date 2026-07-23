package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxCircuitProviderNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitProviderNetworkCreate,
		Read:   resourceNetboxCircuitProviderNetworkRead,
		Update: resourceNetboxCircuitProviderNetworkUpdate,
		Delete: resourceNetboxCircuitProviderNetworkDelete,

		Description: `:meta:subcategory:Circuits:From the [offical documentation](https://netboxlabs.com/docs/netbox/models/circuits/providernetwork/):

> This model can be used to represent the boundary of a provider network, the details of which are unknown or unimportant to the NetBox user. For example, it might represent a provider's regional MPLS network to which multiple circuits provide connectivity.`,

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
			"provider_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"service_id": {
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

func resourceNetboxCircuitProviderNetworkCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableProviderNetwork{}

	data.Comments = d.Get("comments").(string)
	data.Description = d.Get("description").(string)
	name := d.Get("name").(string)
	data.Name = &name
	provider := int64(d.Get("provider_id").(int))
	data.Provider = &provider
	data.ServiceID = d.Get("service_id").(string)

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk("custom_fields")
	if ok {
		data.CustomFields = cf
	}

	params := circuits.NewCircuitsProviderNetworksCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsProviderNetworksCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitProviderNetworkRead(d, m)
}

func resourceNetboxCircuitProviderNetworkRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsProviderNetworksReadParams().WithID(id)

	res, err := api.Circuits.CircuitsProviderNetworksRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsProviderNetworksReadDefault); ok {
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
	d.Set("provider_id", res.GetPayload().Provider.ID)
	d.Set("service_id", res.GetPayload().ServiceID)

	api.readTags(d, res.GetPayload().Tags)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxCircuitProviderNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableProviderNetwork{}

	data.Comments = d.Get("comments").(string)
	data.Description = d.Get("description").(string)
	name := d.Get("name").(string)
	data.Name = &name
	provider := int64(d.Get("provider_id").(int))
	data.Provider = &provider
	data.ServiceID = d.Get("service_id").(string)

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := circuits.NewCircuitsProviderNetworksPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsProviderNetworksPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCircuitProviderNetworkRead(d, m)
}

func resourceNetboxCircuitProviderNetworkDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsProviderNetworksDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsProviderNetworksDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
