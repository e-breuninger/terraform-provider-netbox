package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type objectRefValues struct {
	typ        types.String
	id         types.Int64
	device, vm types.Int64
}

func (objectRefValues *objectRefValues) ref() objectRef {
	return objectRef{
		typeName: "assigned_object_type", idName: "assigned_object_id",
		typ: &objectRefValues.typ, id: &objectRefValues.id,
		aliases: []objectRefAlias{
			{name: "device_interface_id", contentType: "dcim.interface", id: &objectRefValues.device},
			{name: "virtual_machine_interface_id", contentType: "virtualization.vminterface", id: &objectRefValues.vm},
		},
	}
}

var (
	nullStr    = types.StringNull()
	nullInt    = types.Int64Null()
	unknownInt = types.Int64Unknown()
	devType    = types.StringValue("dcim.interface")
	vmType     = types.StringValue("virtualization.vminterface")
	fhrpType   = types.StringValue("ipam.fhrpgroup")
	seven      = types.Int64Value(7)
)

func TestObjectRefPlan(t *testing.T) {
	tests := []struct {
		name      string
		config    objectRefValues
		want      objectRefValues
		wantError string
	}{
		{"nothing set", objectRefValues{nullStr, nullInt, nullInt, nullInt}, objectRefValues{nullStr, nullInt, nullInt, nullInt}, ""},
		{"device alias", objectRefValues{nullStr, nullInt, seven, nullInt}, objectRefValues{devType, seven, seven, nullInt}, ""},
		{"vm alias", objectRefValues{nullStr, nullInt, nullInt, seven}, objectRefValues{vmType, seven, nullInt, seven}, ""},
		{"unknown alias keeps id unknown", objectRefValues{nullStr, nullInt, nullInt, unknownInt}, objectRefValues{vmType, unknownInt, nullInt, unknownInt}, ""},
		{"explicit pair derives alias", objectRefValues{vmType, seven, nullInt, nullInt}, objectRefValues{vmType, seven, nullInt, seven}, ""},
		{"explicit pair without alias", objectRefValues{fhrpType, seven, nullInt, nullInt}, objectRefValues{fhrpType, seven, nullInt, nullInt}, ""},
		{"explicit pair with unknown type", objectRefValues{types.StringUnknown(), seven, nullInt, nullInt}, objectRefValues{types.StringUnknown(), seven, unknownInt, unknownInt}, ""},
		{"two aliases", objectRefValues{nullStr, nullInt, seven, seven}, objectRefValues{}, "Conflicting assignment"},
		{"alias next to pair", objectRefValues{vmType, seven, nullInt, seven}, objectRefValues{}, "Conflicting assignment"},
		{"half a pair", objectRefValues{vmType, nullInt, nullInt, nullInt}, objectRefValues{}, "Incomplete assignment"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// The plan starts as the configuration (what the framework proposes for known values).
			plan := test.config
			var diags diag.Diagnostics
			plan.ref().plan(test.config.ref(), &diags)
			if test.wantError != "" {
				if !diags.HasError() || diags.Errors()[0].Summary() != test.wantError {
					t.Fatalf("want error %q, got %v", test.wantError, diags)
				}
				return
			}
			if diags.HasError() {
				t.Fatalf("unexpected error: %v", diags)
			}
			if plan != test.want {
				t.Errorf("plan = %+v, want %+v", plan, test.want)
			}
		})
	}
}

func TestObjectRefToPair(t *testing.T) {
	m := objectRefValues{nullStr, unknownInt, seven, nullInt}
	m.ref().toPair()
	if want := (objectRefValues{devType, seven, seven, nullInt}); m != want {
		t.Errorf("toPair = %+v, want %+v", m, want)
	}
	m = objectRefValues{vmType, unknownInt, nullInt, unknownInt}
	m.ref().toPair()
	if want := (objectRefValues{vmType, unknownInt, nullInt, unknownInt}); m != want {
		t.Errorf("toPair with unknown alias = %+v, want %+v", m, want)
	}
}

func TestObjectRefFromPair(t *testing.T) {
	tests := []struct{ in, want objectRefValues }{
		{objectRefValues{devType, seven, nullInt, seven}, objectRefValues{devType, seven, seven, nullInt}},
		{objectRefValues{vmType, seven, seven, nullInt}, objectRefValues{vmType, seven, nullInt, seven}},
		{objectRefValues{fhrpType, seven, seven, seven}, objectRefValues{fhrpType, seven, nullInt, nullInt}},
		{objectRefValues{nullStr, nullInt, seven, seven}, objectRefValues{nullStr, nullInt, nullInt, nullInt}},
	}
	for _, test := range tests {
		m := test.in
		m.ref().fromPair()
		if m != test.want {
			t.Errorf("fromPair(%+v) = %+v, want %+v", test.in, m, test.want)
		}
	}
}
