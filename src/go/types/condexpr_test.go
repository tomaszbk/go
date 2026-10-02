package types_test

import (
	"go/ast"
	"go/constant"
	"go/token"
	"slices"
	"testing"

	. "go/types"
)

const condExprSrc = `package p

type (
	MyErr  struct{}
	MyBool bool
)

func (*MyErr) Error() string { return "" }

func gen[T any](x T) T { return x }
func takeInts(...int)   {}
func takeAny(...any)    {}

var (
	c   bool
	me  *MyErr
	f32 float32
	i   int
	sl  []int
	rn  rune
)

var _ error = if c { me } else { nil }
var _ = if c { 1 } else { 2.5 }
const _ = if true { 1 } else { 2.5 }
var _ = f32 + if c { 1 } else { 0.5 }
var _ any = if c { 1 } else { 2.5 }
var _ = float64(if c { 1 } else { 2 })
var _ = if c { i } else { 0 }
var _ MyBool = if c { i == 0 } else { true }
var _ = error(if c { me } else { nil })
var _ error = (if c { me } else { nil })
var _ = gen(if c { me } else { nil })
var _ = func() { takeInts(if c { sl } else { nil }...) }
var _ = func() { takeInts(if c { 1 } else { 2 }) }
var _ int8 = if false { 1 } else { 2 }
var _ = len(if c { [3]int{} } else { [3]int{} })
var _ = i << if c { 1 } else { 2 }
var _ = func() { for range if c { "ab" } else { "cd" } {} }
var _ = append([]error{}, if c { me } else { nil })
var _ = string(if c { rn } else { 'b' })
const _ = string(if true { 'a' } else { 'b' })
var _ = func() { takeAny(if c { f32 } else { 0 }) }
var _ = float64(if c { i } else { 2 })

func _[T any](c bool, a, b T) T { return if c { a } else { b } }
`

// TestCondExprTypes checks the types and values recorded for conditional
// expressions and their branches.
func TestCondExprTypes(t *testing.T) {
	fset := token.NewFileSet()
	f := mustParse(fset, condExprSrc)
	info := &Info{Types: make(map[ast.Expr]TypeAndValue)}
	conf := Config{Importer: defaultImporter(fset)}
	if _, err := conf.Check("p", fset, []*ast.File{f}, info); err != nil {
		t.Fatal(err)
	}

	var conds []*ast.CondExpr
	ast.Inspect(f, func(n ast.Node) bool {
		if x, ok := n.(*ast.CondExpr); ok {
			conds = append(conds, x)
		}
		return true
	})

	// For each conditional expression in source order: the recorded types
	// and constant values (or "" for non-constants) of the expression, its
	// then branch, and its else branch.
	// Unlike types2, go/types records nil as untyped nil in all contexts.
	want := [][6]string{
		{"error", "", "*MyErr", "", "untyped nil", ""},
		{"float64", "", "float64", "1", "float64", "2.5"},
		{"untyped float", "1", "untyped float", "1", "untyped float", "2.5"},
		{"float32", "", "float32", "1", "float32", "0.5"},
		{"any", "", "float64", "1", "float64", "2.5"},
		{"float64", "", "float64", "1", "float64", "2"},
		{"int", "", "int", "", "int", "0"},
		{"MyBool", "", "MyBool", "", "MyBool", "true"},
		{"error", "", "*MyErr", "", "untyped nil", ""},
		{"error", "", "*MyErr", "", "untyped nil", ""},
		{"*MyErr", "", "*MyErr", "", "untyped nil", ""},
		{"[]int", "", "[]int", "", "untyped nil", ""},
		{"int", "", "int", "1", "int", "2"},
		{"int8", "2", "int8", "1", "int8", "2"},
		{"[3]int", "", "[3]int", "", "[3]int", ""},
		{"uint", "", "uint", "1", "uint", "2"},
		// The context does not give a final type to some untyped values.
		{"untyped string", "", "untyped string", `"ab"`, "untyped string", `"cd"`},
		// Arguments of append, delete, and panic, and conversions are targets.
		{"error", "", "*MyErr", "", "untyped nil", ""},
		{"string", "", "rune", "", "untyped rune", "98"},
		{"string", `"a"`, "untyped rune", "97", "untyped rune", "98"},
		{"any", "", "float32", "", "float32", "0"},
		{"float64", "", "int", "", "float64", "2"},
		{"T", "", "T", "", "T", ""},
	}
	if len(conds) != len(want) {
		t.Fatalf("got %d conditional expressions, want %d", len(conds), len(want))
	}

	qf := func(*Package) string { return "" }
	describe := func(x ast.Expr) (string, string) {
		tv, ok := info.Types[x]
		if !ok {
			return "<missing>", ""
		}
		val := ""
		if tv.Value != nil {
			val = tv.Value.String()
		}
		return TypeString(tv.Type, qf), val
	}
	for k, x := range conds {
		var got [6]string
		got[0], got[1] = describe(x)
		got[2], got[3] = describe(x.Then)
		got[4], got[5] = describe(x.Else)
		if got != want[k] {
			t.Errorf("%s:\ngot  %q\nwant %q", ExprString(x), got, want[k])
		}
		if tv := info.Types[x]; tv.IsNil() || tv.Addressable() || !tv.IsValue() {
			t.Errorf("%s: invalid mode", ExprString(x))
		}
	}

	// An untyped floating-point constant has a floating-point value.
	if v := info.Types[conds[2]].Value; v.Kind() != constant.Float {
		t.Errorf("got %s value %s, want a float value", v.Kind(), v)
	}
	if got, want := ExprString(conds[1]), "if c { 1 } else { 2.5 }"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestCondExprConversions checks that a conversion distributes over the
// branches of a conditional expression and that a branch needs an explicit
// conversion exactly if its type is not assignable to the expression type.
func TestCondExprConversions(t *testing.T) {
	const src = `package p
type (
	MyErr struct{}
	MyInt int
)
func (*MyErr) Error() string { return "" }
var (
	c  bool
	i  int
	rn rune
	s  string
	me *MyErr
)
var _ = float64(if c { i } else { 2 })
var _ = string(if c { rn } else { 'b' })
var _ = error(if c { me } else { nil })
var _ = []byte(if c { s } else { "b" })
var _ = MyInt(if c { i } else { 0 })
var _ = (*int)((if c { nil } else { nil }))
`
	info := &Info{Types: make(map[ast.Expr]TypeAndValue)}
	fset := token.NewFileSet()
	f := mustParse(fset, src)
	conf := Config{Importer: defaultImporter(fset)}
	if _, err := conf.Check("p", fset, []*ast.File{f}, info); err != nil {
		t.Fatal(err)
	}
	want := []struct {
		typ                string
		thenConv, elseConv bool // explicit conversion required
	}{
		{"float64", true, false},
		{"string", true, true},
		{"error", false, false},
		{"[]byte", true, true},
		{"MyInt", true, false},
		{"*int", false, false},
	}
	var k int
	ast.Inspect(f, func(n ast.Node) bool {
		x, ok := n.(*ast.CondExpr)
		if !ok {
			return true
		}
		if k >= len(want) {
			t.Fatalf("too many conditional expressions")
		}
		tv := info.Types[x]
		if got := TypeString(tv.Type, func(*Package) string { return "" }); got != want[k].typ {
			t.Errorf("%s: got type %s, want %s", ExprString(x), got, want[k].typ)
		}
		for j, b := range []ast.Expr{x.Then, x.Else} {
			conv := !AssignableTo(info.Types[b].Type, tv.Type)
			if wantConv := []bool{want[k].thenConv, want[k].elseConv}[j]; conv != wantConv {
				t.Errorf("%s: branch %s needs explicit conversion: got %v, want %v", ExprString(x), ExprString(b), conv, wantConv)
			}
		}
		k++
		return true
	})
	if k != len(want) {
		t.Errorf("got %d conditional expressions, want %d", k, len(want))
	}
}

func TestCondExprConstants(t *testing.T) {
	const src = `package p
const intSize = if ^uint(0)>>63 == 1 { 64 } else { 32 }
const s = if intSize == 64 { "64-bit" } else { "32-bit" }
const k int8 = if false { 1 } else { -2 }
var a [if intSize == 64 { 3 } else { 4 }]int
`
	pkg := mustTypecheck(src, &Config{Sizes: SizesFor("gc", "amd64")}, nil)
	for _, test := range []struct{ name, typ, val string }{
		{"intSize", "untyped int", "64"},
		{"s", "untyped string", `"64-bit"`},
		{"k", "int8", "-2"},
	} {
		obj := pkg.Scope().Lookup(test.name).(*Const)
		if got := obj.Type().String(); got != test.typ {
			t.Errorf("%s: got type %s, want %s", test.name, got, test.typ)
		}
		if got := obj.Val().ExactString(); got != test.val {
			t.Errorf("%s: got value %s, want %s", test.name, got, test.val)
		}
	}
	if got := pkg.Scope().Lookup("a").Type().String(); got != "[3]int" {
		t.Errorf("got type %s, want [3]int", got)
	}
}

// References in both branches count as dependencies for initialization.
func TestCondExprInitOrder(t *testing.T) {
	const src = `package p
var x = if c { f() } else { g() }
var c = h()
func f() int { return a }
func g() int { return b }
func h() bool { return true }
var a = 1
var b = 2
`
	info := &Info{}
	mustTypecheck(src, nil, info)
	var order []string
	for _, init := range info.InitOrder {
		for _, v := range init.Lhs {
			order = append(order, v.Name())
		}
	}
	if want := []string{"c", "a", "b", "x"}; !slices.Equal(order, want) {
		t.Errorf("got init order %v, want %v", order, want)
	}
}

func TestCondExprEval(t *testing.T) {
	fset := token.NewFileSet()
	f := mustParse(fset, `package p; var c bool; var i int`)
	pkg, err := new(Config).Check("p", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct{ expr, typ, val string }{
		{"if c { 1 } else { 2.5 }", "untyped float", ""},
		{"if true { 1 } else { 2 }", "untyped int", "1"},
		{"if c { i } else { 0 }", "int", ""},
		{"(if c { i } else { 0 })", "int", ""},
	} {
		tv, err := Eval(fset, pkg, token.NoPos, test.expr)
		if err != nil {
			t.Errorf("Eval(%q): %v", test.expr, err)
			continue
		}
		val := ""
		if tv.Value != nil {
			val = tv.Value.String()
		}
		if got := tv.Type.String(); got != test.typ || val != test.val {
			t.Errorf("Eval(%q) = %s %s, want %s %s", test.expr, got, val, test.typ, test.val)
		}
	}
	if _, err := Eval(fset, pkg, token.NoPos, "if c { i } else { 0.5 }"); err == nil {
		t.Errorf("Eval succeeded for mismatched branches")
	}
}
