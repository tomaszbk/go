// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lib

import "errors"

var Failure = errors.New("export failure")

func Source(fail bool) (int, error) {
	if fail {
		return 99, Failure
	}
	return 7, nil
}

func Pass(fail bool) (int, error) {
	v, err := Source(fail)
	if err != nil {
		return 0, err
	}
	return v + 1, nil
}

func Local(fail bool) (int, error) {
	v, err := Source(fail)
	if err != nil {
		return -1, err
	}
	return v, nil
}

func Generic[T any](v T, fail bool) (T, error) {
	if _, err := Source(fail); err != nil {
		var zero T
		return zero, err
	}
	return v, nil
}
