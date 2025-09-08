package netbox

import (
	"testing"
)

func TestValidatePositiveInt16(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{
			name:     "Valid zero",
			value:    0,
			expected: true,
		},
		{
			name:     "Valid positive",
			value:    1000,
			expected: true,
		},
		{
			name:     "Valid max int16",
			value:    maxInt16,
			expected: true,
		},
		{
			name:     "Invalid negative",
			value:    -1,
			expected: false,
		},
		{
			name:     "Invalid too large",
			value:    maxInt16 + 1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, errors := validatePositiveInt16(tt.value, "test")
			hasErrors := len(errors) > 0
			if hasErrors != !tt.expected {
				t.Errorf("validatePositiveInt16(%v) = %v, expected %v", tt.value, !hasErrors, tt.expected)
			}
		})
	}
}

func TestValidatePositiveInt32(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{
			name:     "Valid zero",
			value:    0,
			expected: true,
		},
		{
			name:     "Valid positive",
			value:    100000,
			expected: true,
		},
		{
			name:     "Valid max int32",
			value:    maxInt32,
			expected: true,
		},
		{
			name:     "Invalid negative",
			value:    -1,
			expected: false,
		},
		{
			name:     "Invalid too large",
			value:    maxInt32 + 1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, errors := validatePositiveInt32(tt.value, "test")
			hasErrors := len(errors) > 0
			if hasErrors != !tt.expected {
				t.Errorf("validatePositiveInt32(%v) = %v, expected %v", tt.value, !hasErrors, tt.expected)
			}
		})
	}
}

func TestConstants(t *testing.T) {
	// Test that constants are set correctly
	if maxUint16 != 65535 {
		t.Errorf("maxUint16 = %d, expected 65535", maxUint16)
	}
	if maxInt16 != 32767 {
		t.Errorf("maxInt16 = %d, expected 32767", maxInt16)
	}
	if maxUint32 != 4294967295 {
		t.Errorf("maxUint32 = %d, expected 4294967295", maxUint32)
	}
	if maxInt32 != 2147483647 {
		t.Errorf("maxInt32 = %d, expected 2147483647", maxInt32)
	}
}
