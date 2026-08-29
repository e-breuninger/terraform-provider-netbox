package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxCircuitGroupAssignmentPriorityOptions = []string{"primary", "secondary", "tertiary", "inactive"}

func resourceNetboxCircuitGroupAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCircuitGroupAssignmentCreate,
		Read:   resourceNetboxCircuitGroupAssignmentRead,
		Update: resourceNetboxCircuitGroupAssignmentUpdate,
		Delete: resourceNetboxCircuitGroupAssignmentDelete,

		Description: `:meta:subcategory:Circuits:From the [official documentation](https://netboxlabs.com/docs/netbox/models/circuits/circuitgroupassignment/):

> Circuit group assignments are used to associate a circuit (or virtual circuit) with a circuit group, optionally indicating its priority within that group.`,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the circuit group this assignment belongs to.",
			},
			"member_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The content type of the assigned member, e.g. `circuits.circuit` or `circuits.virtualcircuit`.",
			},
			"member_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the assigned member (circuit or virtual circuit).",
			},
			"priority": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxCircuitGroupAssignmentPriorityOptions, false),
				Description:  buildValidValueDescription(resourceNetboxCircuitGroupAssignmentPriorityOptions),
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCircuitGroupAssignmentCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableCircuitGroupAssignment{}

	groupID := int64(d.Get("group_id").(int))
	data.Group = &groupID

	memberID := int64(d.Get("member_id").(int))
	data.MemberID = &memberID

	memberType := d.Get("member_type").(string)
	data.MemberType = &memberType

	if priority, ok := d.GetOk("priority"); ok {
		priorityStr := priority.(string)
		data.Priority = &priorityStr
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := circuits.NewCircuitsCircuitGroupAssignmentsCreateParams().WithData(&data)

	res, err := api.Circuits.CircuitsCircuitGroupAssignmentsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCircuitGroupAssignmentRead(d, m)
}

func resourceNetboxCircuitGroupAssignmentRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitGroupAssignmentsReadParams().WithID(id)

	res, err := api.Circuits.CircuitsCircuitGroupAssignmentsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsCircuitGroupAssignmentsReadDefault); ok {
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

	if assignment.Group != nil {
		d.Set("group_id", assignment.Group.ID)
	}

	d.Set("member_id", assignment.MemberID)
	d.Set("member_type", assignment.MemberType)

	if assignment.Priority != nil {
		d.Set("priority", assignment.Priority.Value)
	} else {
		d.Set("priority", nil)
	}

	api.readTags(d, assignment.Tags)

	return nil
}

func resourceNetboxCircuitGroupAssignmentUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableCircuitGroupAssignment{}

	groupID := int64(d.Get("group_id").(int))
	data.Group = &groupID

	memberID := int64(d.Get("member_id").(int))
	data.MemberID = &memberID

	memberType := d.Get("member_type").(string)
	data.MemberType = &memberType

	if priority, ok := d.GetOk("priority"); ok {
		priorityStr := priority.(string)
		data.Priority = &priorityStr
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := circuits.NewCircuitsCircuitGroupAssignmentsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Circuits.CircuitsCircuitGroupAssignmentsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCircuitGroupAssignmentRead(d, m)
}

func resourceNetboxCircuitGroupAssignmentDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := circuits.NewCircuitsCircuitGroupAssignmentsDeleteParams().WithID(id)

	_, err := api.Circuits.CircuitsCircuitGroupAssignmentsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*circuits.CircuitsCircuitGroupAssignmentsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
