package provider

import "testing"

func TestVersionAtLeast(t *testing.T) {
	for _, c := range []struct {
		version string
		want    bool
	}{
		{"4.7.0", true}, {"4.7.5", true}, {"4.8.1", true}, {"5.0.0", true},
		{"4.6.10", false}, {"3.9.9", false}, {"", false}, {"dev", false}, {"4", false},
	} {
		if got := versionAtLeast(c.version, [2]int{4, 7}); got != c.want {
			t.Errorf("versionAtLeast(%q, 4.7) = %v, want %v", c.version, got, c.want)
		}
	}
}
