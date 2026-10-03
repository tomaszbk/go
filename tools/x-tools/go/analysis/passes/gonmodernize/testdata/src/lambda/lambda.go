package lambda

type callback func(int) int

var typed func(int) int = func(x int) int { // want "function literal can use Gon lambda syntax"
	// The body and its comments survive unchanged.
	defer func() {}()
	return x + 1
}

var named callback = func(x int) int { return x } // want "function literal can use Gon lambda syntax"

var grouped func(int, int) int = func(a, b int) int { return a + b } // want "function literal can use Gon lambda syntax"

var blank func(int) = func(_ int) {} // want "function literal can use Gon lambda syntax"

var zero func() = func() {} // want "function literal can use Gon lambda syntax"

var variadic func(...int) int = func(xs ...int) int { return len(xs) } // want "function literal can use Gon lambda syntax"

var nested func() int = func() int { // want "function literal can use Gon lambda syntax"
	var inner func(int) int = func(x int) int { return x + 1 } // want "function literal can use Gon lambda syntax"
	return inner(2)
}

func consume(callback)                       {}
func consumers(...callback)                  {}
func generic[T any](func(T) T)               {}
func genericPair[T any](T, func(int) int)    {}
func acceptsInterface(any)                   {}
func acceptsVariadic(func(...int) int)       {}
func acceptsResult(func() (int, error))      {}
func acceptsNamedResult(func() (result int)) {}
func acceptsNoResult(func())                 {}
func acceptsNamedCallback(f callback)        {}

func calls() {
	consume(func(x int) int { return x + 2 })               // want "function literal can use Gon lambda syntax"
	consumers(func(x int) int { return x + 3 })             // want "function literal can use Gon lambda syntax"
	consumers(nil, func(x int) int { return x + 4 })        // want "function literal can use Gon lambda syntax"
	acceptsVariadic(func(xs ...int) int { return len(xs) }) // want "function literal can use Gon lambda syntax"
	acceptsResult(func() (int, error) { return 1, nil })    // want "function literal can use Gon lambda syntax"
	acceptsNamedResult(func() int { return 1 })             // want "function literal can use Gon lambda syntax"
	acceptsNoResult(func() {})                              // want "function literal can use Gon lambda syntax"
	acceptsNamedCallback(func(x int) int { return x })      // want "function literal can use Gon lambda syntax"
	go consume(func(x int) int { return x })                // want "function literal can use Gon lambda syntax"
	defer consume(func(x int) int { return x })             // want "function literal can use Gon lambda syntax"

	// An inferred generic signature is deliberately not a rewrite target,
	// even when every argument is explicitly instantiated.
	generic(func(x int) int { return x })
	generic[int](func(x int) int { return x })
	genericPair(1, func(x int) int { return x })
	acceptsInterface(func(x int) int { return x })
	consume((func(x int) int { return x }))
	consume(func(int) int { return 1 })
	consume(func(x int) (result int) { result = x; return })
	consume(func(x /* parameter */ int) int { return x })
	consume(func(x int) int /* before block */ { return x })
	_ = callback(func(x int) int { return x })
	_ = func(x int) int { return x }(1)
}

func assignment() {
	var f func(int) int
	f = func(x int) int { return x + 1 }                                          // want "function literal can use Gon lambda syntax"
	f, typed = func(x int) int { return x + 2 }, func(x int) int { return x + 3 } // want "function literal can use Gon lambda syntax" "function literal can use Gon lambda syntax"
	_ = f

	var callbacks [1]callback
	callbacks[0] = func(x int) int { return x } // want "function literal can use Gon lambda syntax"
	var holder struct{ f callback }
	holder.f = func(x int) int { return x } // want "function literal can use Gon lambda syntax"

	var inferred = func(x int) int { return x }
	local := func(x int) int { return x }
	var iface any = func(x int) int { return x }
	iface = func(x int) int { return x }
	_, _, _ = inferred, local, iface
}

func typeParameter[T any]() {
	var f func(T) T = func(x T) T { return x } // want "function literal can use Gon lambda syntax"
	_ = f
}
