package format

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

// Helpers to build syntax trees without positions or parentheses.

func ident(name string) *ast.Ident { return ast.NewIdent(name) }

func condExpr() *ast.CondExpr {
	return &ast.CondExpr{Cond: ident("c"), Then: ident("f"), Else: ident("g")}
}

func call(fun ast.Expr) *ast.CallExpr          { return &ast.CallExpr{Fun: fun} }
func index(x ast.Expr) *ast.IndexExpr          { return &ast.IndexExpr{X: x, Index: ident("k")} }
func sel(x ast.Expr, name string) ast.Expr     { return &ast.SelectorExpr{X: x, Sel: ident(name)} }
func exprStmt(x ast.Expr) *ast.ExprStmt        { return &ast.ExprStmt{X: x} }
func block(list ...ast.Stmt) *ast.BlockStmt    { return &ast.BlockStmt{List: list} }
func assign(lhs, rhs ast.Expr) *ast.AssignStmt { return assignTok(lhs, token.ASSIGN, rhs) }
func assignTok(lhs ast.Expr, tok token.Token, rhs ast.Expr) *ast.AssignStmt {
	return &ast.AssignStmt{Lhs: []ast.Expr{lhs}, Tok: tok, Rhs: []ast.Expr{rhs}}
}

// TestCondExprNoParens checks the printing of syntax trees built by
// programs, without the parentheses that the parser would produce: a
// conditional expression whose "if" would start a statement must be
// parenthesized, and nowhere else.
func TestCondExprNoParens(t *testing.T) {
	for _, test := range []struct {
		stmt ast.Stmt
		want string
	}{
		// at the start of a statement
		{exprStmt(call(condExpr())), "(if c { f } else { g })()"},
		{exprStmt(&ast.CallExpr{Fun: condExpr(), Args: []ast.Expr{ident("x"), ident("y")}}), "(if c { f } else { g })(x, y)"},
		{assign(index(condExpr()), ident("v")), "(if c { f } else { g })[k] = v"},
		{&ast.AssignStmt{Lhs: []ast.Expr{index(condExpr()), ident("x")}, Tok: token.ASSIGN, Rhs: []ast.Expr{ident("v"), ident("w")}},
			"(if c { f } else { g })[k], x = v, w"},
		{assign(sel(sel(condExpr(), "a"), "b"), ident("v")), "(if c { f } else { g }).a.b = v"},
		{exprStmt(call(sel(sel(condExpr(), "a"), "m"))), "(if c { f } else { g }).a.m()"},
		{&ast.IncDecStmt{X: index(condExpr()), Tok: token.INC}, "(if c { f } else { g })[k]++"},
		{&ast.IncDecStmt{X: condExpr(), Tok: token.DEC}, "(if c { f } else { g })--"},
		{&ast.SendStmt{Chan: condExpr(), Value: ident("v")}, "(if c { f } else { g }) <- v"},
		{exprStmt(&ast.BinaryExpr{X: condExpr(), Op: token.ADD, Y: ident("x")}), "(if c { f } else { g }) + x"},
		{exprStmt(&ast.BinaryExpr{X: &ast.BinaryExpr{X: condExpr(), Op: token.MUL, Y: ident("y")}, Op: token.ADD, Y: ident("x")}),
			"(if c { f } else { g })*y + x"}, // spaced like a ParenExpr
		{exprStmt(&ast.ErrorExpr{X: call(condExpr())}), "(if c { f } else { g })()!"},
		{exprStmt(&ast.TypeAssertExpr{X: condExpr(), Type: ident("T")}), "(if c { f } else { g }).(T)"},
		{exprStmt(&ast.SliceExpr{X: condExpr(), Low: ident("i"), High: ident("j")}), "(if c { f } else { g })[i:j]"},
		{exprStmt(&ast.IndexListExpr{X: condExpr(), Indices: []ast.Expr{ident("A"), ident("B")}}), "(if c { f } else { g })[A, B]"},
		{&ast.LabeledStmt{Label: ident("L"), Stmt: exprStmt(call(condExpr()))}, "L:\n(if c { f } else { g })()"},
		{&ast.SwitchStmt{Body: block(&ast.CaseClause{Body: []ast.Stmt{exprStmt(call(condExpr()))}})},
			"switch {\ndefault:\n\t(if c { f } else { g })()\n}"},
		{&ast.SelectStmt{Body: block(&ast.CommClause{Body: []ast.Stmt{exprStmt(call(condExpr()))}})},
			"select {\ndefault:\n\t(if c { f } else { g })()\n}"},
		{&ast.IfStmt{Cond: ident("ok"), Body: block(), Else: block(exprStmt(call(condExpr())))},
			"if ok {\n} else {\n\t(if c { f } else { g })()\n}"},
		{block(exprStmt(call(condExpr()))), "{\n\t(if c { f } else { g })()\n}"},
		{exprStmt(call(&ast.FuncLit{Type: &ast.FuncType{Params: &ast.FieldList{}}, Body: block(exprStmt(call(condExpr())))})),
			"func() {\n\t(if c { f } else { g })()\n}()"},

		// already parenthesized
		{exprStmt(call(&ast.ParenExpr{X: condExpr()})), "(if c { f } else { g })()"},

		// not at the start of a statement
		{exprStmt(call(sel(&ast.BinaryExpr{X: condExpr(), Op: token.ADD, Y: ident("x")}, "m"))), "(if c { f } else { g } + x).m()"},
		{exprStmt(&ast.UnaryExpr{Op: token.ARROW, X: condExpr()}), "<-if c { f } else { g }"},
		{assign(&ast.StarExpr{X: condExpr()}, ident("v")), "*if c { f } else { g } = v"},
		{assignTok(ident("x"), token.DEFINE, call(condExpr())), "x := if c { f } else { g }()"},
		{assignTok(ident("x"), token.DEFINE, &ast.UnaryExpr{Op: token.SUB, X: condExpr()}), "x := -if c { f } else { g }"},
		{assignTok(ident("x"), token.DEFINE, &ast.UnaryExpr{Op: token.NOT, X: condExpr()}), "x := !if c { f } else { g }"},
		{assignTok(ident("x"), token.DEFINE, sel(index(condExpr()), "f")), "x := if c { f } else { g }[k].f"},
		{&ast.GoStmt{Call: call(condExpr())}, "go if c { f } else { g }()"},
		{&ast.DeferStmt{Call: call(condExpr())}, "defer if c { f } else { g }()"},
		{&ast.ReturnStmt{Results: []ast.Expr{call(condExpr())}}, "return if c { f } else { g }()"},
		{exprStmt(&ast.CallExpr{Fun: ident("use"), Args: []ast.Expr{condExpr()}}), "use(if c { f } else { g })"},

		// simple statements in control clauses
		{&ast.IfStmt{Init: exprStmt(call(condExpr())), Cond: ident("ok"), Body: block()}, "if if c { f } else { g }(); ok {\n}"},
		{&ast.IfStmt{Cond: condExpr(), Body: block()}, "if if c { f } else { g } {\n}"},
		{&ast.ForStmt{Post: &ast.IncDecStmt{X: index(condExpr()), Tok: token.INC}, Body: block()}, "for ; ; if c { f } else { g }[k]++ {\n}"},
		{&ast.ForStmt{Init: assignTok(ident("i"), token.DEFINE, condExpr()), Cond: ident("ok"), Body: block()},
			"for i := if c { f } else { g }; ok; {\n}"},
		{&ast.SwitchStmt{Init: assign(index(condExpr()), ident("v")), Body: block()}, "switch if c { f } else { g }[k] = v; {\n}"},
		{&ast.TypeSwitchStmt{Assign: exprStmt(&ast.TypeAssertExpr{X: condExpr()}), Body: block()},
			"switch if c { f } else { g }.(type) {\n}"},
		{&ast.SelectStmt{Body: block(&ast.CommClause{Comm: &ast.SendStmt{Chan: condExpr(), Value: ident("v")}})},
			"select {\ncase if c { f } else { g } <- v:\n}"},
		{&ast.RangeStmt{Key: ident("x"), Tok: token.DEFINE, X: condExpr(), Body: block()}, "for x := range if c { f } else { g } {\n}"},
	} {
		src := printFunc(t, test.stmt)
		// Labels are outdented by one level.
		want := strings.Replace("package p\n\nfunc _() {\n"+indent(test.want)+"\n}\n", "\tL:", "L:", 1)
		if src != want {
			t.Errorf("got:\n%s\nwant:\n%s", src, want)
			continue
		}
		f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.SkipObjectResolution)
		if err != nil {
			t.Errorf("%s: %v", src, err)
			continue
		}
		if got, want := noParensString(f.Decls[0].(*ast.FuncDecl).Body.List[0]), noParensString(test.stmt); got != want {
			t.Errorf("%s: parsed differently:\n%s\nwant:\n%s", src, got, want)
		}
		// Formatting the result changes nothing.
		if again, err := Source([]byte(src)); err != nil || string(again) != src {
			t.Errorf("%s: format is not idempotent: %v\n%s", src, err, again)
		}
	}
}

// TestCondExprNoParensNode checks the printing of single statements and
// expressions.
func TestCondExprNoParensNode(t *testing.T) {
	for _, test := range []struct {
		node any
		want string
	}{
		{exprStmt(call(condExpr())), "(if c { f } else { g })()"},
		{[]ast.Stmt{exprStmt(ident("x")), exprStmt(call(condExpr()))}, "x\n(if c { f } else { g })()"},
		{call(condExpr()), "if c { f } else { g }()"}, // an expression, not a statement
	} {
		var buf bytes.Buffer
		if err := Node(&buf, token.NewFileSet(), test.node); err != nil {
			t.Fatal(err)
		}
		if got := buf.String(); got != test.want {
			t.Errorf("got:\n%s\nwant:\n%s", got, test.want)
		}
	}
}

// TestCondExprInvalidNesting checks that invalid syntax trees with
// conditional expressions are printed without panicking, and that the result
// is rejected rather than read back as a different valid program.
func TestCondExprInvalidNesting(t *testing.T) {
	for _, stmt := range []ast.Stmt{
		assignTok(ident("x"), token.DEFINE, &ast.CondExpr{Cond: condExpr(), Then: ident("a"), Else: ident("b")}),
		assignTok(ident("x"), token.DEFINE, &ast.CondExpr{Cond: ident("c"), Then: condExpr(), Else: ident("b")}),
		assignTok(ident("x"), token.DEFINE, &ast.CondExpr{Cond: ident("c"), Then: ident("a"), Else: condExpr()}),
		exprStmt(call(&ast.CondExpr{Cond: condExpr(), Then: ident("a"), Else: ident("b")})),
		exprStmt(&ast.CompositeLit{Type: condExpr()}), // a conditional expression is not a type
	} {
		src := printFunc(t, stmt)
		if _, err := parser.ParseFile(token.NewFileSet(), "", src, parser.SkipObjectResolution); err == nil {
			t.Errorf("invalid syntax tree printed as valid source:\n%s", src)
		}
	}
}

// printFunc formats a file with a function whose body is stmt.
func printFunc(t *testing.T, stmt ast.Stmt) string {
	t.Helper()
	file := &ast.File{
		Name: ident("p"),
		Decls: []ast.Decl{&ast.FuncDecl{
			Name: ident("_"),
			Type: &ast.FuncType{Params: &ast.FieldList{}},
			Body: block(stmt),
		}},
	}
	var buf bytes.Buffer
	if err := Node(&buf, token.NewFileSet(), file); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

// noParensString returns a description of the syntax tree n that ignores
// positions and parentheses.
func noParensString(n ast.Node) string {
	var b strings.Builder
	var open []bool // for each node being visited, whether it was described
	ast.Inspect(n, func(n ast.Node) bool {
		if n == nil {
			if open[len(open)-1] {
				b.WriteString(")")
			}
			open = open[:len(open)-1]
			return true
		}
		_, paren := n.(*ast.ParenExpr)
		open = append(open, !paren)
		if paren {
			return true
		}
		fmt.Fprintf(&b, "(%T", n)
		switch n := n.(type) {
		case *ast.Ident:
			b.WriteString(" " + n.Name)
		case *ast.BinaryExpr:
			b.WriteString(" " + n.Op.String())
		case *ast.UnaryExpr:
			b.WriteString(" " + n.Op.String())
		case *ast.AssignStmt:
			b.WriteString(" " + n.Tok.String())
		case *ast.IncDecStmt:
			b.WriteString(" " + n.Tok.String())
		}
		return true
	})
	return b.String()
}
