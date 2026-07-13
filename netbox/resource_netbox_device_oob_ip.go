package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceOobIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceOobIPCreate,
		Read:   resourceNetboxDeviceOobIPRead,
		Update: resourceNetboxDeviceOobIPUpdate,
		Delete: resourceNetboxDeviceOobIPDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):This resource is used to define the out-of-band (OOB) IP for a given device. Modelled as a separate resource (like ` + "`netbox_device_primary_ip`" + `) to avoid a dependency cycle between the device and an IP address assigned to one of its own interfaces.`,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_address_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceOobIPCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(strconv.Itoa(d.Get("device_id").(int)))

	return resourceNetboxDeviceOobIPUpdate(d, m)
}

func resourceNetboxDeviceOobIPRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDevicesReadParams().WithID(id)

	res, err := api.Dcim.DcimDevicesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDevicesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if res.GetPayload().OobIP != nil {
		d.Set("ip_address_id", res.GetPayload().OobIP.ID)
	} else {
		// if the device exists, but has no oob ip, consider this element deleted
		d.SetId("")
		return nil
	}
	d.Set("device_id", res.GetPayload().ID)
	return nil
}

func resourceNetboxDeviceOobIPUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	deviceID := int64(d.Get("device_id").(int))
	IPAddressID := int64(d.Get("ip_address_id").(int))

	// because the go-netbox library does not have patch support atm, we have to get the whole object and re-put it

	// first, get the device
	readParams := dcim.NewDcimDevicesReadParams().WithID(deviceID)
	res, err := api.Dcim.DcimDevicesRead(readParams, nil)
	if err != nil {
		return err
	}

	device := res.GetPayload()

	// then update the FULL device with ALL tracked attributes
	data := models.WritableDeviceWithConfigContext{}

	data.Name = device.Name
	data.Tags = device.Tags
	// the netbox API sends the URL property as part of NestedTag, but it does not accept the URL property when we send it back
	// so set it to empty
	// display too
	for _, tag := range data.Tags {
		tag.URL = ""
		tag.Display = ""
	}

	if device.DeviceType != nil {
		data.DeviceType = &device.DeviceType.ID
	}

	if device.Site != nil {
		data.Site = &device.Site.ID
	}

	if device.Role != nil {
		data.Role = &device.Role.ID
	}

	// preserve the primary ips so this partial update does not clear them
	if device.PrimaryIp4 != nil {
		data.PrimaryIp4 = &device.PrimaryIp4.ID
	}

	if device.PrimaryIp6 != nil {
		data.PrimaryIp6 = &device.PrimaryIp6.ID
	}

	// unset oob ip address if -1 is passed as id
	if IPAddressID == -1 {
		data.OobIP = nil
	} else {
		data.OobIP = &IPAddressID
	}

	updateParams := dcim.NewDcimDevicesPartialUpdateParams().WithID(deviceID).WithData(&data)

	_, err = api.Dcim.DcimDevicesPartialUpdate(updateParams, nil)
	if err != nil {
		return err
	}
	return resourceNetboxDeviceOobIPRead(d, m)
}

func resourceNetboxDeviceOobIPDelete(d *schema.ResourceData, m interface{}) error {
	// Set ip_address_id to minus one and go to update. Update will set nil
	d.Set("ip_address_id", -1)
	return resourceNetboxDeviceOobIPUpdate(d, m)
}
