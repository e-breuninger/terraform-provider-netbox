package netbox

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceTagCreate,
		Read:   resourceNetboxDeviceTagRead,
		Delete: resourceNetboxDeviceTagDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):This resource attaches a single tag to a device, similar in purpose to the ` + "`aws_ec2_tag`" + ` resource in the AWS provider.

Use this when a device is created and managed outside of Terraform (e.g. via onboarding automation or the NetBox UI) but you still want Terraform to manage the presence of one or more specific tags on it, without taking ownership of the rest of the device.

Do not use this on a device that is also managed by a ` + "`netbox_device`" + ` resource. Since that resource does not itself declare the tag, it will treat the tag added here as drift and try to remove it on every subsequent plan/apply; this is deterministic, not a race. Unlike the AWS provider, which has an ` + "`ignore_tags`" + ` provider setting for exactly this situation, this provider has no way to make a resource ignore a specific tag, so the two resources cannot currently coexist on the same device.

This resource performs a read-modify-write of the device's tag list. If the same device is targeted by more than one ` + "`netbox_device_tag`" + ` resource (or has its tags modified by anything else) at the same time, concurrent writes can race and one may silently drop the other's change.`,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceTagParseID(id string) (int64, string, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("invalid ID %q, expected format <device_id>:<tag>", id)
	}
	deviceID, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, "", fmt.Errorf("invalid device ID in %q: %w", id, err)
	}
	return deviceID, parts[1], nil
}

// cleanNestedTags strips the URL and Display fields the API sends on read, which it refuses to accept when sent back on write.
func cleanNestedTags(tags []*models.NestedTag) []*models.NestedTag {
	cleaned := make([]*models.NestedTag, 0, len(tags))
	for _, tag := range tags {
		t := *tag
		t.URL = ""
		t.Display = ""
		cleaned = append(cleaned, &t)
	}
	return cleaned
}

// resourceNetboxDeviceTagUpdateDeviceTags overwrites a device's tag list. device_type, role and site are re-sent
// because the NetBox API does not treat them as optional on partial update: omitting them clears them. The
// primary IPs are preserved too, matching resource_netbox_device_oob_ip.go and resource_netbox_device_primary_ip.go.
func resourceNetboxDeviceTagUpdateDeviceTags(api *providerState, device *models.DeviceWithConfigContext, tags []*models.NestedTag) error {
	data := models.WritableDeviceWithConfigContext{}
	data.Name = device.Name
	data.Tags = tags

	if device.DeviceType != nil {
		data.DeviceType = &device.DeviceType.ID
	}
	if device.Site != nil {
		data.Site = &device.Site.ID
	}
	if device.Role != nil {
		data.Role = &device.Role.ID
	}
	if device.PrimaryIp4 != nil {
		data.PrimaryIp4 = &device.PrimaryIp4.ID
	}
	if device.PrimaryIp6 != nil {
		data.PrimaryIp6 = &device.PrimaryIp6.ID
	}

	updateParams := dcim.NewDcimDevicesPartialUpdateParams().WithID(device.ID).WithData(&data)
	_, err := api.Dcim.DcimDevicesPartialUpdate(updateParams, nil)
	return err
}

func resourceNetboxDeviceTagCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	deviceID := int64(d.Get("device_id").(int))
	tagName := d.Get("tag").(string)

	nbTag, err := findTag(api.NetBoxAPI, tagName)
	if err != nil {
		return err
	}

	readParams := dcim.NewDcimDevicesReadParams().WithID(deviceID)
	res, err := api.Dcim.DcimDevicesRead(readParams, nil)
	if err != nil {
		return err
	}
	device := res.GetPayload()

	tags := cleanNestedTags(device.Tags)
	alreadyPresent := false
	for _, tag := range tags {
		if tag.Name != nil && nbTag.Name != nil && *tag.Name == *nbTag.Name {
			alreadyPresent = true
			break
		}
	}
	if !alreadyPresent {
		tags = append(tags, nbTag)
	}

	if err := resourceNetboxDeviceTagUpdateDeviceTags(api, device, tags); err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d:%s", deviceID, tagName))
	return resourceNetboxDeviceTagRead(d, m)
}

func resourceNetboxDeviceTagRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	deviceID, tagName, err := resourceNetboxDeviceTagParseID(d.Id())
	if err != nil {
		return err
	}

	readParams := dcim.NewDcimDevicesReadParams().WithID(deviceID)
	res, err := api.Dcim.DcimDevicesRead(readParams, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDevicesReadDefault); ok {
			if errresp.Code() == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	found := false
	for _, tag := range res.GetPayload().Tags {
		if tag.Name != nil && *tag.Name == tagName {
			found = true
			break
		}
	}

	if !found {
		// the tag was removed from the device out of band
		d.SetId("")
		return nil
	}

	d.Set("device_id", deviceID)
	d.Set("tag", tagName)
	return nil
}

func resourceNetboxDeviceTagDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	deviceID, tagName, err := resourceNetboxDeviceTagParseID(d.Id())
	if err != nil {
		return err
	}

	readParams := dcim.NewDcimDevicesReadParams().WithID(deviceID)
	res, err := api.Dcim.DcimDevicesRead(readParams, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDevicesReadDefault); ok {
			if errresp.Code() == 404 {
				// the device is already gone, nothing left to clean up
				return nil
			}
		}
		return err
	}
	device := res.GetPayload()

	tags := []*models.NestedTag{}
	for _, tag := range cleanNestedTags(device.Tags) {
		if tag.Name != nil && *tag.Name == tagName {
			continue
		}
		tags = append(tags, tag)
	}

	return resourceNetboxDeviceTagUpdateDeviceTags(api, device, tags)
}
