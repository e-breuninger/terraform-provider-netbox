package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxVirtualDisks() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxVirtualDisksCreate,
		ReadContext:   resourceNetboxVirtualDisksRead,
		UpdateContext: resourceNetboxVirtualDisksUpdate,
		DeleteContext: resourceNetboxVirtualDisksDelete,
		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/models/virtualization/virtualdisk/):

		> A virtual disk is used to model discrete virtual hard disks assigned to virtual machines.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"size_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtual_machine_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			customFieldsKey: customFieldsSchema,
			tagsKey:         tagsSchema,
		},
		CustomizeDiff: customFieldsDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVirtualDisksCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	size := d.Get("size_gb").(int)
	virtualMachineID := d.Get("virtual_machine_id").(int)

	data := models.WritableVirtualDisk{
		Name:           &name,
		Size:           int64ToPtr(int64(size)),
		VirtualMachine: int64ToPtr(int64(virtualMachineID)),
	}

	descriptionValue, ok := d.GetOk("description")
	if ok {
		description := descriptionValue.(string)
		data.Description = description
	}

	data.CustomFields = computeCustomFieldsModel(d)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := virtualization.NewVirtualizationVirtualDisksCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationVirtualDisksCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVirtualDisksRead(ctx, d, m)
}

func resourceNetboxVirtualDisksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := virtualization.NewVirtualizationVirtualDisksReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationVirtualDisksRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationVirtualDisksReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	VirtualDisks := res.GetPayload()

	d.Set("name", VirtualDisks.Name)
	d.Set("description", VirtualDisks.Description)

	if VirtualDisks.Size != nil {
		d.Set("size_gb", *VirtualDisks.Size)
	}
	if VirtualDisks.VirtualMachine != nil {
		d.Set("virtual_machine_id", VirtualDisks.VirtualMachine.ID)
	}

	d.Set(customFieldsKey, res.GetPayload().CustomFields)

	d.Set(tagsKey, getTagListFromNestedTagList(VirtualDisks.Tags))
	return nil
}

func resourceNetboxVirtualDisksUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVirtualDisk{}

	name := d.Get("name").(string)
	size := int64(d.Get("size_gb").(int))
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))

	data.Name = &name
	data.Size = &size
	data.VirtualMachine = &virtualMachineID

	data.CustomFields = computeCustomFieldsModel(d)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	if d.HasChanges("description") {
		// check if description is set
		if descriptionValue, ok := d.GetOk("description"); ok {
			data.Description = descriptionValue.(string)
		} else {
			data.Description = " "
		}
	}

	params := virtualization.NewVirtualizationVirtualDisksUpdateParams().WithID(id).WithData(&data)

	_, err := api.Virtualization.VirtualizationVirtualDisksUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxVirtualDisksRead(ctx, d, m)
}

func resourceNetboxVirtualDisksDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationVirtualDisksDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationVirtualDisksDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationVirtualDisksDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	return nil
}
