package inline_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/internal/refactor/inline"
	"strings"
	"testing"
)

func TestGonCalleeSafety(t *testing.T) {
	for _, body := range []string{"return if c { 1 } else { 2 }, nil", "return read()!, nil", "return read() or err { return 0, err }, nil", "return 1, nil"} {
		t.Run(body, func(t *testing.T) {
			src := []byte("package p; func read() (int,error) { return 0,nil }; func f(c bool) (int,error) { " + body + " }")
			fs := token.NewFileSet()
			file, err := parser.ParseFile(fs, "p.go", src, parser.SkipObjectResolution)
			if err != nil {
				t.Fatal(err)
			}
			info := &types.Info{Implicits: map[ast.Node]types.Object{}, Types: map[ast.Expr]types.TypeAndValue{}, Defs: map[*ast.Ident]types.Object{}, Uses: map[*ast.Ident]types.Object{}, Scopes: map[ast.Node]*types.Scope{}, Selections: map[*ast.SelectorExpr]*types.Selection{}, Instances: map[*ast.Ident]types.Instance{}, FileVersions: map[*ast.File]string{}}
			pkg, err := new(types.Config).Check("p", fs, []*ast.File{file}, info)
			if err != nil {
				t.Fatal(err)
			}
			_, err = inline.AnalyzeCallee(t.Logf, fs, pkg, info, file.Decls[1].(*ast.FuncDecl), src)
			if body == "return 1, nil" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), "Gon control-flow") {
				t.Fatalf("got %v, want conservative refusal", err)
			}
		})
	}
}
