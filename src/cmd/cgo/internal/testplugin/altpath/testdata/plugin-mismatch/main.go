package main

// // No C code required.
import "C"

// The common package imported here does not match the common package
// imported by plugin1. A program that attempts to load plugin1 and
// plugin-mismatch should produce an error.
import "testplugin/common"

func ReadCommonX() int {
	return common.X
}
