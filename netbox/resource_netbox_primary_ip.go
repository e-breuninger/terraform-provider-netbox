package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxPrimaryIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPrimaryIPCreate,
		Read:   resourceNetboxPrimaryIPRead,
		Update: resourceNetboxPrimaryIPUpdate,
		Delete: resourceNetboxPrimaryIPDelete,

		Description: `:meta:subcategory:Virtualization:This resource is used to define the primary IP for a given virtual machine. The primary IP is reflected in the Virtual machine Netbox UI, which identifies the Primary IPv4 and IPv6 addresses.`,

		Schema: map[string]*schema.Schema{
			"virtual_machine_id": {
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

func resourceNetboxPrimaryIPCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(strconv.Itoa(d.Get("virtual_machine_id").(int)))

	return resourceNetboxPrimaryIPUpdate(d, m)
}

func resourceNetboxPrimaryIPRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationVirtualMachinesReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationVirtualMachinesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationVirtualMachinesReadDefault); ok {
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
		// if the vm exists, but has no primary ip, consider this element deleted
		d.SetId("")
		return nil
	}
	d.Set("virtual_machine_id", res.GetPayload().ID)
	return nil
}

func resourceNetboxPrimaryIPUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	virtualMachineID := int64(d.Get("virtual_machine_id").(int))
	IPAddressID := int64(d.Get("ip_address_id").(int))
	IPAddressVersion := int64(d.Get("ip_address_version").(int))

	// because the go-netbox library does not have patch support atm, we have to get the whole object and re-put it

	// first, get the vm
	readParams := virtualization.NewVirtualizationVirtualMachinesReadParams().WithID(virtualMachineID)
	res, err := api.Virtualization.VirtualizationVirtualMachinesRead(readParams, nil)
	if err != nil {
		return err
	}

	vm := res.GetPayload()

	// then update the FULL vm with ALL tracked attributes
	data := models.WritableVirtualMachineWithConfigContext{}
	data.Name = vm.Name
	data.Tags = vm.Tags
	// the netbox API sends the URL property as part of NestedTag, but it does not accept the URL property when we send it back
	// so set it to empty
	// display too
	for _, tag := range data.Tags {
		tag.URL = ""
		tag.Display = ""
	}
	data.Comments = vm.Comments
	data.Memory = vm.Memory
	data.Vcpus = vm.Vcpus
	data.Disk = vm.Disk

	if vm.Cluster != nil {
		data.Cluster = &vm.Cluster.ID
	}

	if vm.Site != nil {
		data.Site = &vm.Site.ID
	}

	if vm.PrimaryIp4 != nil {
		data.PrimaryIp4 = &vm.PrimaryIp4.ID
	}
	if vm.PrimaryIp6 != nil {
		data.PrimaryIp6 = &vm.PrimaryIp6.ID
	}

	if vm.Platform != nil {
		data.Platform = &vm.Platform.ID
	}

	if vm.Tenant != nil {
		data.Tenant = &vm.Tenant.ID
	}

	if vm.Role != nil {
		data.Role = &vm.Role.ID
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

	updateParams := virtualization.NewVirtualizationVirtualMachinesUpdateParams().WithID(virtualMachineID).WithData(&data)

	_, err = api.Virtualization.VirtualizationVirtualMachinesUpdate(updateParams, nil)
	if err != nil {
		return err
	}
	return resourceNetboxPrimaryIPRead(d, m)
}

func resourceNetboxPrimaryIPDelete(d *schema.ResourceData, m interface{}) error {
	// Set ip_address_id to minus one and go to update. Update will set nil
	d.Set("ip_address_id", -1)
	return resourceNetboxPrimaryIPUpdate(d, m)
}
