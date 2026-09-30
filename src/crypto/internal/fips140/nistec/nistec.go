// Package nistec implements the elliptic curves from NIST SP 800-186.
//
// This package uses fiat-crypto or specialized assembly and Go code for its
// backend field arithmetic (not math/big) and exposes constant-time, heap
// allocation-free, byte slice-based safe APIs. Group operations use modern and
// safe complete addition formulas where possible. The point at infinity is
// handled and encoded according to SEC 1, Version 2.0, and invalid curve points
// can't be represented.
package nistec

import _ "crypto/internal/fips140/check"

//go:generate go run generate.go
