//go:build darwin || freebsd || netbsd

package tar

import (
	"syscall"
	"time"
)

func statAtime(st *syscall.Stat_t) time.Time {
	return time.Unix(st.Atimespec.Unix())
}

func statCtime(st *syscall.Stat_t) time.Time {
	return time.Unix(st.Ctimespec.Unix())
}
