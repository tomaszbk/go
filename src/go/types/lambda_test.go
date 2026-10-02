package types_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	. "go/types"
	"strings"
	"testing"
)

func TestLambdaTargetsAndInference(t *testing.T) {
	src := `package p
func Map[T,U any](xs []T, f func(T) U) []U { return nil }
func Pick[T any](f func() T,x T) T { return x }
func Pair[T,U any](x T,y U,f func(T) U) U { return y }
func Twice[T any,F ~func(T) T](x T,f F) T { return f(x) }
func Reduce[T,A any](xs []T,a A,f func(A,T) A) A { return a }
var f func(int) int = (x) => x+1
var g func(int) func(int) int = (x) => (y) => x+y
var _ = Map([]int{1}, (x) => x+1)
var _ = Map([]int{1}, (x) => 0.5)
var _ = Pick(() => 1,2.5)
var _ = Pick(() => { return 'a' },0)
var _ = Pair(0,1,(x) => float64(x))
var _ = Twice(3,(x) => x*2)
var _ = Reduce([]int{1},0,(acc,x) => acc+x)
var _ = append([]func(int)int{}, (x) => x)
var _ func(...int) int = (xs) => len(xs)
`
	fset := token.NewFileSet()
	f := mustParse(fset, src)
	conf := Config{Importer: defaultImporter(fset)}
	if _, err := conf.Check("p", fset, []*ast.File{f}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestNilSafetyTypes(t *testing.T) {
	src := `package p
 type S struct { N int; P *int; Next *S; F func() int }
 var p *S
 var _ = p?.N ?? 0
 var _ = p?.P
 var _ = p?.Next?.N ?? 0
 var _ any = p?.P
 var _ = *p?.P ?? 0
 var _ = p?.F?() ?? 0
 var _ = func(){ p?.F(); p?.F?(); p.P ??= new(int) }
 var _ = p ?? new(S)
 `
	fset := token.NewFileSet()
	f := mustParse(fset, src)
	conf := Config{}
	if _, err := conf.Check("p", fset, []*ast.File{f}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestLambdaAndNilSafetyInvalid(t *testing.T) {
	cases := []string{
		`var _ = (x) => x`,
		`var _ any = () => 1`,
		`var _ func(int)int = ((x) => x)`,
		`var _ func(int,int)int = (x) => x`,
		`var _ func(int,int)int = (x,x) => x`,
		`var _ func() = () => 1`,
		`var _ func()int = () => { return }`,
		`func Apply[T any](func(T)T)T { panic(0) }; var _ = Apply((x) => x)`,
		`func Map[T,U any]([]T,func(T)U)U { panic(0) }; var _ = Map([]int{},(x) => {return x})`,
		`var p *struct{N int}; var _ = p?.N`,
		`var p *struct{N int}; var _ = &p?.N`,
		`var p *struct{N int}; func f(){p?.N=1}`,
		`var p *struct{F func()};func f(){defer p?.F()}`,
		`var p *struct{F func()};func f(){go p?.F()}`,
		`var p *struct{N int};var _ = p ?? 0`,
		`var _ = nil ?? new(int)`,
		`var _ = 1 ?? 2`,
		`var a any;var _ = a.(int) ?? 0`,
		`var a int;func f(){a ??= 1}`,
		`type F func();func(F) M(){};var f F;func g(){f?.M()}`,
	}
	for _, src := range cases {
		t.Run(src, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "p.go", strings.NewReader("package p;"+src), 0)
			if err != nil {
				return
			}
			conf := Config{Error: func(error) {}}
			if _, err = conf.Check("p", fset, []*ast.File{f}, nil); err == nil {
				t.Fatal("invalid program accepted")
			}
		})
	}
}

func TestLambdaNilSafetyInfo(t *testing.T) {
	src := `package p
 type E struct{}
 func (*E) Error()string{return ""}
 type S struct{Err *E; Count int}
 var s *S
 var _ error = s?.Err
 var f func(target int)(named string) = (actual) => "ok"
 func Map[T,U any]([]T,func(T)U)U {panic(0)}
 var _ = Map([]int{},(n) => 0.5)
 `
	fset := token.NewFileSet()
	f := mustParse(fset, src)
	info := &Info{Types: map[ast.Expr]TypeAndValue{}, Defs: map[*ast.Ident]Object{}, Scopes: map[ast.Node]*Scope{}}
	if _, err := (&Config{}).Check("p", fset, []*ast.File{f}, info); err != nil {
		t.Fatal(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch e := n.(type) {
		case *ast.LambdaExpr:
			sig := info.Types[e].Type.(*Signature)
			if info.Scopes[e] == nil {
				t.Error("missing lambda scope")
			}
			for i, id := range e.Params {
				v, ok := info.Defs[id].(*Var)
				if !ok || v != sig.Params().At(i) || v.Kind() != ParamVar {
					t.Errorf("parameter identity not preserved: %v", id)
				}
			}
			for i := range sig.Results().Len() {
				if sig.Results().At(i).Name() != "" {
					t.Error("target result names leaked")
				}
			}
		case *ast.NilGuardExpr:
			if info.Types[e].Type.String() != "*p.S" || info.Types[e].Addressable() {
				t.Errorf("guard type/mode: %v", info.Types[e])
			}
		case *ast.SafeNavExpr:
			if info.Types[e].Type.String() != "error" || info.Types[e.X].Type.String() != "*p.E" {
				t.Errorf("chain lost intrinsic or contextual type: %v, %v", info.Types[e], info.Types[e.X])
			}
		}
		return true
	})
}
