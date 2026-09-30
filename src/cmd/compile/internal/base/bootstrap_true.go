//go:build compiler_bootstrap

package base

// CompilerBootstrap reports whether the current compiler binary was
// built with -tags=compiler_bootstrap.
const CompilerBootstrap = true
