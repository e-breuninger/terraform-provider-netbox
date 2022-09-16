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
