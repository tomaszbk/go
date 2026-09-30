// Keep in sync with notboring.go and crypto/internal/boring/boring.go.
//go:build boringcrypto && linux && (amd64 || arm64) && !android && !msan && cgo

package fips140

const boringEnabled = true
