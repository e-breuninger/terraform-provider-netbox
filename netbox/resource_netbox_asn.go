package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxAsn() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAsnCreate,
		Read:   resourceNetboxAsnRead,
		Update: resourceNetboxAsnUpdate,
		Delete: resourceNetboxAsnDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#asn):
> ASN is short for Autonomous System Number. This identifier is used in the BGP protocol to identify which "autonomous system" a particular prefix is originating and transiting through.
>
> The AS number model within NetBox allows you to model some of this real-world relationship.`,

		Schema: map[string]*schema.Schema{
			"asn": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"rir_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxAsnCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableASN{}

	asn := int64(d.Get("asn").(int))
	data.Asn = &asn

	rir := int64(d.Get("rir_id").(int))
	data.Rir = &rir

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamAsnsCreateParams().WithData(&data)

	res, err := api.Ipam.IpamAsnsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxAsnRead(d, m)
}

func resourceNetboxAsnRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamAsnsReadParams().WithID(id)

	res, err := api.Ipam.IpamAsnsRead(params, nil)

	if err != nil {
		errorcode := err.(*ipam.IpamAsnsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("asn", res.GetPayload().Asn)
	d.Set("rir_id", res.GetPayload().Rir)

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxAsnUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableASN{}

	asn := int64(d.Get("asn").(int))
	data.Asn = &asn

	rir := int64(d.Get("rir_id").(int))
	data.Rir = &rir

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamAsnsUpdateParams().WithID(id).WithData(&data)

	_, err := api.Ipam.IpamAsnsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxAsnRead(d, m)
}

func resourceNetboxAsnDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamAsnsDeleteParams().WithID(id)

	_, err := api.Ipam.IpamAsnsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamAsnsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
