package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPrefixCreate,
		Read:   resourceNetboxPrefixRead,
		Update: resourceNetboxPrefixUpdate,
		Delete: resourceNetboxPrefixDelete,

		Schema: map[string]*schema.Schema{
			"prefix": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated", "dhcp"}, false),
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_pool": {
				Type:     schema.TypeBool,
				Optional: true,
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
func resourceNetboxPrefixCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritablePrefix{}

	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	is_pool := d.Get("is_pool").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.Description = description
	data.IsPool = is_pool

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := ipam.NewIpamPrefixesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamPrefixesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxPrefixUpdate(d, m)
}

func resourceNetboxPrefixRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamPrefixesReadParams().WithID(id)

	res, err := api.Ipam.IpamPrefixesRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamPrefixesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("description", res.GetPayload().Description)
	d.Set("is_pool", res.GetPayload().IsPool)
	// FIGURE OUT NESTED VRF AND NESTED VLAN (from maybe interfaces?)

	return nil
}

func resourceNetboxPrefixUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritablePrefix{}
	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	is_pool := d.Get("is_pool").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.Description = description
	data.IsPool = is_pool

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := ipam.NewIpamPrefixesUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamPrefixesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxPrefixRead(d, m)
}

func resourceNetboxPrefixDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamPrefixesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamPrefixesDelete(params, nil)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
