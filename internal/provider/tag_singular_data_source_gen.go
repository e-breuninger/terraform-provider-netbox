// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

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
	_ datasource.DataSource              = (*tagDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*tagDataSource)(nil)
)

// NewTagDataSource returns a new tag data source.
func NewTagDataSource() datasource.DataSource {
	return &tagDataSource{}
}

type tagDataSource struct {
	client *netboxapi.Client
}

// tagDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type tagDataSourceModel struct {
	tagResourceModel
	Color           types.String `tfsdk:"color"`
	ContentType     types.String `tfsdk:"content_type"`
	ContentTypeID   types.Int64  `tfsdk:"content_type_id"`
	ForObjectTypeID types.Int64  `tfsdk:"for_object_type_id"`
	ObjectTypeName  types.Int64  `tfsdk:"object_type_name"`
	OwnerName       types.String `tfsdk:"owner_name"`
	OwnerGroupName  types.String `tfsdk:"owner_group_name"`
	OwnerGroupID    types.Int64  `tfsdk:"owner_group_id"`
	OwnerID         types.Int64  `tfsdk:"owner_id"`
	Q               types.String `tfsdk:"q"`
}

func (d *tagDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (d *tagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox tag (extras.tag). go-netbox has no Writable variant for tags.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/extras/tag/):\n> Tags are user-defined labels which can be applied to a variety of objects within NetBox. They can be used to establish dimensions of organization beyond the relationships built into NetBox. For example, you might create a tag to identify a particular ownership or condition across several types of objects.\n>\n> Each tag has a label, color, and a URL-friendly slug. For example, the slug for a tag named \"Dunder Mifflin, Inc.\" would be dunder-mifflin-inc. The slug is generated automatically and makes tags easier to work with as URL parameters. Each tag can also be assigned a description indicating its purpose.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id of the tag.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from name when not set.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-\w]+$`), ""),
				},
			},
			"color_hex": schema.StringAttribute{
				Computed:    true,
				Description: "Hex color without the leading #. Defaults to 9e9e9e.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Sort weight; tags with a lower weight are listed first. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"object_types": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Object types the tag can be assigned to, e.g. dcim.site.",
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
			"color": schema.StringAttribute{
				Optional:    true,
				Description: "Color to match.",
			},
			"content_type": schema.StringAttribute{
				Optional:    true,
				Description: "Content type to match.",
			},
			"content_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the content type.",
			},
			"for_object_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the for object type.",
			},
			"object_type_name": schema.Int64Attribute{
				Optional:    true,
				Description: "Name of the object types.",
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
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
		},
	}
}

func (d *tagDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *tagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data tagDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasTagsListParams()
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
	if !state.Slug.IsNull() {
		v := state.Slug.ValueString()
		params.SetSlug([]string{v})
		hasInput = true
	}
	if !state.Color.IsNull() {
		v := state.Color.ValueString()
		params.SetColor([]string{v})
		hasInput = true
	}
	if !state.ContentType.IsNull() {
		v := state.ContentType.ValueString()
		params.SetContentType([]string{v})
		hasInput = true
	}
	if !state.ContentTypeID.IsNull() {
		v := state.ContentTypeID.ValueInt64()
		params.SetContentTypeID([]int64{v})
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
	if !state.ForObjectTypeID.IsNull() {
		v := state.ForObjectTypeID.ValueInt64()
		params.SetForObjectTypeID([]int64{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.ObjectTypeName.IsNull() {
		v := state.ObjectTypeName.ValueInt64()
		params.SetObjectTypes([]int64{v})
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
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_tag.")
		return
	}
	res, err := d.client.Extras.ExtrasTagsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_tag", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_tag",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenTag(ctx, netboxapi.TagResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.tagResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
