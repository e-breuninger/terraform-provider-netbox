package provider

import (
	"context"
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

var _ resource.Resource = (*tagResource)(nil)

func NewTagResource() resource.Resource {
	return &tagResource{}
}

type tagResource struct {
	NetboxResource
}

func (r *tagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (r *tagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Tags are user-defined labels which can be applied to a variety of objects within NetBox. They can be used to establish dimensions of organization beyond the relationships built into NetBox. For example, you might create a tag to identify a particular ownership or condition across several types of objects.",
		Attributes: map[string]schema.Attribute{
			"color_hex": schema.StringAttribute{
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
				Description:         "A unique integer value identifying this tag.",
				MarkdownDescription: "A unique integer value identifying this tag.",
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
			"object_types": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile("^[-\\w]+$"), ""),
				},
			},
		},
	}
}

func (r *tagResource) writeAPI(ctx context.Context, data *models.TagTerraformModel) (*netbox.TagRequest, diag.Diagnostics) {
	var diags = diag.Diagnostics{}
	tagRequest := netbox.NewTagRequestWithDefaults()
	tagRequest.Name = data.Name.ValueString()
	if !data.ColorHex.IsUnknown() {
		tagRequest.Color = data.ColorHex.ValueStringPointer()
	}

	if data.Slug.IsUnknown() {
		tagRequest.Slug = getSlug(tagRequest.Name)
	} else {
		tagRequest.Slug = data.Slug.ValueString()
	}
	tagRequest.Description = data.Description.ValueStringPointer()
	if len(data.ObjectTypes.Elements()) > 0 {
		elements := make([]string, 0, len(data.ObjectTypes.Elements()))
		diagObjectTypeErrors := data.ObjectTypes.ElementsAs(ctx, &elements, false)
		diags.Append(diagObjectTypeErrors...)
		tagRequest.ObjectTypes = elements
	}

	return tagRequest, diags
}

func (r *tagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.TagTerraformModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	tagRequest, diags := r.writeAPI(ctx, &data)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	apiRes, _, err := r.provider.client.ExtrasAPI.
		ExtrasTagsCreate(ctx).
		TagRequest(*tagRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error while creating the tag.",
			err.Error(),
		)
		return
	}
	// Example data value setting
	errors := data.ReadAPI(ctx, apiRes)

	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.TagTerraformModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	tag, httpCode, err := r.provider.client.ExtrasAPI.ExtrasTagsRetrieve(ctx, data.Id.ValueInt32()).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to retrieve Tag value.",
				err.Error(),
			)
			return
		}
	}

	errors := data.ReadAPI(ctx, tag)

	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.TagTerraformModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	tagRequest, diags := r.writeAPI(ctx, &data)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	tag, httpCode, err := r.provider.client.ExtrasAPI.
		ExtrasTagsUpdate(ctx, data.Id.ValueInt32()).
		TagRequest(*tagRequest).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.Diagnostics.AddError(
				"Tag no longer exists",
				err.Error())
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to update Tag.",
				err.Error(),
			)
			return
		}
	}

	errors := data.ReadAPI(ctx, tag)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.TagTerraformModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	httpCode, err := r.provider.client.ExtrasAPI.ExtrasTagsDestroy(ctx, data.Id.ValueInt32()).Execute()
	if err != nil {
		if httpCode != nil && httpCode.StatusCode != 404 {
			resp.Diagnostics.AddError(
				"Unable to delete Tag.",
				err.Error(),
			)
			return
		}
	}
}
