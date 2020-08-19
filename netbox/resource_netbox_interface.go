package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceNetboxInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxInterfaceCreate,
		Read:   resourceNetboxInterfaceRead,
		Update: resourceNetboxInterfaceUpdate,
		Delete: resourceNetboxInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_machine_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "virtual",
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)

	name := d.Get("name").(string)
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))
	description := d.Get("description").(string)
	interfaceType := d.Get("type").(string)
	tags := getTagListFromResourceDataSet(d.Get("tags"))

	data := models.WritableVirtualMachineInterface{
		Name:           &name,
		Description:    description,
		VirtualMachine: &virtualMachineID,
		Type:           &interfaceType,
		Tags:           tags,
		TaggedVlans:    []int64{},
	}

	params := virtualization.NewVirtualizationInterfacesCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationInterfacesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxInterfaceUpdate(d, m)
}

func resourceNetboxInterfaceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := virtualization.NewVirtualizationInterfacesReadParams().WithID(id)
	res, err := api.Virtualization.VirtualizationInterfacesRead(params, nil)
	if err != nil {
		errorcode := err.(*runtime.APIError).Response.(runtime.ClientResponse).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band).
			// Just like the destroy callback, the Read function should gracefully handle this case.
			// https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("virtual_machine_id", res.GetPayload().VirtualMachine.ID)
	d.Set("description", res.GetPayload().Description)
	d.Set("tags", res.GetPayload().Tags)
	d.Set("type", res.GetPayload().Type.Value)
	return nil
}

func resourceNetboxInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))
	description := d.Get("description").(string)
	interfaceType := d.Get("type").(string)
	tags := getTagListFromResourceDataSet(d.Get("tags"))

	data := models.WritableVirtualMachineInterface{
		Name:           &name,
		Description:    description,
		VirtualMachine: &virtualMachineID,
		Type:           &interfaceType,
		Tags:           tags,
		TaggedVlans:    []int64{},
	}

	params := virtualization.NewVirtualizationInterfacesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Virtualization.VirtualizationInterfacesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxInterfaceRead(d, m)
}

func resourceNetboxInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationInterfacesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationInterfacesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
