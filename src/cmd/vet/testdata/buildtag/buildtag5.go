// This file contains tests for the buildtag checker.

//go:build !(bad || worse)

package testdata

// +build other // ERROR `misplaced \+build comment`
