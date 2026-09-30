//go:build cmd_go_bootstrap

package modfetch

import "golang.org/x/mod/module"

func useSumDB(mod module.Version) bool {
	return false
}

func lookupSumDB(mod module.Version) (string, []string, error) {
	panic("bootstrap")
}
