// errorcheck -d=panic


// Verify that we get "use of .(type) outside type switch"
// before any other (misleading) errors. Test case from issue.

package p

func f(i interface{}) {
	if x, ok := i.(type); ok { // ERROR "assignment mismatch|outside type switch"
		_ = x
	}
}
