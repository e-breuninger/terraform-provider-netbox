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

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxClusterTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to name if not given
	if !slugOk {
		slug = name
	} else {
		slug = slugValue.(string)
	}

	params := virtualization.NewVirtualizationClusterTypesCreateParams().WithData(
		&models.ClusterType{
			Name:         &name,
			Slug:         &slug,
			CustomFields: d.Get(customFieldsKey),
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
	d.Set(customFieldsKey, res.GetPayload().CustomFields)
	return nil
}

func resourceNetboxClusterTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.ClusterType{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to name if not given
	if !slugOk {
		slug = name
	} else {
		slug = slugValue.(string)
	}

	data.Slug = &slug
	data.Name = &name
	data.CustomFields = d.Get(customFieldsKey)

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
		return err
	}
	return nil
}
