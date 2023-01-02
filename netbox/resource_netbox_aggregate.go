package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxAggregate() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAggregateCreate,
		Read:   resourceNetboxAggregateRead,
		Update: resourceNetboxAggregateUpdate,
		Delete: resourceNetboxAggregateDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#aggregates):

> NetBox allows us to specify the portions of IP space that are interesting to us by defining aggregates. Typically, an aggregate will correspond to either an allocation of public (globally routable) IP space granted by a regional authority, or a private (internally-routable) designation.`,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rir_id": {
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
func resourceNetboxAggregateCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableAggregate{}

	prefix := d.Get("prefix").(string)
	description := d.Get("description").(string)

	data.Prefix = &prefix
	data.Description = description

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if rirID, ok := d.GetOk("rir_id"); ok {
		data.Rir = int64ToPtr(int64(rirID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamAggregatesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamAggregatesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxAggregateRead(d, m)
}

func resourceNetboxAggregateRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamAggregatesReadParams().WithID(id)

	res, err := api.Ipam.IpamAggregatesRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamAggregatesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("description", res.GetPayload().Description)
	if res.GetPayload().Prefix != nil {
		d.Set("prefix", res.GetPayload().Prefix)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if res.GetPayload().Rir != nil {
		d.Set("rir_id", res.GetPayload().Rir.ID)
	} else {
		d.Set("rir_id", nil)
	}

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxAggregateUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableAggregate{}
	prefix := d.Get("prefix").(string)
	description := d.Get("description").(string)

	data.Prefix = &prefix
	data.Description = description

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if rirID, ok := d.GetOk("rir_id"); ok {
		data.Rir = int64ToPtr(int64(rirID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamAggregatesUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamAggregatesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxAggregateRead(d, m)
}

func resourceNetboxAggregateDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamAggregatesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamAggregatesDelete(params, nil)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
