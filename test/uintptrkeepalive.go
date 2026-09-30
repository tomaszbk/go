// errorcheck -std

package p

//go:uintptrkeepalive
func missingNosplit(uintptr) { // ERROR "go:uintptrkeepalive requires go:nosplit"
}
