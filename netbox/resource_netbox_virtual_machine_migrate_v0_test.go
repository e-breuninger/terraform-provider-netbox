package netbox

import (
	"context"
	"reflect"
	"testing"
)

func TestResourceNetboxVirtualMachineStateUpgradeV0(t *testing.T) {
	for _, tt := range []struct {
		name     string
		state    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "Empty",
			state:    map[string]interface{}{"vcpus": ""},
			expected: map[string]interface{}{"vcpus": float64(0)},
		},
		{
			name:     "Zero",
			state:    map[string]interface{}{"vcpus": "0"},
			expected: map[string]interface{}{"vcpus": float64(0)},
		},
		{
			name:     "NonZero",
			state:    map[string]interface{}{"vcpus": "123"},
			expected: map[string]interface{}{"vcpus": float64(123)},
		},
		{
			name:     "Float",
			state:    map[string]interface{}{"vcpus": "4.5"},
			expected: map[string]interface{}{"vcpus": 4.5},
		},
		{
			name:     "Invalid",
			state:    map[string]interface{}{"vcpus": "foo"},
			expected: map[string]interface{}{"vcpus": float64(0)},
		},
		{
			name:     "FloatAsFloat",
			state:    map[string]interface{}{"vcpus": 4.5},
			expected: map[string]interface{}{"vcpus": 4.5},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := resourceNetboxVirtualMachineStateUpgradeV0(context.Background(), tt.state, nil)
			if err != nil {
				t.Fatalf("error migrating state: %s", err)
			}
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}
