//go:build wasm

package x509

// Possible certificate files; stop after finding one.
var certFiles = []string{}

// Possible directories with certificate files; all will be read.
var certDirectories = []string{}
