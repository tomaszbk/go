//go:build aix || linux || dragonfly || openbsd || solaris

package tar

import (
	"syscall"
	"time"
)

func statAtime(st *syscall.Stat_t) time.Time {
	return time.Unix(st.Atim.Unix())
}

func statCtime(st *syscall.Stat_t) time.Time {
	return time.Unix(st.Ctim.Unix())
}
