package cgroup

import (
	_ "unsafe" // for linkname
)

// Functions below pushed from runtime.

//go:linkname throw
func throw(s string)
