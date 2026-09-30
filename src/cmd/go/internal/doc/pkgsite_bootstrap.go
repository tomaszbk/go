//go:build cmd_go_bootstrap

// Don't build the pkgsite code into go_bootstrap because it depends on net.

package doc

import "context"

func doPkgsite(context.Context, string, string) error { return nil }
