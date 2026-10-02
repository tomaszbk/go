package astutil_test

import (
	"go/ast"
	"go/parser"
	"honnef.co/go/tools/go/ast/astutil"
	"testing"
)

type unknownExpr struct{ ast.Expr }

func TestGonTransformSafety(t *testing.T) {
	for _, src := range []string{"if c { 1 } else { 2 }", "f()!", "f() or err { panic(err) }"} {
		a, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatal(err)
		}
		b, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := astutil.CopyExpr(a); ok {
			t.Fatalf("unexpected unsupported copy: %s", src)
		}
		if astutil.Equal(a, b) {
			t.Fatalf("unproven equality: %s", src)
		}
	}
	if _, ok := astutil.CopyExpr(&unknownExpr{}); ok {
		t.Fatal("unknown copy accepted")
	}
	if astutil.Equal(&unknownExpr{}, &unknownExpr{}) {
		t.Fatal("unknown equality accepted")
	}
	a, _ := parser.ParseExpr("x + 1")
	b, ok := astutil.CopyExpr(a)
	if !ok || !astutil.Equal(a, b) {
		t.Fatal("legacy copy/equality regressed")
	}
}
