//go:build unix || wasip1

package os

func newDirFile(fd int, name string) (*File, error) {
	// We use kindNoPoll because we know that this is a directory.
	return newFile(fd, name, kindNoPoll, false), nil
}
