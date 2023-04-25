package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxSiteGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxSiteGroupCreate,
		Read:   resourceNetboxSiteGroupRead,
		Update: resourceNetboxSiteGroupUpdate,
		Delete: resourceNetboxSiteGroupDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/facilities/#site-groups):

> Like regions, site groups can be arranged in a recursive hierarchy for grouping sites. However, whereas regions are intended for geographic organization, site groups may be used for functional grouping. For example, you might classify sites as corporate, branch, or customer sites in addition to where they are physically located.
>
> The use of both regions and site groups affords to independent but complementary dimensions across which sites can be organized.`,

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

func resourceNetboxSiteGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	data := &models.WritableSiteGroup{}
	data.Name = &name
	data.Slug = &slug
	data.Description = description
	data.Tags = []*models.NestedTag{}

	if parent_id != 0 {
		data.Parent = &parent_id
	}

	params := dcim.NewDcimSiteGroupsCreateParams().WithData(data)

	res, err := api.Dcim.DcimSiteGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxSiteGroupRead(d, m)
}

func resourceNetboxSiteGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimSiteGroupsReadParams().WithID(id)

	res, err := api.Dcim.DcimSiteGroupsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimSiteGroupsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	siteGroup := res.GetPayload()
	d.Set("name", siteGroup.Name)
	d.Set("slug", siteGroup.Slug)
	d.Set("description", siteGroup.Description)
	if siteGroup.Parent != nil {
		d.Set("parent_id", siteGroup.Parent.ID)
	}
	return nil
}

func resourceNetboxSiteGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableSiteGroup{}

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
	params := dcim.NewDcimSiteGroupsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimSiteGroupsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxSiteGroupRead(d, m)
}

func resourceNetboxSiteGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimSiteGroupsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimSiteGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimSiteGroupsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
