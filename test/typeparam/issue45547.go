// compile


package p

func f[T any]() (f, g T) { return f, g }

// Tests for generic function instantiation on the right hande side of multi-value
// assignments.

func g() {
	// Multi-value assignment within a function
	var _, _ = f[int]()
}

// Multi-value assignment outside a function.
var _, _ = f[int]()
