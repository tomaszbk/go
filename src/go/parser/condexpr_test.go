package parser

import (
	"fmt"
	"go/ast"
	"go/scanner"
	"go/token"
	"strings"
	"testing"
)

// condExprs returns the conditional expressions in f, in source order.
func condExprs(f ast.Node) []*ast.CondExpr {
	var list []*ast.CondExpr
	ast.Inspect(f, func(n ast.Node) bool {
		if x, ok := n.(*ast.CondExpr); ok {
			list = append(list, x)
		}
		return true
	})
	return list
}

func TestCondExprPositions(t *testing.T) {
	const src = `package p

var a = if c { x } else { y }
var b = if  n > 0  {
	f(n)
}  else  {
	g()
}
`
	fset := token.NewFileSet()
	f, err := ParseFile(fset, "p.go", src, SkipObjectResolution)
	if err != nil {
		t.Fatal(err)
	}
	list := condExprs(f)
	if len(list) != 2 {
		t.Fatalf("found %d conditional expressions, want 2", len(list))
	}
	pos := func(p token.Pos) string {
		p1 := fset.Position(p)
		return fmt.Sprintf("%d:%d", p1.Line, p1.Column)
	}
	for i, want := range []struct {
		if_, lbrace, rbrace, else_, elseLbrace, elseRbrace, end string
		cond, then, els                                         string
	}{
		{"3:9", "3:14", "3:18", "3:20", "3:25", "3:29", "3:30", "*ast.Ident", "*ast.Ident", "*ast.Ident"},
		{"4:9", "4:20", "6:1", "6:4", "6:10", "8:1", "8:2", "*ast.BinaryExpr", "*ast.CallExpr", "*ast.CallExpr"},
	} {
		x := list[i]
		got := []string{pos(x.If), pos(x.Lbrace), pos(x.Rbrace), pos(x.ElsePos), pos(x.ElseLbrace), pos(x.ElseRbrace), pos(x.End())}
		exp := []string{want.if_, want.lbrace, want.rbrace, want.else_, want.elseLbrace, want.elseRbrace, want.end}
		if strings.Join(got, " ") != strings.Join(exp, " ") {
			t.Errorf("%d: positions (If, Lbrace, Rbrace, ElsePos, ElseLbrace, ElseRbrace, End) = %v, want %v", i, got, exp)
		}
		if x.Pos() != x.If {
			t.Errorf("%d: Pos() = %s, want If = %s", i, pos(x.Pos()), pos(x.If))
		}
		if got := fmt.Sprintf("%T %T %T", x.Cond, x.Then, x.Else); got != want.cond+" "+want.then+" "+want.els {
			t.Errorf("%d: operands = %s", i, got)
		}
		// The operands are within their braces.
		if !(x.If < x.Cond.Pos() && x.Cond.End() <= x.Lbrace && x.Lbrace < x.Then.Pos() && x.Then.End() <= x.Rbrace &&
			x.Rbrace < x.ElsePos && x.ElsePos < x.ElseLbrace && x.ElseLbrace < x.Else.Pos() && x.Else.End() <= x.ElseRbrace) {
			t.Errorf("%d: inconsistent positions", i)
		}
	}
}

func TestCondExprValid(t *testing.T) {
	for _, test := range []struct {
		src string
		n   int // number of conditional expressions
	}{
		{`x := if c { 1 } else { 2 }`, 1},
		{"x := if c {\n\t1\n} else {\n\t2\n}", 1},
		{"x := if c { 1 // one\n} else { 2 /* two\n*/ }", 1},
		{"x := if c { 1 } else\n{ 2 }", 1},
		{`f(if c { a } else { b }, if d { 1 } else { 2 })`, 2},
		{`_ = []int{if c { 1 } else { 2 }, 3}`, 1},
		{`_ = T{a: if c { 1 } else { 2 }}`, 1},
		{`return if c { 1 } else { 2 }`, 1},
		{`return if c { 1 } else { 2 }, nil`, 1},
		{`total := price + if taxable { tax } else { 0 }`, 1},
		{`total := if taxable { tax } else { 0 } + price`, 1},
		{`_ = -if c { 1 } else { 2 }`, 1},
		{`_ = if c { a } else { b }.f`, 1},
		{`_ = if c { m1 } else { m2 }[k]`, 1},
		{`(if c { m1 } else { m2 })[k] = v`, 1},
		{`_ = if c { f } else { g }(1, 2)`, 1},
		{`go if c { f } else { g }()`, 1},
		{`defer if c { f } else { g }()`, 1},
		{`ch <- if c { 1 } else { 2 }`, 1},
		{`switch { case if c { true } else { false }: }`, 1},
		{`select { case x := <-if c { ch1 } else { ch2 }: _ = x }`, 1},
		// control clauses
		{`if (if c { a } else { b }) { }`, 1},
		{`if if c { a } else { b } { }`, 1},
		{`if v := if c { T{1} } else { T{2} }; v.x > 0 { }`, 1},
		{`for x := if c { 1 } else { 2 }; x < 10; x++ { }`, 1},
		{`for i := 0; i < (if c { n } else { m }); i++ { }`, 1},
		{`switch if c { 1 } else { 2 } { case 1: }`, 1},
		{`for _, v := range if c { a } else { b } { _ = v }`, 1},
		// composite literals
		{`x := if p == (Point{}) { a } else { b }`, 1},
		{`x := if c { T{1} } else { T{2} }`, 1},
		{`x := if c { {1, 2} } else { {3} }`, 1},
		{`x := if c { []int{1} } else { nil }`, 1},
		// nesting inside other constructs is valid
		{`x := if c { f(if d { 1 } else { 2 }) } else { 3 }`, 2},
		{`x := if c { (if d { 1 } else { 2 }) + 1 } else { 3 }`, 2},
		{`x := if c { if d { 1 } else { 2 } + 1 } else { 3 }`, 2},
		{`x := if (if d { true } else { false }) == true { 1 } else { 2 }`, 2},
		{`x := if c { func() int { return if d { 1 } else { 2 } } } else { nil }`, 2},
		// error propagation and handlers in branches
		{`data := if cached { readCache()! } else { fetch()! }`, 1},
		{`data := if cached { readCache() or err { return err } } else { fetch()! }`, 1},
		{"data := if cached {\n\treadCache()!\n} else {\n\tfetch() or err {\n\t\treturn err\n\t}\n}", 1},
		// declarations
		{`const intSize = if ^uint(0)>>63 == 1 { 64 } else { 32 }`, 1},
		{`var a [if c { 1 } else { 2 }]int`, 1},
	} {
		src := "package p; func _() {\n" + test.src + "\n}"
		f, err := ParseFile(token.NewFileSet(), "", src, AllErrors)
		if err != nil {
			t.Errorf("%s: %v", test.src, err)
			continue
		}
		if n := len(condExprs(f)); n != test.n {
			t.Errorf("%s: found %d conditional expressions, want %d", test.src, n, test.n)
		}
	}
}

// TestCondExprStatement checks that "if" at the start of a statement is
// still an if statement, including else-if chains and composite literals in
// conditions.
func TestCondExprStatement(t *testing.T) {
	for _, src := range []string{
		`if c { a } else { b }`,
		`if c { a } else if d { b } else { c }`,
		`if c { if d { a } else { b } }`,
		`if x == (T{}) { a }`,
		`if x := (T{}); x.y { }`,
		`L: if c { a } else { b }`,
	} {
		f, err := ParseFile(token.NewFileSet(), "", "package p; func _() {\n"+src+"\n}", AllErrors)
		if err != nil {
			t.Errorf("%s: %v", src, err)
			continue
		}
		if n := len(condExprs(f)); n != 0 {
			t.Errorf("%s: found %d conditional expressions, want 0", src, n)
		}
		stmt := f.Decls[0].(*ast.FuncDecl).Body.List[0]
		if l, ok := stmt.(*ast.LabeledStmt); ok {
			stmt = l.Stmt
		}
		if _, ok := stmt.(*ast.IfStmt); !ok {
			t.Errorf("%s: got %T, want *ast.IfStmt", src, stmt)
		}
	}
}

func TestCondExprParseExpr(t *testing.T) {
	x, err := ParseExpr("if c { 1 } else { 2 } + 3")
	if err != nil {
		t.Fatal(err)
	}
	b, ok := x.(*ast.BinaryExpr)
	if !ok {
		t.Fatalf("got %T, want *ast.BinaryExpr", x)
	}
	if _, ok := b.X.(*ast.CondExpr); !ok {
		t.Fatalf("got %T, want *ast.CondExpr as left operand", b.X)
	}
}

// TestCondExprErrors checks all errors reported for invalid conditional
// expressions, to verify both the messages and the absence of follow-on
// errors.
func TestCondExprErrors(t *testing.T) {
	const (
		noElse   = "conditional expression requires an else branch"
		nested   = "conditional expressions cannot be chained or nested; use a switch statement"
		notExpr  = "conditional expression branch must be a single expression"
		hasInit  = "conditional expression cannot have an init statement"
		noLbrace = "expected '{', found "
		noRbrace = "expected '}', found 'else'"
	)
	for _, test := range []struct {
		src  string
		errs []string // "line:col: msg", lines are relative to the test source
	}{
		// missing else
		{`x := if c { 1 }`, []string{"1:16: " + noElse}},
		{"x := if c {\n\t1\n}\ny := 2", []string{"3:2: " + noElse}},
		{`f(if c { 1 }, 2)`, []string{"1:13: " + noElse}},
		{`x := if c { 1 } + 2`, []string{"1:17: " + noElse}},

		// chaining and nesting, with and without parentheses
		{`x := if a { 0 } else if b { 1 } else { 2 }`, []string{"1:22: " + nested}},
		{`x := if a { 0 } else if b { 1 } else if c { 2 } else { 3 }`, []string{"1:22: " + nested}},
		{`x := if a { 0 } else if b { 1 }`, []string{"1:22: " + nested}},
		{`x := if a { 0 } else { if b { 1 } else { 2 } }`, []string{"1:24: " + nested}},
		{`x := if a { 0 } else { (if b { 1 } else { 2 }) }`, []string{"1:25: " + nested}},
		{`x := if a { ((if b { 1 } else { 2 })) } else { 0 }`, []string{"1:15: " + nested}},
		{`x := if if a { b } else { c } { 1 } else { 2 }`, []string{"1:9: " + nested}},
		{`x := if (if a { b } else { c }) { 1 } else { 2 }`, []string{"1:10: " + nested}},
		{`x := if a { if b { 1 } } else { 2 }`, []string{"1:24: " + noElse}},

		// init statements
		{`x := if v, ok := m[k]; ok { v } else { 0 }`, []string{"1:9: " + hasInit}},
		{`x := if v := f(); v > 0 { v } else { 0 }`, []string{"1:9: " + hasInit}},

		// statements and other non-expressions in a branch
		{`x := if c { y = 1 } else { 2 }`, []string{"1:15: " + notExpr}},
		{`x := if c { y := 1 } else { 2 }`, []string{"1:15: " + notExpr}},
		{`x := if c { y++ } else { 2 }`, []string{"1:14: " + notExpr}},
		{`x := if c { ch <- 1 } else { 2 }`, []string{"1:16: " + notExpr}},
		{`x := if c { f(); g() } else { 2 }`, []string{"1:16: " + notExpr}},
		{`x := if c { f(); } else { 2 }`, []string{"1:16: " + notExpr}},
		{`x := if c { a, b } else { 2 }`, []string{"1:14: " + notExpr}},
		{`x := if c { return 1 } else { 2 }`, []string{"1:13: " + notExpr}},
		{`x := if c { return if a { b } else { c } } else { 2 }`, []string{"1:13: " + notExpr}},
		{`x := if c { var y = 1 } else { 2 }`, []string{"1:13: " + notExpr}},
		{`x := if c { for {} } else { 2 }`, []string{"1:13: " + notExpr}},
		{`x := if c { } else { 2 }`, []string{"1:13: " + notExpr}},
		{`x := if c { 1 } else { }`, []string{"1:24: " + notExpr}},
		{`x := if c { ) } else { 2 }`, []string{"1:13: " + notExpr}},
		{"x := if c {\n\tf()\n\tg()\n} else {\n\t2\n}", []string{"3:2: " + notExpr}},
		{"x := if c { y = 1 }\nz := 2", []string{"1:15: " + notExpr}},

		// missing braces
		{"x := if c { 1 } else 2\ny := 3", []string{"1:22: " + noLbrace + "2"}},
		{"x := if c { 1 else { 2 }\ny := 3", []string{"1:15: " + noRbrace}},
		{"x := if c { else { 2 }\ny := 3", []string{"1:13: " + noRbrace}},

		// header errors use the wording of if statements
		{`x := if { 1 } else { 2 }`, []string{"1:9: missing condition in if statement"}},
		{"x := if c\n{ 1 } else { 2 }", []string{"1:10: unexpected newline, expecting { after if clause"}},
	} {
		src := "package p; func _() {\n" + test.src + "\n}"
		_, err := ParseFile(token.NewFileSet(), "", src, AllErrors)
		var got []string
		if list, ok := err.(scanner.ErrorList); ok {
			for _, e := range list {
				got = append(got, fmt.Sprintf("%d:%d: %s", e.Pos.Line-1, e.Pos.Column, e.Msg))
			}
		} else if err != nil {
			t.Errorf("%s: %v", test.src, err)
			continue
		}
		if strings.Join(got, "\n") != strings.Join(test.errs, "\n") {
			t.Errorf("%s:\ngot errors:\n\t%s\nwant:\n\t%s", test.src, strings.Join(got, "\n\t"), strings.Join(test.errs, "\n\t"))
		}
	}
}

// TestCondExprBadOperands checks that the parser does not produce a
// conditional expression whose condition or branch is a conditional
// expression, even after reporting an error.
func TestCondExprBadOperands(t *testing.T) {
	for _, src := range []string{
		`x := if a { 0 } else if b { 1 } else { 2 }`,
		`x := if a { 0 } else { if b { 1 } else { 2 } }`,
		`x := if a { (if b { 1 } else { 2 }) } else { 0 }`,
		`x := if (if a { b } else { c }) { 1 } else { 2 }`,
	} {
		f, _ := ParseFile(token.NewFileSet(), "", "package p; func _() {\n"+src+"\n}", AllErrors)
		list := condExprs(f)
		if len(list) != 1 {
			t.Errorf("%s: found %d conditional expressions, want 1", src, len(list))
			continue
		}
		x := list[0]
		for _, y := range []ast.Expr{x.Cond, x.Then, x.Else} {
			if _, ok := ast.Unparen(y).(*ast.CondExpr); ok {
				t.Errorf("%s: conditional expression operand", src)
			}
		}
		if x.Pos() >= x.End() {
			t.Errorf("%s: invalid positions", src)
		}
	}
}

func TestCondExprResolution(t *testing.T) {
	const src = `package p
func f(c bool, a, b int) int {
	return if c { a } else { b + func(a int) int { return a }(1) }
}`
	f, err := ParseFile(token.NewFileSet(), "", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	params := f.Decls[0].(*ast.FuncDecl).Type.Params.List
	c, a, b := params[0].Names[0], params[1].Names[0], params[1].Names[1]
	x := condExprs(f)[0]
	if id := x.Cond.(*ast.Ident); id.Obj != c.Obj {
		t.Errorf("condition does not refer to parameter c")
	}
	if id := x.Then.(*ast.Ident); id.Obj != a.Obj {
		t.Errorf("then branch does not refer to parameter a")
	}
	sum := x.Else.(*ast.BinaryExpr)
	if id := sum.X.(*ast.Ident); id.Obj != b.Obj {
		t.Errorf("else branch does not refer to parameter b")
	}
	lit := sum.Y.(*ast.CallExpr).Fun.(*ast.FuncLit)
	inner := lit.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.Ident)
	if inner.Obj == a.Obj || inner.Obj != lit.Type.Params.List[0].Names[0].Obj {
		t.Errorf("function literal parameter a is not resolved in its own scope")
	}
}

// TestCondExprDepth checks that deeply nested conditional expressions are
// rejected with an error instead of exhausting the stack.
func TestCondExprDepth(t *testing.T) {
	if testing.Short() {
		t.Skip("test requires significant memory")
	}
	for _, test := range []struct {
		name, left, base, right string
		valid                   bool // valid if not too deep
	}{
		{"cond", "if f(", "c", ") { 1 } else { 2 }", true},
		{"arg", "if c { f(", "1", ") } else { 0 }", true},
		{"branch", "if c { ", "0", " } else { 0 }", false}, // directly nested
	} {
		for _, n := range []int{100, maxNestLev} {
			src := "package p; var x = " + strings.Repeat(test.left, n) + test.base + strings.Repeat(test.right, n)
			_, err := ParseFile(token.NewFileSet(), "", src, SkipObjectResolution)
			if n < maxNestLev {
				if test.valid != (err == nil) {
					t.Errorf("%s, depth %d: got error %v", test.name, n, err)
				}
			} else if err == nil || !strings.HasSuffix(err.Error(), "exceeded max nesting depth") {
				t.Errorf("%s, depth %d: got error %v, want nesting depth error", test.name, n, err)
			}
		}
	}
}
