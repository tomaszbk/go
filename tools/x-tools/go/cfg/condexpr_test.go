package cfg

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// newCondCFG returns the CFG of a function test with the given body.
// Calls of panic do not return.
func newCondCFG(t *testing.T, body string) *CFG {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), "test.go", "package p; func test() {"+body+"}", 0)
	if err != nil {
		t.Fatal(err)
	}
	return New(file.Decls[0].(*ast.FuncDecl).Body, func(call *ast.CallExpr) bool {
		id, ok := call.Fun.(*ast.Ident)
		return !ok || id.Name != "panic"
	})
}

// Only the selected branch of a conditional expression is evaluated, so
// error propagation and handlers in a branch return on that path only.
func TestCondExprControlFlow(t *testing.T) {
	for _, test := range []struct {
		body     string
		noReturn bool
	}{
		{`_ = if c { 1 } else { 2 }; panic(0)`, true},
		{`panic(if c { 1 } else { 2 })`, true},
		{`return if c { 1 } else { 2 }`, false},
		{`_ = if c { f()! } else { 0 }; panic(0)`, false},
		{`_ = if c { 0 } else { f() or err { return } }; panic(0)`, false},
		{`_ = if c { f() or err { panic(err) } } else { g() or err { panic(err) } }; panic(0)`, true},
		{`_ = if f()! { 1 } else { 2 }; panic(0)`, false},
		{`_ = g(1, if c { f()! } else { 0 }); panic(0)`, false},
		{`_ = if c { g(if d { f()! } else { 0 }) } else { 0 }; panic(0)`, false},
		{`_ = if c { func() int { f()!; return 0 } } else { nil }; panic(0)`, true},
		{`defer g(if c { f()! } else { 0 }); panic(0)`, false}, // defer may recover
	} {
		t.Run(test.body, func(t *testing.T) {
			g := newCondCFG(t, test.body)
			if g.NoReturn() != test.noReturn {
				t.Fatalf("NoReturn = %v, want %v\n%s", g.NoReturn(), test.noReturn, g.Format(token.NewFileSet()))
			}
		})
	}
}

// TestCondExprBlocks checks the shape of the graph: the block evaluating
// the condition ends with it and branches to one block per branch, which
// rejoin in the block holding the enclosing statement.
func TestCondExprBlocks(t *testing.T) {
	g := newCondCFG(t, `before(); x := if ok() { a() } else { b()! }; after(x)`)
	fset := token.NewFileSet()
	dump := g.Format(fset) // also checks that the nodes can be printed

	entry := g.Blocks[0]
	if len(entry.Nodes) != 2 || len(entry.Succs) != 2 {
		t.Fatalf("entry block: want 2 nodes and 2 successors\n%s", dump)
	}
	cond, ok := entry.Nodes[1].(*ast.CallExpr)
	if !ok || cond.Fun.(*ast.Ident).Name != "ok" {
		t.Fatalf("entry block does not end with the condition\n%s", dump)
	}
	then, _else := entry.Succs[0], entry.Succs[1]
	if then.Kind != KindCondThen || _else.Kind != KindCondElse {
		t.Fatalf("successors of the condition are %v and %v, want CondThen and CondElse\n%s", then.Kind, _else.Kind, dump)
	}

	// The then branch has no control flow of its own.
	if len(then.Nodes) != 1 || len(then.Succs) != 1 || then.Succs[0].Kind != KindCondDone {
		t.Fatalf("unexpected then block\n%s", dump)
	}
	done := then.Succs[0]

	// The else branch propagates an error: the call ends the block, the
	// handler returns, and the propagation follows in its done block.
	if len(_else.Nodes) != 1 || len(_else.Succs) != 2 {
		t.Fatalf("unexpected else block\n%s", dump)
	}
	handler, errDone := _else.Succs[0], _else.Succs[1]
	if handler.Kind != KindErrorHandler || handler.Return() == nil || errDone.Kind != KindErrorDone {
		t.Fatalf("unexpected error paths in the else branch\n%s", dump)
	}
	if _, ok := errDone.Nodes[0].(*ast.ErrorExpr); !ok || len(errDone.Succs) != 1 || errDone.Succs[0] != done {
		t.Fatalf("else branch does not rejoin the then branch\n%s", dump)
	}

	// The enclosing statement, and those after it, are in the done block.
	if len(done.Nodes) < 2 {
		t.Fatalf("unexpected done block\n%s", dump)
	}
	if assign, ok := done.Nodes[0].(*ast.AssignStmt); !ok || assign.Rhs[0].(*ast.CondExpr).Then != then.Nodes[0] {
		t.Fatalf("done block does not start with the enclosing statement\n%s", dump)
	}

	// The propagating return is reachable from the else branch only.
	if !reaches(_else, handler) || reaches(then, handler) {
		t.Errorf("the propagation in the else branch is reachable from the wrong branch\n%s", dump)
	}
}

// reaches reports whether target is reachable from b.
func reaches(b, target *Block) bool {
	seen := make(map[*Block]bool)
	var search func(*Block) bool
	search = func(b *Block) bool {
		if b == target {
			return true
		}
		if seen[b] {
			return false
		}
		seen[b] = true
		for _, succ := range b.Succs {
			if search(succ) {
				return true
			}
		}
		return false
	}
	return search(b)
}
