package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxService() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxServiceRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"virtual_machine_id"},
			},
			"virtual_machine_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"device_id"},
			},
			"protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ports": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxServiceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamServicesListParams()

	name := d.Get("name").(string)
	params.Name = &name

	if deviceID, ok := d.GetOk("device_id"); ok {
		deviceIDStr := strconv.Itoa(deviceID.(int))
		params.DeviceID = &deviceIDStr
	}

	if vmID, ok := d.GetOk("virtual_machine_id"); ok {
		vmIDStr := strconv.Itoa(vmID.(int))
		params.VirtualMachineID = &vmIDStr
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Ipam.IpamServicesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one service returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no service found matching filter")
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)

	// Since NetBox 4.4 the parent is a generic object reference rather than a
	// dedicated device/virtual_machine field.
	switch result.ParentObjectType {
	case "virtualization.virtualmachine":
		d.Set("virtual_machine_id", result.ParentObjectID)
	case "dcim.device":
		d.Set("device_id", result.ParentObjectID)
	}

	if result.Protocol != nil {
		d.Set("protocol", result.Protocol.Value)
	}
	d.Set("ports", result.Ports)
	d.Set("description", result.Description)
	return nil
}
