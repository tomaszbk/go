//go:build !go1.26

package chacha20poly1305

func fips140Enforced() bool { return false }
