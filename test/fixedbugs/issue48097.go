// errorcheck -complete

package p

func F() // ERROR "missing function body"

//go:noescape
func f() {} // ERROR "can only use //go:noescape with external func implementations"
