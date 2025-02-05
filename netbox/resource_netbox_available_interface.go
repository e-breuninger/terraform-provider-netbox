package netbox

import (
	"context"
	"fmt"
	"maps"
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxAvailableInterface() *schema.Resource {
	// schema is extended from `netbox_device_interface`
	resourceSchema := maps.Clone(resourceNetboxDeviceInterfaceSchema)
	resourceSchema["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	resourceSchema["prefix"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Prefix to use for the new interface",
	}

	return &schema.Resource{
		CreateContext: resourceNetboxAvailableInterfaceCreate,
		ReadContext:   resourceNetboxAvailableInterfaceRead,
		UpdateContext: resourceNetboxAvailableInterfaceUpdate,
		DeleteContext: resourceNetboxAvailableInterfaceDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/device/#interface):

> Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.

This is a special resource that is able to allocate a free interface name based on existing interfaces and keep it stable through the Terraform state.`,
		Schema: resourceSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxAvailableInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	deviceID := int64(d.Get("device_id").(int))
	prefix := d.Get("prefix").(string)

	// secret column that enables natural ordering
	// see https://github.com/netbox-community/netbox/issues/11279#issuecomment-1367394944
	ordering := "_name"

	deviceIDStr := strconv.FormatInt(deviceID, 10)
	p := dcim.DcimInterfacesListParams{
		DeviceID: &deviceIDStr,
		NameIsw:  &prefix, // case-insensitive starts with
		Ordering: &ordering,
	}
	res, err := api.Dcim.DcimInterfacesList(&p, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// look for the first suffix that doesn't clash with an existing interface
	var currentSuffix int64 = 0 // interfaces start with 0
	for _, intf := range res.GetPayload().Results {
		ifSuffix, ok := strings.CutPrefix(*intf.Name, prefix)
		if !ok {
			continue // just for safety, this should not happen
		}
		suffixNum, err := strconv.ParseInt(ifSuffix, 10, 64)
		if err != nil {
			// ignore suffixes that are not integers
			continue
		}
		if suffixNum != currentSuffix {
			// we found a spot
			break
		}
		currentSuffix++
	}

	ifName := fmt.Sprintf("%s%d", prefix, currentSuffix)
	d.Set("name", ifName)

	return resourceNetboxDeviceInterfaceCreate(ctx, d, m)
}

func resourceNetboxAvailableInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetboxDeviceInterfaceRead(ctx, d, m)
}

func resourceNetboxAvailableInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetboxDeviceInterfaceUpdate(ctx, d, m)
}

func resourceNetboxAvailableInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetboxDeviceInterfaceDelete(ctx, d, m)
}
