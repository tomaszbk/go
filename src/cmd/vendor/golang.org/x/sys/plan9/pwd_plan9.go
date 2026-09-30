package plan9

import "syscall"

func fixwd() {
	syscall.Fixwd()
}

func Getwd() (wd string, err error) {
	return syscall.Getwd()
}

func Chdir(path string) error {
	return syscall.Chdir(path)
}
