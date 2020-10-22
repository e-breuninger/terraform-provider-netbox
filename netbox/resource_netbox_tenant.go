package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strconv"
)

func resourceNetboxTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTenantCreate,
		Read:   resourceNetboxTenantRead,
		Update: resourceNetboxTenantUpdate,
		Delete: resourceNetboxTenantDelete,

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

func resourceNetboxTenantCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to name attribute if not given
	if !slugOk {
		slug = name
	} else {
		slug = slugValue.(string)
	}

//	tagsValue := d.Get("tags").(*schema.Set).List()
//	tags := []string{}
//	for _, tag := range tagsValue {
//		tags = append(tags, tag.(string))
//	}
        tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := tenancy.NewTenancyTenantsCreateParams().WithData(
		&models.WritableTenant{
			Name: &name,
			Slug: &slug,
			Tags: tags,
		},
	)

	res, err := api.Tenancy.TenancyTenantsCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTenantRead(d, m)
}

func resourceNetboxTenantRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyTenantsReadParams().WithID(id)

	res, err := api.Tenancy.TenancyTenantsRead(params, nil)
	if err != nil {
		errorcode := err.(*runtime.APIError).Response.(runtime.ClientResponse).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)
	return nil
}

func resourceNetboxTenantUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableTenant{}

	name := d.Get("name").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to name if not given
	if !slugOk {
		slug = name
	} else {
		slug = slugValue.(string)
	}

//	tagsValue := d.Get("tags").(*schema.Set).List()
//	tags := []string{}
//	for _, tag := range tagsValue {
//		tags = append(tags, tag.(string))
//	}
        tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	data.Slug = &slug
	data.Name = &name
	data.Tags = tags

	params := tenancy.NewTenancyTenantsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Tenancy.TenancyTenantsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxTenantRead(d, m)
}

func resourceNetboxTenantDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyTenantsDeleteParams().WithID(id)

	_, err := api.Tenancy.TenancyTenantsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
