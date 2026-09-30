package morestrings

import "strings"

// CutLast is the "last" analogue of [strings.Cut].
func CutLast(s, sep string) (before, after string, ok bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
