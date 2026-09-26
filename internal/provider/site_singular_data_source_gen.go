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
	_ datasource.DataSource              = (*siteDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*siteDataSource)(nil)
)

// NewSiteDataSource returns a new site data source.
func NewSiteDataSource() datasource.DataSource {
	return &siteDataSource{}
}

type siteDataSource struct {
	client *netboxapi.Client
}

// siteDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type siteDataSourceModel struct {
	siteResourceModel
	NameContains       types.String `tfsdk:"name_contains"`
	IDNot              types.Int64  `tfsdk:"id_not"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *siteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site"
}

func (d *siteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A NetBox site (dcim.site).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/site/):\n\n> How you choose to employ sites when modeling your network may vary depending on the nature of your organization, but generally a site will equate to a building or campus. For example, a chain of banks might create a site to represent each of its branches, a site for its corporate headquarters, and two additional sites for its presence in two colocation facilities.\n>\n> Each site must be assigned a unique name and may optionally be assigned to a region and/or tenant.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id of the site.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Full name of the site.",
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
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: planned, staging, active, decommissioning, retired.",
				Validators: []validator.String{
					stringvalidator.OneOf("planned", "staging", "active", "decommissioning", "retired"),
				},
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"facility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local facility ID or description.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"physical_address": schema.StringAttribute{
				Computed: true,
			},
			"shipping_address": schema.StringAttribute{
				Computed: true,
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"timezone": schema.StringAttribute{
				Computed:    true,
				Description: "Time zone name, e.g. Europe/Berlin.",
			},
			"latitude": schema.Float64Attribute{
				Computed:    true,
				Description: "GPS coordinate in decimal format (xx.yyyyyy).",
			},
			"longitude": schema.Float64Attribute{
				Computed:    true,
				Description: "GPS coordinate in decimal format (xx.yyyyyy).",
			},
			"region_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the region.",
			},
			"tenant_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"group_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the site group.",
			},
			"asn_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of ASNs assigned to the site.",
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
			"circuit_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of circuits at the site.",
			},
			"device_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of devices at the site.",
			},
			"prefix_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of prefixes at the site.",
			},
			"rack_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of racks at the site.",
			},
			"virtual_machine_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of virtual machines at the site.",
			},
			"vlan_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of VLANs at the site.",
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
			"id_not": schema.Int64Attribute{
				Optional:    true,
				Description: "Exclude this id.",
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

func (d *siteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *siteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data siteDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimSitesListParams()
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
	if !state.Facility.IsNull() {
		v := state.Facility.ValueString()
		params.SetFacility([]string{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
		hasInput = true
	}
	if !state.IDNot.IsNull() {
		v := state.IDNot.ValueInt64()
		params.SetIDn([]int64{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or slug and/or facility and/or status and/or name_contains and/or id_not and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_site.")
		return
	}
	res, err := d.client.Dcim.DcimSitesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_site", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_site",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenSite(ctx, netboxapi.SiteResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.siteResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
