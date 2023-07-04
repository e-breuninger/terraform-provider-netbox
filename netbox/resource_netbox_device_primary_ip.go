package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxDevicePrimaryIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDevicePrimaryIPCreate,
		Read:   resourceNetboxDevicePrimaryIPRead,
		Update: resourceNetboxDevicePrimaryIPUpdate,
		Delete: resourceNetboxDevicePrimaryIPDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):This resource is used to define the primary IP for a given device. The primary IP is reflected in the device Netbox UI, which identifies the Primary IPv4 and IPv6 addresses.`,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_address_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_address_version": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{4, 6}),
				Optional:     true,
				Default:      4,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDevicePrimaryIPCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(strconv.Itoa(d.Get("device_id").(int)))

	return resourceNetboxDevicePrimaryIPUpdate(d, m)
}

func resourceNetboxDevicePrimaryIPRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
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

	IPAddressVersion := d.Get("ip_address_version")
	d.Set("ip_address_version", IPAddressVersion)

	if IPAddressVersion == 4 && res.GetPayload().PrimaryIp4 != nil {
		d.Set("ip_address_id", res.GetPayload().PrimaryIp4.ID)
	} else if IPAddressVersion == 6 && res.GetPayload().PrimaryIp6 != nil {
		d.Set("ip_address_id", res.GetPayload().PrimaryIp6.ID)
	} else {
		// if the device exists, but has no primary ip, consider this element deleted
		d.SetId("")
		return nil
	}
	d.Set("device_id", res.GetPayload().ID)
	return nil
}

func resourceNetboxDevicePrimaryIPUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	deviceID := int64(d.Get("device_id").(int))
	IPAddressID := int64(d.Get("ip_address_id").(int))
	IPAddressVersion := int64(d.Get("ip_address_version").(int))

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

	data.Comments = device.Comments
	data.Serial = device.Serial

	if device.DeviceType != nil {
		data.DeviceType = &device.DeviceType.ID
	}

	if device.Cluster != nil {
		data.Cluster = &device.Cluster.ID
	}

	if device.Site != nil {
		data.Site = &device.Site.ID
	}

	if device.Location != nil {
		data.Location = &device.Location.ID
	}

	if device.DeviceRole != nil {
		data.DeviceRole = &device.DeviceRole.ID
	}

	if device.PrimaryIp4 != nil {
		data.PrimaryIp4 = &device.PrimaryIp4.ID
	}

	if device.PrimaryIp6 != nil {
		data.PrimaryIp6 = &device.PrimaryIp6.ID
	}

	if device.Platform != nil {
		data.Platform = &device.Platform.ID
	}

	if device.Tenant != nil {
		data.Tenant = &device.Tenant.ID
	}

	if device.Status != nil {
		data.Status = *device.Status.Value
	}

	if device.Rack != nil {
		data.Rack = &device.Rack.ID
	}

	if device.Face != nil {
		data.Face = *device.Face.Value
	}

	if device.Position != nil {
		data.Position = device.Position
	}

	// unset primary ip address if -1 is passed as id
	if IPAddressID == -1 {
		if IPAddressVersion == 4 {
			data.PrimaryIp4 = nil
		} else {
			data.PrimaryIp6 = nil
		}
	} else {
		if IPAddressVersion == 4 {
			data.PrimaryIp4 = &IPAddressID
		} else {
			data.PrimaryIp6 = &IPAddressID
		}
	}

	updateParams := dcim.NewDcimDevicesUpdateParams().WithID(deviceID).WithData(&data)

	_, err = api.Dcim.DcimDevicesUpdate(updateParams, nil)
	if err != nil {
		return err
	}
	return resourceNetboxDevicePrimaryIPRead(d, m)
}

func resourceNetboxDevicePrimaryIPDelete(d *schema.ResourceData, m interface{}) error {
	// Set ip_address_id to minus one and go to update. Update will set nil
	d.Set("ip_address_id", -1)
	return resourceNetboxDevicePrimaryIPUpdate(d, m)
}
