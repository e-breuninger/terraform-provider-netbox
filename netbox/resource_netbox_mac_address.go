package netbox

import (
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxMACAddressObjectTypeOptions = []string{"virtualization.vminterface", "dcim.interface"}

func resourceNetboxMACAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxMACAddressCreate,
		Read:   resourceNetboxMACAddressRead,
		Update: resourceNetboxMACAddressUpdate,
		Delete: resourceNetboxMACAddressDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/macaddress/):

> A MAC address object in NetBox comprises a single Ethernet link layer address, and represents a MAC address as reported by or assigned to a network interface. MAC addresses can be assigned to device and virtual machine interfaces. A MAC address can be specified as the primary MAC address for a given device or VM interface.`,

		Schema: map[string]*schema.Schema{
			"mac_address": {
				Type:     schema.TypeString,
				Required: true,
				// Netbox converts MAC addresses always to uppercase
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},
			"interface_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"object_type"},
			},
			"object_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxMACAddressObjectTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxMACAddressObjectTypeOptions),
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxMACAddressCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.MACAddress{}

	data.MacAddress = strToPtr(d.Get("mac_address").(string))

	data.Description = strToPtr(getOptionalStr(d, "description", false))
	data.Comments = strToPtr(getOptionalStr(d, "comments", false))

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
	// default = mac address is not linked to anything
	default:
		data.AssignedObjectType = strToPtr("")
		data.AssignedObjectID = nil
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimMacAddressesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimMacAddressesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxMACAddressRead(d, m)
}

func resourceNetboxMACAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimMacAddressesReadParams().WithID(id)

	res, err := api.Dcim.DcimMacAddressesRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*dcim.DcimMacAddressesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	macAddress := res.GetPayload()

	if macAddress.AssignedObjectID != nil && macAddress.AssignedObjectType != nil {
		vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
		deviceInterfaceID := getOptionalInt(d, "device_interface_id")
		interfaceID := getOptionalInt(d, "interface_id")

		switch {
		case vmInterfaceID != nil && *macAddress.AssignedObjectType == "virtualization.vminterface":
			d.Set("virtual_machine_interface_id", macAddress.AssignedObjectID)
		case deviceInterfaceID != nil && *macAddress.AssignedObjectType == "dcim.interface":
			d.Set("device_interface_id", macAddress.AssignedObjectID)
		// if interfaceID is given, object_type must be set as well
		case interfaceID != nil:
			d.Set("object_type", macAddress.AssignedObjectType)
			d.Set("interface_id", macAddress.AssignedObjectID)
		}
	} else {
		d.Set("interface_id", nil)
		d.Set("object_type", "")
	}

	d.Set("mac_address", macAddress.MacAddress)
	d.Set("description", macAddress.Description)
	d.Set("comments", macAddress.Comments)
	api.readTags(d, macAddress.Tags)

	cf := flattenCustomFields(macAddress.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxMACAddressUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.MACAddress{}

	data.MacAddress = strToPtr(d.Get("mac_address").(string))

	data.Description = strToPtr(getOptionalStr(d, "description", false))
	data.Comments = strToPtr(getOptionalStr(d, "comments", false))

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
	// default = mac address is not linked to anything
	default:
		data.AssignedObjectType = strToPtr("")
		data.AssignedObjectID = nil
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimMacAddressesPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimMacAddressesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxMACAddressRead(d, m)
}

func resourceNetboxMACAddressDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimMacAddressesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimMacAddressesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimMacAddressesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
