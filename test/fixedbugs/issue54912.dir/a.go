// Test that inlining a function literal that captures both a type
// switch case variable and another local variable works correctly.

package a

func F(p *int, x any) func() {
	switch x := x.(type) {
	case int:
		return func() {
			*p += x
		}
	}
	return nil
}
