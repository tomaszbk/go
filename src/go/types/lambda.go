package types

import (
	"go/ast"
	. "internal/types/errors"
)

// lambdaExpr fixes the signature before entering the function's own scope.
func (check *Checker) lambdaExpr(T *target, x *operand, e *ast.LambdaExpr) {
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

func (check *Checker) lambdaSignature(e *ast.LambdaExpr, targetSig *Signature) *Signature {
	scope := NewScope(check.scope, e.Pos(), endPos(e), "lambda")
	scope.isFunc = true
	check.recordScope(e, scope)
	params := make([]*Var, len(e.Params))
	for i, name := range e.Params {
		params[i] = newVar(ParamVar, name.Pos(), check.pkg, name.Name, targetSig.params.At(i).typ)
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

func (check *Checker) finishLambda(e *ast.LambdaExpr, sig *Signature, checked bool) {
	if check.lambdaTypes == nil {
		check.lambdaTypes = make(map[*ast.LambdaExpr]*Signature)
	}
	check.lambdaTypes[e] = sig
	if checked {
		check.later(func() { check.usage(sig.scope) }).describef(e, "lambda usage")
	}
	body := e.Block
	if body == nil {
		body = &ast.BlockStmt{Lbrace: e.Arrow, Rbrace: endPos(e.Body)}
		if sig.results.Len() == 0 {
			stmt := &ast.ExprStmt{X: e.Body}
			body.List = []ast.Stmt{stmt}
		} else {
			stmt := &ast.ReturnStmt{Return: e.Body.Pos(), Results: []ast.Expr{e.Body}}
			body.List = []ast.Stmt{stmt}
		}
	}

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

// inferLambda checks an expression body once its input types are fixed. The
// unifier supplies output types later; operands retain their untyped values.
func (check *Checker) inferLambda(e *ast.LambdaExpr, sig *Signature) []*operand {
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
func (check *Checker) recoverLambda(e *ast.LambdaExpr, hint *Signature) {
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
