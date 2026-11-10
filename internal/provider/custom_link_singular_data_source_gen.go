// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*customLinkDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customLinkDataSource)(nil)
)

// NewCustomLinkDataSource returns a new custom_link data source.
func NewCustomLinkDataSource() datasource.DataSource {
	return &customLinkDataSource{}
}

type customLinkDataSource struct {
	client *netboxapi.Client
}

// customLinkDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type customLinkDataSourceModel struct {
	customLinkResourceModel
	ObjectType     types.String `tfsdk:"object_type"`
	ObjectTypeID   types.Int64  `tfsdk:"object_type_id"`
	OwnerName      types.String `tfsdk:"owner_name"`
	OwnerGroupName types.String `tfsdk:"owner_group_name"`
	OwnerGroupID   types.Int64  `tfsdk:"owner_group_id"`
	Q              types.String `tfsdk:"q"`
}

func (d *customLinkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_link"
}

func (d *customLinkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A custom link rendered on the pages of the object types it applies to (extras.customlink).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/extras/customlink/):\n\n> Users can add custom links to object views in NetBox to reference external resources. For example, you might create a custom link for devices pointing to a monitoring system. See the [custom links documentation](https://netboxlabs.com/docs/netbox/customization/custom-links/) for more information.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"object_types": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Object types the link is shown on, as app-labelled content types (dcim.device).",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the link is rendered.",
			},
			"link_text": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jinja2 template for the link text; the object is available as {{ object }}.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"link_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Jinja2 template for the link URL; the object is available as {{ object }}.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order within the group. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"group_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Links with the same group name are rendered as a dropdown.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"button_class": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Button style. One of: default, blue, indigo, purple, pink, red, orange, yellow, green, teal, cyan, gray, black, white, ghost-dark.",
				Validators: []validator.String{
					stringvalidator.OneOf("default", "blue", "indigo", "purple", "pink", "red", "orange", "yellow", "green", "teal", "cyan", "gray", "black", "white", "ghost-dark"),
				},
			},
			"new_window": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Open the link in a new browser window.",
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
			"object_type": schema.StringAttribute{
				Optional:    true,
				Description: "Object type to match.",
			},
			"object_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the object type.",
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

func (d *customLinkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *customLinkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customLinkDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasCustomLinksListParams()
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
	if !state.Enabled.IsNull() {
		v := state.Enabled.ValueBool()
		params.SetEnabled(&v)
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.ButtonClass.IsNull() {
		v := state.ButtonClass.ValueString()
		params.SetButtonClass(&v)
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.GroupName.IsNull() {
		v := state.GroupName.ValueString()
		params.SetGroupName([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.LinkText.IsNull() {
		v := state.LinkText.ValueString()
		params.SetLinkText(&v)
		hasInput = true
	}
	if !state.LinkURL.IsNull() {
		v := state.LinkURL.ValueString()
		params.SetLinkURL(&v)
		hasInput = true
	}
	if !state.NewWindow.IsNull() {
		v := state.NewWindow.ValueBool()
		params.SetNewWindow(&v)
		hasInput = true
	}
	if !state.ObjectType.IsNull() {
		v := state.ObjectType.ValueString()
		params.SetObjectType([]string{v})
		hasInput = true
	}
	if !state.ObjectTypeID.IsNull() {
		v := state.ObjectTypeID.ValueInt64()
		params.SetObjectTypeID([]int64{v})
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
	if !state.Weight.IsNull() {
		v := state.Weight.ValueInt64()
		params.SetWeight([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_custom_link.")
		return
	}
	res, err := d.client.Extras.ExtrasCustomLinksListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_custom_link", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_custom_link",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCustomLink(ctx, netboxapi.CustomLinkResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.customLinkResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
