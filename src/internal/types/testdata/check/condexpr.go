// Conditional expressions: if cond { a } else { b }.

package condexpr

import (
	"bytes"
	"io"
	"os"
)

type (
	MyBool bool
	MyInt  int
	MyErr  struct{}
	Point  struct{ X, Y int }
)

func (*MyErr) Error() string { return "" }

var (
	c   bool
	mb  MyBool
	i   int
	i8  int8
	f   float64
	s   string
	p   *int
	e   error
	me  *MyErr
	sl  []int
	arr [2]int
	m   map[string]int
	ch  chan int
	fl  *os.File
	buf *bytes.Buffer
	u   uint
	rn  rune
)

func two() (int, int)      { return 1, 2 }
func none()                {}
func ee() (int, error)     { return 0, nil }
func errOnly() error       { return nil }
func takeErr(error)        {}
func takeAny(...any)       {}
func takeInts(...int)      {}
func gen[T any](x T) T     { return x }
func gen2[T any](T, error) {}

// Conditions

var _ = if c { 1 } else { 2 }
var _ = if mb { 1 } else { 2 }
var _ = if i == 0 { 1 } else { 2 }
var _ = if i /* ERROR "non-boolean condition in conditional expression" */ { 1 } else { 2 }
var _ = if "" /* ERROR "non-boolean condition in conditional expression" */ { 1 } else { 2 }
var _ = if nil /* ERROR "non-boolean condition in conditional expression" */ { 1 } else { 2 }
var _ = if (Point{}) == (Point{1, 2}) { 1 } else { 2 }

// Result type without target type

func _() {
	x1 := if c { i } else { i }
	var _ int = x1
	x2 := if c { 1 } else { 2.5 }
	var _ float64 = x2
	var _ int = x2 /* ERROR "cannot use x2 (variable of type float64) as int value" */
	x3 := if c { i } else { 0 }
	var _ int = x3
	x4 := if c { 0 } else { f }
	var _ float64 = x4
	x5 := if c { p } else { nil }
	var _ *int = x5
	x6 := if c { nil } else { e }
	var _ error = x6
	x7 := if c { me } else { nil }
	var _ *MyErr = x7
	x8 := if c { 'a' } else { 1 }
	var _ rune = x8
	x9 := if c { i == 0 } else { true }
	var _ bool = x9
	x10 := if c { mb } else { i == 0 }
	var _ MyBool = x10
	x11 := if c { sl } else { nil }
	var _ []int = x11
	x12 := if c { any(1) } else { 2 }
	var _ any = x12

	_ = if /* ERROR "mismatched types int and float64 in conditional expression" */ c { i } else { f }
	_ = if /* ERROR "mismatched types int and MyInt in conditional expression" */ c { i } else { MyInt(0) }
	_ = if /* ERROR "mismatched types int and untyped string in conditional expression" */ c { i } else { "a" }
	_ = if /* ERROR "mismatched types untyped string and int in conditional expression" */ c { "a" } else { i }
	_ = if /* ERROR "mismatched types untyped int and untyped string in conditional expression" */ c { 1 } else { "a" }
	_ = if /* ERROR "mismatched types untyped bool and untyped int in conditional expression" */ c { true } else { 1 }
	_ = if /* ERROR "mismatched types untyped int and untyped nil in conditional expression" */ c { 1 } else { nil }
	_ = if /* ERROR "mismatched types int and untyped nil in conditional expression" */ c { i } else { nil }
	_ = if /* ERROR "mismatched types *int and untyped int in conditional expression" */ c { p } else { 0 }
	_ = if /* ERROR "mismatched types error and untyped int in conditional expression" */ c { e } else { 0 }
	_ = if /* ERROR "mismatched types string and untyped rune in conditional expression" */ c { s } else { 'a' }
	_ = if /* ERROR "use of untyped nil in conditional expression" */ c { nil } else { nil }
	_ = p == if /* ERROR "use of untyped nil in conditional expression" */ c { nil } else { nil }

	// The rule is strict: an interface does not absorb an implementation.
	_ = if /* ERROR "mismatched types io.Reader and *bytes.Buffer in conditional expression" */ c { io.Reader(fl) } else { buf }
	_ = if /* ERROR "mismatched types error and *MyErr in conditional expression" */ c { e } else { me }

	// An untyped constant branch must be representable.
	_ = if c { i8 } else { 1000 /* ERROR "cannot use 1000 (untyped int constant) as int8 value in conditional expression (overflows)" */ }
	_ = if c { 2.5 /* ERROR "cannot use 2.5 (untyped float constant) as int value in conditional expression (truncated)" */ } else { i }
}

// Result type with target type

func _() {
	var _ io.Reader = if c { fl } else { buf }
	var _ error = if c { me } else { nil }
	var _ float32 = if c { 1 } else { 2.5 }
	var _ MyInt = if c { 1 } else { 2 }
	var _ []int = if c { nil } else { nil }
	var _ *int = (if c { nil } else { nil })
	var _ int = if c { 1 } else { "a" /* ERROR "cannot use \"a\" (untyped string constant) as int value in conditional expression" */ }
	var _ int = if c { f /* ERROR "cannot use f (variable of type float64) as int value in conditional expression" */ } else { 0 }
	var _ byte = if c { 1 } else { 1000 /* ERROR "cannot use 1000 (untyped int constant) as byte value in conditional expression (overflows)" */ }

	var r io.Reader
	r = if c { fl } else { buf }
	r, e = if c { fl } else { buf }, if c { me } else { nil }
	_ = r

	takeErr(if c { me } else { nil })
	takeInts(if c { 1 } else { 2 }, 3)
	takeInts(if c { sl } else { nil }...)
	takeInts(if c { 1 } else { "a" /* ERROR "cannot use \"a\" (untyped string constant) as int value in conditional expression" */ })

	ch <- if c { 1 } else { 2 }
	m[if c { "a" } else { "b" }] = if c { 1 } else { 2 }
	_ = []error{if c { me } else { nil }, nil}
	_ = struct{ E error }{E: if c { me } else { nil }}
	_ = []Point{if c { {1, 2} } else { {3, 4} }}
	var _ Point = if c { {1, 2} } else { {} }
	_ = func() error { return if c { me } else { nil } }
	_ = func() (int, error) { return if c { 1 } else { 2 }, if c { me } else { nil } }

	// A generic function may be instantiated from the target type.
	var _ func(int) int = if c { gen } else { gen[int] }
	_ = if c { gen /* ERROR "cannot use generic function gen without instantiation" */ } else { gen[int] }
}

// With an interface target, an untyped branch other than nil first gets the
// type it would have without a target; then each branch is converted to the
// target. Typed branches keep their types.

func _() {
	var _ any = if c { 1 } else { 2.5 }
	var _ any = if c { f } else { 0 }
	var _ any = if c { 1 } else { 'a' }
	var _ any = if c { e } else { 0 }
	var _ any = if c { any(1) } else { "a" }
	var _ any = if c { 1 } else { nil }
	var _ any = if c { nil } else { nil }
	var _ any = if c { i == 0 } else { true }
	var _ any = if /* ERROR "mismatched types untyped int and untyped string in conditional expression" */ c { 1 } else { "a" }
	var _ any = if /* ERROR "mismatched types int and untyped string in conditional expression" */ c { i } else { "a" }
	var _ any = if c { i8 } else { 1000 /* ERROR "cannot use 1000 (untyped int constant) as int8 value in conditional expression (overflows)" */ }
	var _ any = if c { 1 /* ERROR "cannot use 1 << 100 (untyped int constant 1267650600228229401496703205376) as int value in conditional expression (overflows)" */ << 100 } else { 2 }
	var _ any = if c { 1 /* ERROR "shifted operand 1 (type float64) must be integer" */ << u } else { 2.5 }
	var _ io.Reader = if /* ERROR "mismatched types *os.File and untyped int in conditional expression" */ c { fl } else { 1 }
	var _ io.Reader = if c { nil } else { 1 /* ERROR "cannot use 1 (constant of type int) as io.Reader value in conditional expression: int does not implement io.Reader" */ }
	var _ error = if c { me } else { nil }

	takeAny(if c { f } else { 0 }, if c { nil } else { 2.5 })
	_ = []any{if c { 1 } else { 2.5 }}
	_ = map[error]any{if c { me } else { nil }: if c { 1 } else { 'a' }}
}

// Conversions distribute over the branches: T(if c { a } else { b }) is
// if c { T(a) } else { T(b) }.

func _() {
	_ = MyInt(if c { i } else { 0 })
	_ = MyInt(if c { MyInt(1) } else { 2 })
	_ = float64(if c { 1 } else { 2 })
	_ = float64(if c { i } else { f })
	_ = error(if c { me } else { nil })
	_ = (*int)(if c { nil } else { nil })
	_ = string(if c { rn } else { 'b' })
	_ = []byte(if c { s } else { "b" })
	_ = []byte(if c { "a" } else { "b" })
	_ = []rune(if c { "a" } else { "b" })
	_ = []byte("x" + if c { "a" } else { "b" })
	_ = any(if c { 1 } else { 2.5 })
	_ = Point(if c { {1, 2} } else { {} })
	_ = (func(int) int)(if c { gen } else { gen[int] })
	_ = int8(if c { i8 } else { 1000 /* ERROR "cannot convert 1000 (untyped int constant) to type int8: 1000 overflows int8" */ })
	_ = int8(if c { 1 } else { 1000 /* ERROR "overflows" */ })
	_ = int(if c { "a" /* ERROR "cannot convert \"a\" (untyped string constant) to type int" */ } else { 1 })
	_ = int(if c { f } else { 2.5 /* ERROR "truncated" */ })
	_ = string(if c { i /* ERROR "cannot convert i (variable of type int) to type string" */ } else { 'b' })
	_ = io.Reader(if /* ERROR "mismatched types *os.File and untyped int in conditional expression" */ c { fl } else { 1 })
	_ = io.Reader(if c { nil } else { 1 /* ERROR "cannot convert 1 (constant of type int) to type io.Reader" */ })
	_ = float64(if c { i } else { s /* ERROR "cannot convert s (variable of type string) to type float64" */ })
}

const _ = float64(if true { 1 } else { 2 })
const _ = string(if true { 'a' } else { 'b' })
const _ = MyInt(if true { 1 } else { 2 }) + MyInt(1)

var _ [int(if true { 2.0 } else { 3 })]int = [2]int{}

// The element arguments of append (without ...), the key argument of delete,
// and the argument of panic are targets for conditional expressions.

func _() {
	var errs []error
	errs = append(errs, if c { me } else { nil })
	errs = append(errs, if c { me } else { nil }, nil, if c { nil } else { e })
	_ = append([]any{}, if c { 1 } else { 2.5 }, if c { f } else { 0 })
	_ = append([]int{}, if c { 1 } else { "a" /* ERROR "cannot use \"a\" (untyped string constant) as int value in conditional expression" */ })
	_ = append(sl, if c { sl } else { nil }...)
	_ = append(if c { sl } else { nil }, 1)
	_ = append([]Point{}, { /* ERROR "missing type in composite literal" */ 1, 2})
	_ = append([]Point{}, if c { { /* ERROR "missing type in composite literal" */ 1, 2} } else { { /* ERROR "missing type in composite literal" */ 3, 4} })
	delete(map[error]int{}, if c { me } else { nil })
	delete(m, if c { "a" } else { 1 /* ERROR "cannot use 1 (untyped int constant) as string value in conditional expression" */ })
	_ = len(if c { "ab" } else { "c" })
}

func _() {
	panic(if c { me } else { nil })
}

func _() {
	panic(if c { 1 } else { 2.5 })
}

// A parameter whose type depends on the callee's type parameters provides no
// target type: a conditional expression with a nil branch whose type differs
// from the inferred interface type is an error.

func pairOf[T, U any](T, U) {}

func _() {
	pairOf[error](if /* ERROR "cannot use conditional expression with nil branch as error value in argument to pairOf[error] (its type *MyErr is fixed before inference; convert a branch explicitly)" */ c { me } else { nil }, 1)
	pairOf[any](if /* ERROR "cannot use conditional expression with nil branch as any value in argument to pairOf[any] (its type error is fixed before inference; convert a branch explicitly)" */ c { e } else { nil }, 1)
	pairOf[error](if c { e } else { nil }, 1)
	pairOf[error](if c { error(me) } else { nil }, 1)
	pairOf[error](if c { me } else { me }, 1)
	pairOf[error, int](if c { me } else { nil }, 1)
	pairOf(if c { me } else { nil }, 1)
	var _ *MyErr = gen(if c { me } else { nil })
}

// Constant results

const intSize = if ^uint(0)>>63 == 1 { 64 } else { 32 }

const (
	_          = intSize
	debug      = false
	kf         = if debug { 1 } else { 2.5 }
	ki         = if !debug { 1 } else { 2.5 }
	ks         = if debug { "a" } else { "b" }
	kb         = if debug { true } else { false }
	kt    int8 = if debug { 1 } else { 2 }
	ku         = if debug { kt } else { 3 }
)

var _ [if debug { 1 } else { 2 }]int
var _ [2]int = [ki + 1]int{}
var _ int8 = ku
var _ = ki /* ERROR "not defined on ki (untyped float constant 1)" */ % 2

const _ = if /* ERROR "is not constant" */ c { 1 } else { 2 }
const _ = if /* ERROR "is not constant" */ debug { i } else { 2 }
const _ = if debug { 1 } else { two /* ERROR "multiple-value" */ () }

const (
	zero = 0
	ten  = 10
)

const _ = if zero != 0 { ten / zero /* ERROR "division by zero" */ } else { 0 }
const _ byte = if true { 1 } else { 1000 /* ERROR "overflows" */ }
const _ byte = if false { 1000 /* ERROR "overflows" */ } else { 1 }

var _ = map[int]string{if true { 1 } else { 2 }: "a", 1 /* ERROR "duplicate key" */ : "b"}

func _() {
	var b8 byte
	_ = b8 + if true { 1 } else { 1000 /* ERROR "overflows" */ }
}

// An inherited constant expression is evaluated again.
const (
	ia = if iota == 0 { 10 } else { 20 }
	ib
)

var _ [ia]int = [10]int{}
var _ [ib]int = [20]int{}

// len and cap of an array are constant if the expression contains no
// function calls or receive operations.
const _ = len(if c { arr } else { arr })
const _ = len /* ERROR "is not constant" */ (if c { arr } else { [2]int{gen(1)} })
const _ = len /* ERROR "is not constant" */ (if <-ch == 0 { arr } else { arr })

// Untyped non-constant results in other contexts

func _() {
	for i, r := range if c { "ab" } else { "cd" } {
		_, _ = i, r
	}
	for i := range if c { 3 } else { 4 } {
		var _ int = i
	}
	var _ byte = (if c { "ab" } else { "cd" })[0]
	var _ string = (if c { "ab" } else { "cd" })[1:]
	var _ int = len(if c { "ab" } else { "c" })
	var _ = arr[if c { 0 } else { 1 }]
	var _ = sl[if c { 0 } else { 1.0 }:]
	var _ = make([]int, if c { 1 } else { 2 })
	var _ = i << if c { 1 } else { 2 }
	var _ = i<<(if c { 1 } else { 2 }) + 1
	var _ = i << if c { 1.0 } else { 2 }
	var _ = i << if c { - /* ERROR "overflows uint" */ 1 } else { 2 }
	var _ = i << if c { 1 } else { 2.5 /* ERROR "truncated to uint" */ }
	var _ = "x" + if c { "a" } else { "b" }
	var _ = if c { 1 } else { 2 } == 1
	var _ = -if c { 1 } else { 2 }
	var _ = !if c { true } else { false }
	switch if c { 1 } else { 2 } {
	case 1:
	}
	switch i {
	case if c { 1 } else { 2 }:
	}
}

// A value, not a variable

func _() {
	_ = &( /* ERROR "cannot take address" */ if c { i } else { i })
	(if c { sl } else { sl })[0] = 1
	(if c { m } else { m })["a"] = 1
	( /* ERROR "cannot assign" */ if c { arr } else { arr })[0] = 1
	( /* ERROR "cannot assign" */ if c { i } else { i }) = 1
	_ = (if c { sl } else { sl })[1:]
	_ = if c { sl } else { sl }[0]
	_ = if c { Point{} } else { Point{} }.X
	_ = if c { gen[int] } else { gen[int] }(1)

	// Statement context
	( /* ERROR "is not used" */ if c { i } else { i })
	( /* ERROR "is not used" */ if c { errOnly() } else { errOnly() })
	defer (if c { none } else { none })()
	go (if c { none } else { none })()
}

// Single-valued branches

func _() {
	_ = if c { two /* ERROR "multiple-value two() (value of type (int, int)) in single-value context" */ () } else { 1 }
	_ = if c { 1 } else { none /* ERROR "none() (no value) used as value" */ () }
	_ = if c { int /* ERROR "int (type) is not an expression" */ } else { 1 }
	_ = if c { len /* ERROR "len (built-in) must be called" */ } else { 1 }
	x, y := if c { two /* ERROR "multiple-value" */ () } else { two /* ERROR "multiple-value" */ () }
	_, _ = x, y
	// comma-ok and comma-err do not propagate through a conditional expression
	v, ok := if /* ERROR "assignment mismatch: 2 variables but 1 value" */ c { m["a"] } else { m["b"] }
	_, _ = v, ok
}

// Error handling in branches

func _() error {
	x := if c { ee()! } else { 0 }
	y := if c {
		ee() or err {
			return err
		}
	} else {
		1
	}
	_, _ = x, y
	_ = if c { errOnly /* ERROR "used as value" */ ()! } else { 1 }
	return nil
}

// Generic type parameters

func _[T any](c bool, a, b T) T {
	x := if c { a } else { b }
	var _ T = x
	return if c { a } else { b }
}

func _[T ~int | ~float64](c bool, a T) T {
	var _ T = if c { a } else { 1 }
	_ = if c { a } else { 1 }
	return if c { 2 } else { a }
}

func _[T ~int8 | ~int16](c bool, a T) {
	_ = if c { a } else { 1000 /* ERROR "cannot use 1000 (untyped int constant) as T value in conditional expression" */ }
}

func _[T any, U any](c bool, a T, b U) {
	_ = if /* ERROR "mismatched types T and U in conditional expression" */ c { a } else { b }
}

// Shifts

func _() {
	var _ int = (if c { 1 } else { 2 }) << u
	var _ int64 = if c { 1 } else { 2 } << u
	var _ float64 = (if /* ERROR "must be integer" */ c { 1 } else { 2 }) << u
	var _ float64 = ((if /* ERROR "must be integer" */ c { 1 } else { 2 }) + 1) << u
	_ = ( /* ERROR "must be integer" */ if c { 1 } else { 2.5 }) << u
}
