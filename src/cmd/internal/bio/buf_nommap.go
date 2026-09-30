//go:build !unix

package bio

func (r *Reader) sliceOS(length uint64) ([]byte, bool) {
	return nil, false
}
