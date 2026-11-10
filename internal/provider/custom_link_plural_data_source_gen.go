// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*customLinksDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customLinksDataSource)(nil)
)

// NewCustomLinksDataSource returns a new custom_links data source, which
// lists custom_link objects matching its filters.
func NewCustomLinksDataSource() datasource.DataSource {
	return &customLinksDataSource{}
}

type customLinksDataSource struct {
	client *netboxapi.Client
}

// customLinksDataSourceModel is the data source model: the filters, the limit and the matching
// custom_links.
type customLinksDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"custom_links"`
}

// customLinkResourceAttrTypes is the attribute type map of customLinkResourceModel.
func customLinkResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":           types.Int64Type,
		"name":         types.StringType,
		"object_types": types.SetType{ElemType: types.StringType},
		"enabled":      types.BoolType,
		"link_text":    types.StringType,
		"link_url":     types.StringType,
		"weight":       types.Int64Type,
		"group_name":   types.StringType,
		"button_class": types.StringType,
		"new_window":   types.BoolType,
		"owner_id":     types.Int64Type,
		"created":      types.StringType,
		"last_updated": types.StringType,
		"url":          types.StringType,
	}
}

func (d *customLinksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_links"
}

func (d *customLinksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists custom_link objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: button_class, button_class__empty, button_class__ic, button_class__ie, button_class__iew, button_class__iregex, button_class__isw, button_class__n, button_class__nic, button_class__nie, button_class__niew, button_class__nisw, button_class__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, enabled, group_name, group_name__empty, group_name__ic, group_name__ie, group_name__iew, group_name__iregex, group_name__isw, group_name__n, group_name__nic, group_name__nie, group_name__niew, group_name__nisw, group_name__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, link_text, link_text__ic, link_text__ie, link_text__iew, link_text__iregex, link_text__isw, link_text__n, link_text__nic, link_text__nie, link_text__niew, link_text__nisw, link_text__regex, link_url, link_url__ic, link_url__ie, link_url__iew, link_url__iregex, link_url__isw, link_url__n, link_url__nic, link_url__nie, link_url__niew, link_url__nisw, link_url__regex, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, new_window, object_type, object_type__ic, object_type__ie, object_type__iew, object_type__iregex, object_type__isw, object_type__n, object_type__nic, object_type__nie, object_type__niew, object_type__nisw, object_type__regex, object_type_id, object_type_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n. Repeating a name sends that parameter once per value.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "Name of the query filter (the API list parameter).",
						},
						"value": schema.StringAttribute{
							Required:    true,
							Description: "Value to filter by.",
						},
					},
				},
			},
			"limit": schema.Int64Attribute{
				Optional:    true,
				Description: "The maximum number of objects to return from the API lookup. Defaults to 1000.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"custom_links": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching custom_links.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"object_types": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Object types the link is shown on, as app-labelled content types (dcim.device).",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the link is rendered.",
						},
						"link_text": schema.StringAttribute{
							Computed:    true,
							Description: "Jinja2 template for the link text; the object is available as {{ object }}.",
						},
						"link_url": schema.StringAttribute{
							Computed:    true,
							Description: "Jinja2 template for the link URL; the object is available as {{ object }}.",
						},
						"weight": schema.Int64Attribute{
							Computed:    true,
							Description: "Order within the group. Must be a whole number.",
						},
						"group_name": schema.StringAttribute{
							Computed:    true,
							Description: "Links with the same group name are rendered as a dropdown.",
						},
						"button_class": schema.StringAttribute{
							Computed:    true,
							Description: "Button style. One of: default, blue, indigo, purple, pink, red, orange, yellow, green, teal, cyan, gray, black, white, ghost-dark.",
						},
						"new_window": schema.BoolAttribute{
							Computed:    true,
							Description: "Open the link in a new browser window.",
						},
						"owner_id": schema.Int64Attribute{
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
					},
				},
			},
		},
	}
}

func (d *customLinksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *customLinksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customLinksDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	limit := int64(1000)
	if !state.Limit.IsNull() && !state.Limit.IsUnknown() {
		limit = state.Limit.ValueInt64()
	}
	_ = limit

	// The filters entries, in no particular order. Entries sharing a name are distinct values of
	// the same API parameter.
	var filters []struct {
		Name  types.String `tfsdk:"name"`
		Value types.String `tfsdk:"value"`
	}
	if !state.Filters.IsNull() {
		resp.Diagnostics.Append(state.Filters.ElementsAs(ctx, &filters, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	_ = filters
	var responseDTOs []*netboxapi.CustomLinkResponseDTO

	params := extras.NewExtrasCustomLinksListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vName []string
	var vNamen []string
	var vNameIc []string
	var vNameNic []string
	var vNameIe []string
	var vNameNie []string
	var vNameIsw []string
	var vNameNisw []string
	var vNameIew []string
	var vNameNiew []string
	var vNameEmpty *bool
	var vNameRegex []string
	var vNameIregex []string
	var vEnabled *bool
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vButtonClass *string
	var vButtonClassn *string
	var vButtonClassIc []string
	var vButtonClassNic []string
	var vButtonClassIe []string
	var vButtonClassNie []string
	var vButtonClassIsw []string
	var vButtonClassNisw []string
	var vButtonClassIew []string
	var vButtonClassNiew []string
	var vButtonClassEmpty *bool
	var vButtonClassRegex []string
	var vButtonClassIregex []string
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vGroupName []string
	var vGroupNamen []string
	var vGroupNameIc []string
	var vGroupNameNic []string
	var vGroupNameIe []string
	var vGroupNameNie []string
	var vGroupNameIsw []string
	var vGroupNameNisw []string
	var vGroupNameIew []string
	var vGroupNameNiew []string
	var vGroupNameEmpty *bool
	var vGroupNameRegex []string
	var vGroupNameIregex []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLinkText *string
	var vLinkTextn *string
	var vLinkTextIc *string
	var vLinkTextNic *string
	var vLinkTextIe *string
	var vLinkTextNie *string
	var vLinkTextIsw *string
	var vLinkTextNisw *string
	var vLinkTextIew *string
	var vLinkTextNiew *string
	var vLinkTextRegex *string
	var vLinkTextIregex *string
	var vLinkURL *string
	var vLinkURLn *string
	var vLinkURLIc *string
	var vLinkURLNic *string
	var vLinkURLIe *string
	var vLinkURLNie *string
	var vLinkURLIsw *string
	var vLinkURLNisw *string
	var vLinkURLIew *string
	var vLinkURLNiew *string
	var vLinkURLRegex *string
	var vLinkURLIregex *string
	var vNewWindow *bool
	var vObjectType []string
	var vObjectTypen []string
	var vObjectTypeIc []string
	var vObjectTypeNic []string
	var vObjectTypeIe []string
	var vObjectTypeNie []string
	var vObjectTypeIsw []string
	var vObjectTypeNisw []string
	var vObjectTypeIew []string
	var vObjectTypeNiew []string
	var vObjectTypeRegex []string
	var vObjectTypeIregex []string
	var vObjectTypeID []int64
	var vObjectTypeIDn []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vQ *string
	var vWeight []int64
	var vWeightn []int64
	var vWeightLt []int64
	var vWeightLte []int64
	var vWeightGt []int64
	var vWeightGte []int64
	var vWeightEmpty *bool
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		switch name {
		case "id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id' takes an integer, got %q.", value))
				return
			}
			vID = append(vID, v)
		case "id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__n' takes an integer, got %q.", value))
				return
			}
			vIDn = append(vIDn, v)
		case "id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lt' takes an integer, got %q.", value))
				return
			}
			vIDLt = append(vIDLt, v)
		case "id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lte' takes an integer, got %q.", value))
				return
			}
			vIDLte = append(vIDLte, v)
		case "id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gt' takes an integer, got %q.", value))
				return
			}
			vIDGt = append(vIDGt, v)
		case "id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gte' takes an integer, got %q.", value))
				return
			}
			vIDGte = append(vIDGte, v)
		case "id__empty":
			if vIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__empty' takes a boolean, got %q.", value))
				return
			}
			vIDEmpty = &v
		case "name":
			v := value
			vName = append(vName, v)
		case "name__n":
			v := value
			vNamen = append(vNamen, v)
		case "name__ic":
			v := value
			vNameIc = append(vNameIc, v)
		case "name__nic":
			v := value
			vNameNic = append(vNameNic, v)
		case "name__ie":
			v := value
			vNameIe = append(vNameIe, v)
		case "name__nie":
			v := value
			vNameNie = append(vNameNie, v)
		case "name__isw":
			v := value
			vNameIsw = append(vNameIsw, v)
		case "name__nisw":
			v := value
			vNameNisw = append(vNameNisw, v)
		case "name__iew":
			v := value
			vNameIew = append(vNameIew, v)
		case "name__niew":
			v := value
			vNameNiew = append(vNameNiew, v)
		case "name__empty":
			if vNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'name__empty' takes a boolean, got %q.", value))
				return
			}
			vNameEmpty = &v
		case "name__regex":
			v := value
			vNameRegex = append(vNameRegex, v)
		case "name__iregex":
			v := value
			vNameIregex = append(vNameIregex, v)
		case "enabled":
			if vEnabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'enabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'enabled' takes a boolean, got %q.", value))
				return
			}
			vEnabled = &v
		case "owner_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id' takes an integer, got %q.", value))
				return
			}
			vOwnerID = append(vOwnerID, v)
		case "owner_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerIDn = append(vOwnerIDn, v)
		case "button_class":
			if vButtonClass != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'button_class' takes a single value.")
				return
			}
			v := value
			vButtonClass = &v
		case "button_class__n":
			if vButtonClassn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'button_class__n' takes a single value.")
				return
			}
			v := value
			vButtonClassn = &v
		case "button_class__ic":
			v := value
			vButtonClassIc = append(vButtonClassIc, v)
		case "button_class__nic":
			v := value
			vButtonClassNic = append(vButtonClassNic, v)
		case "button_class__ie":
			v := value
			vButtonClassIe = append(vButtonClassIe, v)
		case "button_class__nie":
			v := value
			vButtonClassNie = append(vButtonClassNie, v)
		case "button_class__isw":
			v := value
			vButtonClassIsw = append(vButtonClassIsw, v)
		case "button_class__nisw":
			v := value
			vButtonClassNisw = append(vButtonClassNisw, v)
		case "button_class__iew":
			v := value
			vButtonClassIew = append(vButtonClassIew, v)
		case "button_class__niew":
			v := value
			vButtonClassNiew = append(vButtonClassNiew, v)
		case "button_class__empty":
			if vButtonClassEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'button_class__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'button_class__empty' takes a boolean, got %q.", value))
				return
			}
			vButtonClassEmpty = &v
		case "button_class__regex":
			v := value
			vButtonClassRegex = append(vButtonClassRegex, v)
		case "button_class__iregex":
			v := value
			vButtonClassIregex = append(vButtonClassIregex, v)
		case "created":
			v := value
			vCreated = append(vCreated, v)
		case "created__n":
			v := value
			vCreatedn = append(vCreatedn, v)
		case "created__lt":
			v := value
			vCreatedLt = append(vCreatedLt, v)
		case "created__lte":
			v := value
			vCreatedLte = append(vCreatedLte, v)
		case "created__gt":
			v := value
			vCreatedGt = append(vCreatedGt, v)
		case "created__gte":
			v := value
			vCreatedGte = append(vCreatedGte, v)
		case "created__empty":
			v := value
			vCreatedEmpty = append(vCreatedEmpty, v)
		case "group_name":
			v := value
			vGroupName = append(vGroupName, v)
		case "group_name__n":
			v := value
			vGroupNamen = append(vGroupNamen, v)
		case "group_name__ic":
			v := value
			vGroupNameIc = append(vGroupNameIc, v)
		case "group_name__nic":
			v := value
			vGroupNameNic = append(vGroupNameNic, v)
		case "group_name__ie":
			v := value
			vGroupNameIe = append(vGroupNameIe, v)
		case "group_name__nie":
			v := value
			vGroupNameNie = append(vGroupNameNie, v)
		case "group_name__isw":
			v := value
			vGroupNameIsw = append(vGroupNameIsw, v)
		case "group_name__nisw":
			v := value
			vGroupNameNisw = append(vGroupNameNisw, v)
		case "group_name__iew":
			v := value
			vGroupNameIew = append(vGroupNameIew, v)
		case "group_name__niew":
			v := value
			vGroupNameNiew = append(vGroupNameNiew, v)
		case "group_name__empty":
			if vGroupNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'group_name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_name__empty' takes a boolean, got %q.", value))
				return
			}
			vGroupNameEmpty = &v
		case "group_name__regex":
			v := value
			vGroupNameRegex = append(vGroupNameRegex, v)
		case "group_name__iregex":
			v := value
			vGroupNameIregex = append(vGroupNameIregex, v)
		case "last_updated":
			v := value
			vLastUpdated = append(vLastUpdated, v)
		case "last_updated__n":
			v := value
			vLastUpdatedn = append(vLastUpdatedn, v)
		case "last_updated__lt":
			v := value
			vLastUpdatedLt = append(vLastUpdatedLt, v)
		case "last_updated__lte":
			v := value
			vLastUpdatedLte = append(vLastUpdatedLte, v)
		case "last_updated__gt":
			v := value
			vLastUpdatedGt = append(vLastUpdatedGt, v)
		case "last_updated__gte":
			v := value
			vLastUpdatedGte = append(vLastUpdatedGte, v)
		case "last_updated__empty":
			v := value
			vLastUpdatedEmpty = append(vLastUpdatedEmpty, v)
		case "link_text":
			if vLinkText != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text' takes a single value.")
				return
			}
			v := value
			vLinkText = &v
		case "link_text__n":
			if vLinkTextn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__n' takes a single value.")
				return
			}
			v := value
			vLinkTextn = &v
		case "link_text__ic":
			if vLinkTextIc != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__ic' takes a single value.")
				return
			}
			v := value
			vLinkTextIc = &v
		case "link_text__nic":
			if vLinkTextNic != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__nic' takes a single value.")
				return
			}
			v := value
			vLinkTextNic = &v
		case "link_text__ie":
			if vLinkTextIe != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__ie' takes a single value.")
				return
			}
			v := value
			vLinkTextIe = &v
		case "link_text__nie":
			if vLinkTextNie != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__nie' takes a single value.")
				return
			}
			v := value
			vLinkTextNie = &v
		case "link_text__isw":
			if vLinkTextIsw != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__isw' takes a single value.")
				return
			}
			v := value
			vLinkTextIsw = &v
		case "link_text__nisw":
			if vLinkTextNisw != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__nisw' takes a single value.")
				return
			}
			v := value
			vLinkTextNisw = &v
		case "link_text__iew":
			if vLinkTextIew != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__iew' takes a single value.")
				return
			}
			v := value
			vLinkTextIew = &v
		case "link_text__niew":
			if vLinkTextNiew != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__niew' takes a single value.")
				return
			}
			v := value
			vLinkTextNiew = &v
		case "link_text__regex":
			if vLinkTextRegex != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__regex' takes a single value.")
				return
			}
			v := value
			vLinkTextRegex = &v
		case "link_text__iregex":
			if vLinkTextIregex != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_text__iregex' takes a single value.")
				return
			}
			v := value
			vLinkTextIregex = &v
		case "link_url":
			if vLinkURL != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url' takes a single value.")
				return
			}
			v := value
			vLinkURL = &v
		case "link_url__n":
			if vLinkURLn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__n' takes a single value.")
				return
			}
			v := value
			vLinkURLn = &v
		case "link_url__ic":
			if vLinkURLIc != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__ic' takes a single value.")
				return
			}
			v := value
			vLinkURLIc = &v
		case "link_url__nic":
			if vLinkURLNic != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__nic' takes a single value.")
				return
			}
			v := value
			vLinkURLNic = &v
		case "link_url__ie":
			if vLinkURLIe != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__ie' takes a single value.")
				return
			}
			v := value
			vLinkURLIe = &v
		case "link_url__nie":
			if vLinkURLNie != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__nie' takes a single value.")
				return
			}
			v := value
			vLinkURLNie = &v
		case "link_url__isw":
			if vLinkURLIsw != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__isw' takes a single value.")
				return
			}
			v := value
			vLinkURLIsw = &v
		case "link_url__nisw":
			if vLinkURLNisw != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__nisw' takes a single value.")
				return
			}
			v := value
			vLinkURLNisw = &v
		case "link_url__iew":
			if vLinkURLIew != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__iew' takes a single value.")
				return
			}
			v := value
			vLinkURLIew = &v
		case "link_url__niew":
			if vLinkURLNiew != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__niew' takes a single value.")
				return
			}
			v := value
			vLinkURLNiew = &v
		case "link_url__regex":
			if vLinkURLRegex != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__regex' takes a single value.")
				return
			}
			v := value
			vLinkURLRegex = &v
		case "link_url__iregex":
			if vLinkURLIregex != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link_url__iregex' takes a single value.")
				return
			}
			v := value
			vLinkURLIregex = &v
		case "new_window":
			if vNewWindow != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'new_window' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'new_window' takes a boolean, got %q.", value))
				return
			}
			vNewWindow = &v
		case "object_type":
			v := value
			vObjectType = append(vObjectType, v)
		case "object_type__n":
			v := value
			vObjectTypen = append(vObjectTypen, v)
		case "object_type__ic":
			v := value
			vObjectTypeIc = append(vObjectTypeIc, v)
		case "object_type__nic":
			v := value
			vObjectTypeNic = append(vObjectTypeNic, v)
		case "object_type__ie":
			v := value
			vObjectTypeIe = append(vObjectTypeIe, v)
		case "object_type__nie":
			v := value
			vObjectTypeNie = append(vObjectTypeNie, v)
		case "object_type__isw":
			v := value
			vObjectTypeIsw = append(vObjectTypeIsw, v)
		case "object_type__nisw":
			v := value
			vObjectTypeNisw = append(vObjectTypeNisw, v)
		case "object_type__iew":
			v := value
			vObjectTypeIew = append(vObjectTypeIew, v)
		case "object_type__niew":
			v := value
			vObjectTypeNiew = append(vObjectTypeNiew, v)
		case "object_type__regex":
			v := value
			vObjectTypeRegex = append(vObjectTypeRegex, v)
		case "object_type__iregex":
			v := value
			vObjectTypeIregex = append(vObjectTypeIregex, v)
		case "object_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id' takes an integer, got %q.", value))
				return
			}
			vObjectTypeID = append(vObjectTypeID, v)
		case "object_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id__n' takes an integer, got %q.", value))
				return
			}
			vObjectTypeIDn = append(vObjectTypeIDn, v)
		case "owner":
			v := value
			vOwner = append(vOwner, v)
		case "owner__n":
			v := value
			vOwnern = append(vOwnern, v)
		case "owner_group":
			v := value
			vOwnerGroup = append(vOwnerGroup, v)
		case "owner_group__n":
			v := value
			vOwnerGroupn = append(vOwnerGroupn, v)
		case "owner_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupID = append(vOwnerGroupID, v)
		case "owner_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupIDn = append(vOwnerGroupIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "weight":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight' takes an integer, got %q.", value))
				return
			}
			vWeight = append(vWeight, v)
		case "weight__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__n' takes an integer, got %q.", value))
				return
			}
			vWeightn = append(vWeightn, v)
		case "weight__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__lt' takes an integer, got %q.", value))
				return
			}
			vWeightLt = append(vWeightLt, v)
		case "weight__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__lte' takes an integer, got %q.", value))
				return
			}
			vWeightLte = append(vWeightLte, v)
		case "weight__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__gt' takes an integer, got %q.", value))
				return
			}
			vWeightGt = append(vWeightGt, v)
		case "weight__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__gte' takes an integer, got %q.", value))
				return
			}
			vWeightGte = append(vWeightGte, v)
		case "weight__empty":
			if vWeightEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'weight__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__empty' takes a boolean, got %q.", value))
				return
			}
			vWeightEmpty = &v
		default:
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
	}
	if len(vID) > 0 {
		params.SetID(vID)
	}
	if len(vIDn) > 0 {
		params.SetIDn(vIDn)
	}
	if len(vIDLt) > 0 {
		params.SetIDLt(vIDLt)
	}
	if len(vIDLte) > 0 {
		params.SetIDLte(vIDLte)
	}
	if len(vIDGt) > 0 {
		params.SetIDGt(vIDGt)
	}
	if len(vIDGte) > 0 {
		params.SetIDGte(vIDGte)
	}
	if vIDEmpty != nil {
		params.SetIDEmpty(vIDEmpty)
	}
	if len(vName) > 0 {
		params.SetName(vName)
	}
	if len(vNamen) > 0 {
		params.SetNamen(vNamen)
	}
	if len(vNameIc) > 0 {
		params.SetNameIc(vNameIc)
	}
	if len(vNameNic) > 0 {
		params.SetNameNic(vNameNic)
	}
	if len(vNameIe) > 0 {
		params.SetNameIe(vNameIe)
	}
	if len(vNameNie) > 0 {
		params.SetNameNie(vNameNie)
	}
	if len(vNameIsw) > 0 {
		params.SetNameIsw(vNameIsw)
	}
	if len(vNameNisw) > 0 {
		params.SetNameNisw(vNameNisw)
	}
	if len(vNameIew) > 0 {
		params.SetNameIew(vNameIew)
	}
	if len(vNameNiew) > 0 {
		params.SetNameNiew(vNameNiew)
	}
	if vNameEmpty != nil {
		params.SetNameEmpty(vNameEmpty)
	}
	if len(vNameRegex) > 0 {
		params.SetNameRegex(vNameRegex)
	}
	if len(vNameIregex) > 0 {
		params.SetNameIregex(vNameIregex)
	}
	if vEnabled != nil {
		params.SetEnabled(vEnabled)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
	}
	if vButtonClass != nil {
		params.SetButtonClass(vButtonClass)
	}
	if vButtonClassn != nil {
		params.SetButtonClassn(vButtonClassn)
	}
	if len(vButtonClassIc) > 0 {
		params.SetButtonClassIc(vButtonClassIc)
	}
	if len(vButtonClassNic) > 0 {
		params.SetButtonClassNic(vButtonClassNic)
	}
	if len(vButtonClassIe) > 0 {
		params.SetButtonClassIe(vButtonClassIe)
	}
	if len(vButtonClassNie) > 0 {
		params.SetButtonClassNie(vButtonClassNie)
	}
	if len(vButtonClassIsw) > 0 {
		params.SetButtonClassIsw(vButtonClassIsw)
	}
	if len(vButtonClassNisw) > 0 {
		params.SetButtonClassNisw(vButtonClassNisw)
	}
	if len(vButtonClassIew) > 0 {
		params.SetButtonClassIew(vButtonClassIew)
	}
	if len(vButtonClassNiew) > 0 {
		params.SetButtonClassNiew(vButtonClassNiew)
	}
	if vButtonClassEmpty != nil {
		params.SetButtonClassEmpty(vButtonClassEmpty)
	}
	if len(vButtonClassRegex) > 0 {
		params.SetButtonClassRegex(vButtonClassRegex)
	}
	if len(vButtonClassIregex) > 0 {
		params.SetButtonClassIregex(vButtonClassIregex)
	}
	if len(vCreated) > 0 {
		params.SetCreated(vCreated)
	}
	if len(vCreatedn) > 0 {
		params.SetCreatedn(vCreatedn)
	}
	if len(vCreatedLt) > 0 {
		params.SetCreatedLt(vCreatedLt)
	}
	if len(vCreatedLte) > 0 {
		params.SetCreatedLte(vCreatedLte)
	}
	if len(vCreatedGt) > 0 {
		params.SetCreatedGt(vCreatedGt)
	}
	if len(vCreatedGte) > 0 {
		params.SetCreatedGte(vCreatedGte)
	}
	if len(vCreatedEmpty) > 0 {
		params.SetCreatedEmpty(vCreatedEmpty)
	}
	if len(vGroupName) > 0 {
		params.SetGroupName(vGroupName)
	}
	if len(vGroupNamen) > 0 {
		params.SetGroupNamen(vGroupNamen)
	}
	if len(vGroupNameIc) > 0 {
		params.SetGroupNameIc(vGroupNameIc)
	}
	if len(vGroupNameNic) > 0 {
		params.SetGroupNameNic(vGroupNameNic)
	}
	if len(vGroupNameIe) > 0 {
		params.SetGroupNameIe(vGroupNameIe)
	}
	if len(vGroupNameNie) > 0 {
		params.SetGroupNameNie(vGroupNameNie)
	}
	if len(vGroupNameIsw) > 0 {
		params.SetGroupNameIsw(vGroupNameIsw)
	}
	if len(vGroupNameNisw) > 0 {
		params.SetGroupNameNisw(vGroupNameNisw)
	}
	if len(vGroupNameIew) > 0 {
		params.SetGroupNameIew(vGroupNameIew)
	}
	if len(vGroupNameNiew) > 0 {
		params.SetGroupNameNiew(vGroupNameNiew)
	}
	if vGroupNameEmpty != nil {
		params.SetGroupNameEmpty(vGroupNameEmpty)
	}
	if len(vGroupNameRegex) > 0 {
		params.SetGroupNameRegex(vGroupNameRegex)
	}
	if len(vGroupNameIregex) > 0 {
		params.SetGroupNameIregex(vGroupNameIregex)
	}
	if len(vLastUpdated) > 0 {
		params.SetLastUpdated(vLastUpdated)
	}
	if len(vLastUpdatedn) > 0 {
		params.SetLastUpdatedn(vLastUpdatedn)
	}
	if len(vLastUpdatedLt) > 0 {
		params.SetLastUpdatedLt(vLastUpdatedLt)
	}
	if len(vLastUpdatedLte) > 0 {
		params.SetLastUpdatedLte(vLastUpdatedLte)
	}
	if len(vLastUpdatedGt) > 0 {
		params.SetLastUpdatedGt(vLastUpdatedGt)
	}
	if len(vLastUpdatedGte) > 0 {
		params.SetLastUpdatedGte(vLastUpdatedGte)
	}
	if len(vLastUpdatedEmpty) > 0 {
		params.SetLastUpdatedEmpty(vLastUpdatedEmpty)
	}
	if vLinkText != nil {
		params.SetLinkText(vLinkText)
	}
	if vLinkTextn != nil {
		params.SetLinkTextn(vLinkTextn)
	}
	if vLinkTextIc != nil {
		params.SetLinkTextIc(vLinkTextIc)
	}
	if vLinkTextNic != nil {
		params.SetLinkTextNic(vLinkTextNic)
	}
	if vLinkTextIe != nil {
		params.SetLinkTextIe(vLinkTextIe)
	}
	if vLinkTextNie != nil {
		params.SetLinkTextNie(vLinkTextNie)
	}
	if vLinkTextIsw != nil {
		params.SetLinkTextIsw(vLinkTextIsw)
	}
	if vLinkTextNisw != nil {
		params.SetLinkTextNisw(vLinkTextNisw)
	}
	if vLinkTextIew != nil {
		params.SetLinkTextIew(vLinkTextIew)
	}
	if vLinkTextNiew != nil {
		params.SetLinkTextNiew(vLinkTextNiew)
	}
	if vLinkTextRegex != nil {
		params.SetLinkTextRegex(vLinkTextRegex)
	}
	if vLinkTextIregex != nil {
		params.SetLinkTextIregex(vLinkTextIregex)
	}
	if vLinkURL != nil {
		params.SetLinkURL(vLinkURL)
	}
	if vLinkURLn != nil {
		params.SetLinkURLn(vLinkURLn)
	}
	if vLinkURLIc != nil {
		params.SetLinkURLIc(vLinkURLIc)
	}
	if vLinkURLNic != nil {
		params.SetLinkURLNic(vLinkURLNic)
	}
	if vLinkURLIe != nil {
		params.SetLinkURLIe(vLinkURLIe)
	}
	if vLinkURLNie != nil {
		params.SetLinkURLNie(vLinkURLNie)
	}
	if vLinkURLIsw != nil {
		params.SetLinkURLIsw(vLinkURLIsw)
	}
	if vLinkURLNisw != nil {
		params.SetLinkURLNisw(vLinkURLNisw)
	}
	if vLinkURLIew != nil {
		params.SetLinkURLIew(vLinkURLIew)
	}
	if vLinkURLNiew != nil {
		params.SetLinkURLNiew(vLinkURLNiew)
	}
	if vLinkURLRegex != nil {
		params.SetLinkURLRegex(vLinkURLRegex)
	}
	if vLinkURLIregex != nil {
		params.SetLinkURLIregex(vLinkURLIregex)
	}
	if vNewWindow != nil {
		params.SetNewWindow(vNewWindow)
	}
	if len(vObjectType) > 0 {
		params.SetObjectType(vObjectType)
	}
	if len(vObjectTypen) > 0 {
		params.SetObjectTypen(vObjectTypen)
	}
	if len(vObjectTypeIc) > 0 {
		params.SetObjectTypeIc(vObjectTypeIc)
	}
	if len(vObjectTypeNic) > 0 {
		params.SetObjectTypeNic(vObjectTypeNic)
	}
	if len(vObjectTypeIe) > 0 {
		params.SetObjectTypeIe(vObjectTypeIe)
	}
	if len(vObjectTypeNie) > 0 {
		params.SetObjectTypeNie(vObjectTypeNie)
	}
	if len(vObjectTypeIsw) > 0 {
		params.SetObjectTypeIsw(vObjectTypeIsw)
	}
	if len(vObjectTypeNisw) > 0 {
		params.SetObjectTypeNisw(vObjectTypeNisw)
	}
	if len(vObjectTypeIew) > 0 {
		params.SetObjectTypeIew(vObjectTypeIew)
	}
	if len(vObjectTypeNiew) > 0 {
		params.SetObjectTypeNiew(vObjectTypeNiew)
	}
	if len(vObjectTypeRegex) > 0 {
		params.SetObjectTypeRegex(vObjectTypeRegex)
	}
	if len(vObjectTypeIregex) > 0 {
		params.SetObjectTypeIregex(vObjectTypeIregex)
	}
	if len(vObjectTypeID) > 0 {
		params.SetObjectTypeID(vObjectTypeID)
	}
	if len(vObjectTypeIDn) > 0 {
		params.SetObjectTypeIDn(vObjectTypeIDn)
	}
	if len(vOwner) > 0 {
		params.SetOwner(vOwner)
	}
	if len(vOwnern) > 0 {
		params.SetOwnern(vOwnern)
	}
	if len(vOwnerGroup) > 0 {
		params.SetOwnerGroup(vOwnerGroup)
	}
	if len(vOwnerGroupn) > 0 {
		params.SetOwnerGroupn(vOwnerGroupn)
	}
	if len(vOwnerGroupID) > 0 {
		params.SetOwnerGroupID(vOwnerGroupID)
	}
	if len(vOwnerGroupIDn) > 0 {
		params.SetOwnerGroupIDn(vOwnerGroupIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vWeight) > 0 {
		params.SetWeight(vWeight)
	}
	if len(vWeightn) > 0 {
		params.SetWeightn(vWeightn)
	}
	if len(vWeightLt) > 0 {
		params.SetWeightLt(vWeightLt)
	}
	if len(vWeightLte) > 0 {
		params.SetWeightLte(vWeightLte)
	}
	if len(vWeightGt) > 0 {
		params.SetWeightGt(vWeightGt)
	}
	if len(vWeightGte) > 0 {
		params.SetWeightGte(vWeightGte)
	}
	if vWeightEmpty != nil {
		params.SetWeightEmpty(vWeightEmpty)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Extras.ExtrasCustomLinksListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_custom_links", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CustomLinkResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]customLinkResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model customLinkResourceModel
		resp.Diagnostics.Append(flattenCustomLink(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: customLinkResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
