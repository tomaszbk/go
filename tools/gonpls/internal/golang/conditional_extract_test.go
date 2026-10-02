package golang

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
	"testing"

	"golang.org/x/tools/go/ast/inspector"
)

func TestConditionalExtraction(t *testing.T) {
	for _, test := range []struct {
		name, body, selected, want string
		all                        bool
	}{
		{"then", `return if flag { touch(1) } else { touch(2) }`, "touch(1)", "lazy conditional", false},
		{"else", `return if flag { touch(1) } else { touch(2) }`, "touch(2)", "lazy conditional", false},
		{"nested operand", `return if flag { touch(touch(1)) } else { 0 }`, "touch(1)", "lazy conditional", false},
		{"all occurrences", `_ = touch(1); return if flag { touch(1) } else { 0 }`, "touch(1)", "lazy conditional", true},
		{"whole expression", `return (if flag { 1 } else { 2 })`, "(if flag { 1 } else { 2 })", "target type", false},
		{"interface target", `var p *int; var value any = if flag { p } else { nil }; _ = value; return 0`, "if flag { p } else { nil }", "target type", false},
		{"ordinary statement", `if flag { return touch(1) }; return 2`, "touch(1)", "", false},
		{"function boundary", `return if flag { func() int { return touch(1) }() } else { 0 }`, "touch(1)", "", false},
	} {
		t.Run(test.name, func(t *testing.T) {
			src := "package p\nfunc touch(x int) int { return x }\nfunc f(flag bool) int { " + test.body + " }"
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "p.go", src, 0)
			if err != nil {
				t.Fatal(err)
			}
			info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue), Uses: make(map[*ast.Ident]types.Object), Defs: make(map[*ast.Ident]types.Object)}
			if _, err := new(types.Config).Check("p", fset, []*ast.File{file}, info); err != nil {
				t.Fatal(err)
			}
			curFile, _ := inspector.New([]*ast.File{file}).Root().FindNode(file)
			start := fset.File(file.Pos()).Pos(strings.Index(src, test.selected))
			_, err = canExtractVariable(info, curFile, start, start+token.Pos(len(test.selected)), test.all)
			if test.want == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
}
