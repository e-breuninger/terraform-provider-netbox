package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestVirtualMachineCustomFieldsAreNotSentForUnrelatedUpdates(t *testing.T) {
	t.Parallel()

	d := schema.TestResourceDataRaw(t, resourceNetboxVirtualMachine().Schema, map[string]interface{}{
		"name":       "vm01",
		"cluster_id": 1,
	})

	// A value loaded from state is present, but no custom-field configuration
	// change occurred. The update path must therefore omit custom_fields.
	require.Nil(t, getCustomFieldsForUpdate(d))

	d.Set("description", "changed")
	require.Nil(t, getCustomFieldsForUpdate(d))

	changed := schema.TestResourceDataRaw(t, resourceNetboxVirtualMachine().Schema, map[string]interface{}{
		"name":        "vm01",
		"cluster_id":  1,
		customFieldsKey: map[string]interface{}{"pve_balloon": "2"},
	})
	require.Equal(t, map[string]interface{}{"pve_balloon": "2"}, getCustomFieldsForUpdate(changed))
}
