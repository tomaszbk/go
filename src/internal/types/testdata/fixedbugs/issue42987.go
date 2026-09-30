// Check that there is only one error (no follow-on errors).

package p
var _ = [ ... /* ERROR "invalid use of [...] array" */ ]byte("foo")