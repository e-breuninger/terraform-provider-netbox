package provider

import (
	"context"
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

type regionResourceModel struct {
	CustomFields types.Dynamic `tfsdk:"custom_fields"`
	Description  types.String  `tfsdk:"description"`
	Id           types.Int32   `tfsdk:"id"`
	Name         types.String  `tfsdk:"name"`
	Parent       types.Int32   `tfsdk:"parent"`
	Slug         types.String  `tfsdk:"slug"`
	Tags         types.List    `tfsdk:"tags"`
}

func (r *regionResource) readAPI(ctx context.Context, data *regionResourceModel, region *netbox.Region) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	data.Id = types.Int32Value(region.Id)
	data.Name = types.StringValue(region.Name)
	data.Slug = types.StringValue(region.Slug)
	data.Description = types.StringPointerValue(region.Description)

	if region.Parent.Get() != nil {
		data.Parent = types.Int32Value(region.Parent.Get().Id)
	} else {
		data.Parent = types.Int32Null()
	}

	customFieldResults, diags := readCustomFieldsFromAPI(region.CustomFields)
	if diags.HasError() {
		return diags
	}
	data.CustomFields = types.DynamicValue(customFieldResults)

	tags := readTags(region.Tags)
	tagsdata, diagdata := types.ListValueFrom(ctx, types.StringType, tags)
	if diagdata.HasError() {
		diags.Append(diagdata...)
		return diags
	}
	data.Tags = tagsdata
	return nil
}

func (r *regionResource) writeAPI(ctx context.Context, data *regionResourceModel) *netbox.WritableRegionRequest {
	regionRequest := netbox.NewWritableRegionRequestWithDefaults()
	regionRequest.Name = data.Name.ValueString()
	regionRequest.Slug = data.Slug.ValueString()
	regionRequest.Description = data.Description.ValueStringPointer()

	var tagList []netbox.NestedTagRequest
	for _, element := range data.Tags.Elements() {
		tag := netbox.NewNestedTagRequestWithDefaults()
		tag.Name = element.String()
		tagList = append(tagList, *tag)
	}
	regionRequest.Tags = tagList

	regionRequest.CustomFields = readCustomFieldsFromTerraform(data.CustomFields)

	if !data.Parent.IsUnknown() {
		regionRequest.Parent = *netbox.NewNullableInt32(data.Parent.ValueInt32Pointer())
	}

	return regionRequest
}

func (r *regionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

func (r *regionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"custom_fields": schema.DynamicAttribute{
				Optional: true,
				Computed: true,
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
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (r *regionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data regionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	regionRequest := r.writeAPI(ctx, &data)

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
	errors := r.readAPI(ctx, &data, api_res)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *regionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data regionResourceModel

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
	errors := r.readAPI(ctx, &data, region)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *regionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data regionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	regionRequest := r.writeAPI(ctx, &data)

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

	errors := r.readAPI(ctx, &data, region)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *regionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data regionResourceModel

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
