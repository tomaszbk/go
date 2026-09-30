//go:generate go run mkbuiltin.go

package typecheck

import "cmd/compile/internal/ir"

// Target is the package being compiled.
var Target *ir.Package
