//go:build !unix

package runtime

func checkfds() {
	// Nothing to do on non-Unix platforms.
}
