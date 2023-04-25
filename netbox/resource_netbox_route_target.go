package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxRouteTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRouteTargetCreate,
		Read:   resourceNetboxRouteTargetRead,
		Update: resourceNetboxRouteTargetUpdate,
		Delete: resourceNetboxRouteTargetDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/routetarget/):

> A route target is a particular type of extended BGP community used to control the redistribution of routes among VRF tables in a network. Route targets can be assigned to individual VRFs in NetBox as import or export targets (or both) to model this exchange in an L3VPN. Each route target must be given a unique name, which should be in a format prescribed by RFC 4364, similar to a VR route distinguisher.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 21),
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxRouteTargetCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableRouteTarget{}

	name := d.Get("name").(string)

	data.Name = &name
	data.Tags = []*models.NestedTag{}

	params := ipam.NewIpamRouteTargetsCreateParams().WithData(&data)
	res, err := api.Ipam.IpamRouteTargetsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRouteTargetUpdate(d, m)
}

func resourceNetboxRouteTargetRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamRouteTargetsReadParams().WithID(id)

	res, err := api.Ipam.IpamRouteTargetsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamRouteTargetsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if res.GetPayload().Name != nil {
		d.Set("name", res.GetPayload().Name)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	}

	if res.GetPayload().Description != "" {
		d.Set("description", res.GetPayload().Description)
	}

	if res.GetPayload().Tags != nil {
		d.Set(tagsKey, res.GetPayload().Tags)
	}

	return nil
}

func resourceNetboxRouteTargetUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableRouteTarget{}

	name := d.Get("name").(string)
	tenant_id := int64(d.Get("tenant_id").(int))
	description := d.Get("description").(string)

	data.Name = &name
	data.Description = description
	data.Tenant = &tenant_id
	data.Tags = []*models.NestedTag{}

	params := ipam.NewIpamRouteTargetsUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamRouteTargetsUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxRouteTargetRead(d, m)
}

func resourceNetboxRouteTargetDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamRouteTargetsDeleteParams().WithID(id)
	_, err := api.Ipam.IpamRouteTargetsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamRouteTargetsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.SetId("")
	return nil
}
