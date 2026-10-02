package cfg

import (
	"go/token"
	"testing"
)

func TestGonFeatureControlFlow(t *testing.T) {
	for _, test := range []struct {
		body     string
		noReturn bool
	}{
		{`var f func() error = () => { g()!; return nil }; panic(0)`, true},
		{`var f func() error = () => g()!; panic(0)`, true},
		{`_ = p?.M(g()!); panic(0)`, false},
		{`_ = p?.M(g() or err { panic(err) }); panic(0)`, true},
		{`_ = p ?? g()!; panic(0)`, false},
		{`p ??= g()!; panic(0)`, false},
		{`_ = **p ?? g()!; panic(0)`, false},
	} {
		t.Run(test.body, func(t *testing.T) {
			g := newCondCFG(t, test.body)
			if g.NoReturn() != test.noReturn {
				t.Fatalf("NoReturn = %v, want %v\n%s", g.NoReturn(), test.noReturn, g.Format(token.NewFileSet()))
			}
		})
	}
}

func TestGonNilAbsentPaths(t *testing.T) {
	g := newCondCFG(t, `_ = p?.M(arg()!) ?? fallback()!`)
	entry := g.Blocks[0]
	if len(entry.Succs) != 2 || entry.Succs[1].Kind != KindNilFallback {
		t.Fatalf("nil receiver must go directly to fallback\n%s", g.Format(token.NewFileSet()))
	}
	present, absent := entry.Succs[0], entry.Succs[1]
	if len(present.Succs) != 2 || present.Succs[0].Kind != KindErrorHandler {
		t.Fatal("missing argument propagation")
	}
	if reaches(absent, present.Succs[0]) {
		t.Fatal("nil receiver evaluates guarded argument")
	}
	if !reaches(present, absent) {
		t.Fatal("nil result of a present call cannot reach fallback")
	}

	g = newCondCFG(t, `_ = **p ?? fallback()`)
	entry = g.Blocks[0]
	if len(entry.Succs) != 2 || entry.Succs[1].Kind != KindNilFallback ||
		len(entry.Succs[0].Succs) != 2 || entry.Succs[0].Succs[1] != entry.Succs[1] {
		t.Fatalf("both guarded dereferences must share the fallback\n%s", g.Format(token.NewFileSet()))
	}
}
