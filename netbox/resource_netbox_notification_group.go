package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxNotificationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxNotificationGroupCreate,
		Read:   resourceNetboxNotificationGroupRead,
		Update: resourceNetboxNotificationGroupUpdate,
		Delete: resourceNetboxNotificationGroupDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/features/notifications/#notification-groups):

> A notification group defines a set of users and/or groups to be notified when an associated event is triggered. Notification groups can be referenced by event rules and subscriptions to designate recipients for a change notification.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"user_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxNotificationGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := &models.WritableNotificationGroup{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Description = getOptionalStr(d, "description", false)
	data.Groups = toInt64List(d.Get("group_ids"))
	data.Users = toInt64List(d.Get("user_ids"))

	params := extras.NewExtrasNotificationGroupsCreateParams().WithData(data)

	res, err := api.Extras.ExtrasNotificationGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxNotificationGroupRead(d, m)
}

func resourceNetboxNotificationGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasNotificationGroupsReadParams().WithID(id)

	res, err := api.Extras.ExtrasNotificationGroupsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasNotificationGroupsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	notificationGroup := res.GetPayload()
	d.Set("name", notificationGroup.Name)
	d.Set("description", notificationGroup.Description)
	d.Set("group_ids", getIDsFromNestedGroup(notificationGroup.Groups))
	d.Set("user_ids", getIDsFromNestedUser(notificationGroup.Users))

	return nil
}

func resourceNetboxNotificationGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableNotificationGroup{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Description = getOptionalStr(d, "description", true)
	data.Groups = toInt64List(d.Get("group_ids"))
	data.Users = toInt64List(d.Get("user_ids"))

	params := extras.NewExtrasNotificationGroupsUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasNotificationGroupsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxNotificationGroupRead(d, m)
}

func resourceNetboxNotificationGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasNotificationGroupsDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasNotificationGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasNotificationGroupsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}

func getIDsFromNestedUser(nestedUsers []*models.NestedUser) []int64 {
	var userIDs []int64
	for _, user := range nestedUsers {
		userIDs = append(userIDs, user.ID)
	}
	return userIDs
}
