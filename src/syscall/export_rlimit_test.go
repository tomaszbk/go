//go:build unix

package syscall

import "sync/atomic"

func OrigRlimitNofile() *Rlimit {
	return origRlimitNofile.Load()
}

func GetInternalOrigRlimitNofile() *atomic.Pointer[Rlimit] {
	return &origRlimitNofile
}
