package netbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomFields(t *testing.T) {
	// Test with valid map
	input := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}
	result := getCustomFields(input)
	assert.Equal(t, input, result)

	// Test with empty map
	emptyInput := map[string]interface{}{}
	result2 := getCustomFields(emptyInput)
	assert.Nil(t, result2)

	// Test with nil
	result3 := getCustomFields(nil)
	assert.Nil(t, result3)

	// Test with non-map type
	result4 := getCustomFields("not a map")
	assert.Nil(t, result4)
}
