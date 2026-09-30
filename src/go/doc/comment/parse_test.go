package comment

import "testing"

// See https://golang.org/issue/52353
func Test52353(t *testing.T) {
	ident("𫕐ﯯ")
}
