package inspector

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/ast/edge"
)

func TestGonFeatureEdges(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "p.go", `package p
var f func(int) int = (x) => x+1
var g func() = () => { callback?(p?.N ?? 1) }
`, 0)
	if err != nil {
		t.Fatal(err)
	}
	in := New([]*ast.File{f})
	want := []ast.Node{}
	for node := range ast.Preorder(f) {
		want = append(want, node)
	}
	i := 0
	in.Preorder(nil, func(n ast.Node) {
		if i >= len(want) || want[i] != n {
			t.Fatalf("node %d: got %T", i, n)
		}
		i++
	})
	if i != len(want) {
		t.Fatalf("visited %d of %d nodes", i, len(want))
	}
	seen := map[edge.Kind]bool{}
	for cur := range in.Root().Preorder((*ast.LambdaExpr)(nil), (*ast.NilGuardExpr)(nil), (*ast.SafeNavExpr)(nil)) {
		var edges []edge.Kind
		switch cur.Node().(type) {
		case *ast.LambdaExpr:
			edges = []edge.Kind{edge.LambdaExpr_Params}
			if cur.Node().(*ast.LambdaExpr).Body != nil {
				edges = append(edges, edge.LambdaExpr_Body)
			} else {
				edges = append(edges, edge.LambdaExpr_Block)
			}
		case *ast.NilGuardExpr:
			edges = []edge.Kind{edge.NilGuardExpr_X}
		case *ast.SafeNavExpr:
			edges = []edge.Kind{edge.SafeNavExpr_X}
		}
		for _, kind := range edges {
			index := -1
			if kind == edge.LambdaExpr_Params {
				index = 0
				if len(cur.Node().(*ast.LambdaExpr).Params) == 0 {
					continue
				}
			}
			node := kind.Get(cur.Node(), index)
			if node == nil {
				continue
			}
			child := cur.ChildAt(kind, index)
			if child.Node() != node || child.Parent() != cur || child.ParentEdgeKind() != kind {
				t.Fatalf("bad %s", kind)
			}
			seen[kind] = true
		}
	}
	if len(seen) != 5 {
		t.Fatalf("missing feature edges: %v", seen)
	}
}
