package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strconv"
)

func resourceNetboxIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIPAddressCreate,
		Read:   resourceNetboxIPAddressRead,
		Update: resourceNetboxIPAddressUpdate,
		Delete: resourceNetboxIPAddressDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"interface_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated", "dhcp"}, false),
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

func resourceNetboxIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)

	data := models.WritableIPAddress{}
	ipAddress := d.Get("ip_address").(string)
	data.Address = &ipAddress
	data.Status = d.Get("status").(string)

	if interfaceID, ok := d.GetOk("interface_id"); ok {
		tmpInterfaceID := int64(interfaceID.(int))
		data.Interface = &tmpInterfaceID
	}

	tagsValue := d.Get("tags").(*schema.Set).List()
	tags := []string{}
	for _, tag := range tagsValue {
		tags = append(tags, tag.(string))
	}
	data.Tags = tags

	params := ipam.NewIpamIPAddressesCreateParams().WithData(&data)

	res, err := api.Ipam.IpamIPAddressesCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIPAddressRead(d, m)
}

func resourceNetboxIPAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamIPAddressesReadParams().WithID(id)

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		errorcode := err.(*runtime.APIError).Response.(runtime.ClientResponse).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	if res.GetPayload().Interface != nil {
		d.Set("interface_id", res.GetPayload().Interface.ID)
	}

	d.Set("ip_address", res.GetPayload().Address)
	d.Set("status", res.GetPayload().Status.Value)
	d.Set("tags", res.GetPayload().Tags)
	return nil
}

func resourceNetboxIPAddressUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPAddress{}

	ipAddress := d.Get("ip_address").(string)
	status := d.Get("status").(string)

	data.Status = status
	data.Address = &ipAddress

	if interfaceID, ok := d.GetOk("interface_id"); ok {
		tmpInterfaceID := int64(interfaceID.(int))
		data.Interface = &tmpInterfaceID
	}

	tagsValue := d.Get("tags").(*schema.Set).List()
	tags := []string{}
	for _, tag := range tagsValue {
		tags = append(tags, tag.(string))
	}

	data.Tags = tags

	params := ipam.NewIpamIPAddressesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Ipam.IpamIPAddressesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIPAddressRead(d, m)
}

func resourceNetboxIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBox)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamIPAddressesDeleteParams().WithID(id)

	_, err := api.Ipam.IpamIPAddressesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
