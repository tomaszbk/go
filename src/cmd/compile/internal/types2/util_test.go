// This file exports various functionality of util.go
// so that it can be used in (package-external) tests.

package types2

import (
	"cmd/compile/internal/syntax"
)

func CmpPos(p, q syntax.Pos) int { return cmpPos(p, q) }

func ScopeComment(s *Scope) string         { return s.comment }
func ObjectScopePos(obj Object) syntax.Pos { return obj.scopePos() }
