// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*customFieldChoiceSetDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customFieldChoiceSetDataSource)(nil)
)

// NewCustomFieldChoiceSetDataSource returns a new custom_field_choice_set data source.
func NewCustomFieldChoiceSetDataSource() datasource.DataSource {
	return &customFieldChoiceSetDataSource{}
}

type customFieldChoiceSetDataSource struct {
	client *netboxapi.Client
}

// customFieldChoiceSetDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type customFieldChoiceSetDataSourceModel struct {
	customFieldChoiceSetResourceModel
	Choice         types.String `tfsdk:"choice"`
	ChoiceColor    types.String `tfsdk:"choice_color"`
	OwnerName      types.String `tfsdk:"owner_name"`
	OwnerGroupName types.String `tfsdk:"owner_group_name"`
	OwnerGroupID   types.Int64  `tfsdk:"owner_group_id"`
	Q              types.String `tfsdk:"q"`
}

func (d *customFieldChoiceSetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_field_choice_set"
}

func (d *customFieldChoiceSetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox custom field choice set (extras.customfieldchoiceset).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/extras/customfieldchoiceset/):\n\nSingle- and multi-selection custom fields must define a set of valid choices from which the user may choose when defining the field value. These choices are defined as sets that may be reused among multiple custom fields.\n\nA choice set must define a base choice set and/or a set of arbitrary extra choices.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id of the choice set.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"base_choices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Predefined set of choices to start from. One of: IATA, ISO_3166, UN_LOCODE.",
				Validators: []validator.String{
					stringvalidator.OneOf("IATA", "ISO_3166", "UN_LOCODE"),
				},
			},
			"extra_choices": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value": schema.StringAttribute{
							Computed: true,
						},
						"label": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Computed:    true,
				Description: "Additional choices as value/label pairs.",
			},
			"choice_colors": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Colors of choices, by choice value. One of blue, indigo, purple, pink, red, orange, yellow, green, teal, cyan, gray, black or white each.",
			},
			"order_alphabetically": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sort choices alphabetically instead of by definition order.",
			},
			"choices_count": schema.Int64Attribute{
				Computed: true,
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"choice": schema.StringAttribute{
				Optional:    true,
				Description: "Choice to match.",
			},
			"choice_color": schema.StringAttribute{
				Optional:    true,
				Description: "Choice colors.",
			},
			"owner_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner.",
			},
			"owner_group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner group.",
			},
			"owner_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner group.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
		},
	}
}

func (d *customFieldChoiceSetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *customFieldChoiceSetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customFieldChoiceSetDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasCustomFieldChoiceSetsListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.BaseChoices.IsNull() {
		v := state.BaseChoices.ValueString()
		params.SetBaseChoices(&v)
		hasInput = true
	}
	if !state.Choice.IsNull() {
		v := state.Choice.ValueString()
		params.SetChoice([]string{v})
		hasInput = true
	}
	if !state.ChoiceColor.IsNull() {
		v := state.ChoiceColor.ValueString()
		params.SetChoiceColors([]string{v})
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.OrderAlphabetically.IsNull() {
		v := state.OrderAlphabetically.ValueBool()
		params.SetOrderAlphabetically(&v)
		hasInput = true
	}
	if !state.OwnerName.IsNull() {
		v := state.OwnerName.ValueString()
		params.SetOwner([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupName.IsNull() {
		v := state.OwnerGroupName.ValueString()
		params.SetOwnerGroup([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupID.IsNull() {
		v := state.OwnerGroupID.ValueInt64()
		params.SetOwnerGroupID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_custom_field_choice_set.")
		return
	}
	res, err := d.client.Extras.ExtrasCustomFieldChoiceSetsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_custom_field_choice_set", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_custom_field_choice_set",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCustomFieldChoiceSet(ctx, netboxapi.CustomFieldChoiceSetResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.customFieldChoiceSetResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
