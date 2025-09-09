package netbox

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxConfigTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxConfigTemplateCreate,
		ReadContext:   resourceNetboxConfigTemplateRead,
		UpdateContext: resourceNetboxConfigTemplateUpdate,
		DeleteContext: resourceNetboxConfigTemplateDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/configtemplate/):

> Configuration templates can be used to render device configurations from context data. Templates are written in the Jinja2 language and can be associated with devices roles, platforms, and/or individual devices.

> Context data is made available to devices and/or virtual machines based on their relationships to other objects in NetBox. For example, context data can be associated only with devices assigned to a particular site, or only to virtual machines in a certain cluster.`,
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
			"template_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_params": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "{}",
				ValidateFunc: validation.StringIsJSON,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxConfigTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	templateCode := d.Get("template_code").(string)

	tags, _ := getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))

	data := models.WritableConfigTemplate{
		Name:         &name,
		Description:  description,
		TemplateCode: &templateCode,
		Tags:         tags,
	}

	// Unmarshal environment_params and add it to data if valid
	environmentParamsJSON, ok := d.GetOk("environment_params")

	if ok {
		var environmentParams any
		err := json.Unmarshal([]byte(environmentParamsJSON.(string)), &environmentParams)
		if err != nil {
			return diag.FromErr(err)
		}

		data.EnvironmentParams = environmentParams
	}

	params := extras.NewExtrasConfigTemplatesCreateParams().WithData(&data)

	res, err := api.Extras.ExtrasConfigTemplatesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return diags
}

func resourceNetboxConfigTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	var diags diag.Diagnostics

	params := extras.NewExtrasConfigTemplatesReadParams().WithID(id)

	res, err := api.Extras.ExtrasConfigTemplatesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasConfigTemplatesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	tmpl := res.GetPayload()

	d.Set("name", tmpl.Name)
	d.Set("description", tmpl.Description)
	d.Set("template_code", tmpl.TemplateCode)

	if tmpl.EnvironmentParams != nil {
		environmentParamsJSON, err := json.Marshal(tmpl.EnvironmentParams)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("environment_params", string(environmentParamsJSON))
	} else {
		d.Set("environment_params", "{}")
	}

	return diags
}

func resourceNetboxConfigTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	templateCode := d.Get("template_code").(string)

	tags, _ := getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))

	data := models.WritableConfigTemplate{
		Name:         &name,
		Description:  description,
		TemplateCode: &templateCode,
		Tags:         tags,
	}

	// Unmarshal environment_params and add it to data if valid
	environmentParamsJSON, ok := d.GetOk("environment_params")

	if ok {
		var environmentParams any
		err := json.Unmarshal([]byte(environmentParamsJSON.(string)), &environmentParams)
		if err != nil {
			return diag.FromErr(err)
		}

		data.EnvironmentParams = environmentParams
	}

	params := extras.NewExtrasConfigTemplatesUpdateParams().WithID(id).WithData(&data)
	_, err := api.Extras.ExtrasConfigTemplatesUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetboxConfigTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasConfigTemplatesDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasConfigTemplatesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*extras.ExtrasConfigTemplatesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return nil
}
