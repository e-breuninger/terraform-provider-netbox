// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// contactResourceModel is the Terraform state/plan model of the "contact" resource.
type contactResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Title        types.String `tfsdk:"title"`
	Phone        types.String `tfsdk:"phone"`
	Email        types.String `tfsdk:"email"`
	Address      types.String `tfsdk:"address"`
	Link         types.String `tfsdk:"link"`
	GroupIds     types.Set    `tfsdk:"group_ids"`
	Description  types.String `tfsdk:"description"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandContact converts the Terraform model of type "contact" into an API request.
func expandContact(ctx context.Context, model *contactResourceModel) (*netboxapi.ContactRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ContactRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Title = conv.StringPtr(model.Title)
	requestDTO.Phone = conv.StringPtr(model.Phone)
	requestDTO.Email = conv.StringPtr(model.Email)
	requestDTO.Address = conv.StringPtr(model.Address)
	requestDTO.Link = conv.StringPtr(model.Link)
	requestDTO.Groups = conv.SetTo[int64](ctx, model.GroupIds, &diags)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenContact refreshes the Terraform model of type "contact" from an API response.
func flattenContact(ctx context.Context, responseDTO *netboxapi.ContactResponseDTO, model *contactResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Title = conv.FromStringPtr(responseDTO.Title)
	model.Phone = conv.FromStringPtr(responseDTO.Phone)
	model.Email = conv.FromStringPtr(responseDTO.Email)
	model.Address = conv.FromStringPtr(responseDTO.Address)
	model.Link = conv.FromStringPtr(responseDTO.Link)
	model.GroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Groups, model.GroupIds.IsNull() || model.GroupIds.IsUnknown(), &diags)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
