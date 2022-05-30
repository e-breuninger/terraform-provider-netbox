package netbox

import (
	"regexp"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxInterface() *schema.Resource {
	validModes := []string{"access", "tagged", "tagged-all"}

	return &schema.Resource{
		Create: resourceNetboxInterfaceCreate,
		Read:   resourceNetboxInterfaceRead,
		Update: resourceNetboxInterfaceUpdate,
		Delete: resourceNetboxInterfaceDelete,

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
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^([A-Z0-9]{2}:){5}[A-Z0-9]{2}$"),
					"Must be like AA:AA:AA:AA:AA"),
				ForceNew: true,
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
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	mode := d.Get("mode").(string)
	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))
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
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxInterfaceUpdate(d, m)
}

func resourceNetboxInterfaceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := virtualization.NewVirtualizationInterfacesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationInterfacesRead(params, nil)
	if err != nil {
		errorcode := err.(*virtualization.VirtualizationInterfacesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	iface := res.GetPayload()

	d.Set("name", iface.Name)
	d.Set("description", iface.Description)
	d.Set("enabled", iface.Enabled)
	d.Set("mac_address", iface.MacAddress)
	d.Set("mtu", iface.Mtu)
	d.Set("tags", getTagListFromNestedTagList(iface.Tags))
	d.Set("tagged_vlans", getIDsFromNestedVLAN(iface.TaggedVlans))
	d.Set("virtual_machine_id", iface.VirtualMachine.ID)

	if iface.Mode != nil {
		d.Set("mode", iface.Mode.Value)
	}
	if iface.UntaggedVlan != nil {
		d.Set("untagged_vlan", iface.UntaggedVlan.ID)
	}

	return nil
}

func resourceNetboxInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	mode := d.Get("mode").(string)
	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))
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
		return err
	}

	return resourceNetboxInterfaceRead(d, m)
}

func resourceNetboxInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationInterfacesDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationInterfacesDelete(params, nil)
	if err != nil {
		return err
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
