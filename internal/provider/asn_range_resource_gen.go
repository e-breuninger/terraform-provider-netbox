// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*asnRangeResource)(nil)
	_ resource.ResourceWithConfigure   = (*asnRangeResource)(nil)
	_ resource.ResourceWithImportState = (*asnRangeResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*asnRangeResource)(nil)
)

// NewAsnRangeResource returns a new asn_range resource.
func NewAsnRangeResource() resource.Resource {
	return &asnRangeResource{}
}

type asnRangeResource struct {
	client *netboxapi.Client
}

func (r *asnRangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asn_range"
}

func (r *asnRangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox ASN range (ipam.asnrange): a contiguous block of AS numbers under one RIR.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/asnrange/):\n\n> Ranges can be defined to group [AS numbers](https://netboxlabs.com/docs/netbox/models/ipam/asn/) numerically and to facilitate their automatic provisioning. Each range must be assigned to a [RIR](https://netboxlabs.com/docs/netbox/models/ipam/rir/).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from name when not set.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"rir_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the RIR the range is allocated by.",
			},
			"start": schema.Int64Attribute{
				Required:    true,
				Description: "First AS number in the range.",
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"end": schema.Int64Attribute{
				Required:    true,
				Description: "Last AS number in the range.",
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant the range belongs to.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Optional: true,
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"asn_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of ASNs in the range.",
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Slugs of the tags assigned to the object (the provider's default_tags are added on top, see tags_all).",
			},
			"tags_all": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of all tags on the object, including the provider's default_tags.",
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *asnRangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

// ModifyPlan runs on every plan, after Terraform builds the proposed new state. It is the
// resource-level hook for changing the plan itself: filling computed attributes already known at
// plan time (avoiding "known after apply" noise), or vetoing the plan with diagnostics. It is
// generated when something needs it, each part only where the spec asks for it: a backend plan
// fragment, the companion file's modify_plan hook, the sync that keeps attributes marked as
// deprecated equal to their canonical ones and the stage that decides the volatile attributes.
func (r *asnRangeResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var config, plan asnRangeResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.Slug.IsNull() && !plan.Name.IsUnknown() && !plan.Name.IsNull() {
			plan.Slug = types.StringValue(netboxapi.Slugify(plan.Name.ValueString()))
		}
		if r.client != nil && !plan.Tags.IsUnknown() {
			configured := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
			plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, netboxapi.MergeTags(configured, r.client.DefaultTags), false, &resp.Diagnostics)
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}()
	if resp.Diagnostics.HasError() {
		return
	}
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state asnRangeResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		if config.AsnCount.IsNull() {
			plan.AsnCount = state.AsnCount
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		if config.AsnCount.IsNull() {
			plan.AsnCount = types.Int64Unknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *asnRangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data asnRangeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandAsnRange(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_asn_range custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Ipam.IpamAsnRangesCreateContext(ctx, ipam.NewIpamAsnRangesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_asn_range", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenAsnRange(ctx, netboxapi.AsnRangeResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *asnRangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data asnRangeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Ipam.IpamAsnRangesRetrieveContext(ctx, ipam.NewIpamAsnRangesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_asn_range", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenAsnRange(ctx, netboxapi.AsnRangeResponseDTOFromGoNetbox(res.Payload), state)...)
	allTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	state.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	state.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	state.CustomFields = keepPriorCustomFields(priorCustomFields, state.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *asnRangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data asnRangeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandAsnRange(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior asnRangeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_asn_range custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Ipam.IpamAsnRangesUpdateContext(ctx, ipam.NewIpamAsnRangesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_asn_range", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenAsnRange(ctx, netboxapi.AsnRangeResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *asnRangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data asnRangeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Ipam.IpamAsnRangesDestroyContext(ctx, ipam.NewIpamAsnRangesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_asn_range", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *asnRangeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Expected an integer id, got %q.", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
