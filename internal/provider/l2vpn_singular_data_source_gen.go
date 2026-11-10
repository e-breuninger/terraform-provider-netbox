// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*l2vpnDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*l2vpnDataSource)(nil)
)

// NewL2vpnDataSource returns a new l2vpn data source.
func NewL2vpnDataSource() datasource.DataSource {
	return &l2vpnDataSource{}
}

type l2vpnDataSource struct {
	client *netboxapi.Client
}

// l2vpnDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type l2vpnDataSourceModel struct {
	l2vpnResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	Contact            types.Int64  `tfsdk:"contact"`
	ContactGroupSlug   types.String `tfsdk:"contact_group_slug"`
	ContactRole        types.Int64  `tfsdk:"contact_role"`
	ExportTargetName   types.String `tfsdk:"export_target_name"`
	ExportTargetID     types.Int64  `tfsdk:"export_target_id"`
	ImportTargetName   types.String `tfsdk:"import_target_name"`
	ImportTargetID     types.Int64  `tfsdk:"import_target_id"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String `tfsdk:"tenant_group_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *l2vpnDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2vpn"
}

func (d *l2vpnDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:A layer 2 VPN (vpn.l2vpn).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/vpn/l2vpn/):\n\n> A L2VPN object is NetBox is a representation of a layer 2 bridge technology such as VXLAN, VPLS, or EPL. Each L2VPN can be identified by name as well as by an optional unique identifier (VNI would be an example). Once created, L2VPNs can be terminated to [interfaces](https://netboxlabs.com/docs/netbox/models/dcim/interface/) and [VLANs](https://netboxlabs.com/docs/netbox/models/ipam/vlan/).",
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
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from name when not set.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "L2VPN technology. One of: vpws, vpls, vxlan, vxlan-evpn, mpls-evpn, pbb-evpn, epl, evpl, ep-lan, evp-lan, ep-tree, evp-tree, evpn-vpws, spb.",
				Validators: []validator.String{
					stringvalidator.OneOf("vpws", "vpls", "vxlan", "vxlan-evpn", "mpls-evpn", "pbb-evpn", "epl", "evpl", "ep-lan", "evp-lan", "ep-tree", "evp-tree", "evpn-vpws", "spb"),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: active, planned, decommissioning.",
				Validators: []validator.String{
					stringvalidator.OneOf("active", "planned", "decommissioning"),
				},
			},
			"identifier": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Numeric identifier (e.g. the VNI or VPLS id), not the NetBox id. Must be a whole number.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"import_target_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of imported route targets.",
			},
			"export_target_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of exported route targets.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
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
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"contact": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact.",
			},
			"contact_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the contact group.",
			},
			"contact_role": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact role.",
			},
			"export_target_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the export target.",
			},
			"export_target_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the export target.",
			},
			"import_target_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the import target.",
			},
			"import_target_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the import target.",
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
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"tenant_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant.",
			},
			"tenant_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant group.",
			},
			"tenant_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the tenant group.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *l2vpnDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *l2vpnDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data l2vpnDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := vpn.NewVpnL2vpnsListParams()
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
	if !state.Slug.IsNull() {
		v := state.Slug.ValueString()
		params.SetSlug([]string{v})
		hasInput = true
	}
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.Identifier.IsNull() {
		v := state.Identifier.ValueInt64()
		params.SetIdentifier([]int64{v})
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
	if !state.Contact.IsNull() {
		v := state.Contact.ValueInt64()
		params.SetContact([]int64{v})
		hasInput = true
	}
	if !state.ContactGroupSlug.IsNull() {
		v := state.ContactGroupSlug.ValueString()
		params.SetContactGroup([]string{v})
		hasInput = true
	}
	if !state.ContactRole.IsNull() {
		v := state.ContactRole.ValueInt64()
		params.SetContactRole([]int64{v})
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
	if !state.ExportTargetName.IsNull() {
		v := state.ExportTargetName.ValueString()
		params.SetExportTarget([]string{v})
		hasInput = true
	}
	if !state.ExportTargetID.IsNull() {
		v := state.ExportTargetID.ValueInt64()
		params.SetExportTargetID([]int64{v})
		hasInput = true
	}
	if !state.ImportTargetName.IsNull() {
		v := state.ImportTargetName.ValueString()
		params.SetImportTarget([]string{v})
		hasInput = true
	}
	if !state.ImportTargetID.IsNull() {
		v := state.ImportTargetID.ValueInt64()
		params.SetImportTargetID([]int64{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
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
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.TenantSlug.IsNull() {
		v := state.TenantSlug.ValueString()
		params.SetTenant([]string{v})
		hasInput = true
	}
	if !state.TenantGroupSlug.IsNull() {
		v := state.TenantGroupSlug.ValueString()
		params.SetTenantGroup([]string{v})
		hasInput = true
	}
	if !state.TenantGroupID.IsNull() {
		v := state.TenantGroupID.ValueString()
		params.SetTenantGroupID([]string{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_l2vpn.")
		return
	}
	res, err := d.client.Vpn.VpnL2vpnsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_l2vpn", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_l2vpn",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenL2vpn(ctx, netboxapi.L2vpnResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.l2vpnResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
