package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxFhrpGroupAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxFhrpGroupAssignmentCreate,
		Read:   resourceNetboxFhrpGroupAssignmentRead,
		Update: resourceNetboxFhrpGroupAssignmentUpdate,
		Delete: resourceNetboxFhrpGroupAssignmentDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/fhrpgroup/):`,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Group ID",
			},
			"interface_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Interface ID",
			},
			"priority": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Priority",
			},
			"interface_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Interface type",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxFhrpGroupAssignmentCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableFHRPGroupAssignment{}

	group_id := int64(d.Get("group_id").(int))
	data.Group = &group_id
	interface_id := int64(d.Get("interface_id").(int))
	data.InterfaceID = &interface_id
	priority := int64(d.Get("priority").(int))
	data.Priority = &priority
	interface_type := d.Get("interface_type").(string)
	data.InterfaceType = &interface_type

	params := ipam.NewIpamFhrpGroupAssignmentsCreateParams().WithData(&data)

	var err error
	res, err := api.Ipam.IpamFhrpGroupAssignmentsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxFhrpGroupAssignmentRead(d, m)
}

func resourceNetboxFhrpGroupAssignmentRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamFhrpGroupAssignmentsReadParams().WithID(id)

	res, err := api.Ipam.IpamFhrpGroupAssignmentsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*ipam.IpamAsnsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	assignment := res.GetPayload()
	
	d.Set("group_id", assignment.Group.GroupID)
	d.Set("interface_id", assignment.Interface)
	d.Set("interface_type", assignment.InterfaceType)
	d.Set("priority", assignment.Priority)

	return nil
}

func resourceNetboxFhrpGroupAssignmentUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableFHRPGroupAssignment{}

	group_id := int64(d.Get("group_id").(int))
	data.Group = &group_id
	interface_id := int64(d.Get("interface_id").(int))
	data.InterfaceID = &interface_id
	priority := int64(d.Get("priority").(int))
	data.Priority = &priority
	interface_type := d.Get("interface_type").(string)
	data.InterfaceType = &interface_type

	params := ipam.NewIpamFhrpGroupAssignmentsUpdateParams().WithID(id).WithData(&data)

	var err error

	_, err = api.Ipam.IpamFhrpGroupAssignmentsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxFhrpGroupAssignmentRead(d, m)
}

func resourceNetboxFhrpGroupAssignmentDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamFhrpGroupAssignmentsDeleteParams().WithID(id)

	_, err := api.Ipam.IpamFhrpGroupAssignmentsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamFhrpGroupAssignmentsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
