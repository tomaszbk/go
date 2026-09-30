package diagnostics

import (
	"strings"
	"testing"

	. "golang.org/x/tools/gopls/internal/test/integration"
)

func TestIssue44866(t *testing.T) {
	src := `
-- go.mod --
module mod.com

go 1.12
-- a.go --
package a

const (
	c = iota
)
`
	Run(t, src, func(t *testing.T, env *Env) {
		env.OpenFile("a.go")
		loc := env.FirstDefinition(env.RegexpSearch("a.go", "iota"))
		if !strings.HasSuffix(string(loc.URI), "builtin.go") {
			t.Fatalf("jumped to %q, want builtin.go", loc.URI)
		}
		env.AfterChange(NoDiagnostics(ForFile("builtin.go")))
	})
}
