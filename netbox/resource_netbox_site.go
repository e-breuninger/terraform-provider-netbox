package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxSiteCreate,
		Read:   resourceNetboxSiteRead,
		Update: resourceNetboxSiteUpdate,
		Delete: resourceNetboxSiteDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"planned", "staging", "active", "decommissioning", "retired"}, false),
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"facility": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},
			"longitude": &schema.Schema{
				Type:         schema.TypeFloat,
				Optional:     true,
			},
			"latitude": &schema.Schema{
				Type:         schema.TypeFloat,
				Optional:     true,
			},
			"region_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeInt,
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
			"timezone": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
			},
			"asn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxSiteCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableSite{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to name if not given
	if !slugOk {
		data.Slug = strToPtr(name)
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Status = d.Get("status").(string)

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	if facility, ok := d.GetOk("facility"); ok {
		data.Facility = facility.(string)
	}

	latitudeValue, ok := d.GetOk("latitude")
	if ok {
		data.Latitude = float64ToPtr(float64(latitudeValue.(float64)))
	}

	longitudeValue, ok := d.GetOk("longitude")
	if ok {
		data.Longitude = float64ToPtr(float64(longitudeValue.(float64)))
	}

	regionIDValue, ok := d.GetOk("region_id")
	if ok {
		data.Region = int64ToPtr(int64(regionIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	if timezone, ok := d.GetOk("timezone"); ok {
		data.TimeZone = timezone.(string)
	}

	asnValue, ok := d.GetOk("asn")
	if ok {
		data.Asn = int64ToPtr(int64(asnValue.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimSitesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimSitesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxSiteRead(d, m)
}

func resourceNetboxSiteRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimSitesReadParams().WithID(id)

	res, err := api.Dcim.DcimSitesRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimSitesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)
	d.Set("status", res.GetPayload().Status.Value)
	d.Set("description", res.GetPayload().Description)
	d.Set("facility", res.GetPayload().Facility)
	d.Set("longitude", res.GetPayload().Longitude)
	d.Set("latitude", res.GetPayload().Latitude)
	d.Set("timezone", res.GetPayload().TimeZone)
	d.Set("asn", res.GetPayload().Asn)

	if res.GetPayload().Region != nil {
		d.Set("region_id", res.GetPayload().Region.ID)
	} else {
		d.Set("region_id", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxSiteUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableSite{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to name if not given
	if !slugOk {
		data.Slug = strToPtr(name)
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Status = d.Get("status").(string)

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	if facility, ok := d.GetOk("facility"); ok {
		data.Facility = facility.(string)
	}

	latitudeValue, ok := d.GetOk("latitude")
	if ok {
		data.Latitude = float64ToPtr(float64(latitudeValue.(float64)))
	}

	longitudeValue, ok := d.GetOk("longitude")
	if ok {
		data.Longitude = float64ToPtr(float64(longitudeValue.(float64)))
	}

	regionIDValue, ok := d.GetOk("region_id")
	if ok {
		data.Region = int64ToPtr(int64(regionIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	if timezone, ok := d.GetOk("timezone"); ok {
		data.TimeZone = timezone.(string)
	}

	asnValue, ok := d.GetOk("asn")
	if ok {
		data.Asn = int64ToPtr(int64(asnValue.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimSitesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimSitesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxSiteRead(d, m)
}

func resourceNetboxSiteDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimSitesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimSitesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
