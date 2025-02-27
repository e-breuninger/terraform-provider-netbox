package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/netbox-community/go-netbox/v4"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*tagResource)(nil)

func NewTagResource() resource.Resource {
	return &tagResource{}
}

type tagResource struct {
	provider *netboxProvider
}

type tagResourceModel struct {
	Color       types.String `tfsdk:"color"`
	Description types.String `tfsdk:"description"`
	Id          types.Int32  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	ObjectTypes types.List   `tfsdk:"object_types"`
	Slug        types.String `tfsdk:"slug"`
	ColorHex    types.String `tfsdk:"color_hex"`
}

func (r *tagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (r *tagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"color": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 6),
					stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-f]{6}$"), ""),
				},
			},
			"color_hex": schema.StringAttribute{
				Optional:           true,
				Description:        "**Deprecated** Use color instead.",
				DeprecationMessage: "Use the *color* attribute instead.",
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRoot("color"),
					}...),
				},
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

func (r *tagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(*netboxProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *netbox.apiCLient, got: %T, Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.provider = provider
}

func (r *tagResource) readAPI(ctx context.Context, data *tagResourceModel, tag *netbox.Tag) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	data.Id = types.Int32Value(tag.Id)
	data.Name = types.StringValue(tag.Name)
	data.Slug = types.StringValue(tag.Slug)
	data.Color = types.StringPointerValue(tag.Color)
	data.Description = types.StringPointerValue(tag.Description)
	listObjectTypes, err := types.ListValueFrom(ctx, types.StringType, tag.ObjectTypes)
	if err != nil {
		return err
	}
	data.ObjectTypes = listObjectTypes

	return diags
}

func (r *tagResource) writeAPI(ctx context.Context, data *tagResourceModel) (*netbox.TagRequest, diag.Diagnostics) {
	var diags = diag.Diagnostics{}
	tagRequest := netbox.NewTagRequestWithDefaults()
	tagRequest.Name = data.Name.ValueString()
	tagRequest.Color = data.Color.ValueStringPointer()
	tagRequest.Slug = data.Slug.ValueString()
	tagRequest.Description = data.Description.ValueStringPointer()
	elements := make([]string, 0, len(data.ObjectTypes.Elements()))
	diagObjectType := data.ObjectTypes.ElementsAs(ctx, &elements, false)
	diags.Append(diagObjectType...)
	tagRequest.ObjectTypes = elements
	if tagRequest.ObjectTypes == nil {
		tagRequest.ObjectTypes = []string{}
	}
	return tagRequest, diags
}

func (r *tagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to import Webhook. This method requires a integer.",
			err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *tagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data tagResourceModel

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
	api_res, _, err := r.provider.client.ExtrasAPI.
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
	errors := r.readAPI(ctx, &data, api_res)

	if errors != nil {
		resp.Diagnostics.Append(errors...)
		return
	}
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data tagResourceModel

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
				"Unable to retrieve Webhook value.",
				err.Error(),
			)
			return
		}
	}

	errors := r.readAPI(ctx, &data, tag)

	if errors != nil {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data tagResourceModel

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

	errors := r.readAPI(ctx, &data, tag)
	if errors != nil {
		resp.Diagnostics.Append(errors...)
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data tagResourceModel

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
				"Unable to update Tag.",
				err.Error(),
			)
			return
		}
	}
}
