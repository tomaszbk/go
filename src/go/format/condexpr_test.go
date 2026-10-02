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

// TestCondExpr checks the formatting of conditional expressions: a
// conditional expression written on one line stays on one line; otherwise it
// is laid out like an if statement. The result must be stable.
func TestCondExpr(t *testing.T) {
	for _, test := range []struct{ src, want string }{
		// single line, normalized spacing
		{"x := if c {1} else {2}", "x := if c { 1 } else { 2 }"},
		{"x := if   n==1   {  \"item\"  }   else   { \"items\" }", `x := if n == 1 { "item" } else { "items" }`},
		{"x := if c { a*b } else { a+b }", "x := if c { a * b } else { a + b }"},
		{"x := if p == (Point{}) { a } else { b }", "x := if p == (Point{}) { a } else { b }"},
		{"x := price+if c { tax } else { 0 }", "x := price + if c { tax } else { 0 }"},
		{"f(a, price+if c { tax } else { 0 })", "f(a, price + if c { tax } else { 0 })"},
		{"x := if c { tax } else { 0 }+price", "x := if c { tax } else { 0 } + price"},
		{"x := if c { 1 /* one */ } else { 2 }", "x := if c { 1 /* one */ } else { 2 }"},
		{"x := if c { read()! } else { 0 }", "x := if c { read()! } else { 0 }"},

		// multiple lines, laid out like an if statement
		{"x := if c {\n1\n} else {\n2\n}", "x := if c {\n\t1\n} else {\n\t2\n}"},
		{"x := if c { 1\n} else { 2 }", "x := if c {\n\t1\n} else {\n\t2\n}"},
		{"x := if c { 1 } else\n{ 2 }", "x := if c {\n\t1\n} else {\n\t2\n}"},
		{"x := if c { 1 // one\n} else { 2 }", "x := if c {\n\t1 // one\n} else {\n\t2\n}"},
		{"f(a, if c {\n1\n} else {\n2\n})", "f(a, if c {\n\t1\n} else {\n\t2\n})"},
		{"f(a,\nif c {\n1\n} else {\n2\n})", "f(a,\n\tif c {\n\t\t1\n\t} else {\n\t\t2\n\t})"},
		{"x := price +\nif c {\ntax\n} else {\n0\n}", "x := price +\n\tif c {\n\t\ttax\n\t} else {\n\t\t0\n\t}"},

		// a branch that cannot be printed on one line
		{"x := if c { read() or err { return err } } else { 0 }",
			"x := if c {\n\tread() or err {\n\t\treturn err\n\t}\n} else {\n\t0\n}"},

		// parentheses are kept in control clauses
		{"if (if c { a } else { b }) {\n}", "if (if c { a } else { b }) {\n}"},
		{"if (x == if c { a } else { b }) {\n}", "if (x == if c { a } else { b }) {\n}"},
		{"for (if c { a } else { b }) {\n}", "for (if c { a } else { b }) {\n}"},
		{"switch (if c { a } else { b }) {\n}", "switch (if c { a } else { b }) {\n}"},
		{"for range (if c { a } else { b }) {\n}", "for range (if c { a } else { b }) {\n}"},
		// but not otherwise
		{"if (x == y) {\n}", "if x == y {\n}"},

		// "if" at the start of a statement is an if statement
		{"if c { 1 } else { 2 }", "if c {\n\t1\n} else {\n\t2\n}"},
	} {
		src := "package p\n\nfunc _() {\n" + test.src + "\n}\n"
		want := "package p\n\nfunc _() {\n" + indent(test.want) + "\n}\n"
		got, err := Source([]byte(src))
		if err != nil {
			t.Errorf("%q: %v", test.src, err)
			continue
		}
		if string(got) != want {
			t.Errorf("%q:\ngot:\n%s\nwant:\n%s", test.src, got, want)
			continue
		}
		checkStable(t, test.src, src, got)
	}
}

// TestCondExprDecls checks package-level and constant uses.
func TestCondExprDecls(t *testing.T) {
	const src = `package p

const intSize = if ^uint(0)>>63==1 {64} else {32}

var (
	a = if c {1} else {2} // a
	bb = 3 // bb
)
`
	const want = `package p

const intSize = if ^uint(0)>>63 == 1 { 64 } else { 32 }

var (
	a  = if c { 1 } else { 2 } // a
	bb = 3                     // bb
)
`
	got, err := Source([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
	checkStable(t, "decls", src, got)
}

// TestCondExprNode checks the printing of conditional expressions without
// position information, as built by programs.
func TestCondExprNode(t *testing.T) {
	cond := &ast.CondExpr{
		Cond: &ast.BinaryExpr{X: ast.NewIdent("n"), Op: token.EQL, Y: &ast.BasicLit{Kind: token.INT, Value: "1"}},
		Then: &ast.BasicLit{Kind: token.STRING, Value: `"item"`},
		Else: &ast.BasicLit{Kind: token.STRING, Value: `"items"`},
	}
	for _, test := range []struct {
		node ast.Node
		want string
	}{
		{cond, `if n == 1 { "item" } else { "items" }`},
		{&ast.BinaryExpr{X: ast.NewIdent("x"), Op: token.ADD, Y: cond}, `x + if n == 1 { "item" } else { "items" }`},
		{&ast.CondExpr{
			Cond: ast.NewIdent("c"),
			Then: &ast.ErrorExpr{
				X:    &ast.CallExpr{Fun: ast.NewIdent("read")},
				Err:  ast.NewIdent("err"),
				Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent("err")}}}},
			},
			Else: ast.NewIdent("x"),
		}, "if c {\n\tread() or err {\n\t\treturn err\n\t}\n} else {\n\tx\n}"},
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

// checkStable checks that formatted, the result of formatting src, is valid,
// describes the same syntax tree as src (ignoring positions and parentheses
// around control clauses), and does not change when formatted again.
func checkStable(t *testing.T, name, src string, formatted []byte) {
	t.Helper()
	again, err := Source(formatted)
	if err != nil {
		t.Errorf("%s: reformat: %v\n%s", name, err, formatted)
		return
	}
	if !bytes.Equal(formatted, again) {
		t.Errorf("%s: format is not idempotent:\n%s\nformatted again:\n%s", name, formatted, again)
	}
	if a, b := treeString(t, src), treeString(t, string(formatted)); a != b {
		t.Errorf("%s: formatting changed the syntax tree:\n%s\nformatted:\n%s", name, a, b)
	}
}

// treeString returns a description of the syntax tree of src that does not
// depend on positions. Parentheses around control clause expressions, which
// gofmt may remove, are ignored.
func treeString(t *testing.T, src string) string {
	t.Helper()
	f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatal(err)
	}
	unparen := func(x *ast.Expr) {
		if *x != nil {
			*x = ast.Unparen(*x)
		}
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.IfStmt:
			unparen(&n.Cond)
		case *ast.ForStmt:
			unparen(&n.Cond)
		case *ast.SwitchStmt:
			unparen(&n.Tag)
		case *ast.RangeStmt:
			unparen(&n.X)
		}
		return true
	})
	var b strings.Builder
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			b.WriteString(")")
			return true
		}
		fmt.Fprintf(&b, "(%T", n)
		switch n := n.(type) {
		case *ast.Ident:
			b.WriteString(" " + n.Name)
		case *ast.BasicLit:
			b.WriteString(" " + n.Value)
		case *ast.BinaryExpr:
			b.WriteString(" " + n.Op.String())
		case *ast.UnaryExpr:
			b.WriteString(" " + n.Op.String())
		case *ast.AssignStmt:
			b.WriteString(" " + n.Tok.String())
		}
		return true
	})
	return b.String()
}

func indent(s string) string {
	return "\t" + strings.ReplaceAll(s, "\n", "\n\t")
}
