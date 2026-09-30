// Check correctness of various closure corner cases
// that are expected to be inlined

package a

func f() bool               { return true }
func G() func() func() bool { return func() func() bool { return f } }
