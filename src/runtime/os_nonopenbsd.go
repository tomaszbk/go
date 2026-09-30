//go:build !openbsd

package runtime

// osStackAlloc performs OS-specific initialization before s is used
// as stack memory.
func osStackAlloc(s *mspan) {
}

// osStackFree undoes the effect of osStackAlloc before s is returned
// to the heap.
func osStackFree(s *mspan) {
}
