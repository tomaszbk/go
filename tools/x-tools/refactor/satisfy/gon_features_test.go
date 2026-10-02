package satisfy

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"slices"
	"testing"
)

func TestGonFeatureConstraints(t *testing.T) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "p.go", `package p
type I interface { M() }
type A struct{}; func (A) M() {}
type B struct{}; func (B) M() {}
type C struct{}; func (*C) M() {}
type D struct{}; func (*D) M() {}
type E struct{}; func (*E) M() {}
type F struct{}; func (*F) M() {}
type S struct { C *C }
var expr func() I = () => A{}
var block func() I = () => { return B{} }
func use(p *S, d *D) {
 var i I = p?.C
 var j I = d ?? &E{}
 i ??= &F{}
 _, _ = i, j
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
	pkg, err := new(types.Config).Check("p", fs, []*ast.File{f}, info)
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
	want := []string{"I <- *C", "I <- *D", "I <- *E", "I <- *F", "I <- A", "I <- B"}
	if !slices.Equal(got, want) {
		t.Fatalf("constraints: got %q, want %q", got, want)
	}
}
