package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxDeviceInterface() *schema.Resource {
	validModes := []string{"access", "tagged", "tagged-all"}

	return &schema.Resource{
		CreateContext: resourceNetboxDeviceInterfaceCreate,
		ReadContext:   resourceNetboxDeviceInterfaceRead,
		UpdateContext: resourceNetboxDeviceInterfaceUpdate,
		DeleteContext: resourceNetboxDeviceInterfaceDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/device/#interface):

> Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_id": {
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
				ForceNew:     true,
			},
			"mgmtonly": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Type:     schema.TypeString,
				Required: true,
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

func resourceNetboxDeviceInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	interface_type := d.Get("type").(string)
	enabled := d.Get("enabled").(bool)
	mgmtonly := d.Get("mgmtonly").(bool)
	mode := d.Get("mode").(string)
	tags, diagnostics := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	if diagnostics != nil {
		diags = append(diags, diagnostics...)
	}
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	deviceID := int64(d.Get("device_id").(int))

	data := models.WritableInterface{
		Name:         &name,
		Description:  description,
		Type:         &interface_type,
		Enabled:      enabled,
		MgmtOnly:     mgmtonly,
		Mode:         mode,
		Tags:         tags,
		TaggedVlans:  taggedVlans,
		Device:       &deviceID,
		WirelessLans: []int64{},
		Vdcs:         []int64{},
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

	params := dcim.NewDcimInterfacesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimInterfacesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxDeviceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := dcim.NewDcimInterfacesReadParams().WithID(id)

	res, err := api.Dcim.DcimInterfacesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimInterfacesReadDefault); ok {
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
	d.Set("type", iface.Type.Value)
	d.Set("enabled", iface.Enabled)
	d.Set("mgmtonly", iface.MgmtOnly)
	d.Set("mac_address", iface.MacAddress)
	d.Set("mtu", iface.Mtu)
	d.Set(tagsKey, getTagListFromNestedTagList(iface.Tags))
	d.Set("tagged_vlans", getIDsFromNestedVLANDevice(iface.TaggedVlans))
	d.Set("device_id", iface.Device.ID)

	if iface.Mode != nil {
		d.Set("mode", iface.Mode.Value)
	}
	if iface.UntaggedVlan != nil {
		d.Set("untagged_vlan", iface.UntaggedVlan.ID)
	}

	return diags
}

func resourceNetboxDeviceInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	interface_type := d.Get("type").(string)
	enabled := d.Get("enabled").(bool)
	mgmtonly := d.Get("mgmtonly").(bool)
	mode := d.Get("mode").(string)
	tags, diagnostics := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	if diagnostics != nil {
		diags = append(diags, diagnostics...)
	}
	taggedVlans := toInt64List(d.Get("tagged_vlans"))
	deviceID := int64(d.Get("device_id").(int))

	data := models.WritableInterface{
		Name:         &name,
		Description:  description,
		Type:         &interface_type,
		Enabled:      enabled,
		MgmtOnly:     mgmtonly,
		Mode:         mode,
		Tags:         tags,
		TaggedVlans:  taggedVlans,
		Device:       &deviceID,
		WirelessLans: []int64{},
		Vdcs:         []int64{},
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

	params := dcim.NewDcimInterfacesPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Dcim.DcimInterfacesPartialUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxDeviceInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimInterfacesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimInterfacesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimInterfacesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}

func getIDsFromNestedVLANDevice(nestedvlans []*models.NestedVLAN) []int64 {
	var vlans []int64
	for _, vlan := range nestedvlans {
		vlans = append(vlans, vlan.ID)
	}
	return vlans
}
