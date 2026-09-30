//go:build tools

package tools

// Arrange to vendor the bisect command for use
// by the internal/godebug package test.
import _ "golang.org/x/tools/cmd/bisect"
