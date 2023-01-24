package netbox

import (
	"testing"
)

func TestSlugGeneration(t *testing.T) {
	for _, tt := range []struct {
		name, input, expected string
	}{
		{
			name:     "LowerCase",
			input:    "FOO",
			expected: "foo",
		},
		{
			name:     "SpecialChars",
			input:    `f^o!o"§$%&/-_()=?b`,
			expected: "foo-_b",
		},
		{
			name:     "Multidash",
			input:    "--d-a---s---h------",
			expected: "d-a-s-h",
		},
		{
			name:     "Trailing",
			input:    "foo&  $",
			expected: "foo",
		},
		{
			name:     "Full",
			input:    "Foo & 33 bar -- yes-",
			expected: "foo-33-bar-yes",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := getSlug(tt.input)
			if actual != tt.expected {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}
