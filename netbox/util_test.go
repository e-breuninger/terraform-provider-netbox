package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestJoinStringWithFinalConjunction(t *testing.T) {
	for _, tt := range []struct {
		name     string
		list     []string
		sep      string
		con      string
		expected string
	}{
		{
			name:     "Full",
			list:     []string{"foo", "bar", "baz"},
			sep:      ", ",
			con:      "and",
			expected: "foo, bar and baz",
		},
		{
			name:     "OnlyTwoItems",
			list:     []string{"foo", "bar"},
			sep:      ", ",
			con:      "and",
			expected: "foo and bar",
		},
		{
			name:     "OneItem",
			list:     []string{"foo"},
			sep:      ", ",
			con:      "and",
			expected: "foo",
		},
		{
			name:     "Empty",
			list:     []string{},
			sep:      ", ",
			con:      "and",
			expected: "",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinStringWithFinalConjunction(tt.list, tt.sep, tt.con)
			if actual != tt.expected {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}

func TestBuildValidValueDescription(t *testing.T) {
	for _, tt := range []struct {
		name     string
		list     []string
		expected string
	}{
		{
			name:     "Full",
			list:     []string{"foo", "bar", "baz"},
			expected: "Valid values are `foo`, `bar` and `baz`",
		},
		{
			name:     "TwoItems",
			list:     []string{"foo", "bar"},
			expected: "Valid values are `foo` and `bar`",
		},
		{
			name:     "OneItem",
			list:     []string{"foo"},
			expected: "Valid values are `foo`",
		},
		{
			name:     "Empty",
			list:     []string{},
			expected: "Valid values are ",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := buildValidValueDescription(tt.list)
			if actual != tt.expected {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}

func TestStrToPtr(t *testing.T) {
	input := "test"
	result := strToPtr(input)
	if *result != input {
		t.Fatalf("expected %q, got %q", input, *result)
	}
}

func TestInt64ToPtr(t *testing.T) {
	input := int64(42)
	result := int64ToPtr(input)
	if *result != input {
		t.Fatalf("expected %d, got %d", input, *result)
	}
}

func TestFloat64ToPtr(t *testing.T) {
	input := 3.14
	result := float64ToPtr(input)
	if *result != input {
		t.Fatalf("expected %f, got %f", input, *result)
	}
}

func TestToStringList(t *testing.T) {
	set := schema.NewSet(schema.HashString, []interface{}{"a", "b", "c"})
	result := toStringList(set)
	// Since sets are unordered, we need to check that all expected values are present
	expected := map[string]bool{"a": true, "b": true, "c": true}
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for _, v := range result {
		if !expected[v] {
			t.Fatalf("unexpected value %q", v)
		}
	}
}

func TestToInt64List(t *testing.T) {
	// Use a custom hash function that can handle both int and int64
	hashFunc := func(v interface{}) int {
		switch val := v.(type) {
		case int:
			return val
		case int64:
			return int(val)
		default:
			return 0
		}
	}
	set := schema.NewSet(hashFunc, []interface{}{int(1), int64(2), int(3)})
	result := toInt64List(set)
	// Since sets are unordered, we need to check that all expected values are present
	expected := map[int64]bool{1: true, 2: true, 3: true}
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for _, v := range result {
		if !expected[v] {
			t.Fatalf("unexpected value %d", v)
		}
	}
}

func TestToInt64PtrList(t *testing.T) {
	// Use a custom hash function that can handle both int and int64
	hashFunc := func(v interface{}) int {
		switch val := v.(type) {
		case int:
			return val
		case int64:
			return int(val)
		default:
			return 0
		}
	}
	set := schema.NewSet(hashFunc, []interface{}{int(1), int64(2)})
	result := toInt64PtrList(set)
	// Since sets are unordered, we need to check that all expected values are present
	expected := map[int64]bool{1: true, 2: true}
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for _, v := range result {
		if !expected[*v] {
			t.Fatalf("unexpected value %d", *v)
		}
	}
}

func TestGetOptionalStr(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"test_key": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]interface{}{
		"test_key": "value",
	})

	result := getOptionalStr(d, "test_key", false)
	if result != "value" {
		t.Fatalf("expected 'value', got %q", result)
	}

	// Test with key not set
	d2 := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"test_key": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]interface{}{})

	result2 := getOptionalStr(d2, "test_key", false)
	if result2 != "" {
		t.Fatalf("expected empty string, got %q", result2)
	}
}

func TestGetOptionalInt(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"test_key": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}, map[string]interface{}{
		"test_key": 42,
	})

	result := getOptionalInt(d, "test_key")
	if result == nil || *result != 42 {
		t.Fatalf("expected 42, got %v", result)
	}

	// Test with key not set
	d2 := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"test_key": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}, map[string]interface{}{})

	result2 := getOptionalInt(d2, "test_key")
	if result2 != nil {
		t.Fatalf("expected nil, got %v", result2)
	}
}

func TestGetOptionalFloat(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"test_key": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
	}, map[string]interface{}{
		"test_key": 3.14,
	})

	result := getOptionalFloat(d, "test_key")
	if result == nil || *result != 3.14 {
		t.Fatalf("expected 3.14, got %v", result)
	}

	// Test with key not set
	d2 := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"test_key": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
	}, map[string]interface{}{})

	result2 := getOptionalFloat(d2, "test_key")
	if result2 != nil {
		t.Fatalf("expected nil, got %v", result2)
	}
}

func TestJsonSemanticCompareEqual(t *testing.T) {
	a := `{"a": [{ "b": [1, 2, 3]}]}`
	b := `{"a":[{"b":[1,2,3]}]}`

	equal, err := jsonSemanticCompare(a, b)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !equal {
		t.Errorf("expected 'a' and 'b' to be semantically equal\n\na: %s\nb: %s\n", a, b)
	}
}

func TestJsonSemanticCompareUnequal(t *testing.T) {
	a := `{"a": [{ "b": [1, 2, 3]}]}`
	b := `{"a": [{ "b": [1, 2, 4]}]}`

	equal, err := jsonSemanticCompare(a, b)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if equal {
		t.Errorf("expected 'a' and 'b' to be semantically unequal\n\na: %s\nb: %s\n", a, b)
	}
}

func TestExtractSemanticVersionFromString(t *testing.T) {
	for _, tt := range []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Incomplete",
			input:       "v1.3",
			expected:    "",
			expectError: true,
		},
		{
			name:        "SimpleWithV",
			input:       "v1.2.3",
			expected:    "1.2.3",
			expectError: false,
		},
		{
			name:        "SimpleWithoutV",
			input:       "1.2.3",
			expected:    "1.2.3",
			expectError: false,
		},
		{
			name:        "Docker",
			input:       "v4.5.6-Docker-3.2",
			expected:    "4.5.6",
			expectError: false,
		},
		{
			name:        "EmptyString",
			input:       "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "NoVersion",
			input:       "some-random-string",
			expected:    "",
			expectError: true,
		},
		{
			name:        "ComplexVersion",
			input:       "v10.20.30-beta.1+build.2",
			expected:    "10.20.30",
			expectError: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := extractSemanticVersionFromString(tt.input)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if actual != tt.expected {
					t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
				}
			}
		})
	}
}
