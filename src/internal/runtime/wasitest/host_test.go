package wasi_test

import "flag"

var target string

func init() {
	// The dist test runner passes -target when running this as a host test.
	flag.StringVar(&target, "target", "", "")
}
