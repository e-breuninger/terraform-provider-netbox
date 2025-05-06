package provider

import (
	"context"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider/helpers"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
	"regexp"
)

var _ resource.Resource = (*regionResource)(nil)

func NewRegionResource() resource.Resource {
	return &regionResource{}
}

type regionResource struct {
	NetboxResource
}

func (r *regionResource) writeAPI(ctx context.Context, data *models.RegionTerraformModel) (*netbox.WritableRegionRequest, diag.Diagnostics) {
	regionRequest := netbox.NewWritableRegionRequestWithDefaults()
	regionRequest.Name = data.Name.ValueString()
	regionRequest.Slug = data.Slug.ValueString()
	regionRequest.Description = data.Description.ValueStringPointer()

	var tagList []netbox.NestedTagRequest
	tagList, diags := helpers.WriteTagsToApi(ctx, *r.provider.client, data.Tags)

	if diags.HasError() {
		return nil, diags
	}
	regionRequest.Tags = tagList

	regionRequest.CustomFields = helpers.ReadCustomFieldsFromTerraform(data.CustomFields)

	if !data.Parent.IsUnknown() {
		regionRequest.Parent = *netbox.NewNullableInt32(data.Parent.ValueInt32Pointer())
	}

	return regionRequest, nil
}

func (r *regionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

func (r *regionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Attributes: map[string]schema.Attribute{
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"id": schema.Int32Attribute{
				Computed:            true,
				Description:         "A unique integer value identifying this region.",
				MarkdownDescription: "A unique integer value identifying this region.",
				PlanModifiers: []planmodifier.Int32{
					int32planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"parent": schema.Int32Attribute{
				Optional: true,
				Computed: true,
			},
			"slug": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile("^[-a-zA-Z0-9_]+$"), ""),
				},
			},
			"tags": schema.ListAttribute{
				ElementType: types.Int32Type,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (r *regionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.RegionTerraformModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	regionRequest, diags := r.writeAPI(ctx, &data)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	api_res, _, err := r.provider.client.DcimAPI.DcimRegionsCreate(ctx).
		WritableRegionRequest(*regionRequest).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error while creating the region.",
			err.Error(),
		)
		return
	}

	// Read API call logic
	errors := data.ReadAPI(ctx, api_res)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *regionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.RegionTerraformModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	region, httpCode, err := r.provider.client.DcimAPI.DcimRegionsRetrieve(ctx, data.Id.ValueInt32()).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to retrieve region value.",
				err.Error(),
			)
			return
		}
	}

	// Read API call logic
	errors := data.ReadAPI(ctx, region)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *regionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.RegionTerraformModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	regionRequest, diags := r.writeAPI(ctx, &data)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	region, httpCode, err := r.provider.client.DcimAPI.DcimRegionsUpdate(ctx, data.Id.ValueInt32()).WritableRegionRequest(*regionRequest).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.Diagnostics.AddError(
				"Region no longer exists",
				err.Error())
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to update Region.",
				err.Error(),
			)
			return
		}
	}

	errors := data.ReadAPI(ctx, region)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *regionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.RegionTerraformModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	httpCode, err := r.provider.client.DcimAPI.DcimRegionsDestroy(ctx, data.Id.ValueInt32()).Execute()
	if err != nil {
		if httpCode != nil && httpCode.StatusCode != 404 {
			resp.Diagnostics.AddError(
				"Unable to delete Region.",
				err.Error(),
			)
			return
		}
	}
}
