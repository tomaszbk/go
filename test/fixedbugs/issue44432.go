// errorcheck -d=panic

package p

var m = map[string]int{
	"a": 1,
	1:   1, // ERROR "cannot use 1.*as.*string.*in map"
	2:   2, // ERROR "cannot use 2.*as.*string.*in map"
}
