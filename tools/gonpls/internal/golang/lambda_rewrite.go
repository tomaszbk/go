package golang

import (
	"bytes"
	"context"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/gopls/internal/protocol"
	"golang.org/x/tools/internal/typesinternal"
)

// refactorRewriteLambda expands a lambda only when its inferred signature can
// be written at its original position, and no enclosing generic invocation can
// change inference when an explicitly typed literal replaces it.
func refactorRewriteLambda(ctx context.Context, req *codeActionsRequest) error {
	if len(req.pkg.ParseErrors()) != 0 || len(req.pkg.TypeErrors()) != 0 {
		return nil
	}
	cur, ok := req.pgf.Cursor().FindByPos(req.start, req.end)
	if !ok {
		return nil
	}
	var lambda *ast.LambdaExpr
	var literal *ast.FuncLit
	var functionCur inspector.Cursor
	for c := range cur.Enclosing() {
		switch n := c.Node().(type) {
		case *ast.LambdaExpr:
			if lambda == nil && literal == nil {
				lambda, functionCur = n, c
			}
		case *ast.FuncLit:
			if lambda == nil && literal == nil {
				literal, functionCur = n, c
			}
		case *ast.CallExpr:
			var id *ast.Ident
			fun := ast.Unparen(n.Fun)
			for {
				if index, ok := fun.(*ast.IndexExpr); ok {
					fun = ast.Unparen(index.X)
					continue
				}
				if index, ok := fun.(*ast.IndexListExpr); ok {
					fun = ast.Unparen(index.X)
					continue
				}
				break
			}
			switch f := fun.(type) {
			case *ast.Ident:
				id = f
			case *ast.SelectorExpr:
				id = f.Sel
			}
			if id != nil && req.pkg.TypesInfo().Instances[id].TypeArgs != nil {
				return nil
			}
		}
	}
	if literal != nil {
		return rewriteFunctionLiteral(req, functionCur, literal)
	}
	if lambda == nil {
		return nil
	}
	typ := req.pkg.TypesInfo().TypeOf(lambda)
	if typ == nil {
		return nil
	}
	sig, ok := typ.Underlying().(*types.Signature)
	if !ok || sig.TypeParams().Len() != 0 {
		return nil
	}
	// Preserve comments in the signature and between the arrow and the body
	// by declining a rewrite that would need to relocate them.
	bodyPos := lambda.End()
	if lambda.Body != nil {
		bodyPos = lambda.Body.Pos()
	}
	if lambda.Block != nil {
		bodyPos = lambda.Block.Pos()
	}
	for _, group := range req.pgf.File.Comments {
		if lambda.Pos() <= group.Pos() && group.Pos() < bodyPos {
			return nil
		}
	}
	qual := typesinternal.FileQualifier(req.pgf.File, req.pkg.Types())
	// Eval checks both accessibility and shadowing for every type we spell.
	// Parameter names only scope the body, so evaluate just before the lambda.
	for _, tuple := range []*types.Tuple{sig.Params(), sig.Results()} {
		for v := range tuple.Variables() {
			tv, err := types.Eval(req.pkg.FileSet(), req.pkg.Types(), lambda.Pos()-1, types.TypeString(v.Type(), qual))
			if err != nil || !tv.IsType() || !types.Identical(tv.Type, v.Type()) {
				return nil
			}
		}
	}
	syntax, err := parser.ParseExpr(types.TypeString(sig, qual))
	if err != nil {
		return nil
	}
	var header bytes.Buffer
	if err := format.Node(&header, token.NewFileSet(), syntax); err != nil {
		return nil
	}
	replacement := header.String()
	body := ast.Node(lambda.Body)
	if lambda.Block != nil {
		body = lambda.Block
	}
	start, end, err := req.pgf.NodeOffsets(body)
	if err != nil {
		return err
	}
	source := string(req.pgf.Src[start:end])
	if lambda.Block != nil {
		replacement += " " + source
	} else {
		prefix := " { "
		if sig.Results().Len() > 0 {
			prefix += "return "
		}
		replacement += prefix + source + "\n}"
	}
	// Validate/format through a file so interior comments are preserved.
	const prefix = "package p\nvar _ = "
	formatted, err := format.Source([]byte(prefix + replacement + "\n"))
	if err != nil {
		return nil
	}
	replacement = strings.TrimSuffix(strings.TrimPrefix(string(formatted), "package p\n\nvar _ = "), "\n")
	rng, err := req.pgf.NodeRange(lambda)
	if err != nil {
		return err
	}
	req.addEditAction("Convert lambda to function literal", nil, protocol.DocumentChangeEdit(req.fh, []protocol.TextEdit{{Range: rng, NewText: replacement}}))
	return nil
}

// rewriteFunctionLiteral keeps the block body and only removes a signature
// when a direct typed declaration or assignment supplies the identical one.
// Interfaces, inferred declarations, parentheses, and named results are excluded.
func rewriteFunctionLiteral(req *codeActionsRequest, cur inspector.Cursor, literal *ast.FuncLit) error {
	info := req.pkg.TypesInfo()
	actual, ok := info.TypeOf(literal).(*types.Signature)
	if !ok {
		return nil
	}
	var target types.Type
	switch parent := cur.Parent().Node().(type) {
	case *ast.ValueSpec:
		if parent.Type == nil {
			return nil
		}
		target = info.TypeOf(parent.Type)
	case *ast.AssignStmt:
		if parent.Tok != token.ASSIGN || len(parent.Lhs) != len(parent.Rhs) {
			return nil
		}
		for i, rhs := range parent.Rhs {
			if rhs == literal {
				target = info.TypeOf(parent.Lhs[i])
			}
		}
	default:
		return nil
	}
	if target == nil {
		return nil
	}
	if _, ok := target.Underlying().(*types.Signature); !ok || !types.Identical(target.Underlying(), actual) {
		return nil
	}
	if literal.Type.Results != nil {
		for _, field := range literal.Type.Results.List {
			if len(field.Names) != 0 {
				return nil
			}
		}
	}
	// Removing a signature must not remove an import's last use. Dot imports
	// require additional binding analysis, so this conservative rewrite declines
	// them. Ordinary import bindings can be checked directly in Info.Uses.
	for _, spec := range req.pgf.File.Imports {
		if spec.Name != nil && spec.Name.Name == "." {
			return nil
		}
	}
	removedImports := make(map[*types.PkgName]bool)
	for id, obj := range info.Uses {
		if pkg, ok := obj.(*types.PkgName); ok && literal.Type.Pos() <= id.Pos() && id.Pos() < literal.Type.End() {
			removedImports[pkg] = true
		}
	}
	for id, obj := range info.Uses {
		if pkg, ok := obj.(*types.PkgName); ok && (id.Pos() < literal.Type.Pos() || literal.Type.End() <= id.Pos()) {
			delete(removedImports, pkg)
		}
	}
	if len(removedImports) != 0 {
		return nil
	}
	var names []string
	for _, field := range literal.Type.Params.List {
		if len(field.Names) == 0 {
			return nil
		}
		for _, name := range field.Names {
			names = append(names, name.Name)
		}
	}
	for _, group := range req.pgf.File.Comments {
		if literal.Pos() <= group.Pos() && group.Pos() < literal.Body.Pos() {
			return nil
		}
	}
	start, end, err := req.pgf.NodeOffsets(literal.Body)
	if err != nil {
		return err
	}
	replacement := "(" + strings.Join(names, ", ") + ") => " + string(req.pgf.Src[start:end])
	rng, err := req.pgf.NodeRange(literal)
	if err != nil {
		return err
	}
	req.addEditAction("Convert function literal to lambda", nil, protocol.DocumentChangeEdit(req.fh, []protocol.TextEdit{{Range: rng, NewText: replacement}}))
	return nil
}
