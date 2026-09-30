//go:build js || wasip1

package toolchain

import "cmd/go/internal/base"

func execGoToolchain(gotoolchain, dir, exe string) {
	base.Fatalf("execGoToolchain unsupported")
}
