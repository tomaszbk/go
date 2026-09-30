//go:build !go1.21

package cgocall

import "go/types"

func setGoVersion(tc *types.Config, pkg *types.Package) {
	// no types.Package.GoVersion until Go 1.21
}
