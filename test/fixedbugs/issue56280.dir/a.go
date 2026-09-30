package a

func F() { // ERROR "can inline F"
	g(0) // ERROR "inlining call to g\[go.shape.int\]"
}

func g[T any](_ T) {} // ERROR "can inline g\[int\]" "can inline g\[go.shape.int\]" "inlining call to g\[go.shape.int\]"
