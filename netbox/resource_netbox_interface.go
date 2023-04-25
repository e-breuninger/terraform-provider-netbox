package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxInterface() *schema.Resource {
	validModes := []string{"access", "tagged", "tagged-all"}

	return &schema.Resource{
		CreateContext: resourceNetboxInterfaceCreate,
		ReadContext:   resourceNetboxInterfaceRead,
		UpdateContext: resourceNetboxInterfaceUpdate,
		DeleteContext: resourceNetboxInterfaceDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#interfaces):

> Virtual machine interfaces behave similarly to device interfaces, and can be assigned to VRFs, and may have IP addresses, VLANs, and services attached to them. However, given their virtual nature, they lack properties pertaining to physical attributes. For example, VM interfaces do not have a physical type and cannot have cables attached to them.`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_machine_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"mac_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsMACAddress,
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(validModes, false),
			},
			"mtu": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 65536),
			},
			"type": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "This attribute is not supported by netbox any longer. It will be removed in future versions of this provider.",
			},
			tagsKey: tagsSchema,
			"tagged_vlans": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"untagged_vlan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	mode := d.Get("mode").(string)
	tags, diagnostics := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	if diagnostics != nil {
		diags = append(diags, diagnostics...)
	}
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))

	data := models.WritableVMInterface{
		Name:           &name,
		Description:    description,
		Enabled:        enabled,
		Mode:           mode,
		Tags:           tags,
		TaggedVlans:    taggedVlans,
		VirtualMachine: &virtualMachineID,
	}
	if macAddress := d.Get("mac_address").(string); macAddress != "" {
		data.MacAddress = &macAddress
	}
	if mtu, ok := d.Get("mtu").(int); ok && mtu != 0 {
		data.Mtu = int64ToPtr(int64(mtu))
	}
	if untaggedVlan, ok := d.Get("untagged_vlan").(int); ok && untaggedVlan != 0 {
		data.UntaggedVlan = int64ToPtr(int64(untaggedVlan))
	}
	params := virtualization.NewVirtualizationInterfacesCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationInterfacesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := virtualization.NewVirtualizationInterfacesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationInterfacesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationInterfacesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	iface := res.GetPayload()

	d.Set("name", iface.Name)
	d.Set("description", iface.Description)
	d.Set("enabled", iface.Enabled)
	d.Set("mac_address", iface.MacAddress)
	d.Set("mtu", iface.Mtu)
	d.Set(tagsKey, getTagListFromNestedTagList(iface.Tags))
	d.Set("tagged_vlans", getIDsFromNestedVLAN(iface.TaggedVlans))
	d.Set("virtual_machine_id", iface.VirtualMachine.ID)

	if iface.Mode != nil {
		d.Set("mode", iface.Mode.Value)
	}
	if iface.UntaggedVlan != nil {
		d.Set("untagged_vlan", iface.UntaggedVlan.ID)
	}

	return diags
}

func resourceNetboxInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	mode := d.Get("mode").(string)
	tags, diagnostics := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	if diagnostics != nil {
		diags = append(diags, diagnostics...)
	}
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	virtualMachineID := int64(d.Get("virtual_machine_id").(int))

	data := models.WritableVMInterface{
		Name:           &name,
		Description:    description,
		Enabled:        enabled,
		Mode:           mode,
		Tags:           tags,
		TaggedVlans:    taggedVlans,
		VirtualMachine: &virtualMachineID,
	}

	if d.HasChange("mac_address") {
		macAddress := d.Get("mac_address").(string)
		data.MacAddress = &macAddress
	}
	if d.HasChange("mtu") {
		mtu := int64(d.Get("mtu").(int))
		data.Mtu = &mtu
	}
	if d.HasChange("untagged_vlan") {
		untaggedvlan := int64(d.Get("untagged_vlan").(int))
		data.UntaggedVlan = &untaggedvlan
	}

	params := virtualization.NewVirtualizationInterfacesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Virtualization.VirtualizationInterfacesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationInterfacesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationInterfacesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationInterfacesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}

func getIDsFromNestedVLAN(nestedvlans []*models.NestedVLAN) []int64 {
	var vlans []int64
	for _, vlan := range nestedvlans {
		vlans = append(vlans, vlan.ID)
	}
	return vlans
}
