package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxCircuit() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitCreate,
		Read:   resourceNetboxCircuitRead,
		Update: resourceNetboxCircuitUpdate,
		Delete: resourceNetboxCircuitDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuits_1):

> A communications circuit represents a single physical link connecting exactly two endpoints, commonly referred to as its A and Z terminations. A circuit in NetBox may have zero, one, or two terminations defined. It is common to have only one termination defined when you don't necessarily care about the details of the provider side of the circuit, e.g. for Internet access circuits. Both terminations would likely be modeled for circuits which connect one customer site to another.
>
> Each circuit is associated with a provider and a user-defined type. For example, you might have Internet access circuits delivered to each site by one provider, and private MPLS circuits delivered by another. Each circuit must be assigned a circuit ID, each of which must be unique per provider.`,

		Schema: map[string]*schema.Schema{
			"provider_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"planned", "provisioning", "active", "offline", "deprovisioning", "decommissioning"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCircuitCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableCircuit{}

	cid := d.Get("cid").(string)
	data.Cid = &cid

	data.Status = d.Get("status").(string)

	providerIDValue, ok := d.GetOk("provider_id")
	if ok {
		data.Provider = int64ToPtr(int64(providerIDValue.(int)))
	}

	typeIDValue, ok := d.GetOk("type_id")
	if ok {
		data.Type = int64ToPtr(int64(typeIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	data.Tags = []*models.NestedTag{}

	params := circuits.NewCircuitsCircuitsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsCircuitsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitRead(d, m)
}

func resourceNetboxCircuitRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitsRead(params, nil)

	if err != nil {
		errorcode := err.(*circuits.CircuitsCircuitsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("cid", res.GetPayload().Cid)
	d.Set("status", res.GetPayload().Status.Value)

	if res.GetPayload().Provider != nil {
		d.Set("provider_id", res.GetPayload().Provider.ID)
	} else {
		d.Set("provider_id", nil)
	}

	if res.GetPayload().Type != nil {
		d.Set("type_id", res.GetPayload().Type.ID)
	} else {
		d.Set("type_id", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	return nil
}

func resourceNetboxCircuitUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableCircuit{}

	cid := d.Get("cid").(string)
	data.Cid = &cid

	data.Status = d.Get("status").(string)

	providerIDValue, ok := d.GetOk("provider_id")
	if ok {
		data.Provider = int64ToPtr(int64(providerIDValue.(int)))
	}

	typeIDValue, ok := d.GetOk("type_id")
	if ok {
		data.Type = int64ToPtr(int64(typeIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	data.Tags = []*models.NestedTag{}

	params := circuits.NewCircuitsCircuitsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Circuits.CircuitsCircuitsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCircuitRead(d, m)
}

func resourceNetboxCircuitDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsCircuitsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsCircuitsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
