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

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#virtual-routing-and-forwarding-vrf):

> A VRF object in NetBox represents a virtual routing and forwarding (VRF) domain. Each VRF is essentially a separate routing table. VRFs are commonly used to isolate customers or organizations from one another within a network, or to route overlapping address space (e.g. multiple instances of the 10.0.0.0/8 space). Each VRF may be assigned to a specific tenant to aid in organizing the available IP space by customer or internal user.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVrfCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableVRF{}

	name := d.Get("name").(string)
	tenant_id := int64(d.Get("tenant_id").(int))

	data.Name = &name
	if tenant_id != 0 {
		data.Tenant = &tenant_id
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	data.ExportTargets = []int64{}
	data.ImportTargets = []int64{}

	params := ipam.NewIpamVrfsCreateParams().WithData(&data)

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
	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}
	return nil
}

func resourceNetboxVrfUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVRF{}

	name := d.Get("name").(string)

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	data.Name = &name
	data.Tags = tags
	data.ExportTargets = []int64{}
	data.ImportTargets = []int64{}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}
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
