// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*contactAssignmentDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*contactAssignmentDataSource)(nil)
)

// NewContactAssignmentDataSource returns a new contact_assignment data source.
func NewContactAssignmentDataSource() datasource.DataSource {
	return &contactAssignmentDataSource{}
}

type contactAssignmentDataSource struct {
	client *netboxapi.Client
}

// contactAssignmentDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type contactAssignmentDataSourceModel struct {
	contactAssignmentResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *contactAssignmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contact_assignment"
}

func (d *contactAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Tenancy:A NetBox contact assignment (tenancy.contactassignment): a contact attached to an object in a role.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/contacts#contactassignments_1):\n\n> Much like tenancy, contact assignment enables you to track ownership of resources modeled in NetBox.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"object_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the object the contact is assigned to (dcim.site, dcim.device, tenancy.tenant, ...).",
			},
			"object_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the object the contact is assigned to.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"contact_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the contact.",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the contact role. Once set it cannot be removed, only changed: NetBox rejects a null. If you drop it from the config, keep a depends_on on the role so it is destroyed after the assignment.",
			},
			"priority": schema.StringAttribute{
				Computed:    true,
				Description: "Priority of the contact for this object. One of: primary, secondary, tertiary, inactive.",
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
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of the tags assigned to the object (the provider's default_tags are added on top, see tags_all).",
			},
			"tags_all": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of all tags on the object, including the provider's default_tags.",
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *contactAssignmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *contactAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data contactAssignmentDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := tenancy.NewTenancyContactAssignmentsListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.ObjectType.IsNull() {
		v := state.ObjectType.ValueString()
		params.SetObjectType([]string{v})
		hasInput = true
	}
	if !state.ObjectID.IsNull() {
		v := state.ObjectID.ValueInt64()
		params.SetObjectID([]int64{v})
		hasInput = true
	}
	if !state.ContactID.IsNull() {
		v := state.ContactID.ValueInt64()
		params.SetContactID([]int64{v})
		hasInput = true
	}
	if !state.RoleID.IsNull() {
		v := state.RoleID.ValueInt64()
		params.SetRoleID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or object_type and/or object_id and/or contact_id and/or role_id and/or tag_slug and/or custom_field_filters to look up a netbox_contact_assignment.")
		return
	}
	res, err := d.client.Tenancy.TenancyContactAssignmentsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_contact_assignment", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_contact_assignment",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenContactAssignment(ctx, netboxapi.ContactAssignmentResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.contactAssignmentResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
