// compile


// Issue 30430: isGoConst returned true for non-const variables,
// resulting in ICE.

package p

func f() {
	var s string
	_ = map[string]string{s: ""}
}

const s = ""
