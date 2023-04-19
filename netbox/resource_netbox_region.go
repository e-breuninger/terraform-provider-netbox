package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxRegion() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRegionCreate,
		Read:   resourceNetboxRegionRead,
		Update: resourceNetboxRegionUpdate,
		Delete: resourceNetboxRegionDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/sites-and-racks/#regions):

> Sites can be arranged geographically using regions. A region might represent a continent, country, city, campus, or other area depending on your use case. Regions can be nested recursively to construct a hierarchy. For example, you might define several country regions, and within each of those several state or city regions to which sites are assigned.
>
> Each region must have a name that is unique within its parent region, if any.`,

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
			"parent_region_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxRegionCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableRegion{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	parentRegionIDValue, ok := d.GetOk("parent_region_id")
	if ok {
		data.Parent = int64ToPtr(int64(parentRegionIDValue.(int)))
	}

	data.Tags = []*models.NestedTag{}

	params := dcim.NewDcimRegionsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimRegionsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRegionRead(d, m)
}

func resourceNetboxRegionRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRegionsReadParams().WithID(id)

	res, err := api.Dcim.DcimRegionsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimRegionsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)
	if res.GetPayload().Parent != nil {
		d.Set("parent_region_id", res.GetPayload().Parent.ID)
	} else {
		d.Set("parent_region_id", nil)
	}
	d.Set("description", res.GetPayload().Description)
	return nil
}

func resourceNetboxRegionUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableRegion{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	parentRegionIDValue, ok := d.GetOk("parent_region_id")
	if ok {
		data.Parent = int64ToPtr(int64(parentRegionIDValue.(int)))
	}

	data.Tags = []*models.NestedTag{}

	params := dcim.NewDcimRegionsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimRegionsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxRegionRead(d, m)
}

func resourceNetboxRegionDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRegionsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimRegionsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRegionsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
