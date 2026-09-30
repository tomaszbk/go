//go:build !unix

package base

import (
	"io"
	"os"
)

func MapFile(f *os.File, offset, length int64) (string, error) {
	buf := make([]byte, length)
	_, err := io.ReadFull(io.NewSectionReader(f, offset, length), buf)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
