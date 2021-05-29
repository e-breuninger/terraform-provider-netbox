package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxService() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxServiceCreate,
		Read:   resourceNetboxServiceRead,
		Update: resourceNetboxServiceUpdate,
		Delete: resourceNetboxServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_machine_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}
func resourceNetboxServiceCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableService{}

	dataName := d.Get("name").(string)
	data.Name = &dataName

	dataProtocol := d.Get("protocol").(string)
	data.Protocol = &dataProtocol

	dataPort := int64(d.Get("port").(int))
	data.Port = &dataPort

	dataVirtualMachineID := int64(d.Get("virtual_machine_id").(int))
	data.VirtualMachine = &dataVirtualMachineID

	data.Tags = []*models.NestedTag{}
	data.Ipaddresses = []int64{}

	params := ipam.NewIpamServicesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamServicesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxServiceUpdate(d, m)
}

func resourceNetboxServiceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamServicesReadParams().WithID(id)

	res, err := api.Ipam.IpamServicesRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamServicesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}
	d.Set("name", res.GetPayload().Name)
	d.Set("protocol", res.GetPayload().Protocol.Value)
	d.Set("port", res.GetPayload().Port)
	d.Set("virtual_machine_id", res.GetPayload().VirtualMachine.ID)

	return nil
}

func resourceNetboxServiceUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableService{}

	dataName := d.Get("name").(string)
	data.Name = &dataName

	dataProtocol := d.Get("protocol").(string)
	data.Protocol = &dataProtocol

	dataPort := int64(d.Get("port").(int))
	data.Port = &dataPort

	data.Tags = []*models.NestedTag{}
	data.Ipaddresses = []int64{}

	dataVirtualMachineID := int64(d.Get("virtual_machine_id").(int))
	data.VirtualMachine = &dataVirtualMachineID

	params := ipam.NewIpamServicesUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamServicesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxServiceRead(d, m)
}

func resourceNetboxServiceDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamServicesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamServicesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
