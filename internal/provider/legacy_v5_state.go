package provider

// State written by provider v5 that the v6 schemas cannot read as is:
//   - netbox_device and netbox_custom_field stored an unset JSON attribute
//     (local_context_data, default) as "", which jsontypes.Normalized rejects
//     ("Invalid JSON String Value"), so every plan of a migrated workspace fails;
//   - netbox_primary_ip was VM-only and its id was the bare VM id, not "vm/<id>";
//   - netbox_interface was renamed to netbox_virtual_machine_interface, so a
//     moved block needs a state mover.
// The device and custom field upgraders keep only the id and let the following
// Read fill in the rest, the same as the virtual machine upgrader.

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
)

var (
	_ resource.ResourceWithUpgradeState = (*deviceResource)(nil)
	_ resource.ResourceWithUpgradeState = (*customFieldResource)(nil)
	_ resource.ResourceWithUpgradeState = (*primaryIPResource)(nil)
	_ resource.ResourceWithMoveState    = (*virtualMachineInterfaceResource)(nil)
)

// priorNumericID reads the id of raw prior state; v5 stored it as a string.
func priorNumericID(raw []byte) (int64, error) {
	var prior struct {
		ID any `json:"id"`
	}
	if err := json.Unmarshal(raw, &prior); err != nil {
		return 0, err
	}
	switch v := prior.ID.(type) {
	case float64:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("unexpected id type %T", prior.ID)
	}
}

// idOnlyUpgrader replaces the prior state with its id alone.
func idOnlyUpgrader(schemaFunc func(context.Context, resource.SchemaRequest, *resource.SchemaResponse)) resource.StateUpgrader {
	return resource.StateUpgrader{
		StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
			if req.RawState == nil {
				resp.Diagnostics.AddError("Unable to upgrade the prior state", "The prior state is empty.")
				return
			}
			id, err := priorNumericID(req.RawState.JSON)
			if err != nil {
				resp.Diagnostics.AddError("Unable to upgrade the prior state", "The prior state carries no numeric id: "+err.Error())
				return
			}
			var sr resource.SchemaResponse
			schemaFunc(ctx, resource.SchemaRequest{}, &sr)
			obj := conv.IDOnlyObject(ctx, sr.Schema.Type().(types.ObjectType), "id", types.Int64Value(id))
			resp.Diagnostics.Append(resp.State.Set(ctx, obj)...)
		},
	}
}

func (r *deviceResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{0: idOnlyUpgrader(r.Schema)}
}

func (r *customFieldResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{0: idOnlyUpgrader(r.Schema)}
}

// UpgradeState turns a v5 id (the bare virtual machine id) into "vm/<id>". State
// written by v6 at version 0 already has the prefixed id and passes through.
func (r *primaryIPResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior struct {
					ID               string `json:"id"`
					DeviceID         *int64 `json:"device_id"`
					VirtualMachineID *int64 `json:"virtual_machine_id"`
					IPAddressID      *int64 `json:"ip_address_id"`
					Family           *int64 `json:"ip_address_version"`
				}
				if req.RawState == nil || json.Unmarshal(req.RawState.JSON, &prior) != nil {
					resp.Diagnostics.AddError("Unable to upgrade the prior state", "The prior state of netbox_primary_ip could not be read.")
					return
				}
				id := prior.ID
				if _, err := parsePrimaryTarget(id); err != nil {
					if prior.VirtualMachineID == nil {
						resp.Diagnostics.AddError("Unable to upgrade the prior state",
							fmt.Sprintf("The prior id %q is neither device/<id> nor vm/<id>, and no virtual_machine_id is set.", id))
						return
					}
					id = primaryTarget{kind: "vm", id: *prior.VirtualMachineID}.String()
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, primaryIPResourceModel{
					ID:               types.StringValue(id),
					DeviceID:         types.Int64PointerValue(prior.DeviceID),
					VirtualMachineID: types.Int64PointerValue(prior.VirtualMachineID),
					IPAddressID:      types.Int64PointerValue(prior.IPAddressID),
					Family:           types.Int64PointerValue(prior.Family),
				})...)
			},
		},
	}
}

// MoveState accepts netbox_interface, v5's name for this resource, as the
// source of a moved block. Only the id is carried over; Read fills in the rest.
func (r *virtualMachineInterfaceResource) MoveState(ctx context.Context) []resource.StateMover {
	return []resource.StateMover{{
		StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
			if req.SourceTypeName != "netbox_interface" || !strings.HasSuffix(req.SourceProviderAddress, "/netbox") {
				return
			}
			if req.SourceRawState == nil {
				resp.Diagnostics.AddError("Unable to move the prior state", "The netbox_interface state is empty.")
				return
			}
			id, err := priorNumericID(req.SourceRawState.JSON)
			if err != nil {
				resp.Diagnostics.AddError("Unable to move the prior state", "The netbox_interface state carries no numeric id: "+err.Error())
				return
			}
			var sr resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &sr)
			obj := conv.IDOnlyObject(ctx, sr.Schema.Type().(types.ObjectType), "id", types.Int64Value(id))
			resp.Diagnostics.Append(resp.TargetState.Set(ctx, obj)...)
		},
	}}
}
