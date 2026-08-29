package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxCircuitGroupAssignment() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxCircuitGroupAssignmentRead,
		Description: `:meta:subcategory:Circuits:`,
		Schema: map[string]*schema.Schema{
			"id": {
				Required: true,
				Type:     schema.TypeString,
			},
			"group_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"member_type": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"member_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"priority": {
				Computed: true,
				Type:     schema.TypeString,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxCircuitGroupAssignmentRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	params := circuits.NewCircuitsCircuitGroupAssignmentsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitGroupAssignmentsRead(params, nil)
	if err != nil {
		return err
	}

	assignment := res.GetPayload()

	d.SetId(strconv.FormatInt(assignment.ID, 10))

	if assignment.Group != nil {
		d.Set("group_id", assignment.Group.ID)
	}

	d.Set("member_id", assignment.MemberID)
	d.Set("member_type", assignment.MemberType)

	if assignment.Priority != nil {
		d.Set("priority", assignment.Priority.Value)
	}

	d.Set(tagsKey, getTagListFromNestedTagList(assignment.Tags))
	return nil
}
