package netbox

import (
	"regexp"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^([A-Z0-9]{2}:){5}[A-Z0-9]{2}$"),
					"Must be like AA:AA:AA:AA:AA"),
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.",
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
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))
	description := d.Get("description").(string)
	macAddress := d.Get("mac_address").(string)
	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	data := models.WritableVMInterface{
		Name:           &name,
		Description:    description,
		VirtualMachine: &virtualMachineID,
		Tags:           tags,
		TaggedVlans:    []int64{},
	}
	if macAddress != "" {
		data.MacAddress = &macAddress
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
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := virtualization.NewVirtualizationInterfacesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationInterfacesRead(params, nil)
	if err != nil {
		errorcode := err.(*virtualization.VirtualizationInterfacesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("virtual_machine_id", res.GetPayload().VirtualMachine.ID)
	d.Set("description", res.GetPayload().Description)
	d.Set("mac_address", res.GetPayload().MacAddress)
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))
	description := d.Get("description").(string)
	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	data := models.WritableVMInterface{
		Name:           &name,
		Description:    description,
		VirtualMachine: &virtualMachineID,
		Tags:           tags,
		TaggedVlans:    []int64{},
	}

	params := virtualization.NewVirtualizationInterfacesPartialUpdateParams().WithID(id).WithData(&data)
	if d.HasChange("mac_address") {
		macAddress := d.Get("mac_address").(string)
		data.MacAddress = &macAddress
	}
	_, err := api.Virtualization.VirtualizationInterfacesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxInterfaceRead(d, m)
}

func resourceNetboxInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationInterfacesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationInterfacesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
