// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*inventoryItemTemplateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*inventoryItemTemplateDataSource)(nil)
)

// NewInventoryItemTemplateDataSource returns a new inventory_item_template data source.
func NewInventoryItemTemplateDataSource() datasource.DataSource {
	return &inventoryItemTemplateDataSource{}
}

type inventoryItemTemplateDataSource struct {
	client *netboxapi.Client
}

// inventoryItemTemplateDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type inventoryItemTemplateDataSourceModel struct {
	inventoryItemTemplateResourceModel
	NameContains types.String `tfsdk:"name_contains"`
}

func (d *inventoryItemTemplateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_item_template"
}

func (d *inventoryItemTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):An inventory item template of a device type (dcim.inventoryitemtemplate).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/inventoryitemtemplate/):\n\n> **Deprecation Warning:** Beginning in NetBox v4.3, the use of inventory items has been deprecated. They are planned for removal in a future NetBox release. Users are strongly encouraged to begin using [modules](https://netboxlabs.com/docs/netbox/models/dcim/module/) and [module types](https://netboxlabs.com/docs/netbox/models/dcim/moduletype/) in place of inventory items. Modules provide enhanced functionality and can be configured with user-defined attributes.\n>\n> A template for an inventory item that will be automatically created when instantiating a new device. All attributes of this object will be copied to the new inventory item, including the associations with a parent item and assigned component, if any. See the [inventory item](https://netboxlabs.com/docs/netbox/models/dcim/inventoryitem/) documentation for more detail.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"device_type_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device type. NetBox refuses to move a template to another type.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"parent_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the parent inventory item template.",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the inventory item role.",
			},
			"manufacturer_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the manufacturer.",
			},
			"part_id": schema.StringAttribute{
				Computed:    true,
				Description: "Manufacturer part number.",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"component_type": schema.StringAttribute{
				Computed:    true,
				Description: "Type of the component template the item is bound to, e.g. dcim.interfacetemplate. Requires component_id.",
			},
			"component_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the component template named by component_type.",
			},
			"description": schema.StringAttribute{
				Computed: true,
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
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
			},
		},
	}
}

func (d *inventoryItemTemplateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *inventoryItemTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data inventoryItemTemplateDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimInventoryItemTemplatesListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.DeviceTypeID.IsNull() {
		v := state.DeviceTypeID.ValueInt64()
		params.SetDeviceTypeID([]int64{v})
		hasInput = true
	}
	if !state.Label.IsNull() {
		v := state.Label.ValueString()
		params.SetLabel([]string{v})
		hasInput = true
	}
	if !state.RoleID.IsNull() {
		v := state.RoleID.ValueInt64()
		params.SetRoleID([]int64{v})
		hasInput = true
	}
	if !state.ManufacturerID.IsNull() {
		v := state.ManufacturerID.ValueInt64()
		params.SetManufacturerID([]int64{v})
		hasInput = true
	}
	if !state.ParentID.IsNull() {
		v := state.ParentID.ValueInt64()
		params.SetParentID([]int64{v})
		hasInput = true
	}
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or device_type_id and/or label and/or role_id and/or manufacturer_id and/or parent_id and/or name_contains to look up a netbox_inventory_item_template.")
		return
	}
	res, err := d.client.Dcim.DcimInventoryItemTemplatesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_inventory_item_template", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_inventory_item_template",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenInventoryItemTemplate(ctx, netboxapi.InventoryItemTemplateResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.inventoryItemTemplateResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
