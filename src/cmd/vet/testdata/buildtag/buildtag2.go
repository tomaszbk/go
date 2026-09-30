// This file contains tests for the buildtag checker.

// ERRORNEXT `possible malformed [+]build comment`
// +builder
// +build !ignore

package testdata

// ERRORNEXT `misplaced \+build comment`
// +build toolate
// ERRORNEXT "misplaced //go:build comment"
//go:build toolate

var _ = 3

var _ = `
// +build notacomment
`
