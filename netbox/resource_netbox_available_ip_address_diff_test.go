package netbox

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// unknownConfigValue mirrors hcl2shim.UnknownVariableValue, which lives in an
// internal package: the placeholder for a config value that is not known until
// apply.
const unknownConfigValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

func TestAvailableIPAddressAllocationSourceForceNew(t *testing.T) {
	tests := []struct {
		name           string
		statePrefixID  string
		configPrefixID interface{}
		wantForceNew   bool
	}{
		{"moved to another prefix", "3758", 3146, true},
		{"unchanged", "3758", 3758, false},
		{"imported, source unknown to state", "0", 3758, false},
		// A prefix id read from a data source whose read defers to apply.
		// Replacing on it would re-allocate the address on every plan.
		{"config value unknown at plan", "3758", unknownConfigValue, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &terraform.InstanceState{
				ID: "23",
				Attributes: map[string]string{
					"id":         "23",
					"prefix_id":  tt.statePrefixID,
					"ip_address": "10.0.0.1/24",
				},
			}
			config := terraform.NewResourceConfigRaw(map[string]interface{}{
				"prefix_id": tt.configPrefixID,
			})

			diff, err := resourceNetboxAvailableIPAddress().Diff(context.Background(), state, config, nil)
			if err != nil {
				t.Fatalf("Diff: %s", err)
			}

			forceNew := diff != nil && diff.RequiresNew()
			if forceNew != tt.wantForceNew {
				t.Errorf("RequiresNew() = %t, want %t", forceNew, tt.wantForceNew)
			}
		})
	}
}
