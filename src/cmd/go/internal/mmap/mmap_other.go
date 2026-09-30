//go:build (js && wasm) || wasip1 || plan9

package mmap

import (
	"io"
	"os"
)

// mmapFile on other systems doesn't mmap the file. It just reads everything.
func mmapFile(f *os.File) (Data, error) {
	b, err := io.ReadAll(f)
	if err != nil {
		return Data{}, err
	}
	return Data{f, b}, nil
}
