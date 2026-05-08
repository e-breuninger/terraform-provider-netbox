package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxVirtualMachineInterfacePrimaryMACAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVirtualMachineInterfacePrimaryMACAddressCreate,
		Read:   resourceNetboxVirtualMachineInterfacePrimaryMACAddressRead,
		Update: resourceNetboxVirtualMachineInterfacePrimaryMACAddressUpdate,
		Delete: resourceNetboxVirtualMachineInterfacePrimaryMACAddressDelete,

		Description: `:meta:subcategory:Virtualization:This resource is used to define the primary MAC for a given virtual machine interface. The primary MAC is reflected in the Interface Netbox UI.`,

		Schema: map[string]*schema.Schema{
			"interface_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"mac_address_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVirtualMachineInterfacePrimaryMACAddressCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(strconv.Itoa(d.Get("interface_id").(int)))

	return resourceNetboxVirtualMachineInterfacePrimaryMACAddressUpdate(d, m)
}

func resourceNetboxVirtualMachineInterfacePrimaryMACAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
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
		return err
	}

	if res.GetPayload().PrimaryMacAddress != nil {
		d.Set("mac_address_id", res.GetPayload().PrimaryMacAddress.ID)
		d.Set("interface_id", res.GetPayload().ID)
	} else {
		d.SetId("")
	}
	return nil
}

func resourceNetboxVirtualMachineInterfacePrimaryMACAddressUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	interfaceID := int64(d.Get("interface_id").(int))
	macAddressID := int64(d.Get("mac_address_id").(int))

	// because the go-netbox library does not have patch support atm, we have to get the whole object and re-put it

	// get the interface
	readParams := virtualization.NewVirtualizationInterfacesReadParams().WithID(interfaceID)
	res, err := api.Virtualization.VirtualizationInterfacesRead(readParams, nil)
	if err != nil {
		return err
	}
	iface := res.GetPayload()

	// then update the FULL interface with ALL tracked attributes
	data := &models.WritableVMInterface{
		Name:           iface.Name,
		VirtualMachine: &iface.VirtualMachine.ID, // Allowed to set directly as field is required
		CustomFields:   iface.CustomFields,
		Description:    iface.Description,
		Enabled:        iface.Enabled,
		ID:             iface.ID,
		MacAddress:     iface.MacAddress,
		Mtu:            iface.Mtu,
		Tags:           iface.Tags,
		URL:            iface.URL,
	}

	// the netbox API sends the URL property as part of NestedTag, but it does not accept the URL property when we send it back
	// so set it to empty
	// display too
	for _, tag := range data.Tags {
		tag.URL = ""
		tag.Display = ""
	}

	if iface.Bridge != nil {
		data.Bridge = &iface.Bridge.ID
	}

	if iface.Vrf != nil {
		data.Vrf = &iface.Vrf.ID
	}

	if iface.UntaggedVlan != nil {
		data.UntaggedVlan = &iface.UntaggedVlan.ID
	}

	if iface.Parent != nil {
		data.Parent = &iface.Parent.ID
	}

	if iface.Mode != nil {
		data.Mode = *iface.Mode.Value
	}

	if iface.L2vpnTermination != nil {
		data.L2vpnTermination = strconv.FormatInt(iface.L2vpnTermination.ID, 10)
	}

	if iface.TaggedVlans != nil {
		vlanTags := make([]int64, len(iface.TaggedVlans))
		for i, vlan := range iface.TaggedVlans {
			vlanTags[i] = vlan.ID
		}
		data.TaggedVlans = vlanTags
	}

	if macAddressID == -1 {
		data.PrimaryMacAddress = nil
	} else {
		data.PrimaryMacAddress = &macAddressID
	}

	updateParams := virtualization.NewVirtualizationInterfacesUpdateParams().WithID(interfaceID).WithData(data)

	_, err = api.Virtualization.VirtualizationInterfacesUpdate(updateParams, nil)
	if err != nil {
		return err
	}
	return resourceNetboxVirtualMachineInterfacePrimaryMACAddressRead(d, m)
}

func resourceNetboxVirtualMachineInterfacePrimaryMACAddressDelete(d *schema.ResourceData, m interface{}) error {
	// Set mac_address_id to minus one and go to update. Update will set nil
	d.Set("mac_address_id", -1)
	return resourceNetboxVirtualMachineInterfacePrimaryMACAddressUpdate(d, m)
}
