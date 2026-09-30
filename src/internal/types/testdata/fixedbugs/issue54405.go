// Test that we don't see spurious errors for ==
// for values with invalid types due to prior errors.

package p

var x struct {
	f *NotAType /* ERROR "undefined" */
}
var _ = x.f == nil // no error expected here

var y *NotAType  /* ERROR "undefined" */
var _ = y == nil // no error expected here
