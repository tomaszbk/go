package satisfy

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"slices"
	"testing"
)

// Each branch of a conditional expression is converted to the type of the
// expression, which is the target type of its context if there is one.
func TestCondExpr(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", `package p
type I interface{ M() }
type J interface{ M(); N() }
type A struct{}
func (A) M() {}
type B struct{}
func (B) M() {}
type C struct{}
func (C) M() {}
func (C) N() {}
func sink(I) {}
func read() (int, error) { return 0, nil }
var global I = if len("x") > 0 { A{} } else { nil }
func f(c bool) (I, error) {
	var x I = if c { A{} } else { B{} }
	sink(if c { x } else { C{} })
	same := if c { B{} } else { B{} }
	_ = same
	n := if c { read() or err { sink(&B{}); return nil, err } } else { 0 }
	_ = n
	return if c { J(C{}) } else { nil }, nil
}
`, 0)
	if err != nil {
		t.Fatal(err)
	}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	pkg, err := new(types.Config).Check("p", fset, []*ast.File{f}, info)
	if err != nil {
		t.Fatal(err)
	}
	var finder Finder
	finder.Find(info, []*ast.File{f})

	var got []string
	qual := types.RelativeTo(pkg)
	for c := range finder.Result {
		got = append(got, types.TypeString(c.LHS, qual)+" <- "+types.TypeString(c.RHS, qual))
	}
	slices.Sort(got)
	want := []string{
		"I <- *B", // the handler in a branch
		"I <- A",  // global and x
		"I <- B",  // x
		"I <- C",  // the call argument
		"I <- J",  // the result
		"J <- C",  // the conversion in a branch
	}
	if !slices.Equal(got, want) {
		t.Errorf("constraints:\ngot  %q\nwant %q", got, want)
	}
}
