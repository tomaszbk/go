//go:build aix || dragonfly || freebsd || (js && wasm) || linux || netbsd || openbsd || solaris || wasip1

package x509

func (c *Certificate) systemVerify(opts *VerifyOptions) (chains [][]*Certificate, err error) {
	return nil, nil
}
