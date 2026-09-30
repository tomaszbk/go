package inspector

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/ast/edge"
)

func TestErrorHandling(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "p.go", `package p
func f() error {
 x := read()!
 y := read() or err { return err }
 _, _ = x, y
 return nil
}`, 0)
	if err != nil {
		t.Fatal(err)
	}
	in := New([]*ast.File{f})
	count := 0
	for cursor := range in.Root().Preorder((*ast.ErrorExpr)(nil)) {
		count++
		expr := cursor.Node().(*ast.ErrorExpr)
		children := []struct {
			kind edge.Kind
			node ast.Node
		}{{edge.ErrorExpr_X, expr.X}}
		if expr.Body != nil {
			children = append(children, struct {
				kind edge.Kind
				node ast.Node
			}{edge.ErrorExpr_Err, expr.Err}, struct {
				kind edge.Kind
				node ast.Node
			}{edge.ErrorExpr_Body, expr.Body})
		}
		for _, child := range children {
			c := cursor.ChildAt(child.kind, -1)
			if c.Node() != child.node || c.ParentEdgeKind() != child.kind || child.kind.Get(expr, -1) != child.node {
				t.Errorf("incorrect %s edge", child.kind)
			}
		}
	}
	if count != 2 {
		t.Fatalf("found %d error expressions, want 2", count)
	}
	var want, got []ast.Node
	ast.Inspect(f, func(n ast.Node) bool {
		if n != nil {
			want = append(want, n)
		}
		return true
	})
	in.Preorder(nil, func(n ast.Node) { got = append(got, n) })
	if len(got) != len(want) {
		t.Fatalf("visited %d nodes, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("node %d: got %T, want %T", i, got[i], want[i])
		}
	}
}
