package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxIPAddressAssignmentObjectTypeOptions = []string{"virtualization.vminterface", "dcim.interface"}

func resourceNetboxIPAddressAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIPAddressAssignmentCreate,
		Read:   resourceNetboxIPAddressAssignmentRead,
		Update: resourceNetboxIPAddressAssignmentUpdate,
		Delete: resourceNetboxIPAddressAssignmentDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-addresses):

> Assigns a NetBox Device, physical or virtual, to an already constructed IP address.
>
> In cases where the device assigned to the IP Address is not yet known when constructing the IP address (using either netbox_available_ip_address or netbox_ip_address), this resource allows assigning it independently.
>
> A typical scenario is when you statically allocate IP's to virtual machines and use netbox_available_ip_address to fetch that IP, but where the netbox_virtual_machine or netbox_interface can only be constructed after having started the virtual machine.`,

		Schema: map[string]*schema.Schema{
			"ip_address_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"interface_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"object_type"},
			},
			"object_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIPAddressAssignmentObjectTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIPAddressAssignmentObjectTypeOptions),
				RequiredWith: []string{"interface_id"},
			},
			"virtual_machine_interface_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"interface_id", "device_interface_id"},
			},
			"device_interface_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"interface_id", "virtual_machine_interface_id"},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxIPAddressAssignmentCreate(d *schema.ResourceData, m interface{}) error {
	id := d.Get("ip_address_id").(int)

	d.SetId(strconv.Itoa(id))

	return resourceNetboxIPAddressAssignmentUpdate(d, m)
}

func resourceNetboxIPAddressAssignmentRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamIPAddressesReadParams().WithID(id)

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamIPAddressesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	ipAddress := res.GetPayload()
	if ipAddress.AssignedObjectID != nil {
		vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
		deviceInterfaceID := getOptionalInt(d, "device_interface_id")
		interfaceID := getOptionalInt(d, "interface_id")

		switch {
		case vmInterfaceID != nil:
			d.Set("virtual_machine_interface_id", ipAddress.AssignedObjectID)
		case deviceInterfaceID != nil:
			d.Set("device_interface_id", ipAddress.AssignedObjectID)
		// if interfaceID is given, object_type must be set as well
		case interfaceID != nil:
			d.Set("object_type", ipAddress.AssignedObjectType)
			d.Set("interface_id", ipAddress.AssignedObjectID)
		}
	} else {
		d.Set("interface_id", nil)
		d.Set("object_type", "")
	}

	d.Set("ip_address_id", id)

	return nil
}

func resourceNetboxIPAddressAssignmentUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := ipam.NewIpamIPAddressesReadParams().WithID(id)

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamIPAddressesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	ipAddress := res.GetPayload()
	data := models.WritableIPAddress{}

	data.Address = ipAddress.Address
	// if ipAddress.Status != nil {
	// 	data.Status = *ipAddress.Status.Value
	// }

	// data.Description = ipAddress.Description
	if ipAddress.Role != nil {
		data.Role = *ipAddress.Role.Value
	}
	// data.DNSName = ipAddress.DNSName
	if ipAddress.Vrf != nil {
		data.Vrf = &ipAddress.Vrf.ID
	}
	if ipAddress.Tenant != nil {
		data.Tenant = &ipAddress.Tenant.ID
	}
	if ipAddress.NatInside != nil {
		data.NatInside = &ipAddress.NatInside.ID
	}

	tags := make([]*models.NestedTag, len(ipAddress.Tags))
	for i, t := range ipAddress.Tags {
		tags[i] = &models.NestedTag{Name: t.Name, Slug: t.Slug, Color: t.Color}
	}
	data.Tags = tags

	outsideNat := make([]*models.NestedIPAddress, len(ipAddress.NatOutside))
	for i, t := range ipAddress.NatOutside {
		outsideNat[i] = &models.NestedIPAddress{Address: t.Address}
	}
	data.NatOutside = outsideNat

	vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
	deviceInterfaceID := getOptionalInt(d, "device_interface_id")
	interfaceID := getOptionalInt(d, "interface_id")

	switch {
	case vmInterfaceID != nil:
		data.AssignedObjectType = strToPtr("virtualization.vminterface")
		data.AssignedObjectID = vmInterfaceID
	case deviceInterfaceID != nil:
		data.AssignedObjectType = strToPtr("dcim.interface")
		data.AssignedObjectID = deviceInterfaceID
	// if interfaceID is given, object_type must be set as well
	case interfaceID != nil:
		data.AssignedObjectType = strToPtr(d.Get("object_type").(string))
		data.AssignedObjectID = interfaceID
	// default = ip is not linked to anything
	default:
		data.AssignedObjectType = strToPtr("")
		data.AssignedObjectID = nil
	}

	params2 := ipam.NewIpamIPAddressesPartialUpdateParams().WithID(id).WithData(&data)

	_, err2 := api.Ipam.IpamIPAddressesPartialUpdate(params2, nil)
	if err2 != nil {
		return err2
	}

	return nil
}

func resourceNetboxIPAddressAssignmentDelete(d *schema.ResourceData, m interface{}) error {
	d.Set("interface_id", nil)
	d.Set("object_type", "")
	d.Set("virtual_machine_interface_id", nil)
	d.Set("device_interface_id", nil)

	return resourceNetboxIPAddressAssignmentUpdate(d, m)
}
