package netbox

import (
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxEventRuleActionTypeOptions = []string{"webhook"}

func resourceNetboxEventRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxEventRuleCreate,
		Read:   resourceNetboxEventRuleRead,
		Update: resourceNetboxEventRuleUpdate,
		Delete: resourceNetboxEventRuleDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/features/event-rules/):

> NetBox can be configured via Event Rules to transmit outgoing webhooks to remote systems in response to internal object changes. The receiver can act on the data in these webhook messages to perform related tasks.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_types": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"trigger_on_create": {
				Type:         schema.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"trigger_on_create", "trigger_on_update", "trigger_on_delete", "trigger_on_job_start", "trigger_on_job_end"},
			},
			"trigger_on_update": {
				Type:         schema.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"trigger_on_create", "trigger_on_update", "trigger_on_delete", "trigger_on_job_start", "trigger_on_job_end"},
			},
			"trigger_on_delete": {
				Type:         schema.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"trigger_on_create", "trigger_on_update", "trigger_on_delete", "trigger_on_job_start", "trigger_on_job_end"},
			},
			"trigger_on_job_start": {
				Type:         schema.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"trigger_on_create", "trigger_on_update", "trigger_on_delete", "trigger_on_job_start", "trigger_on_job_end"},
			},
			"trigger_on_job_end": {
				Type:         schema.TypeBool,
				Optional:     true,
				AtLeastOneOf: []string{"trigger_on_create", "trigger_on_update", "trigger_on_delete", "trigger_on_job_start", "trigger_on_job_end"},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"conditions": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					equal, _ := jsonSemanticCompare(oldValue, newValue)
					return equal
				},
				DiffSuppressOnRefresh: true,
			},
			"action_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxEventRuleActionTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxEventRuleActionTypeOptions),
			},
			"action_object_id": {
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

func resourceNetboxEventRuleCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := &models.WritableEventRule{}

	name := d.Get("name").(string)
	data.Name = &name
	actionType := d.Get("action_type").(string)
	data.ActionType = actionType
	data.Description = getOptionalStr(d, "description", false)

	// Currently, we just support the webhook action type
	data.ActionObjectType = strToPtr("extras.webhook")

	triggerOnCreate := d.Get("trigger_on_create").(bool)
	data.TypeCreate = triggerOnCreate
	triggerOnUpdate := d.Get("trigger_on_update").(bool)
	data.TypeUpdate = triggerOnUpdate
	triggerOnDelete := d.Get("trigger_on_delete").(bool)
	data.TypeDelete = triggerOnDelete
	triggerOnJobStart := d.Get("trigger_on_job_start").(bool)
	data.TypeJobStart = triggerOnJobStart
	triggerOnJobEnd := d.Get("trigger_on_job_end").(bool)
	data.TypeJobEnd = triggerOnJobEnd
	enabled := d.Get("enabled").(bool)
	data.Enabled = enabled
	data.ActionObjectID = getOptionalInt(d, "action_object_id")

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags

	ctypes := d.Get("content_types").(*schema.Set).List()
	objectTypes := make([]string, 0, len(ctypes))
	for _, contentType := range d.Get("content_types").(*schema.Set).List() {
		objectTypes = append(objectTypes, contentType.(string))
	}
	data.ObjectTypes = objectTypes

	if conditionsData, ok := d.GetOk("conditions"); ok {
		var conditions any
		err := json.Unmarshal([]byte(conditionsData.(string)), &conditions)
		if err != nil {
			return err
		}
		data.Conditions = conditions
	}

	params := extras.NewExtrasEventRulesCreateParams().WithData(data)

	res, err := api.Extras.ExtrasEventRulesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxEventRuleRead(d, m)
}

func resourceNetboxEventRuleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasEventRulesReadParams().WithID(id)

	res, err := api.Extras.ExtrasEventRulesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasEventRulesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	eventRule := res.GetPayload()
	d.Set("name", eventRule.Name)
	d.Set("description", eventRule.Description)
	d.Set("action_type", eventRule.ActionType.Value)
	d.Set("content_types", eventRule.ObjectTypes)

	d.Set("trigger_on_create", eventRule.TypeCreate)
	d.Set("trigger_on_update", eventRule.TypeUpdate)
	d.Set("trigger_on_delete", eventRule.TypeDelete)
	d.Set("trigger_on_job_start", eventRule.TypeJobStart)
	d.Set("trigger_on_job_end", eventRule.TypeJobEnd)
	d.Set("enabled", eventRule.Enabled)
	d.Set("action_object_id", eventRule.ActionObjectID)

	if eventRule.Conditions != nil {
		conditions, err := json.Marshal(eventRule.Conditions)
		if err != nil {
			return err
		}
		d.Set("conditions", string(conditions))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(eventRule.Tags))

	return nil
}

func resourceNetboxEventRuleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableEventRule{}

	name := d.Get("name").(string)
	data.Name = &name
	actionType := d.Get("action_type").(string)
	data.ActionType = actionType
	data.Description = getOptionalStr(d, "description", true)

	// Currently, we just support the webhook action type
	data.ActionObjectType = strToPtr("extras.webhook")

	triggerOnCreate := d.Get("trigger_on_create").(bool)
	data.TypeCreate = triggerOnCreate
	triggerOnUpdate := d.Get("trigger_on_update").(bool)
	data.TypeUpdate = triggerOnUpdate
	triggerOnDelete := d.Get("trigger_on_delete").(bool)
	data.TypeDelete = triggerOnDelete
	triggerOnJobStart := d.Get("trigger_on_job_start").(bool)
	data.TypeJobStart = triggerOnJobStart
	triggerOnJobEnd := d.Get("trigger_on_job_end").(bool)
	data.TypeJobEnd = triggerOnJobEnd
	enabled := d.Get("enabled").(bool)
	data.Enabled = enabled
	data.ActionObjectID = getOptionalInt(d, "action_object_id")

	if conditionsData, ok := d.GetOk("conditions"); ok {
		var conditions any
		err := json.Unmarshal([]byte(conditionsData.(string)), &conditions)
		if err != nil {
			return err
		}
		data.Conditions = conditions
	}

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags

	ctypes := d.Get("content_types").(*schema.Set).List()
	objectTypes := make([]string, 0, len(ctypes))
	for _, contentType := range d.Get("content_types").(*schema.Set).List() {
		objectTypes = append(objectTypes, contentType.(string))
	}
	data.ObjectTypes = objectTypes

	params := extras.NewExtrasEventRulesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasEventRulesUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxEventRuleRead(d, m)
}

func resourceNetboxEventRuleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasEventRulesDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasEventRulesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasEventRulesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
