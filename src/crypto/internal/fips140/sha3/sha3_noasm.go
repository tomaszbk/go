//go:build (!amd64 && !arm64 && !s390x) || purego

package sha3

func keccakF1600(a *[200]byte) {
	keccakF1600Generic(a)
}

func (d *Digest) write(p []byte) (n int, err error) {
	return d.writeGeneric(p)
}
func (d *Digest) read(out []byte) (n int, err error) {
	return d.readGeneric(out)
}
func (d *Digest) sum(b []byte) []byte {
	return d.sumGeneric(b)
}
