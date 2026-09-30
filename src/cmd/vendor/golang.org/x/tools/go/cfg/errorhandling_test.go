package cfg

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestErrorHandlingControlFlow(t *testing.T) {
	for _, test := range []struct {
		body     string
		noReturn bool
	}{
		{`f()!; panic(0)`, false},
		{`f() or err { return }; panic(0)`, false},
		{`f() or err { panic(err) }; panic(0)`, true},
		{`f() or err {}; return`, false},
		{`_ = func() { f()! }; panic(0)`, true},
		{`g(f()!)!; panic(0)`, false},
	} {
		t.Run(test.body, func(t *testing.T) {
			file, err := parser.ParseFile(token.NewFileSet(), "test.go", "package p; func test() {"+test.body+"}", 0)
			if err != nil {
				t.Fatal(err)
			}
			g := New(file.Decls[0].(*ast.FuncDecl).Body, func(call *ast.CallExpr) bool {
				id, ok := call.Fun.(*ast.Ident)
				return !ok || id.Name != "panic"
			})
			if g.NoReturn() != test.noReturn {
				t.Fatalf("NoReturn = %v, want %v\n%s", g.NoReturn(), test.noReturn, g.Format(token.NewFileSet()))
			}
		})
	}
}
