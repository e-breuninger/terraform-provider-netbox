// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*fhrpGroupDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*fhrpGroupDataSource)(nil)
)

// NewFhrpGroupDataSource returns a new fhrp_group data source.
func NewFhrpGroupDataSource() datasource.DataSource {
	return &fhrpGroupDataSource{}
}

type fhrpGroupDataSource struct {
	client *netboxapi.Client
}

// fhrpGroupDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type fhrpGroupDataSourceModel struct {
	fhrpGroupResourceModel
	NameContains       types.String `tfsdk:"name_contains"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *fhrpGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fhrp_group"
}

func (d *fhrpGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A first-hop redundancy protocol group (ipam.fhrpgroup).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/fhrpgroup/):\n\n> A first-hop redundancy protocol (FHRP) enables multiple physical interfaces to present a virtual [IP address](https://netboxlabs.com/docs/netbox/models/ipam/ipaddress/) (VIP) in a redundant manner. Examples of such protocols include:\n>\n> * [Hot Standby Router Protocol](https://en.wikipedia.org/wiki/Hot_Standby_Router_Protocol) (HSRP)\n> * [Virtual Router Redundancy Protocol](https://en.wikipedia.org/wiki/Virtual_Router_Redundancy_Protocol) (VRRP)\n> * [Common Address Redundancy Protocol](https://en.wikipedia.org/wiki/Common_Address_Redundancy_Protocol) (CARP)\n> * [Gateway Load Balancing Protocol](https://en.wikipedia.org/wiki/Gateway_Load_Balancing_Protocol) (GLBP)\n>\n> When creating a new FHRP group, the user may optionally create a VIP as well. This IP address will be automatically assigned to the new group. (Virtual IP addresses can also be assigned after the group has been created.)",
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
					stringvalidator.LengthAtMost(100),
				},
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Redundancy protocol. One of: vrrp2, vrrp3, carp, clusterxl, hsrp, glbp, other.",
				Validators: []validator.String{
					stringvalidator.OneOf("vrrp2", "vrrp3", "carp", "clusterxl", "hsrp", "glbp", "other"),
				},
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol group number (the VRRP/HSRP group id, not a reference to another object).",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"auth_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication type. One of: plaintext, md5.",
				Validators: []validator.String{
					stringvalidator.OneOf("plaintext", "md5"),
				},
			},
			"auth_key": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "Authentication key.",
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
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
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
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

func (d *fhrpGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *fhrpGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data fhrpGroupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamFhrpGroupsListParams()
	hasInput := false
	queryParams := url.Values{}
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
	if !state.Protocol.IsNull() {
		v := state.Protocol.ValueString()
		params.SetProtocol([]string{v})
		hasInput = true
	}
	if !state.GroupID.IsNull() {
		v := state.GroupID.ValueInt64()
		params.SetGroupID([]int64{v})
		hasInput = true
	}
	if !state.AuthType.IsNull() {
		v := state.AuthType.ValueString()
		params.SetAuthType([]string{v})
		hasInput = true
	}
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or protocol and/or group_id and/or auth_type and/or name_contains and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_fhrp_group.")
		return
	}
	res, err := d.client.Ipam.IpamFhrpGroupsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_fhrp_group", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_fhrp_group",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenFhrpGroup(ctx, netboxapi.FhrpGroupResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.fhrpGroupResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
