// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*locationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*locationDataSource)(nil)
)

// NewLocationDataSource returns a new location data source.
func NewLocationDataSource() datasource.DataSource {
	return &locationDataSource{}
}

type locationDataSource struct {
	client *netboxapi.Client
}

// locationDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type locationDataSourceModel struct {
	locationResourceModel
	Site               types.String `tfsdk:"site"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	Ancestor           types.String `tfsdk:"ancestor"`
	AncestorID         types.String `tfsdk:"ancestor_id"`
	Contact            types.Int64  `tfsdk:"contact"`
	ContactGroupSlug   types.String `tfsdk:"contact_group_slug"`
	ContactRole        types.Int64  `tfsdk:"contact_role"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	ParentSlug         types.String `tfsdk:"parent_slug"`
	Q                  types.String `tfsdk:"q"`
	RegionSlug         types.String `tfsdk:"region_slug"`
	RegionID           types.String `tfsdk:"region_id"`
	SiteGroupSlug      types.String `tfsdk:"site_group_slug"`
	SiteGroupID        types.String `tfsdk:"site_group_id"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	TenantGroupSlug    types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String `tfsdk:"tenant_group_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *locationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_location"
}

func (d *locationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A NetBox location within a site (dcim.location).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/location/):\n\n> Racks and devices can be grouped by location within a site. A location may represent a floor, room, cage, or similar organizational unit. Locations can be nested to form a hierarchy. For example, you may have floors within a site, and rooms within a floor.\n\nEach location must have a name that is unique within its parent site and location, if any.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id of the location.",
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
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the site the location belongs to.",
			},
			"parent_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the parent location.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: planned, staging, active, decommissioning, retired.",
				Validators: []validator.String{
					stringvalidator.OneOf("planned", "staging", "active", "decommissioning", "retired"),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"facility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local facility ID or description.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
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
			"depth": schema.Int64Attribute{
				Computed:    true,
				Description: "Nesting depth within the site.",
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
			"rack_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of racks in the location.",
			},
			"device_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of devices in the location.",
			},
			"prefix_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of prefixes in the location.",
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
			"site": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site.",
			},
			"tenant_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"ancestor": schema.StringAttribute{
				Optional:    true,
				Description: "Ancestor to match.",
			},
			"ancestor_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the ancestor.",
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
			"parent_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the parent location.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
			},
			"region_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"site_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site group.",
			},
			"site_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the site group.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
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

func (d *locationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *locationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data locationDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimLocationsListParams()
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
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.ParentID.IsNull() {
		v := state.ParentID.ValueInt64()
		params.SetParentID([]int64{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.Site.IsNull() {
		v := state.Site.ValueString()
		params.SetSite([]string{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.TenantSlug.IsNull() {
		v := state.TenantSlug.ValueString()
		params.SetTenant([]string{v})
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
	if !state.Ancestor.IsNull() {
		v := state.Ancestor.ValueString()
		params.SetAncestor([]string{v})
		hasInput = true
	}
	if !state.AncestorID.IsNull() {
		v := state.AncestorID.ValueString()
		params.SetAncestorID([]string{v})
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
	if !state.Facility.IsNull() {
		v := state.Facility.ValueString()
		params.SetFacility([]string{v})
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
	if !state.ParentSlug.IsNull() {
		v := state.ParentSlug.ValueString()
		params.SetParent([]string{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueString()
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.SiteGroupSlug.IsNull() {
		v := state.SiteGroupSlug.ValueString()
		params.SetSiteGroup([]string{v})
		hasInput = true
	}
	if !state.SiteGroupID.IsNull() {
		v := state.SiteGroupID.ValueString()
		params.SetSiteGroupID([]string{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_location.")
		return
	}
	res, err := d.client.Dcim.DcimLocationsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_location", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_location",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenLocation(ctx, netboxapi.LocationResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.locationResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
