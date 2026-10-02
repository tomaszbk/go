package main

import (
	"go/parser"
	"testing"
)

func TestGonFunctionBoundary(t *testing.T) {
	for _, test := range []struct {
		src                         string
		propagation, handler, block bool
	}{
		{`p?.M(read()!)`, true, false, false},
		{`p ?? read()!`, true, false, false},
		{`() => read()!`, false, false, false},
		{`() => { read()! }`, false, false, true},
		{`() => read() or err { return err }`, false, false, false},
		{`p ?? read() or err { return err }`, false, true, true},
	} {
		e, err := parser.ParseExpr(test.src)
		if err != nil {
			t.Fatal(err)
		}
		if got := hasPropagation(e); got != test.propagation {
			t.Errorf("%q: propagation=%v", test.src, got)
		}
		if got, _ := hasErrorHandler(e); got != test.handler {
			t.Errorf("%q: handler=%v", test.src, got)
		}
		if got, _ := hasNestedBlock(e); got != test.block {
			t.Errorf("%q: block=%v", test.src, got)
		}
	}
}
