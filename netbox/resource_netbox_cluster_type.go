package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxClusterType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxClusterTypeCreate,
		Read:   resourceNetboxClusterTypeRead,
		Update: resourceNetboxClusterTypeUpdate,
		Delete: resourceNetboxClusterTypeDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#cluster-types):

> A cluster type represents a technology or mechanism by which a cluster is formed. For example, you might create a cluster type named "VMware vSphere" for a locally hosted cluster or "DigitalOcean NYC3" for one hosted by a cloud provider.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxClusterTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	params := virtualization.NewVirtualizationClusterTypesCreateParams().WithData(
		&models.ClusterType{
			Name: &name,
			Slug: &slug,
			Tags: []*models.NestedTag{},
		},
	)

	res, err := api.Virtualization.VirtualizationClusterTypesCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxClusterTypeRead(d, m)
}

func resourceNetboxClusterTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationClusterTypesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationClusterTypesRead(params, nil)
	if err != nil {
		errorcode := err.(*virtualization.VirtualizationClusterTypesReadDefault).Code()
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

func resourceNetboxClusterTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.ClusterType{}

	name := d.Get("name").(string)
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
	data.Tags = []*models.NestedTag{}

	params := virtualization.NewVirtualizationClusterTypesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Virtualization.VirtualizationClusterTypesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxClusterTypeRead(d, m)
}

func resourceNetboxClusterTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationClusterTypesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationClusterTypesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationClusterTypesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
