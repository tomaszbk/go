package types2

import (
	"strings"
	"testing"
)

var myInt = func() Type {
	tname := NewTypeName(nopos, nil, "myInt", nil)
	return NewNamed(tname, Typ[Int], nil)
}()

var testTerms = map[string]*term{
	"∅":       nil,
	"𝓤":       {},
	"int":     {false, Typ[Int]},
	"~int":    {true, Typ[Int]},
	"string":  {false, Typ[String]},
	"~string": {true, Typ[String]},
	"myInt":   {false, myInt},
}

func TestTermString(t *testing.T) {
	for want, x := range testTerms {
		if got := x.String(); got != want {
			t.Errorf("%v.String() == %v; want %v", x, got, want)
		}
	}
}

func split(s string, n int) []string {
	r := strings.Split(s, " ")
	if len(r) != n {
		panic("invalid test case: " + s)
	}
	return r
}

func testTerm(name string) *term {
	r, ok := testTerms[name]
	if !ok {
		panic("invalid test argument: " + name)
	}
	return r
}

func TestTermEqual(t *testing.T) {
	for _, test := range []string{
		"∅ ∅ T",
		"𝓤 𝓤 T",
		"int int T",
		"~int ~int T",
		"myInt myInt T",
		"∅ 𝓤 F",
		"∅ int F",
		"∅ ~int F",
		"𝓤 int F",
		"𝓤 ~int F",
		"𝓤 myInt F",
		"int ~int F",
		"int myInt F",
		"~int myInt F",
	} {
		args := split(test, 3)
		x := testTerm(args[0])
		y := testTerm(args[1])
		want := args[2] == "T"
		if got := x.equal(y); got != want {
			t.Errorf("%v.equal(%v) = %v; want %v", x, y, got, want)
		}
		// equal is symmetric
		x, y = y, x
		if got := x.equal(y); got != want {
			t.Errorf("%v.equal(%v) = %v; want %v", x, y, got, want)
		}
	}
}

func TestTermUnion(t *testing.T) {
	for _, test := range []string{
		"∅ ∅ ∅ ∅",
		"∅ 𝓤 𝓤 ∅",
		"∅ int int ∅",
		"∅ ~int ~int ∅",
		"∅ myInt myInt ∅",
		"𝓤 𝓤 𝓤 ∅",
		"𝓤 int 𝓤 ∅",
		"𝓤 ~int 𝓤 ∅",
		"𝓤 myInt 𝓤 ∅",
		"int int int ∅",
		"int ~int ~int ∅",
		"int string int string",
		"int ~string int ~string",
		"int myInt int myInt",
		"~int ~string ~int ~string",
		"~int myInt ~int ∅",

		// union is symmetric, but the result order isn't - repeat symmetric cases explicitly
		"𝓤 ∅ 𝓤 ∅",
		"int ∅ int ∅",
		"~int ∅ ~int ∅",
		"myInt ∅ myInt ∅",
		"int 𝓤 𝓤 ∅",
		"~int 𝓤 𝓤 ∅",
		"myInt 𝓤 𝓤 ∅",
		"~int int ~int ∅",
		"string int string int",
		"~string int ~string int",
		"myInt int myInt int",
		"~string ~int ~string ~int",
		"myInt ~int ~int ∅",
	} {
		args := split(test, 4)
		x := testTerm(args[0])
		y := testTerm(args[1])
		want1 := testTerm(args[2])
		want2 := testTerm(args[3])
		if got1, got2 := x.union(y); !got1.equal(want1) || !got2.equal(want2) {
			t.Errorf("%v.union(%v) = %v, %v; want %v, %v", x, y, got1, got2, want1, want2)
		}
	}
}

func TestTermIntersection(t *testing.T) {
	for _, test := range []string{
		"∅ ∅ ∅",
		"∅ 𝓤 ∅",
		"∅ int ∅",
		"∅ ~int ∅",
		"∅ myInt ∅",
		"𝓤 𝓤 𝓤",
		"𝓤 int int",
		"𝓤 ~int ~int",
		"𝓤 myInt myInt",
		"int int int",
		"int ~int int",
		"int string ∅",
		"int ~string ∅",
		"int string ∅",
		"~int ~string ∅",
		"~int myInt myInt",
	} {
		args := split(test, 3)
		x := testTerm(args[0])
		y := testTerm(args[1])
		want := testTerm(args[2])
		if got := x.intersect(y); !got.equal(want) {
			t.Errorf("%v.intersect(%v) = %v; want %v", x, y, got, want)
		}
		// intersect is symmetric
		x, y = y, x
		if got := x.intersect(y); !got.equal(want) {
			t.Errorf("%v.intersect(%v) = %v; want %v", x, y, got, want)
		}
	}
}

func TestTermIncludes(t *testing.T) {
	for _, test := range []string{
		"∅ int F",
		"𝓤 int T",
		"int int T",
		"~int int T",
		"~int myInt T",
		"string int F",
		"~string int F",
		"myInt int F",
	} {
		args := split(test, 3)
		x := testTerm(args[0])
		y := testTerm(args[1]).typ
		want := args[2] == "T"
		if got := x.includes(y); got != want {
			t.Errorf("%v.includes(%v) = %v; want %v", x, y, got, want)
		}
	}
}

func TestTermSubsetOf(t *testing.T) {
	for _, test := range []string{
		"∅ ∅ T",
		"𝓤 𝓤 T",
		"int int T",
		"~int ~int T",
		"myInt myInt T",
		"∅ 𝓤 T",
		"∅ int T",
		"∅ ~int T",
		"∅ myInt T",
		"𝓤 int F",
		"𝓤 ~int F",
		"𝓤 myInt F",
		"int ~int T",
		"int myInt F",
		"~int myInt F",
		"myInt int F",
		"myInt ~int T",
	} {
		args := split(test, 3)
		x := testTerm(args[0])
		y := testTerm(args[1])
		want := args[2] == "T"
		if got := x.subsetOf(y); got != want {
			t.Errorf("%v.subsetOf(%v) = %v; want %v", x, y, got, want)
		}
	}
}

func TestTermDisjoint(t *testing.T) {
	for _, test := range []string{
		"int int F",
		"~int ~int F",
		"int ~int F",
		"int string T",
		"int ~string T",
		"int myInt T",
		"~int ~string T",
		"~int myInt F",
		"string myInt T",
		"~string myInt T",
	} {
		args := split(test, 3)
		x := testTerm(args[0])
		y := testTerm(args[1])
		want := args[2] == "T"
		if got := x.disjoint(y); got != want {
			t.Errorf("%v.disjoint(%v) = %v; want %v", x, y, got, want)
		}
		// disjoint is symmetric
		x, y = y, x
		if got := x.disjoint(y); got != want {
			t.Errorf("%v.disjoint(%v) = %v; want %v", x, y, got, want)
		}
	}
}
