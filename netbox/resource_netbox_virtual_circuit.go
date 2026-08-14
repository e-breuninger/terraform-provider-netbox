package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxVirtualCircuitStatusOptions = []string{"planned", "provisioning", "active", "offline", "deprovisioning", "decommissioned"}

func resourceNetboxVirtualCircuit() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVirtualCircuitCreate,
		Read:   resourceNetboxVirtualCircuitRead,
		Update: resourceNetboxVirtualCircuitUpdate,
		Delete: resourceNetboxVirtualCircuitDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/models/circuits/virtualcircuit/):

> A virtual circuit represents a discrete germane layer two channel or connection multiplexed onto a physical circuit, provider network, or set of physical circuits, e.g. an EVPL/EVC. Virtual circuits are commonly used to represent provider connectivity which is logically separated from the underlying physical infrastructure, such as MPLS or VXLAN tunnels.`,

		Schema: map[string]*schema.Schema{
			"cid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_network_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"provider_account_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "active",
				ValidateFunc: validation.StringInSlice(resourceNetboxVirtualCircuitStatusOptions, false),
				Description:  buildValidValueDescription(resourceNetboxVirtualCircuitStatusOptions),
			},
			"tenant_id": {
				Type:     schema.TypeInt,
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

func resourceNetboxVirtualCircuitCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableVirtualCircuit{}

	cid := d.Get("cid").(string)
	data.Cid = &cid
	data.Status = d.Get("status").(string)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	providerNetworkIDValue, ok := d.GetOk("provider_network_id")
	if ok {
		data.ProviderNetwork = int64ToPtr(int64(providerNetworkIDValue.(int)))
	}

	data.ProviderAccount = getOptionalInt(d, "provider_account_id")

	typeIDValue, ok := d.GetOk("type_id")
	if ok {
		data.Type = int64ToPtr(int64(typeIDValue.(int)))
	}

	data.Tenant = getOptionalInt(d, "tenant_id")

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := circuits.NewCircuitsVirtualCircuitsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsVirtualCircuitsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualCircuitRead(d, m)
}

func resourceNetboxVirtualCircuitRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsVirtualCircuitsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsVirtualCircuitsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsVirtualCircuitsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	virtualCircuit := res.GetPayload()

	d.Set("cid", virtualCircuit.Cid)
	d.Set("description", virtualCircuit.Description)
	d.Set("comments", virtualCircuit.Comments)

	if virtualCircuit.Status != nil {
		d.Set("status", virtualCircuit.Status.Value)
	} else {
		d.Set("status", nil)
	}

	if virtualCircuit.ProviderNetwork != nil {
		d.Set("provider_network_id", virtualCircuit.ProviderNetwork.ID)
	} else {
		d.Set("provider_network_id", nil)
	}

	if virtualCircuit.ProviderAccount != nil {
		d.Set("provider_account_id", virtualCircuit.ProviderAccount.ID)
	} else {
		d.Set("provider_account_id", nil)
	}

	if virtualCircuit.Type != nil {
		d.Set("type_id", virtualCircuit.Type.ID)
	} else {
		d.Set("type_id", nil)
	}

	if virtualCircuit.Tenant != nil {
		d.Set("tenant_id", virtualCircuit.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	api.readTags(d, virtualCircuit.Tags)

	cf := getCustomFields(virtualCircuit.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxVirtualCircuitUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualCircuit{}

	cid := d.Get("cid").(string)
	data.Cid = &cid
	data.Status = d.Get("status").(string)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	providerNetworkIDValue, ok := d.GetOk("provider_network_id")
	if ok {
		data.ProviderNetwork = int64ToPtr(int64(providerNetworkIDValue.(int)))
	}

	data.ProviderAccount = getOptionalInt(d, "provider_account_id")

	typeIDValue, ok := d.GetOk("type_id")
	if ok {
		data.Type = int64ToPtr(int64(typeIDValue.(int)))
	}

	data.Tenant = getOptionalInt(d, "tenant_id")

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := circuits.NewCircuitsVirtualCircuitsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsVirtualCircuitsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVirtualCircuitRead(d, m)
}

func resourceNetboxVirtualCircuitDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsVirtualCircuitsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsVirtualCircuitsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsVirtualCircuitsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
