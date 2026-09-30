//go:build !linux
// +build !linux

package main

func gettid() int {
	return 0
}

func tidExists(tid int) (exists, supported bool, err error) {
	return false, false, nil
}

func getcwd() (string, error) {
	return "", nil
}

func unshareFs() error {
	return nil
}

func chdir(path string) error {
	return nil
}
