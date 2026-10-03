package gonerrors

import . "io"

func dotImportedZero() (int, error) {
	// want +1 "replace error check with a Gon or handler"
	x, err := read()
	if err != nil {
		return SeekStart, err
	}
	return x, nil
}
