// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// tokenResourceModel is the Terraform state/plan model of the "token" resource.
type tokenResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	UserID       types.Int64  `tfsdk:"user_id"`
	Key          types.String `tfsdk:"key"`
	Token        types.String `tfsdk:"token"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	Version      types.Int64  `tfsdk:"version"`
	PepperID     types.Int64  `tfsdk:"pepper_id"`
	Description  types.String `tfsdk:"description"`
	WriteEnabled types.Bool   `tfsdk:"write_enabled"`
	Expires      types.String `tfsdk:"expires"`
	LastUsed     types.String `tfsdk:"last_used"`
	Created      types.String `tfsdk:"created"`
	URL          types.String `tfsdk:"url"`
}

// expandToken converts the Terraform model of type "token" into an API request.
func expandToken(ctx context.Context, model *tokenResourceModel) (*netboxapi.TokenRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.TokenRequestDTO{}
	requestDTO.User = conv.Int64Ptr(model.UserID)
	requestDTO.Token = conv.StringPtr(model.Token)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.Version = conv.Int64Ptr(model.Version)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.WriteEnabled = conv.BoolPtr(model.WriteEnabled)
	requestDTO.Expires = conv.StringPtr(model.Expires)
	return requestDTO, diags
}

// flattenToken refreshes the Terraform model of type "token" from an API response.
func flattenToken(ctx context.Context, responseDTO *netboxapi.TokenResponseDTO, model *tokenResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.UserID = conv.FromInt64Ptr(responseDTO.User)
	model.Key = conv.FromStringPtr(responseDTO.Key)
	model.Token = conv.FromStringPtr(responseDTO.Token)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.Version = conv.FromInt64Ptr(responseDTO.Version)
	model.PepperID = conv.FromInt64Ptr(responseDTO.PepperID)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.WriteEnabled = conv.FromBoolPtr(responseDTO.WriteEnabled)
	model.Expires = conv.FromStringPtr(responseDTO.Expires)
	model.LastUsed = conv.FromStringPtr(responseDTO.LastUsed)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
