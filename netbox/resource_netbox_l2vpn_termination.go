package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxL2VPNTermination() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxL2VPNTerminationCreate,
		Read:   resourceNetboxL2VPNTerminationRead,
		Update: resourceNetboxL2VPNTerminationUpdate,
		Delete: resourceNetboxL2VPNTerminationDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/l2vpntermination/):

> This model represents the termination of an L2VPN on a device interface, virtual machine interface, or VLAN.`,

		Schema: map[string]*schema.Schema{
			"l2vpn_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"assigned_object_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "One of `dcim.interface`, `virtualization.vminterface` or `ipam.vlan`.",
			},
			"assigned_object_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxL2VPNTerminationCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	l2vpnID := int64(d.Get("l2vpn_id").(int))
	assignedObjectType := d.Get("assigned_object_type").(string)
	assignedObjectID := int64(d.Get("assigned_object_id").(int))

	data := models.WritableL2VPNTermination{
		L2vpn:              &l2vpnID,
		AssignedObjectType: &assignedObjectType,
		AssignedObjectID:   &assignedObjectID,
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := ipam.NewIpamL2vpnTerminationsCreateParams().WithData(&data)
	res, err := api.Ipam.IpamL2vpnTerminationsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxL2VPNTerminationRead(d, m)
}

func resourceNetboxL2VPNTerminationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamL2vpnTerminationsReadParams().WithID(id)

	res, err := api.Ipam.IpamL2vpnTerminationsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamL2vpnTerminationsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	termination := res.GetPayload()

	if termination.L2vpn != nil {
		d.Set("l2vpn_id", termination.L2vpn.ID)
	}

	if termination.AssignedObjectType != nil {
		d.Set("assigned_object_type", termination.AssignedObjectType)
	}

	if termination.AssignedObjectID != nil {
		d.Set("assigned_object_id", termination.AssignedObjectID)
	}

	api.readTags(d, termination.Tags)

	return nil
}

func resourceNetboxL2VPNTerminationUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	l2vpnID := int64(d.Get("l2vpn_id").(int))
	assignedObjectType := d.Get("assigned_object_type").(string)
	assignedObjectID := int64(d.Get("assigned_object_id").(int))

	data := models.WritableL2VPNTermination{
		L2vpn:              &l2vpnID,
		AssignedObjectType: &assignedObjectType,
		AssignedObjectID:   &assignedObjectID,
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := ipam.NewIpamL2vpnTerminationsPartialUpdateParams().WithID(id).WithData(&data)
	_, err = api.Ipam.IpamL2vpnTerminationsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxL2VPNTerminationRead(d, m)
}

func resourceNetboxL2VPNTerminationDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamL2vpnTerminationsDeleteParams().WithID(id)

	_, err := api.Ipam.IpamL2vpnTerminationsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamL2vpnTerminationsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.SetId("")
	return nil
}
