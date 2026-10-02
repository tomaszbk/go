package types2

import (
	"cmd/compile/internal/syntax"
	. "internal/types/errors"
)

// lambdaExpr fixes the signature before entering the function's own scope.
func (check *Checker) lambdaExpr(T *target, x *operand, e *syntax.LambdaExpr) {
	x.expr = e
	if T != nil && T.kind == inferTarget {
		x.mode_, x.typ_ = value, T.typ
		return
	}
	if T == nil || T.sig() == nil {
		check.error(e, InvalidLambda, "lambda requires a function type from context")
		check.recoverLambda(e, nil)
		return
	}
	u, _ := commonUnder(T.typ, nil)
	targetSig := u.(*Signature)
	if len(e.Params) != targetSig.params.Len() {
		check.errorf(e, InvalidLambda, "lambda has %d parameters, but %s has %d", len(e.Params), targetSig, targetSig.params.Len())
		check.recoverLambda(e, targetSig)
		return
	}
	sig := check.lambdaSignature(e, targetSig)
	check.finishLambda(e, sig, false)
	x.mode_, x.typ_ = value, sig
}

func (check *Checker) lambdaSignature(e *syntax.LambdaExpr, targetSig *Signature) *Signature {
	scope := NewScope(check.scope, e.Pos(), endPos(e), "lambda")
	scope.isFunc = true
	check.recordScope(e, scope)
	params := make([]*Var, len(e.Params))
	for i, name := range e.Params {
		params[i] = newVar(ParamVar, name.Pos(), check.pkg, name.Value, targetSig.params.At(i).typ)
		check.declare(scope, name, params[i], e.Arrow)
	}
	results := make([]*Var, targetSig.results.Len())
	for i := range targetSig.results.Len() {
		v := targetSig.results.At(i)
		results[i] = newVar(ResultVar, e.Pos(), check.pkg, "", v.typ)
	}
	sig := &Signature{scope: scope, params: NewTuple(params...), results: NewTuple(results...), variadic: targetSig.variadic}
	scope.funcSig, scope.funcBodyPos = sig, e.Arrow
	return sig
}

func (check *Checker) finishLambda(e *syntax.LambdaExpr, sig *Signature, checked bool) {
	if check.lambdaTypes == nil {
		check.lambdaTypes = make(map[*syntax.LambdaExpr]*Signature)
	}
	check.lambdaTypes[e] = sig
	if checked {
		check.later(func() { check.usage(sig.scope) }).describef(e, "lambda usage")
	}
	body := e.Block
	if body == nil {
		body = &syntax.BlockStmt{Rbrace: endPos(e.Body)}
		body.SetPos(e.Arrow)
		if sig.results.Len() == 0 {
			stmt := &syntax.ExprStmt{X: e.Body}
			stmt.SetPos(e.Body.Pos())
			body.List = []syntax.Stmt{stmt}
		} else {
			stmt := &syntax.ReturnStmt{Results: e.Body}
			stmt.SetPos(e.Body.Pos())
			body.List = []syntax.Stmt{stmt}
		}
	}
	check.lowerLambda(e, sig, body)
	if !checked && !check.conf.IgnoreFuncBodies {
		decl, iota := check.decl, check.iota
		check.later(func() {
			if e.Body != nil && sig.results.Len() == 0 {
				saved := check.environment
				check.environment = environment{decl: decl, scope: sig.scope, version: check.version, iota: iota, sig: sig}
				var out operand
				check.rawExpr(nil, &out, e.Body, false)
				if out.isValid() && out.mode() != novalue {
					check.errorf(e.Body, InvalidLambda, "lambda body returns %s, but %s has no results", out.typ(), sig)
				}
				check.usage(sig.scope)
				check.environment = saved
			} else {
				check.funcBody(decl, "<lambda>", sig, body, iota)
			}
		}).describef(e, "lambda")
	}
}

func (check *Checker) lowerLambda(e *syntax.LambdaExpr, sig *Signature, body *syntax.BlockStmt) {
	ft := new(syntax.FuncType)
	ft.SetPos(e.Pos())
	field := func(v *Var, name *syntax.Name) *syntax.Field {
		typ := syntax.NewName(e.Pos(), "<inferred>")
		check.recordTypeAndValue(typ, typexpr, v.typ, nil)
		f := &syntax.Field{Name: name, Type: typ}
		f.SetPos(e.Pos())
		return f
	}
	for i := range sig.params.Len() {
		v := sig.params.At(i)
		ft.ParamList = append(ft.ParamList, field(v, e.Params[i]))
	}
	for i := range sig.results.Len() {
		v := sig.results.At(i)
		ft.ResultList = append(ft.ResultList, field(v, nil))
	}
	check.recordTypeAndValue(ft, typexpr, sig, nil)
	lit := &syntax.FuncLit{Type: ft, Body: body}
	lit.SetPos(e.Pos())
	check.recordTypeAndValue(lit, value, sig, nil)
	check.recordScope(ft, sig.scope)
	e.Lowered = lit
}

// inferLambda checks an expression body once its input types are fixed. The
// unifier supplies output types later; operands retain their untyped values.
func (check *Checker) inferLambda(e *syntax.LambdaExpr, sig *Signature) []*operand {
	saved := check.environment
	prior := check.inferLambdaSig
	check.inferLambdaSig = sig
	defer func() { check.inferLambdaSig = prior }()
	check.environment = environment{decl: check.decl, scope: sig.scope, version: check.version, iota: check.iota, sig: sig}
	defer func() { check.environment = saved }()
	values, _ := check.multiExpr(e.Body, false)
	return values
}

// recoverLambda keeps parameter bindings and the function boundary available
// to editor clients even when no usable contextual signature exists.
func (check *Checker) recoverLambda(e *syntax.LambdaExpr, hint *Signature) {
	params := make([]*Var, len(e.Params))
	for i, n := range e.Params {
		typ := Type(Typ[Invalid])
		if hint != nil && i < hint.params.Len() {
			typ = hint.params.At(i).typ
		}
		params[i] = newVar(ParamVar, n.Pos(), check.pkg, "", typ)
	}
	results := NewTuple(newVar(ResultVar, e.Pos(), check.pkg, "", Typ[Invalid]))
	if hint != nil {
		results = hint.results
	}
	sig := check.lambdaSignature(e, &Signature{params: NewTuple(params...), results: results})
	check.finishLambda(e, sig, false)
}
