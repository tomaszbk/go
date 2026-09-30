// NOTE: If you change this file you must run "go generate"
// in cmd/compile/internal/typecheck
// to update builtin.go. This is not done automatically
// to avoid depending on having a working compiler binary.

//go:build ignore

package coverage

func initHook(istest bool)
