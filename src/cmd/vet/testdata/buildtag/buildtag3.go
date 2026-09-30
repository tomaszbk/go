// This file contains tests for the buildtag checker.

//go:build good
// ERRORNEXT "[+]build lines do not match //go:build condition"
// +build bad

package testdata

var _ = `
// +build notacomment
`
