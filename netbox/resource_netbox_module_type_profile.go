package netbox

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxModuleTypeProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxModuleTypeProfileCreate,
		ReadContext:   resourceNetboxModuleTypeProfileRead,
		UpdateContext: resourceNetboxModuleTypeProfileUpdate,
		DeleteContext: resourceNetboxModuleTypeProfileDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletypeprofile/):

> Each module type may optionally be assigned a profile according to its classification. A profile can extend module types with user-configured attributes. For example, you might want to specify the input current and voltage of a power supply, or the clock speed and number of cores for a processor.

Module type profiles were introduced in NetBox 4.3.`,
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
				Description:  "JSON schema for the attributes of module types assigned to this profile. NetBox requires the schema to define at least one property. Removing this attribute from the configuration does not clear an already-set schema.",
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					equal, _ := jsonSemanticCompare(oldValue, newValue)
					return equal
				},
				DiffSuppressOnRefresh: true,
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

func resourceNetboxModuleTypeProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	name := d.Get("name").(string)

	data := models.ModuleTypeProfile{
		Name:        &name,
		Description: d.Get("description").(string),
		Comments:    d.Get("comments").(string),
	}

	if schemaJSON, ok := d.GetOk("schema"); ok {
		var schemaObj any
		if err := json.Unmarshal([]byte(schemaJSON.(string)), &schemaObj); err != nil {
			return diag.FromErr(err)
		}
		data.Schema = schemaObj
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return diag.FromErr(err)
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimModuleTypeProfilesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimModuleTypeProfilesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxModuleTypeProfileRead(ctx, d, m)
}

func resourceNetboxModuleTypeProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimModuleTypeProfilesReadParams().WithID(id)

	res, err := api.Dcim.DcimModuleTypeProfilesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimModuleTypeProfilesReadDefault); ok {
			if errresp.Code() == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	profile := res.GetPayload()

	d.Set("name", profile.Name)
	d.Set("description", profile.Description)
	d.Set("comments", profile.Comments)

	if profile.Schema != nil {
		schemaJSON, err := json.Marshal(profile.Schema)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("schema", string(schemaJSON))
	} else {
		d.Set("schema", nil)
	}

	cf := getCustomFields(profile.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, profile.Tags)

	return nil
}

func resourceNetboxModuleTypeProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)

	data := models.ModuleTypeProfile{
		Name:        &name,
		Description: d.Get("description").(string),
		Comments:    d.Get("comments").(string),
	}

	if schemaJSON, ok := d.GetOk("schema"); ok {
		var schemaObj any
		if err := json.Unmarshal([]byte(schemaJSON.(string)), &schemaObj); err != nil {
			return diag.FromErr(err)
		}
		data.Schema = schemaObj
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return diag.FromErr(err)
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimModuleTypeProfilesUpdateParams().WithID(id).WithData(&data)
	_, err = api.Dcim.DcimModuleTypeProfilesUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxModuleTypeProfileRead(ctx, d, m)
}

func resourceNetboxModuleTypeProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}
	return nil
}
