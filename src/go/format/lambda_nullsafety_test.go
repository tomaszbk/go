package format

import "testing"

func TestLambdaNullSafetyFormat(t *testing.T) {
	for _, test := range []struct{ src, want string }{
		{`f((a,b)=>a*b+c)`, `f((a, b) => a*b + c)`},
		{`f(()=>{ return 1 })`, `f(() => { return 1 })`},
		{`f((x,)=>x)`, `f((x) => x)`},
		{"f((x)=>\nx+1)", "f((x) =>\n\tx + 1)"},
		{`x:=p?.Next?.M(a)[i].Field??fallback`, `x := p?.Next?.M(a)[i].Field ?? fallback`},
		{`x:=(p?.Next).Field`, `x := (p?.Next).Field`},
		{`x:=f?(g())`, `x := f?(g())`},
		{`x:=a??b??c`, `x := a ?? b ?? c`},
		{`x:=(a??b)??c`, `x := (a ?? b) ?? c`},
		{`m[key()]??=init()`, `m[key()] ??= init()`},
		{"x:=p?.\nNext??\nfallback", "x := p?.\n\tNext ??\n\tfallback"},
		{"x:=f?(\ng(),\n)", "x := f?(\n\tg(),\n)"},
		{`x:=a??(x)=>x+1`, `x := a ?? (x) => x + 1`},
		{`if (p?.Enabled??false)&&ready {}`, "if (p?.Enabled ?? false) && ready {\n}"},
	} {
		got, err := Source([]byte(test.src))
		if err != nil {
			t.Errorf("%q: %v", test.src, err)
			continue
		}
		if string(got) != test.want {
			t.Errorf("%q: got %q, want %q", test.src, got, test.want)
			continue
		}
		stable, err := Source(got)
		if err != nil || string(stable) != string(got) {
			t.Errorf("%q: unstable format %q: %v", test.src, stable, err)
		}
	}
}
