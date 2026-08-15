package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxVirtualMachineType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVirtualMachineTypeCreate,
		Read:   resourceNetboxVirtualMachineTypeRead,
		Update: resourceNetboxVirtualMachineTypeUpdate,
		Delete: resourceNetboxVirtualMachineTypeDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#virtual-machines):

> A virtual machine type represents a technology or template by which a virtual machine is instantiated. For example, you might create a virtual machine type named "t3.medium" or "n2-standard-4" to model a particular cloud instance size.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceNetboxVirtualMachineTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableVirtualMachineType{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := virtualization.NewVirtualizationVirtualMachineTypesCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationVirtualMachineTypesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualMachineTypeRead(d, m)
}

func resourceNetboxVirtualMachineTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationVirtualMachineTypesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationVirtualMachineTypesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationVirtualMachineTypesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	vmType := res.GetPayload()
	d.Set("name", vmType.Name)
	d.Set("slug", vmType.Slug)
	d.Set("description", vmType.Description)
	d.Set("comments", vmType.Comments)

	api.readTags(d, vmType.Tags)

	cf := getCustomFields(vmType.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxVirtualMachineTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualMachineType{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := virtualization.NewVirtualizationVirtualMachineTypesPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Virtualization.VirtualizationVirtualMachineTypesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVirtualMachineTypeRead(d, m)
}

func resourceNetboxVirtualMachineTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationVirtualMachineTypesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationVirtualMachineTypesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationVirtualMachineTypesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
