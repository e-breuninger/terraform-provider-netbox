// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

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
	_ datasource.DataSource              = (*ipsecProposalDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ipsecProposalDataSource)(nil)
)

// NewIpsecProposalDataSource returns a new ipsec_proposal data source.
func NewIpsecProposalDataSource() datasource.DataSource {
	return &ipsecProposalDataSource{}
}

type ipsecProposalDataSource struct {
	client *netboxapi.Client
}

// ipsecProposalDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type ipsecProposalDataSourceModel struct {
	ipsecProposalResourceModel
	NameContains       types.String `tfsdk:"name_contains"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *ipsecProposalDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsec_proposal"
}

func (d *ipsecProposalDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:An IPSec proposal (vpn.ipsecproposal): the algorithms and lifetimes of one IPSec security association.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/vpn/ipsecproposal/):\n\n> An [IPSec](https://en.wikipedia.org/wiki/IPsec) proposal defines a set of parameters used in negotiating security associations for IPSec tunnels. IPSec proposals defined in NetBox can be referenced by [IPSec policies](https://netboxlabs.com/docs/netbox/models/vpn/ipsecpolicy/), which are in turn employed by [IPSec profiles](https://netboxlabs.com/docs/netbox/models/vpn/ipsecprofile/).",
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
			"encryption_algorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption algorithm. NetBox requires an encryption algorithm, an authentication algorithm, or both, and does not allow clearing either once set; leaving it unset keeps the current value. One of: aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, 3des-cbc, des-cbc.",
				Validators: []validator.String{
					stringvalidator.OneOf("aes-128-cbc", "aes-128-gcm", "aes-192-cbc", "aes-192-gcm", "aes-256-cbc", "aes-256-gcm", "3des-cbc", "des-cbc"),
				},
			},
			"authentication_algorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication algorithm; not needed for GCM encryption algorithms. NetBox does not allow clearing it once set; leaving it unset keeps the current value. One of: hmac-sha1, hmac-sha256, hmac-sha384, hmac-sha512, hmac-md5.",
				Validators: []validator.String{
					stringvalidator.OneOf("hmac-sha1", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-md5"),
				},
			},
			"sa_lifetime_seconds": schema.Int64Attribute{
				Computed:    true,
				Description: "Security association lifetime in seconds.",
			},
			"sa_lifetime_data": schema.Int64Attribute{
				Computed:    true,
				Description: "Security association lifetime in kilobytes.",
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

func (d *ipsecProposalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ipsecProposalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ipsecProposalDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := vpn.NewVpnIpsecProposalsListParams()
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
	if !state.EncryptionAlgorithm.IsNull() {
		v := state.EncryptionAlgorithm.ValueString()
		params.SetEncryptionAlgorithm([]string{v})
		hasInput = true
	}
	if !state.AuthenticationAlgorithm.IsNull() {
		v := state.AuthenticationAlgorithm.ValueString()
		params.SetAuthenticationAlgorithm([]string{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or encryption_algorithm and/or authentication_algorithm and/or name_contains and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_ipsec_proposal.")
		return
	}
	res, err := d.client.Vpn.VpnIpsecProposalsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_ipsec_proposal", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_ipsec_proposal",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenIpsecProposal(ctx, netboxapi.IpsecProposalResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.ipsecProposalResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
