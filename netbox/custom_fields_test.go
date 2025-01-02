package netbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomFields_normalize(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "MapNil",
			input:    nil,
			expected: map[string]interface{}{},
		},
		{
			name:     "MapEmpty",
			input:    map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			name: "MapFull",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "MapMixed",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": "",
				"key3": nil,
			},
			expected: map[string]interface{}{
				"key1": "value1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeCustomFields(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCustomFields_merge(t *testing.T) {
	tests := []struct {
		name     string
		state    map[string]interface{}
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "MapNilBoth",
			state:    nil,
			input:    nil,
			expected: map[string]interface{}{},
		},
		{
			name:     "MapNilNew",
			state:    map[string]interface{}{},
			input:    nil,
			expected: map[string]interface{}{},
		},
		{
			name:     "MapNilOld",
			state:    nil,
			input:    map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			name: "MapUnsetWithMissing",
			state: map[string]interface{}{
				"key1": "value1",
			},
			input: map[string]interface{}{},
			expected: map[string]interface{}{
				"key1": "",
			},
		},
		{
			name: "MapUnsetWithNil",
			state: map[string]interface{}{
				"key1": "value1",
			},
			input: map[string]interface{}{
				"key1": nil,
			},
			expected: map[string]interface{}{
				"key1": "",
			},
		},
		{
			name: "MapUnsetWithEmpty",
			state: map[string]interface{}{
				"key1": "value1",
			},
			input: map[string]interface{}{},
			expected: map[string]interface{}{
				"key1": "",
			},
		},
		{
			name:  "MapUnsetWithNilNotSet",
			state: map[string]interface{}{},
			input: map[string]interface{}{
				"key1": nil,
			},
			expected: map[string]interface{}{
				"key1": "",
			},
		},
		{
			name:  "MapUnsetWithEmptyNotSet",
			state: map[string]interface{}{},
			input: map[string]interface{}{
				"key1": "",
			},
			expected: map[string]interface{}{
				"key1": "",
			},
		},
		{
			name:  "MapSetNew",
			state: map[string]interface{}{},
			input: map[string]interface{}{
				"key1": "test",
			},
			expected: map[string]interface{}{
				"key1": "test",
			},
		},
		{
			name: "MapSetExisting",
			state: map[string]interface{}{
				"key1": " test",
			},
			input: map[string]interface{}{
				"key1": "testnew",
			},
			expected: map[string]interface{}{
				"key1": "testnew",
			},
		},
		{
			name: "MapMixed",
			state: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "",
				"key4": nil,
			},
			input: map[string]interface{}{
				"key1":   "valuenew1",
				"key2":   nil,
				"keynew": "valuenew2",
			},
			expected: map[string]interface{}{
				"key1":   "valuenew1",
				"key2":   "",
				"key3":   "",
				"key4":   "",
				"keynew": "valuenew2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeCustomFields(tt.state, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
