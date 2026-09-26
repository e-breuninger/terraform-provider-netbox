// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*configContextResource)(nil)
	_ resource.ResourceWithConfigure   = (*configContextResource)(nil)
	_ resource.ResourceWithImportState = (*configContextResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*configContextResource)(nil)
)

// NewConfigContextResource returns a new config_context resource.
func NewConfigContextResource() resource.Resource {
	return &configContextResource{}
}

type configContextResource struct {
	client *netboxapi.Client
}

func (r *configContextResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_context"
}

func (r *configContextResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox config context (extras.configcontext): JSON data merged into the rendered context of the devices and virtual machines its scope selects.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/extras/configcontext/):\n\n> Context data is made available to devices and/or virtual machines based on their relationships to other objects in NetBox. For example, context data can be associated only with devices assigned to a particular site, or only to virtual machines in a certain cluster.",
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
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Merge order; higher weights are applied later and win.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"is_active": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the context is applied.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"data": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Required:    true,
				Description: "The context data as JSON text (use jsonencode()).",
			},
			"profile_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the config context profile whose schema the data must satisfy.",
			},
			"region_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the regions the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"regions": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use region_ids instead.",
				Description:        "Deprecated alias of region_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"site_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the site groups the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"site_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use site_group_ids instead.",
				Description:        "Deprecated alias of site_group_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"site_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the sites the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"sites": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use site_ids instead.",
				Description:        "Deprecated alias of site_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"location_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the locations the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"locations": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use location_ids instead.",
				Description:        "Deprecated alias of location_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"device_type_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the device types the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"device_types": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use device_type_ids instead.",
				Description:        "Deprecated alias of device_type_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"device_role_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the device roles the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"roles": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use device_role_ids instead.",
				Description:        "Deprecated alias of device_role_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"platform_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the platforms the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"platforms": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use platform_ids instead.",
				Description:        "Deprecated alias of platform_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_type_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the cluster types the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_types": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use cluster_type_ids instead.",
				Description:        "Deprecated alias of cluster_type_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the cluster groups the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use cluster_group_ids instead.",
				Description:        "Deprecated alias of cluster_group_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the clusters the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"clusters": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use cluster_ids instead.",
				Description:        "Deprecated alias of cluster_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"tenant_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the tenant groups the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"tenant_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use tenant_group_ids instead.",
				Description:        "Deprecated alias of tenant_group_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"tenant_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the tenants the context applies to; empty means all.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"tenants": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use tenant_ids instead.",
				Description:        "Deprecated alias of tenant_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"tag_slugs": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Slugs of the tags an object must carry for the context to apply; empty means any. This is a scope, not the context's own tags.",
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
		},
	}
}

func (r *configContextResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *configContextResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Alias sync, after the backend fragment and the companion hook (both rewrite resp.Plan
	// wholesale from req.Plan, so an earlier sync would be clobbered): fold each configured
	// deprecated alias into its canonical attribute and keep the two equal in the plan.
	if !resp.Diagnostics.HasError() && !resp.Plan.Raw.IsNull() {
		var config, plan configContextResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		switch {
		case !config.Regions.IsNull() && !config.RegionIds.IsNull():
			if !config.Regions.IsUnknown() && !config.RegionIds.IsUnknown() && !config.Regions.Equal(config.RegionIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "regions is a deprecated alias of region_ids; both are set and the values differ.")
			}
		case !config.Regions.IsNull():
			plan.RegionIds = plan.Regions
		case !config.RegionIds.IsNull():
			plan.Regions = plan.RegionIds
		default:
			// region_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.RegionIds = types.SetNull(types.Int64Type)
			plan.Regions = plan.RegionIds
		}
		switch {
		case !config.SiteGroups.IsNull() && !config.SiteGroupIds.IsNull():
			if !config.SiteGroups.IsUnknown() && !config.SiteGroupIds.IsUnknown() && !config.SiteGroups.Equal(config.SiteGroupIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "site_groups is a deprecated alias of site_group_ids; both are set and the values differ.")
			}
		case !config.SiteGroups.IsNull():
			plan.SiteGroupIds = plan.SiteGroups
		case !config.SiteGroupIds.IsNull():
			plan.SiteGroups = plan.SiteGroupIds
		default:
			// site_group_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.SiteGroupIds = types.SetNull(types.Int64Type)
			plan.SiteGroups = plan.SiteGroupIds
		}
		switch {
		case !config.Sites.IsNull() && !config.SiteIds.IsNull():
			if !config.Sites.IsUnknown() && !config.SiteIds.IsUnknown() && !config.Sites.Equal(config.SiteIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "sites is a deprecated alias of site_ids; both are set and the values differ.")
			}
		case !config.Sites.IsNull():
			plan.SiteIds = plan.Sites
		case !config.SiteIds.IsNull():
			plan.Sites = plan.SiteIds
		default:
			// site_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.SiteIds = types.SetNull(types.Int64Type)
			plan.Sites = plan.SiteIds
		}
		switch {
		case !config.Locations.IsNull() && !config.LocationIds.IsNull():
			if !config.Locations.IsUnknown() && !config.LocationIds.IsUnknown() && !config.Locations.Equal(config.LocationIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "locations is a deprecated alias of location_ids; both are set and the values differ.")
			}
		case !config.Locations.IsNull():
			plan.LocationIds = plan.Locations
		case !config.LocationIds.IsNull():
			plan.Locations = plan.LocationIds
		default:
			// location_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.LocationIds = types.SetNull(types.Int64Type)
			plan.Locations = plan.LocationIds
		}
		switch {
		case !config.DeviceTypes.IsNull() && !config.DeviceTypeIds.IsNull():
			if !config.DeviceTypes.IsUnknown() && !config.DeviceTypeIds.IsUnknown() && !config.DeviceTypes.Equal(config.DeviceTypeIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "device_types is a deprecated alias of device_type_ids; both are set and the values differ.")
			}
		case !config.DeviceTypes.IsNull():
			plan.DeviceTypeIds = plan.DeviceTypes
		case !config.DeviceTypeIds.IsNull():
			plan.DeviceTypes = plan.DeviceTypeIds
		default:
			// device_type_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.DeviceTypeIds = types.SetNull(types.Int64Type)
			plan.DeviceTypes = plan.DeviceTypeIds
		}
		switch {
		case !config.Roles.IsNull() && !config.DeviceRoleIds.IsNull():
			if !config.Roles.IsUnknown() && !config.DeviceRoleIds.IsUnknown() && !config.Roles.Equal(config.DeviceRoleIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "roles is a deprecated alias of device_role_ids; both are set and the values differ.")
			}
		case !config.Roles.IsNull():
			plan.DeviceRoleIds = plan.Roles
		case !config.DeviceRoleIds.IsNull():
			plan.Roles = plan.DeviceRoleIds
		default:
			// device_role_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.DeviceRoleIds = types.SetNull(types.Int64Type)
			plan.Roles = plan.DeviceRoleIds
		}
		switch {
		case !config.Platforms.IsNull() && !config.PlatformIds.IsNull():
			if !config.Platforms.IsUnknown() && !config.PlatformIds.IsUnknown() && !config.Platforms.Equal(config.PlatformIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "platforms is a deprecated alias of platform_ids; both are set and the values differ.")
			}
		case !config.Platforms.IsNull():
			plan.PlatformIds = plan.Platforms
		case !config.PlatformIds.IsNull():
			plan.Platforms = plan.PlatformIds
		default:
			// platform_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.PlatformIds = types.SetNull(types.Int64Type)
			plan.Platforms = plan.PlatformIds
		}
		switch {
		case !config.ClusterTypes.IsNull() && !config.ClusterTypeIds.IsNull():
			if !config.ClusterTypes.IsUnknown() && !config.ClusterTypeIds.IsUnknown() && !config.ClusterTypes.Equal(config.ClusterTypeIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "cluster_types is a deprecated alias of cluster_type_ids; both are set and the values differ.")
			}
		case !config.ClusterTypes.IsNull():
			plan.ClusterTypeIds = plan.ClusterTypes
		case !config.ClusterTypeIds.IsNull():
			plan.ClusterTypes = plan.ClusterTypeIds
		default:
			// cluster_type_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.ClusterTypeIds = types.SetNull(types.Int64Type)
			plan.ClusterTypes = plan.ClusterTypeIds
		}
		switch {
		case !config.ClusterGroups.IsNull() && !config.ClusterGroupIds.IsNull():
			if !config.ClusterGroups.IsUnknown() && !config.ClusterGroupIds.IsUnknown() && !config.ClusterGroups.Equal(config.ClusterGroupIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "cluster_groups is a deprecated alias of cluster_group_ids; both are set and the values differ.")
			}
		case !config.ClusterGroups.IsNull():
			plan.ClusterGroupIds = plan.ClusterGroups
		case !config.ClusterGroupIds.IsNull():
			plan.ClusterGroups = plan.ClusterGroupIds
		default:
			// cluster_group_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.ClusterGroupIds = types.SetNull(types.Int64Type)
			plan.ClusterGroups = plan.ClusterGroupIds
		}
		switch {
		case !config.Clusters.IsNull() && !config.ClusterIds.IsNull():
			if !config.Clusters.IsUnknown() && !config.ClusterIds.IsUnknown() && !config.Clusters.Equal(config.ClusterIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "clusters is a deprecated alias of cluster_ids; both are set and the values differ.")
			}
		case !config.Clusters.IsNull():
			plan.ClusterIds = plan.Clusters
		case !config.ClusterIds.IsNull():
			plan.Clusters = plan.ClusterIds
		default:
			// cluster_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.ClusterIds = types.SetNull(types.Int64Type)
			plan.Clusters = plan.ClusterIds
		}
		switch {
		case !config.TenantGroups.IsNull() && !config.TenantGroupIds.IsNull():
			if !config.TenantGroups.IsUnknown() && !config.TenantGroupIds.IsUnknown() && !config.TenantGroups.Equal(config.TenantGroupIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "tenant_groups is a deprecated alias of tenant_group_ids; both are set and the values differ.")
			}
		case !config.TenantGroups.IsNull():
			plan.TenantGroupIds = plan.TenantGroups
		case !config.TenantGroupIds.IsNull():
			plan.TenantGroups = plan.TenantGroupIds
		default:
			// tenant_group_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.TenantGroupIds = types.SetNull(types.Int64Type)
			plan.TenantGroups = plan.TenantGroupIds
		}
		switch {
		case !config.Tenants.IsNull() && !config.TenantIds.IsNull():
			if !config.Tenants.IsUnknown() && !config.TenantIds.IsUnknown() && !config.Tenants.Equal(config.TenantIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "tenants is a deprecated alias of tenant_ids; both are set and the values differ.")
			}
		case !config.Tenants.IsNull():
			plan.TenantIds = plan.Tenants
		case !config.TenantIds.IsNull():
			plan.Tenants = plan.TenantIds
		default:
			// tenant_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.TenantIds = types.SetNull(types.Int64Type)
			plan.Tenants = plan.TenantIds
		}
		if resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state configContextResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *configContextResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data configContextResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	requestDTO, diags := expandConfigContext(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Extras.ExtrasConfigContextsCreateContext(ctx, extras.NewExtrasConfigContextsCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_config_context", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenConfigContext(ctx, netboxapi.ConfigContextResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *configContextResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data configContextResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	res, err := r.client.Extras.ExtrasConfigContextsRetrieveContext(ctx, extras.NewExtrasConfigContextsRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_config_context", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenConfigContext(ctx, netboxapi.ConfigContextResponseDTOFromGoNetbox(res.Payload), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *configContextResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data configContextResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	requestDTO, diags := expandConfigContext(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	res, err := r.client.Extras.ExtrasConfigContextsUpdateContext(ctx, extras.NewExtrasConfigContextsUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_config_context", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenConfigContext(ctx, netboxapi.ConfigContextResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *configContextResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data configContextResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Extras.ExtrasConfigContextsDestroyContext(ctx, extras.NewExtrasConfigContextsDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_config_context", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *configContextResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
