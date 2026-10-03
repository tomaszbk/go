package gonerrors

import "io"

func importedZero() (int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		return io.SeekStart, err
	}
	return x, nil
}
