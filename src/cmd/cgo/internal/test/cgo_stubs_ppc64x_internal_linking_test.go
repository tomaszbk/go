//go:build (ppc64 || ppc64le) && internal

package cgotest

import "testing"

// If gcc is used, and linking internally, __mulsc3 and __muldc3
// will be linked in from libgcc which make several R_PPC64_TOC16_DS
// relocations which may not be resolvable with the internal linker.
func test8694(t *testing.T) { t.Skip("not supported on ppc64/ppc64le with internal linking") }
func test9510(t *testing.T) { t.Skip("not supported on ppc64/ppc64le with internal linking") }
