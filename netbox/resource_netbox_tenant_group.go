package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxTenantGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTenantGroupCreate,
		Read:   resourceNetboxTenantGroupRead,
		Update: resourceNetboxTenantGroupUpdate,
		Delete: resourceNetboxTenantGroupDelete,

		Description: `:meta:subcategory:Tenancy:From the [official documentation](https://docs.netbox.dev/en/stable/features/tenancy/#tenant-groups):

> Tenants can be organized by custom groups. For instance, you might create one group called "Customers" and one called "Departments." The assignment of a tenant to a group is optional.
>
> Tenant groups may be nested recursively to achieve a multi-level hierarchy. For example, you might have a group called "Customers" containing subgroups of individual tenants grouped by product or account team.`,

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
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxTenantGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	parent_id := int64(d.Get("parent_id").(int))
	description := d.Get("description").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data := &models.WritableTenantGroup{}
	data.Name = &name
	data.Slug = &slug
	data.Description = description
	data.Tags = []*models.NestedTag{}

	if parent_id != 0 {
		data.Parent = &parent_id
	}

	params := tenancy.NewTenancyTenantGroupsCreateParams().WithData(data)

	res, err := api.Tenancy.TenancyTenantGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTenantGroupRead(d, m)
}

func resourceNetboxTenantGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := tenancy.NewTenancyTenantGroupsReadParams().WithID(id)

	res, err := api.Tenancy.TenancyTenantGroupsRead(params, nil)
	if err != nil {
		errorcode := err.(*tenancy.TenancyTenantGroupsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)
	d.Set("description", res.GetPayload().Description)
	if res.GetPayload().Parent != nil {
		d.Set("parent", res.GetPayload().Parent.ID)
	}
	return nil
}

func resourceNetboxTenantGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableTenantGroup{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	parent_id := int64(d.Get("parent_id").(int))

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Slug = &slug
	data.Name = &name
	data.Description = description
	data.Tags = []*models.NestedTag{}

	if parent_id != 0 {
		data.Parent = &parent_id
	}
	params := tenancy.NewTenancyTenantGroupsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Tenancy.TenancyTenantGroupsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxTenantGroupRead(d, m)
}

func resourceNetboxTenantGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyTenantGroupsDeleteParams().WithID(id)

	_, err := api.Tenancy.TenancyTenantGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*tenancy.TenancyTenantGroupsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
