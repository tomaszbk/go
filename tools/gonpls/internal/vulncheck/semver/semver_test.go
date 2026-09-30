package semver

import (
	"testing"
)

func TestCanonicalize(t *testing.T) {
	for _, test := range []struct {
		v    string
		want string
	}{
		{"v1.2.3", "v1.2.3"},
		{"1.2.3", "v1.2.3"},
		{"go1.2.3", "v1.2.3"},
	} {
		got := CanonicalizeSemverPrefix(test.v)
		if got != test.want {
			t.Errorf("want %s; got %s", test.want, got)
		}
	}
}
