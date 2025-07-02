package netbox

import (
	"testing"
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
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := buildValidValueDescription(tt.list)
			if actual != tt.expected {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
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
		name     string
		input    string
		expected string
	}{
		{
			name:     "Incomplete",
			input:    "v1.3",
			expected: "",
		},
		{
			name:     "SimpleWithV",
			input:    "v1.2.3",
			expected: "1.2.3",
		},
		{
			name:     "SimpleWithoutV",
			input:    "1.2.3",
			expected: "1.2.3",
		},
		{
			name:     "Docker",
			input:    "v4.5.6-Docker-3.2",
			expected: "4.5.6",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := extractSemanticVersionFromString(tt.input)
			if actual != tt.expected {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}
