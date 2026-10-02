package ir_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"testing"

	"honnef.co/go/tools/go/ir"
	"honnef.co/go/tools/go/ir/irutil"
)

// TestGonConditional builds the conditional-expression legacy/modern pair
// of the Gon repository with IR sanity checks, and checks that the
// branches of a conditional expression are joined with a phi.
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
			mode := ir.SanityCheckFunctions | ir.InstantiateGenerics | ir.GlobalDebug
			p, _, err := irutil.BuildPackage(&types.Config{}, fset, types.NewPackage("main", "main"), []*ast.File{f}, mode)
			if err != nil {
				t.Fatal(err)
			}
			fn := p.Func("order")
			if fn == nil {
				t.Fatal("missing order IR")
			}
			phis := 0
			for _, b := range fn.Blocks {
				for _, instr := range b.Instrs {
					if phi, ok := instr.(*ir.Phi); ok && phi.Comment() == "condexpr" {
						phis++
					}
				}
			}
			if want := map[string]int{"conditional_legacy": 0, "conditional_modern": 1}[name]; phis != want {
				t.Errorf("order has %d conditional-expression phis, want %d", phis, want)
			}
		})
	}
}

// TestGonSwitchTag checks switch statements with constant cases whose tag
// creates blocks. The builder used to emit the switch into the block where
// the tag evaluation started.
func TestGonSwitchTag(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", `package p

func and(a, b bool) int {
	switch a && b {
	case true:
		return 1
	}
	return 0
}

func cond(c bool) int {
	switch if c { "a" } else { "b" } {
	case "a":
		return 1
	case "b":
		return 2
	}
	return 0
}
`, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := irutil.BuildPackage(&types.Config{}, fset, types.NewPackage("p", "p"), []*ast.File{f}, ir.SanityCheckFunctions); err != nil {
		t.Fatal(err)
	}
}
