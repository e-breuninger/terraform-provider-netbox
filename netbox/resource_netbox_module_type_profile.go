package netbox

import (
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxModuleTypeProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxModuleTypeProfileCreate,
		Read:   resourceNetboxModuleTypeProfileRead,
		Update: resourceNetboxModuleTypeProfileUpdate,
		Delete: resourceNetboxModuleTypeProfileDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletypeprofile/):

> A module type profile defines a schema which specifies a set of attributes common to a group of module types. For example, all network interface cards might share a profile which captures information such as their port count or supported speeds.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schema": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
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

func resourceNetboxModuleTypeProfileCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)

	data := &models.ModuleTypeProfile{
		Name:        &name,
		Description: getOptionalStr(d, "description", false),
		Comments:    getOptionalStr(d, "comments", false),
	}

	schemaValue, ok := d.GetOk("schema")
	if ok {
		var jsonObj interface{}
		if err := json.Unmarshal([]byte(schemaValue.(string)), &jsonObj); err == nil {
			data.Schema = jsonObj
		}
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimModuleTypeProfilesCreateParams().WithData(data)

	res, err := api.Dcim.DcimModuleTypeProfilesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxModuleTypeProfileRead(d, m)
}

func resourceNetboxModuleTypeProfileRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimModuleTypeProfilesReadParams().WithID(id)

	res, err := api.Dcim.DcimModuleTypeProfilesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimModuleTypeProfilesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	moduleTypeProfile := res.GetPayload()
	d.Set("name", moduleTypeProfile.Name)
	d.Set("description", moduleTypeProfile.Description)
	d.Set("comments", moduleTypeProfile.Comments)

	if moduleTypeProfile.Schema != nil {
		if jsonBytes, err := json.Marshal(moduleTypeProfile.Schema); err == nil {
			d.Set("schema", string(jsonBytes))
		}
	} else {
		d.Set("schema", nil)
	}

	cf := getCustomFields(moduleTypeProfile.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, moduleTypeProfile.Tags)

	return nil
}

func resourceNetboxModuleTypeProfileUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.ModuleTypeProfile{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Description = getOptionalStr(d, "description", true)
	data.Comments = getOptionalStr(d, "comments", true)

	schemaValue, ok := d.GetOk("schema")
	if ok {
		var jsonObj interface{}
		if err := json.Unmarshal([]byte(schemaValue.(string)), &jsonObj); err == nil {
			data.Schema = jsonObj
		}
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimModuleTypeProfilesPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimModuleTypeProfilesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxModuleTypeProfileRead(d, m)
}

func resourceNetboxModuleTypeProfileDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModuleTypeProfilesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimModuleTypeProfilesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimModuleTypeProfilesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
