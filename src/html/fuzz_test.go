package html

import "testing"

func FuzzEscapeUnescape(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		e := EscapeString(v)
		u := UnescapeString(e)
		if u != v {
			t.Errorf("EscapeString(%q) = %q, UnescapeString(%q) = %q, want %q", v, e, e, u, v)
		}

		// As per the documentation, this isn't always equal to v, so it makes
		// no sense to check for equality. It can still be interesting to find
		// panics in it though.
		EscapeString(UnescapeString(v))
	})
}
