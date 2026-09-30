//go:build !windows

package os

import (
	"internal/testlog"
)

func (f *File) lstatat(name string) (FileInfo, error) {
	if stathook != nil {
		fi, err := stathook(f, name)
		if fi != nil || err != nil {
			return fi, err
		}
	}
	if log := testlog.Logger(); log != nil {
		log.Stat(joinPath(f.Name(), name))
	}
	return f.lstatatNolog(name)
}
