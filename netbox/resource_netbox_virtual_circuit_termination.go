package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxVirtualCircuitTerminationRoleOptions = []string{"peer", "hub", "spoke"}

func resourceNetboxVirtualCircuitTermination() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVirtualCircuitTerminationCreate,
		Read:   resourceNetboxVirtualCircuitTerminationRead,
		Update: resourceNetboxVirtualCircuitTerminationUpdate,
		Delete: resourceNetboxVirtualCircuitTerminationDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://docs.netbox.dev/en/stable/models/circuits/virtualcircuittermination/):

> Each virtual circuit is comprised of two or more terminations, each of which connects the virtual circuit to a device or virtual machine interface. Unlike a physical circuit termination, a virtual circuit termination is always attached directly to an interface rather than to a site or provider network.`,

		Schema: map[string]*schema.Schema{
			"virtual_circuit_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"interface_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxVirtualCircuitTerminationRoleOptions, false),
				Description:  buildValidValueDescription(resourceNetboxVirtualCircuitTerminationRoleOptions),
			},
			"description": {
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

func resourceNetboxVirtualCircuitTerminationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableVirtualCircuitTermination{}

	data.Description = d.Get("description").(string)
	data.Role = d.Get("role").(string)

	virtualCircuitIDValue, ok := d.GetOk("virtual_circuit_id")
	if ok {
		data.VirtualCircuit = int64ToPtr(int64(virtualCircuitIDValue.(int)))
	}

	interfaceIDValue, ok := d.GetOk("interface_id")
	if ok {
		data.Interface = int64ToPtr(int64(interfaceIDValue.(int)))
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

	params := circuits.NewCircuitsVirtualCircuitTerminationsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsVirtualCircuitTerminationsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualCircuitTerminationRead(d, m)
}

func resourceNetboxVirtualCircuitTerminationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsVirtualCircuitTerminationsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsVirtualCircuitTerminationsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsVirtualCircuitTerminationsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	term := res.GetPayload()

	d.Set("description", term.Description)

	if term.Role != nil {
		d.Set("role", term.Role.Value)
	} else {
		d.Set("role", nil)
	}

	if term.VirtualCircuit != nil {
		d.Set("virtual_circuit_id", term.VirtualCircuit.ID)
	} else {
		d.Set("virtual_circuit_id", nil)
	}

	if term.Interface != nil {
		d.Set("interface_id", term.Interface.ID)
	} else {
		d.Set("interface_id", nil)
	}

	api.readTags(d, term.Tags)

	cf := getCustomFields(term.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxVirtualCircuitTerminationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualCircuitTermination{}

	data.Description = d.Get("description").(string)
	data.Role = d.Get("role").(string)

	virtualCircuitIDValue, ok := d.GetOk("virtual_circuit_id")
	if ok {
		data.VirtualCircuit = int64ToPtr(int64(virtualCircuitIDValue.(int)))
	}

	interfaceIDValue, ok := d.GetOk("interface_id")
	if ok {
		data.Interface = int64ToPtr(int64(interfaceIDValue.(int)))
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

	params := circuits.NewCircuitsVirtualCircuitTerminationsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsVirtualCircuitTerminationsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVirtualCircuitTerminationRead(d, m)
}

func resourceNetboxVirtualCircuitTerminationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsVirtualCircuitTerminationsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsVirtualCircuitTerminationsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsVirtualCircuitTerminationsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
