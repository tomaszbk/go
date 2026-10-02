package syntax

import (
	"fmt"
	"strings"
	"testing"
)

// Valid uses of conditional expressions. Each source is a list of
// declarations following "package p; ".
var validCondExprs = []string{
	// basic forms
	`var _ = if c { a } else { b }`,
	`var _ = if c {a} else {b}`,
	"var _ = if c {\n\ta\n} else {\n\tb\n}",
	"var _ = if c { a\n} else { b\n}",
	"var _ = if c {\n\ta // comment\n} else {\n\tb /* comment */\n}",
	"var _ = if c { a /*\n*/ } else { b }",
	"var _ = if c {\n\tf()!\n} else {\n\tg()!\n}",
	`var _ = if !c { !d } else { ! e }`,
	`var _ = if a == b && c != d { -x } else { <-ch }`,
	`var _ = if x.ok() { y.v } else { z[0] }`,

	// operands of other operations and calls
	`var _ = price + if taxable { tax } else { 0 }`,
	`var _ = if taxable { tax } else { 0 } + price`,
	`var _ = -if c { 1 } else { 2 } * 3`,
	`var _ = f(a(), if c() { b() } else { d() }, e())`,
	`var _ = f(if c { a } else { b }...)`,
	`var _ = (if c { s } else { t })[i]`,
	`var _ = if c { s } else { t }[i:j]`,
	`var _ = if c { f } else { g }(x)`,
	`var _ = if c { p } else { q }.f`,
	`var _ = if c { x } else { y }.(T)`,
	`var _ = &T{if c { 1 } else { 2 }}`,

	// composite literals
	`var _ = []int{if c { 1 } else { 2 }, 3}`,
	`var _ = T{X: if c { 1 } else { 2 }}`,
	`var _ = map[string]int{if c { "a" } else { "b" }: if d { 1 } else { 2 }}`,
	`var _ = if c { Point{1, 2} } else { Point{} }`,
	`var _ = if c { {1, 2} } else { {3, 4} }`,
	`var _ = if p == (Point{}) { a } else { b }`,
	`var _ = if (Point{}) == p { a } else { b }`,
	`var _ = if c { []int{1} } else { nil }`,

	// nested conditional expressions as separate operands
	`var _ = if c { f(if d { 1 } else { 2 }) } else { 3 }`,
	`var _ = if c { (if d { 1 } else { 2 }) + 1 } else { 3 }`,
	`var _ = if c { if d { 1 } else { 2 } + 1 } else { 3 }`,
	`var _ = if c { []int{if d { 1 } else { 2 }} } else { nil }`,
	`var _ = if f(if d { true } else { false }) { 1 } else { 2 }`,
	`var _ = if (if d { x } else { y }) == z { 1 } else { 2 }`,
	`var _ = if c { func() int { return if d { 1 } else { 2 } } } else { nil }`,

	// statements and control clauses
	`func _() int { return if c { 1 } else { 2 } }`,
	`func _() (int, int) { return if c { 1 } else { 2 }, if d { 3 } else { 4 } }`,
	`func _() { x := if c { 1 } else { 2 }; x = if c { x } else { 3 }; x += if c { 1 } else { 0 } }`,
	`func _() { if (if c { a } else { b }) { } }`,
	`func _() { if if c { a } else { b } { } }`,
	`func _() { if if c { T{} } else { T{} } == x { } }`,
	`func _() { if x := if c { 1 } else { 2 }; x > 0 { } }`,
	`func _() { for x := if c { 1 } else { 2 }; x < 10; x++ { } }`,
	`func _() { for if c { true } else { false } { } }`,
	`func _() { for _, v := range if c { a } else { b } { _ = v } }`,
	`func _() { switch if c { 1 } else { 2 } { case 1: } }`,
	`func _() { switch x := if c { 1 } else { 2 }; x { } }`,
	`func _() { ch <- if c { 1 } else { 2 } }`,
	`func _() { (if c { s } else { t })[0] = 1 }`,
	`func _() { defer (if c { f } else { g })() }`,
	`func _() { go if c { f } else { g }() }`,
	`func _() { m[if c { 1 } else { 2 }] = if d { 3 } else { 4 } }`,
	`const _ = if ^uint(0)>>63 == 1 { 64 } else { 32 }`,
	`var _ [if c { 1 } else { 2 }]int`,

	// error handling in branches
	`func _() error { _ = if c { f()! } else { g()! }; return nil }`,
	`func _() { _ = if c { f() or err { return } } else { 0 } }`,
	"func _() { _ = if c {\n\tf() or err {\n\t\treturn\n\t}\n} else {\n\tg()!\n} }",

	// or is still an ordinary identifier
	`var _ = if or { or } else { or + 1 }`,

	// existing if statements are unchanged
	`func _() { if c { a } else { b } }`,
	`func _() { if c { a } else if d { b } else { c } }`,
	`func _() { if x := f(); x { } else { } }`,
	`func _() { if p == (Point{}) { } }`,
}

func TestCondExprSyntax(t *testing.T) {
	for _, src := range validCondExprs {
		f, err := Parse(NewFileBase("test.go"), strings.NewReader("package p; "+src), nil, nil, CheckBranches)
		if err != nil {
			t.Errorf("%s: %v", src, err)
			continue
		}
		verifyPrint(t, "test.go", f)
		Inspect(f, func(n Node) bool {
			if e, ok := n.(*CondExpr); ok {
				for _, x := range []Expr{e.Cond, e.Then, e.Else} {
					if x == nil {
						t.Errorf("%s: incomplete CondExpr %s", src, String(e))
					} else if _, ok := x.(*BadExpr); ok {
						t.Errorf("%s: unexpected BadExpr in %s", src, String(e))
					}
				}
				if start, end := StartPos(e), EndPos(e); start != e.Pos() || !end.IsKnown() || end.Cmp(start) <= 0 {
					t.Errorf("%s: invalid extent %s - %s of %s", src, start, end, String(e))
				}
			}
			return true
		})
	}
}

// At the start of a statement, if always begins an if statement.
func TestCondExprStatement(t *testing.T) {
	const src = "package p; func _() { if c { a } else { b }; if c { } else if d { } }"
	f, err := Parse(NewFileBase("test.go"), strings.NewReader(src), nil, nil, CheckBranches)
	if err != nil {
		t.Fatal(err)
	}
	body := f.DeclList[0].(*FuncDecl).Body.List
	for _, s := range body {
		if _, ok := s.(*IfStmt); !ok {
			t.Errorf("got %T, want *IfStmt", s)
		}
	}
	Inspect(f, func(n Node) bool {
		if _, ok := n.(*CondExpr); ok {
			t.Errorf("unexpected CondExpr %s", String(n))
		}
		return true
	})
}

func TestCondExprShape(t *testing.T) {
	const src = "package p; var _ = if c {\n\ta\n} else {\n\tb\n}"
	f, err := Parse(NewFileBase("test.go"), strings.NewReader(src), nil, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	e, ok := f.DeclList[0].(*VarDecl).Values.(*CondExpr)
	if !ok {
		t.Fatalf("got %T, want *CondExpr", f.DeclList[0].(*VarDecl).Values)
	}
	check := func(what string, got Pos, line, col uint) {
		t.Helper()
		if got.Line() != line || got.Col() != col {
			t.Errorf("%s at %d:%d, want %d:%d", what, got.Line(), got.Col(), line, col)
		}
	}
	check("if", e.Pos(), 1, 20)
	check("cond", e.Cond.Pos(), 1, 23)
	check("then", e.Then.Pos(), 2, 2)
	check("else", e.Else.Pos(), 4, 2)
	check("rbrace", e.Rbrace, 5, 1)
	check("EndPos", EndPos(e), 5, 1)
}

var condExprPositions = []test{
	{"CondExpr", `@if c { a } else { b }`},
	{"CondExpr", `@if (T{}) == x { T{} } else { T{} }`},
	{"Operation", `@-if c { a } else { b }`},
	{"Operation", `if c { a } else { b } @+ x`},
	{"IndexExpr", `if c { a } else { b }@[i]`},
	{"CallExpr", `if c { a } else { b }@()`},
}

func TestCondExprPos(t *testing.T) {
	testPos(t, condExprPositions, "package p; var _ = T{ ", " }",
		func(f *File) Node { return f.DeclList[0].(*VarDecl).Values.(*CompositeLit).ElemList[0] },
	)
}

func TestCondExprString(t *testing.T) {
	for _, test := range [][2]string{
		dup("if c { a } else { b }"),
		{"if c {\n a\n} else {\n b\n}", "if c { a } else { b }"},
		{"if c { []int{1} } else { func() {} }", "if c { []int{…} } else { func() {} }"},
		dup("x + if c { a } else { b }"),
		dup("(if c { a } else { b })[i]"),
		dup("f(if c { a } else { b })"),
	} {
		src := "package p; var _ = " + test[0]
		f, err := Parse(nil, strings.NewReader(src), nil, nil, 0)
		if err != nil {
			t.Errorf("%s: %s", test[0], err)
			continue
		}
		x := f.DeclList[0].(*VarDecl).Values
		if got := String(x); got != test[1] {
			t.Errorf("%s: got %s, want %s", test[0], got, test[1])
		}
	}
}

// Each test source must report exactly one syntax error, at the position
// marked by @ (in the source after "package p; "), with the given message
// following the "syntax error: " prefix.
var condExprErrors = []struct{ src, msg string }{
	// missing else
	{`var _ = if c { a }@; var _ = 1`, "conditional expression requires an else branch"},
	{`var _ = f(if c { a }@)`, "conditional expression requires an else branch"},
	{`func _() { x := if c { a }@; _ = x }`, "conditional expression requires an else branch"},

	// chaining and nesting
	{`var _ = if a { 1 } else @if b { 2 } else { 3 }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},
	{`var _ = if a { 1 } else @if b { 2 }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},
	{`var _ = if a { 1 } else @if b { 2 } else if c { 3 } else { 4 }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},
	{`var _ = if a { 1 } else { @if b { 2 } else { 3 } }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},
	{`var _ = if a { 1 } else { (@if b { 2 } else { 3 }) }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},
	{`var _ = if a { ((@if b { 2 } else { 3 })) } else { 1 }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},
	{"var _ = if a {\n\t@if b { 2 } else { 3 }\n} else {\n\t1\n}", "conditional expressions cannot be chained or nested; use a switch statement"},
	{`var _ = if @if a { true } else { false } { 1 } else { 2 }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},
	{`var _ = if (@if a { true } else { false }) { 1 } else { 2 }; var _ = 1`, "conditional expressions cannot be chained or nested; use a switch statement"},

	// statements in a branch
	{`func _() { _ = if c { x @:= 1 } else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { x @= 1 } else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { x@++ } else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { 1 } else { @return 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { @var x int } else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { @} else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { 1 } else { @} }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { a@; b } else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { a@; } else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { @; a } else { 2 } }`, "conditional expression branch must be a single expression"},
	{"func _() { _ = if c {\n\ty @:= 1\n\ty\n} else {\n\t2\n} }", "conditional expression branch must be a single expression"},
	{"func _() { _ = if c {\n\tf()\n\t@g()\n} else {\n\t2\n} }", "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { @for { } } else { 2 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { a @b } else { 2 } }`, "conditional expression branch must be a single expression"},

	// only the first error of a conditional expression is reported
	{`func _() { _ = if @x := f(); x { 1 } }`, "conditional expression cannot have an init statement"},
	{`func _() { _ = if c { if d { 1 } @} else { 2 } }`, "conditional expression requires an else branch"},
	{`func _() { _ = if c { @return } else if d { 1 } }`, "conditional expression branch must be a single expression"},
	{`func _() { _ = if c { 1 } else { x @:= 2 }; _ = if d { 1 } else { 2 } }`, "conditional expression branch must be a single expression"},

	// init statement
	{`func _() { _ = if @v, ok := m[k]; ok { v } else { 0 } }`, "conditional expression cannot have an init statement"},
	{`func _() { _ = if @x := f(); x { 1 } else { 2 } }`, "conditional expression cannot have an init statement"},
	{`func _() { _ = if @f(); c { 1 } else { 2 } }`, "conditional expression cannot have an init statement"},
}

func TestCondExprErrors(t *testing.T) {
	for _, test := range condExprErrors {
		src, index := stripAt("package p; " + test.src)
		var errs []Error
		Parse(NewFileBase("test.go"), strings.NewReader(src), func(err error) { errs = append(errs, err.(Error)) }, nil, CheckBranches)
		if len(errs) != 1 {
			t.Errorf("%s: got %d errors, want 1: %v", test.src, len(errs), errs)
			continue
		}
		err := errs[0]
		if err.Msg != "syntax error: "+test.msg {
			t.Errorf("%s: got %q, want %q", test.src, err.Msg, test.msg)
		}
		if index >= 0 {
			line, col := lineCol(src, index)
			if err.Pos.Line() != line || err.Pos.Col() != col {
				t.Errorf("%s: error at %d:%d, want %d:%d", test.src, err.Pos.Line(), err.Pos.Col(), line, col)
			}
		}
	}
}

// lineCol returns the 1-based line and column of the byte offset in src.
func lineCol(src string, offset int) (line, col uint) {
	line = 1 + uint(strings.Count(src[:offset], "\n"))
	col = 1 + uint(offset-(strings.LastIndex(src[:offset], "\n")+1))
	return
}

// Invalid conditional expressions must produce errors and a complete
// syntax tree, never a crash, whatever follows them.
func TestCondExprRecovery(t *testing.T) {
	for _, src := range []string{
		`var _ = if`,
		`var _ = if c`,
		`var _ = if c {`,
		`var _ = if c { a`,
		`var _ = if c { a }`,
		`var _ = if c { a } else`,
		`var _ = if c { a } else {`,
		`var _ = if c { a } else { b`,
		`var _ = if { a } else { b }`,
		`var _ = if c a else b`,
		`var _ = if c { a } else b`,
		`var _ = if c { a } else if`,
		`var _ = if c { a } else if d`,
		`var _ = if c { { x := 1 } } else { b }`,
		`var _ = if c { a ) } else { b }`,
		`var _ = if c { a } else { b ] }`,
		`var _ = if var x = 1; x { a } else { b }`,
		`func _() { x := if c { a }
else { b } }`,
		`func _() { x := if c {
	return
} else {
	1
}; _ = x }`,
	} {
		var errs []Error
		f, _ := Parse(NewFileBase("test.go"), strings.NewReader("package p; "+src), func(err error) { errs = append(errs, err.(Error)) }, nil, CheckBranches)
		if len(errs) == 0 {
			t.Errorf("%s: no error reported", src)
		}
		if f == nil {
			continue
		}
		Inspect(f, func(n Node) bool {
			if e, ok := n.(*CondExpr); ok {
				if e.Cond == nil || e.Then == nil || e.Else == nil {
					t.Errorf("%s: incomplete CondExpr", src)
				}
				_ = fmt.Sprint(String(e), StartPos(e), EndPos(e))
			}
			return true
		})
	}
}

// The approved newline rule for ! also applies inside branches.
func TestCondExprNegationNewline(t *testing.T) {
	src := "package p; var _ = if c { !\nd } else { e }"
	if _, err := Parse(NewFileBase("test.go"), strings.NewReader(src), nil, nil, 0); err == nil {
		t.Errorf("accepted newline after prefix ! in a branch")
	}
}
