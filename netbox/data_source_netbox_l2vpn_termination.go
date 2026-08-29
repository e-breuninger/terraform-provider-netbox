package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxL2VPNTermination() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxL2VPNTerminationRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"l2vpn_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"assigned_object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assigned_object_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxL2VPNTerminationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id := int64(d.Get("id").(int))
	params := ipam.NewIpamL2vpnTerminationsReadParams().WithID(id)

	res, err := api.Ipam.IpamL2vpnTerminationsRead(params, nil)
	if err != nil {
		return err
	}

	result := res.GetPayload()
	d.SetId(strconv.FormatInt(result.ID, 10))

	if result.L2vpn != nil {
		d.Set("l2vpn_id", result.L2vpn.ID)
	}

	if result.AssignedObjectType != nil {
		d.Set("assigned_object_type", result.AssignedObjectType)
	}

	if result.AssignedObjectID != nil {
		d.Set("assigned_object_id", result.AssignedObjectID)
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
