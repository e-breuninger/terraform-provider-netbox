// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// customLinkResourceModel is the Terraform state/plan model of the "custom_link" resource.
type customLinkResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	ObjectTypes types.Set    `tfsdk:"object_types"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	LinkText    types.String `tfsdk:"link_text"`
	LinkURL     types.String `tfsdk:"link_url"`
	Weight      types.Int64  `tfsdk:"weight"`
	GroupName   types.String `tfsdk:"group_name"`
	ButtonClass types.String `tfsdk:"button_class"`
	NewWindow   types.Bool   `tfsdk:"new_window"`
	OwnerID     types.Int64  `tfsdk:"owner_id"`
	Created     types.String `tfsdk:"created"`
	LastUpdated types.String `tfsdk:"last_updated"`
	URL         types.String `tfsdk:"url"`
}

// expandCustomLink converts the Terraform model of type "custom_link" into an API request.
func expandCustomLink(ctx context.Context, model *customLinkResourceModel) (*netboxapi.CustomLinkRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CustomLinkRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.ObjectTypes = conv.SetTo[string](ctx, model.ObjectTypes, &diags)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.LinkText = conv.StringPtr(model.LinkText)
	requestDTO.LinkURL = conv.StringPtr(model.LinkURL)
	requestDTO.Weight = conv.Int64Ptr(model.Weight)
	requestDTO.GroupName = conv.StringPtr(model.GroupName)
	requestDTO.ButtonClass = conv.StringPtr(model.ButtonClass)
	requestDTO.NewWindow = conv.BoolPtr(model.NewWindow)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	return requestDTO, diags
}

// flattenCustomLink refreshes the Terraform model of type "custom_link" from an API response.
func flattenCustomLink(ctx context.Context, responseDTO *netboxapi.CustomLinkResponseDTO, model *customLinkResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.ObjectTypes = conv.SetFrom(ctx, types.StringType, responseDTO.ObjectTypes, false, &diags)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.LinkText = conv.FromStringPtr(responseDTO.LinkText)
	model.LinkURL = conv.FromStringPtr(responseDTO.LinkURL)
	model.Weight = conv.FromInt64Ptr(responseDTO.Weight)
	model.GroupName = conv.FromStringPtr(responseDTO.GroupName)
	model.ButtonClass = conv.FromStringPtr(responseDTO.ButtonClass)
	model.NewWindow = conv.FromBoolPtr(responseDTO.NewWindow)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
