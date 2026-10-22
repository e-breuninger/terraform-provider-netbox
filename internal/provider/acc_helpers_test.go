//go:build acctest

// Check helpers shared by the suite's test files.
package provider_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// testAccStateID returns the NetBox id of a resource in state.
func testAccStateID(state *terraform.State, name string) (int64, error) {
	resource := state.RootModule().Resources[name]
	if resource == nil {
		return 0, fmt.Errorf("%s not in state", name)
	}
	return strconv.ParseInt(resource.Primary.ID, 10, 64)
}

// stateID remembers a resource's id from one step so a later step's PreConfig can act on it.
func stateID(name string, into *int64) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resource, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("%s not in state", name)
		}
		_, err := fmt.Sscan(resource.Primary.ID, into)
		return err
	}
}

// testAccCheckConfigContextContains asserts that the resource's config_context (rendered server
// side from scoped netbox_config_context objects, merged with local_context_data) decodes as JSON
// and contains at least the given key/value pairs, regardless of what else it holds or how the
// server orders its keys.
func testAccCheckConfigContextContains(resourceName string, want map[string]any) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(resourceName, "config_context", func(value string) error {
		var got map[string]any
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			return fmt.Errorf("%s.config_context is not valid JSON (%v): %s", resourceName, err, value)
		}
		for k, v := range want {
			if !reflect.DeepEqual(got[k], v) {
				return fmt.Errorf("%s.config_context[%q] = %v, want %v (full value: %s)", resourceName, k, got[k], v, value)
			}
		}
		return nil
	})
}

// testAccCheckStateIDIs asserts that the resource's NetBox id equals (same) or differs from (!same)
// *want, read at check time.
func testAccCheckStateIDIs(name string, want *int64, same bool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		id, err := testAccStateID(state, name)
		if err != nil {
			return err
		}
		if same && id != *want {
			return fmt.Errorf("expected %s to keep NetBox id %d, got %d", name, *want, id)
		}
		if !same && id == *want {
			return fmt.Errorf("expected %s to be replaced, but it still has NetBox id %d", name, *want)
		}
		return nil
	}
}
