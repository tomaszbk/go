// compile

// Make sure VARDEF can be a top-level statement.

package p

func f() {
	var s string
	var as []string
	switch false && (s+"a"+as[0]+s+as[0]+s == "") {
	}
}
