package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxLocationCreate,
		Read:   resourceNetboxLocationRead,
		Update: resourceNetboxLocationUpdate,
		Delete: resourceNetboxLocationDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/core-functionality/sites-and-racks/#locations):

> Racks and devices can be grouped by location within a site. A location may represent a floor, room, cage, or similar organizational unit. Locations can be nested to form a hierarchy. For example, you may have floors within a site, and rooms within a floor.

Each location must have a name that is unique within its parent site and location, if any.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			customFieldsKey: customFieldSchema,
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxLocationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableLocation{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to name if not given
	if !slugOk {
		data.Slug = strToPtr(name)
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		data.Site = int64ToPtr(int64(siteIDValue.(int)))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	if resourceCustomFields, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = convertCustomFieldsFromTerraformToAPI(nil, resourceCustomFields.(*schema.Set).List())
	}

	params := dcim.NewDcimLocationsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimLocationsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxLocationRead(d, m)
}

func resourceNetboxLocationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimLocationsReadParams().WithID(id)

	res, err := api.Dcim.DcimLocationsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimLocationsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	location := res.GetPayload()

	d.Set("name", location.Name)
	d.Set("slug", location.Slug)

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if d.Get(customFieldsKey) != nil {
		customFields := updateCustomFieldsFromAPI(
			d.Get(customFieldsKey).(*schema.Set).List(),
			res.GetPayload().CustomFields,
		)
		d.Set(customFieldsKey, customFields)
	} else {
		d.Set(customFieldsKey, nil)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxLocationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableLocation{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to name if not given
	if !slugOk {
		data.Slug = strToPtr(name)
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		data.Tenant = int64ToPtr(int64(tenantIDValue.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	if d.HasChange(customFieldsKey) {
		stateCustomFields, resourceCustomFields := d.GetChange(customFieldsKey)
		data.CustomFields = convertCustomFieldsFromTerraformToAPI(stateCustomFields.(*schema.Set).List(), resourceCustomFields.(*schema.Set).List())
	}

	params := dcim.NewDcimLocationsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimLocationsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxLocationRead(d, m)
}

func resourceNetboxLocationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimLocationsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimLocationsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
