package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxTokenCreate,
		ReadContext:   resourceNetboxTokenRead,
		UpdateContext: resourceNetboxTokenUpdate,
		DeleteContext: resourceNetboxTokenDelete,

		Description: `:meta:subcategory:Authentication:From the [official documentation](https://docs.netbox.dev/en/stable/rest-api/authentication/#tokens):

> A token is a unique identifier mapped to a NetBox user account. Each user may have one or more tokens which he or she can use for authentication when making REST API requests. To create a token, navigate to the API tokens page under your user profile.`,

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"key": {
				Type:         schema.TypeString,
				Sensitive:    true,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(40, 256),
				Deprecated:   "NetBox 4.6 made `key` read-only. Set `token` instead; this value is no longer sent.",
			},
			"token": {
				Type:         schema.TypeString,
				Sensitive:    true,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(40, 256),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"allowed_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsCIDR,
				},
				Deprecated: "NetBox 4.6 removed `allowed_ips` from the API. This value is no longer sent and has no effect.",
			},
			"write_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"last_used": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	data := models.WritableToken{}

	userid := int64(d.Get("user_id").(int))

	data.User = &userid

	// NetBox 4.6 made `key` read-only and dropped `allowed_ips`; `token` is
	// the write field. `key` is still read back for pre-4.6 servers.
	if token, ok := d.GetOk("token"); ok {
		data.Token = token.(string)
	}

	data.Enabled = d.Get("enabled").(bool)
	data.WriteEnabled = d.Get("write_enabled").(bool)
	data.Description = d.Get("description").(string)

	if expiresRaw, ok := d.GetOk("expires"); ok {
		expires, err := strfmt.ParseDateTime(expiresRaw.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		data.Expires = &expires
	}

	params := users.NewUsersTokensCreateParams().WithData(&data)
	res, err := api.Users.UsersTokensCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTokenUpdate(ctx, d, m)
}

func resourceNetboxTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersTokensReadParams().WithID(id)

	res, err := api.Users.UsersTokensRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersTokensReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	token := res.GetPayload()

	if token.User != nil {
		d.Set("user_id", token.User.ID)
	}

	// Since NetBox 4.3.0, ALLOW_TOKEN_RETRIEVAL is disabled by default
	// This means we will usually not get a Key value from the API
	if token.Key != nil && *token.Key != "" {
		d.Set("key", *token.Key)
	}
	if token.Token != "" {
		d.Set("token", token.Token)
	}
	d.Set("enabled", token.Enabled)
	d.Set("last_used", token.LastUsed)
	if token.Expires != nil {
		d.Set("expires", token.Expires.String())
	}
	d.Set("write_enabled", token.WriteEnabled)
	d.Set("description", token.Description)

	return nil
}

func resourceNetboxTokenUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableToken{}

	userid := int64(d.Get("user_id").(int))

	data.User = &userid

	if token, ok := d.GetOk("token"); ok {
		data.Token = token.(string)
	}

	data.Enabled = d.Get("enabled").(bool)
	data.WriteEnabled = d.Get("write_enabled").(bool)
	data.Description = d.Get("description").(string)

	if expiresRaw, ok := d.GetOk("expires"); ok {
		expires, err := strfmt.ParseDateTime(expiresRaw.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		data.Expires = &expires
	}

	params := users.NewUsersTokensUpdateParams().WithID(id).WithData(&data)
	_, err := api.Users.UsersTokensUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceNetboxTokenRead(ctx, d, m)
}

func resourceNetboxTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersTokensDeleteParams().WithID(id)
	_, err := api.Users.UsersTokensDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersTokensDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
