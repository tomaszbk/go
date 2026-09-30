//go:build aix || darwin || dragonfly || freebsd || wasip1 || linux || netbsd || openbsd || solaris

package os

import (
	"internal/syscall/unix"
)

func (f *File) lstatatNolog(name string) (FileInfo, error) {
	var fs fileStat
	if err := f.pfd.Fstatat(name, &fs.sys, unix.AT_SYMLINK_NOFOLLOW); err != nil {
		err = f.wrapErr("fstatat", err)
		if pe, ok := err.(*PathError); ok {
			pe.Path = pe.Path + string(PathSeparator) + name
		}
		return nil, err
	}
	fillFileStatFromSys(&fs, name)
	return &fs, nil
}
