package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceInterfacePrimaryMACAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceInterfacePrimaryMACAddressCreate,
		Read:   resourceNetboxDeviceInterfacePrimaryMACAddressRead,
		Update: resourceNetboxDeviceInterfacePrimaryMACAddressUpdate,
		Delete: resourceNetboxDeviceInterfacePrimaryMACAddressDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):This resource is used to define the primary MAC for a given device interface. The primary MAC is reflected in the device interface Netbox UI.`,

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

func resourceNetboxDeviceInterfacePrimaryMACAddressCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(strconv.Itoa(d.Get("interface_id").(int)))

	return resourceNetboxDeviceInterfacePrimaryMACAddressUpdate(d, m)
}

func resourceNetboxDeviceInterfacePrimaryMACAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
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

func resourceNetboxDeviceInterfacePrimaryMACAddressUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	interfaceID := int64(d.Get("interface_id").(int))
	macAddressID := int64(d.Get("mac_address_id").(int))

	// because the go-netbox library does not have patch support atm, we have to get the whole object and re-put it

	// get the interface
	readParams := dcim.NewDcimInterfacesReadParams().WithID(interfaceID)
	res, err := api.Dcim.DcimInterfacesRead(readParams, nil)
	if err != nil {
		return err
	}
	iface := res.GetPayload()

	// then update the FULL interface with ALL tracked attributes
	data := &models.WritableInterface{
		Device:             &iface.Device.ID, // Allowed to set directly as field is required
		Name:               iface.Name,
		Type:               iface.Type.Value,
		CustomFields:       iface.CustomFields,
		Description:        iface.Description,
		Enabled:            iface.Enabled,
		Label:              iface.Label,
		Mtu:                iface.Mtu,
		MarkConnected:      iface.MarkConnected,
		MgmtOnly:           iface.MgmtOnly,
		RfChannelFrequency: iface.RfChannelFrequency,
		RfChannelWidth:     iface.RfChannelWidth,
		Speed:              iface.Speed,
		TxPower:            iface.TxPower,
		Wwn:                iface.Wwn,
		Tags:               iface.Tags,
		Vdcs:               []int64{},
		WirelessLans:       []int64{},
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

	if iface.Lag != nil {
		data.Lag = &iface.Lag.ID
	}

	if iface.Module != nil {
		data.Module = &iface.Module.ID
	}

	if iface.PoeMode != nil {
		data.PoeMode = *iface.PoeMode.Value
	}

	if iface.PoeType != nil {
		data.PoeType = *iface.PoeType.Value
	}

	if iface.RfChannel != nil {
		data.RfChannel = *iface.RfChannel.Value
	}

	if iface.RfRole != nil {
		data.RfRole = *iface.RfRole.Value
	}

	if iface.Vrf != nil {
		data.Vrf = &iface.Vrf.ID
	}

	if iface.WirelessLink != nil {
		data.WirelessLink = &iface.WirelessLink.ID
	}

	for _, vdc := range iface.Vdcs {
		data.Vdcs = append(data.Vdcs, vdc.ID)
	}

	for _, wlan := range iface.WirelessLans {
		data.WirelessLans = append(data.WirelessLans, wlan.ID)
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

	updateParams := dcim.NewDcimInterfacesUpdateParams().WithID(interfaceID).WithData(data)

	_, err = api.Dcim.DcimInterfacesUpdate(updateParams, nil)
	if err != nil {
		return err
	}
	return resourceNetboxDeviceInterfacePrimaryMACAddressRead(d, m)
}

func resourceNetboxDeviceInterfacePrimaryMACAddressDelete(d *schema.ResourceData, m interface{}) error {
	// Set mac_address_id to minus one and go to update. Update will set nil
	d.Set("mac_address_id", -1)
	return resourceNetboxDeviceInterfacePrimaryMACAddressUpdate(d, m)
}
