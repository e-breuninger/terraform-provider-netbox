package netbox

import (
	"context"
	"reflect"
	"testing"
)

func TestResourceNetboxVirtualMachineStateUpgradeV1(t *testing.T) {
	for _, tt := range []struct {
		name     string
		state    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "Zero",
			state:    map[string]interface{}{"disk_size_gb": float64(0)},
			expected: map[string]interface{}{"disk_size_mb": float64(0)},
		},
		{
			name:     "NonZero",
			state:    map[string]interface{}{"disk_size_gb": float64(123)},
			expected: map[string]interface{}{"disk_size_mb": float64(123000)},
		},
		{
			name:     "Invalid",
			state:    map[string]interface{}{"disk_size_gb": "foo"},
			expected: map[string]interface{}{"disk_size_mb": float64(0)},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := resourceNetboxVirtualMachineStateUpgradeV1(context.Background(), tt.state, nil)
			if err != nil {
				t.Fatalf("error migrating state: %s", err)
			}
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}
