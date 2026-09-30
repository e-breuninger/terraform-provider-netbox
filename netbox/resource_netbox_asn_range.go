package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxAsnRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAsnRangeCreate,
		Read:   resourceNetboxAsnRangeRead,
		Update: resourceNetboxAsnRangeUpdate,
		Delete: resourceNetboxAsnRangeDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/asnrange/):

> Autonomous System Numbers (ASNs) are number blocks used in the BGP protocol to identify networks. NetBox models ASN ranges as blocks of ASNs which can be later utilized within organization to easily assign new ASNs.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"rir_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"start": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 4294967295),
			},
			"end": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 4294967295),
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
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxAsnRangeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.WritableASNRange{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Name = &name
	data.Slug = &slug

	rirID := int64(d.Get("rir_id").(int))
	data.Rir = &rirID

	start := int64(d.Get("start").(int))
	data.Start = &start

	end := int64(d.Get("end").(int))
	data.End = &end

	data.Tenant = getOptionalInt(d, "tenant_id")
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamAsnRangesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamAsnRangesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxAsnRangeRead(d, m)
}

func resourceNetboxAsnRangeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamAsnRangesReadParams().WithID(id)

	res, err := api.Ipam.IpamAsnRangesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamAsnRangesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	asnRange := res.GetPayload()

	d.Set("name", asnRange.Name)
	d.Set("slug", asnRange.Slug)
	d.Set("start", asnRange.Start)
	d.Set("end", asnRange.End)
	d.Set("description", asnRange.Description)
	d.Set("comments", asnRange.Comments)

	if asnRange.Rir != nil {
		d.Set("rir_id", asnRange.Rir.ID)
	}

	if asnRange.Tenant != nil {
		d.Set("tenant_id", asnRange.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	api.readTags(d, asnRange.Tags)

	cf := getCustomFields(asnRange.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxAsnRangeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableASNRange{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Name = &name
	data.Slug = &slug

	rirID := int64(d.Get("rir_id").(int))
	data.Rir = &rirID

	start := int64(d.Get("start").(int))
	data.Start = &start

	end := int64(d.Get("end").(int))
	data.End = &end

	data.Tenant = getOptionalInt(d, "tenant_id")
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamAsnRangesUpdateParams().WithID(id).WithData(&data)
	_, err = api.Ipam.IpamAsnRangesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxAsnRangeRead(d, m)
}

func resourceNetboxAsnRangeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamAsnRangesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamAsnRangesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamAsnRangesDeleteDefault); ok {
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
