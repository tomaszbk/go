//go:build plan9

package x509

// Possible certificate files; stop after finding one.
var certFiles = []string{
	"/sys/lib/tls/ca.pem",
}

var certDirectories = []string{}

func (c *Certificate) systemVerify(opts *VerifyOptions) (chains [][]*Certificate, err error) {
	return nil, nil
}
