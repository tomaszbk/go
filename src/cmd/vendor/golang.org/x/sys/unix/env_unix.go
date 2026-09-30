//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

// Unix environment variables.

package unix

import "syscall"

func Getenv(key string) (value string, found bool) {
	return syscall.Getenv(key)
}

func Setenv(key, value string) error {
	return syscall.Setenv(key, value)
}

func Clearenv() {
	syscall.Clearenv()
}

func Environ() []string {
	return syscall.Environ()
}

func Unsetenv(key string) error {
	return syscall.Unsetenv(key)
}
