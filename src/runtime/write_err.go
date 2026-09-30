//go:build !android

package runtime

//go:nosplit
func writeErr(b []byte) {
	if len(b) > 0 {
		writeErrData(&b[0], int32(len(b)))
	}
}
