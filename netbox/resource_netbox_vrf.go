package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxVrf() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVrfCreate,
		Read:   resourceNetboxVrfRead,
		Update: resourceNetboxVrfUpdate,
		Delete: resourceNetboxVrfDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceNetboxVrfCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := ipam.NewIpamVrfsCreateParams().WithData(
		&models.WritableVRF{
			Name: &name,
			Tags: tags,
		},
	)

	res, err := api.Ipam.IpamVrfsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVrfRead(d, m)
}

func resourceNetboxVrfRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVrfsReadParams().WithID(id)

	res, err := api.Ipam.IpamVrfsRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamVrfsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	return nil
}

func resourceNetboxVrfUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVRF{}

	name := d.Get("name").(string)

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	data.Name = &name
	data.Tags = tags

	params := ipam.NewIpamVrfsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Ipam.IpamVrfsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVrfRead(d, m)
}

func resourceNetboxVrfDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVrfsDeleteParams().WithID(id)

	_, err := api.Ipam.IpamVrfsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
