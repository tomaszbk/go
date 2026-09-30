//go:build go1.21

package cgocall

import "go/types"

func setGoVersion(tc *types.Config, pkg *types.Package) {
	tc.GoVersion = pkg.GoVersion()
}
