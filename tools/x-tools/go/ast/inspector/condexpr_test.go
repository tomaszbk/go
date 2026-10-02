package inspector

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/ast/edge"
)

func TestCondExpr(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "p.go", `package p
func f(c bool, s string) (int, error) {
	n := if c { 1 } else { len(s) }
	m := g(if n > 0 { n } else { -n }, if c { h()! } else { 0 })
	return n + m + if c { func() int { return 1 }() } else { 2 }, nil
}`, 0)
	if err != nil {
		t.Fatal(err)
	}
	in := New([]*ast.File{f})
	count := 0
	for cursor := range in.Root().Preorder((*ast.CondExpr)(nil)) {
		count++
		expr := cursor.Node().(*ast.CondExpr)
		for _, child := range []struct {
			kind edge.Kind
			node ast.Node
		}{
			{edge.CondExpr_Cond, expr.Cond},
			{edge.CondExpr_Then, expr.Then},
			{edge.CondExpr_Else, expr.Else},
		} {
			c := cursor.ChildAt(child.kind, -1)
			if c.Node() != child.node || c.ParentEdgeKind() != child.kind || child.kind.Get(expr, -1) != child.node || c.Parent() != cursor {
				t.Errorf("incorrect %s edge", child.kind)
			}
		}
	}
	if count != 4 {
		t.Fatalf("found %d conditional expressions, want 4", count)
	}
	for kind, want := range map[edge.Kind]string{
		edge.CondExpr_Cond: "CondExpr.Cond",
		edge.CondExpr_Then: "CondExpr.Then",
		edge.CondExpr_Else: "CondExpr.Else",
	} {
		if got := kind.String(); got != want {
			t.Errorf("edge kind %d is %s, want %s", kind, got, want)
		}
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

	// The calls in branches are found by a filtered traversal too.
	calls := 0
	in.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(ast.Node) { calls++ })
	if calls != 4 { // len(s), g(...), h(), func() int {...}()
		t.Errorf("found %d calls, want 4", calls)
	}
}
