//go:build !ppc64 && !ppc64le

package runtime

func prepGoExitFrame(sp uintptr) {
}
