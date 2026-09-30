package midway

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/syntax"
)

// CheckPositions checks that all nodes in the files have known positions.
// This converts lack-of-Pos into an early fatal error instead of a later
// weird downstream error (e.g., in the linker, in debugging information).
func CheckPositions(files []*syntax.File, phase string) {
	for _, file := range files {
		syntax.Inspect(file, func(n syntax.Node) bool {
			if n == nil {
				return true
			}
			if !n.Pos().IsKnown() {
				base.Fatalf("Phase %s, Node without known position: %T\n", phase, n)
			}
			return true
		})
	}
}
