package ssa_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/interp"
	"golang.org/x/tools/go/ssa/ssautil"
)

// TestGonConditional executes the conditional-expression legacy/modern pair
// of the Gon repository in the SSA interpreter.
func TestGonConditional(t *testing.T) {
	root := os.Getenv("GON_ROOT")
	if root == "" {
		t.Skip("set GON_ROOT to run compiler/analysis executable pairs")
	}
	for _, name := range []string{"conditional_legacy", "conditional_modern"} {
		t.Run(name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, filepath.Join(root, "misc/gon/analysisfixtures", name+".go"), nil, parser.SkipObjectResolution)
			if err != nil {
				t.Fatal(err)
			}
			mode := ssa.SanityCheckFunctions | ssa.InstantiateGenerics | ssa.GlobalDebug
			p, _, err := ssautil.BuildPackage(&types.Config{}, fset, types.NewPackage("main", "main"), []*ast.File{f}, mode)
			if err != nil {
				t.Fatal(err)
			}
			if code := interp.Interpret(p, 0, types.SizesFor("gc", "arm64"), "main", nil); code != 0 {
				t.Fatalf("SSA execution failed: %d", code)
			}
		})
	}
}

// TestGonConditionalBlocks checks the SSA form of conditional expressions:
// each branch is evaluated in its own block, converted to the type of the
// expression, and joined with a phi; constant ones are constants.
func TestGonConditionalBlocks(t *testing.T) {
	const src = `package p

type E struct{}

func (*E) Error() string { return "" }

func a() int { return 1 }
func b() int { return 2 }

func lazy(c bool) int { return if c { a() } else { b() } }

func iface(c bool, p *E) error { return if c { p } else { nil } }

func untyped(c bool, s string) bool {
	if if c { s == "" } else { true } {
		return len(if c { "ab" } else { "c" }) > 1
	}
	return false
}

func constant() int { return if 1 < 2 { 3 } else { 4 } }
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatal(err)
	}
	p, _, err := ssautil.BuildPackage(&types.Config{}, fset, types.NewPackage("p", "p"), []*ast.File{f}, ssa.SanityCheckFunctions)
	if err != nil {
		t.Fatal(err)
	}
	phis := func(fn *ssa.Function) (list []*ssa.Phi) {
		for _, b := range fn.Blocks {
			for _, instr := range b.Instrs {
				if phi, ok := instr.(*ssa.Phi); ok && phi.Comment == "condexpr" {
					list = append(list, phi)
				}
			}
		}
		return list
	}
	callee := func(v ssa.Value) string {
		if call, ok := v.(*ssa.Call); ok {
			if callee := call.Call.StaticCallee(); callee != nil {
				return callee.Name()
			}
		}
		return ""
	}

	lazy := phis(p.Func("lazy"))
	if len(lazy) != 1 || len(lazy[0].Edges) != 2 || callee(lazy[0].Edges[0]) != "a" || callee(lazy[0].Edges[1]) != "b" {
		t.Fatalf("lazy: unexpected phis %v", lazy)
	}
	for i, want := range []string{"condexpr.then", "condexpr.else"} {
		if block := lazy[0].Edges[i].(*ssa.Call).Block(); block.Comment != want {
			t.Errorf("lazy: call %d is in block %s, want %s", i, block.Comment, want)
		}
	}

	iface := phis(p.Func("iface"))
	if len(iface) != 1 || iface[0].Type().String() != "error" {
		t.Fatalf("iface: unexpected phis %v", iface)
	}
	if _, ok := iface[0].Edges[0].(*ssa.MakeInterface); !ok {
		t.Errorf("iface: then edge is %T, want *ssa.MakeInterface", iface[0].Edges[0])
	}
	if c, ok := iface[0].Edges[1].(*ssa.Const); !ok || !c.IsNil() || c.Type().String() != "error" {
		t.Errorf("iface: else edge is %v, want a nil error", iface[0].Edges[1])
	}

	untyped := phis(p.Func("untyped"))
	if len(untyped) != 2 || untyped[0].Type() != types.Typ[types.Bool] || untyped[1].Type() != types.Typ[types.String] {
		t.Fatalf("untyped: unexpected phis %v", untyped)
	}

	fn := p.Func("constant")
	if len(phis(fn)) != 0 || len(fn.Blocks) != 1 {
		t.Fatalf("constant: unexpected blocks")
	}
	if ret := fn.Blocks[0].Instrs[0].(*ssa.Return); ret.Results[0].(*ssa.Const).Int64() != 3 {
		t.Errorf("constant: returns %v", ret.Results[0])
	}
}
