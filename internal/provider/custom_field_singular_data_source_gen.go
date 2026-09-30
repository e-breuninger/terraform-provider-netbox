// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*customFieldDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customFieldDataSource)(nil)
)

// NewCustomFieldDataSource returns a new custom_field data source.
func NewCustomFieldDataSource() datasource.DataSource {
	return &customFieldDataSource{}
}

type customFieldDataSource struct {
	client *netboxapi.Client
}

func (d *customFieldDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_field"
}

func (d *customFieldDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox custom field definition (extras.customfield). Values are set per object via the custom_fields attribute.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/customization/custom-fields/#custom-fields):\n\n> Each model in NetBox is represented in the database as a discrete table, and each attribute of a model exists as a column within its table. For example, sites are stored in the dcim_site table, which has columns named name, facility, physical_address, and so on. As new attributes are added to objects throughout the development of NetBox, tables are expanded to include new rows.\n>\n> However, some users might want to store additional object attributes that are somewhat esoteric in nature, and that would not make sense to include in the core NetBox database schema. For instance, suppose your organization needs to associate each device with a ticket number correlating it with an internal support system record. This is certainly a legitimate use for NetBox, but it's not a common enough need to warrant including a field for every NetBox installation. Instead, you can create a custom field to hold this data.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id of the custom field.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Internal field name (also the key in custom_fields).",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 50),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9_]+$`), ""),
				},
			},
			"object_types": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Object types the field applies to, e.g. dcim.site.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Field data type. One of: text, longtext, integer, decimal, boolean, date, datetime, url, json, select, multiselect, object, multiobject.",
			},
			"related_object_type": schema.StringAttribute{
				Computed:    true,
				Description: "Content type of the related objects for object and multiobject fields, e.g. dcim.device.",
			},
			"related_object_filter": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Query parameters that narrow the objects offered for object and multiobject fields, as JSON text, e.g. jsonencode({ status = \"active\" }).",
			},
			"label": schema.StringAttribute{
				Computed:    true,
				Description: "Human-readable label.",
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"group_name": schema.StringAttribute{
				Computed:    true,
				Description: "Custom fields in the same group are shown together.",
			},
			"required": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether a value is required when creating objects.",
			},
			"unique": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the value must be unique among the objects the field is assigned to.",
			},
			"filter_logic": schema.StringAttribute{
				Computed:    true,
				Description: "How the field is matched when filtering. One of: disabled, loose, exact.",
			},
			"weight": schema.Int64Attribute{
				Computed:    true,
				Description: "Order within a group.",
			},
			"search_weight": schema.Int64Attribute{
				Computed:    true,
				Description: "Weight of the field in the global search; lower values rank higher and 0 leaves the field out.",
			},
			"ui_visible": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the field is displayed in the UI. One of: always, if-set, hidden.",
			},
			"ui_editable": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the field value can be edited in the UI. One of: yes, no, hidden.",
			},
			"is_cloneable": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the value is copied when an object is cloned.",
			},
			"validation_minimum": schema.Float64Attribute{
				Computed:    true,
				Description: "Minimum allowed value (numeric fields).",
			},
			"validation_maximum": schema.Float64Attribute{
				Computed:    true,
				Description: "Maximum allowed value (numeric fields).",
			},
			"validation_regex": schema.StringAttribute{
				Computed:    true,
				Description: "Regular expression text values must match.",
			},
			"validation_schema": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "JSON schema that values of a json field must match, as JSON text (use jsonencode()).",
			},
			"default": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Default value as JSON text, e.g. jsonencode(\"foo\") or jsonencode(5).",
			},
			"choice_set_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the choice set (select and multiselect fields).",
			},
			"owner_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *customFieldDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected data source configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = c
}

func (d *customFieldDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customFieldResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasCustomFieldsListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name to look up a netbox_custom_field.")
		return
	}
	res, err := d.client.Extras.ExtrasCustomFieldsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_custom_field", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_custom_field",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCustomField(ctx, netboxapi.CustomFieldResponseDTOFromGoNetbox(res.Payload.Results[0]), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
